package handlers

import (
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
	


	start := models.ToMinutes(schedule.StartTime)
	end := models.ToMinutes(schedule.EndTime)
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
	models.CreateSlots(newSchedule)
	c.JSON(201, gin.H{"schedule" : newSchedule})
}

