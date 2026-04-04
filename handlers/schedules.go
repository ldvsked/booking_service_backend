package handlers

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/internships-backend/test-backend-ldvsked/models"
)

type ScheduleRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
	RoomId uuid.UUID `json:"roomId" binding:"required"`

	DaysOfWeek []int `json:"daysOfWeek" binding:"required,min=1,max=7,dive,min=1,max=7"`

	StartTime string `json:"startTime" binding:"required"`
	EndTime string `json:"endTime" binding:"required"`
}

func CreateShedule(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(403, map[string]any{"error" : "You don't have rights to do it"})
		return
	}

	var schedule ScheduleRequest
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(400, map[string]any{"error" : "invalid data"})
		return
	}

	pathRoomId, err := uuid.Parse(c.Param("roomId"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid roomId in path"})
		return
	}

	if pathRoomId != schedule.RoomId {
		c.JSON(400, gin.H{"error" : "roomId in path and body must match"})
		return
	}
	


	start := toMinutes(schedule.StartTime)
	end := toMinutes(schedule.EndTime)
	if start == -1 || end == -1 || start >= end || (end - start) % 30 != 0 {
		c.JSON(400, gin.H{"error" : "invalid time"})
		return
	}

	for _, scheduleExists := range models.Schedules {
		if schedule.RoomId == scheduleExists.RoomId {
			c.JSON(409, gin.H{"error" : "for this room schedule already exists"})
			return
		}
	}
	
	var check bool = false
	for _, room := range models.Rooms {
		if schedule.RoomId == room.Id {
			check = true
			break
		}
	}
	if !check {
		c.JSON(404, gin.H{"error" : "we don't have this room"})
		return
	}

	var newSchedule = models.Schedule{Id:schedule.Id,
		RoomId: schedule.RoomId,
		DaysOfWeek: schedule.DaysOfWeek,
		StartTime: schedule.StartTime,
		EndTime: schedule.EndTime,
	}

	models.Schedules = append(models.Schedules, newSchedule)
	CreateSlots(newSchedule)
	c.JSON(201, gin.H{"schedule" : newSchedule})
}

func CreateSlots(schedule models.Schedule) {
	now := time.Now().UTC()
	for i := 0; i < 31; i++ {
		cur := now.AddDate(0, 0, i)
		weekday := int(cur.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		var check bool = false 
		for _, dayInSchedule := range schedule.DaysOfWeek {
			if weekday == dayInSchedule {
				check = true
				break
			}
			if dayInSchedule > weekday {
				break
			}
		}
		if !check { //такого нет в расписании
			continue
		}
		year, month, day := cur.Date()
		for curMIn := toMinutes(schedule.StartTime); curMIn < toMinutes(schedule.EndTime); curMIn += 30 {
			startDate := time.Date(year, month, day, curMIn / 60, curMIn % 60, 0, 0, time.UTC)
			endDate := time.Date(year, month, day, (curMIn + 30) / 60, (curMIn + 30) % 60, 0, 0, time.UTC)
			newSlot := models.Slot{
				Id : uuid.New(), 
				RoomId: schedule.RoomId,
				Start:	startDate, 
				End: endDate,
			}
			models.Slots = append(models.Slots, newSlot)
		}
	}
}

func toMinutes(t string) int {
	var h, m int 
	n,_ := fmt.Sscanf(t, "%d:%d", &h, &m)
	if n != 2 {
		return -1
	}
	if h < 0 || h >= 24 || m < 0 || m >= 60 {
		return -1
	}
	return h * 60 + m;
}