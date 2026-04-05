package handlers

import (
	"time"

	"github.com/internships-backend/test-backend-ldvsked/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetSlots(c *gin.Context) {
	roomId, err := uuid.Parse(c.Param("roomId"))
	if err != nil{
		c.JSON(400, gin.H{"error": "invalid roomId"})
		return
	}
	date, err := time.ParseInLocation("2006-01-02", c.Query("date"), time.UTC)

	if err != nil {
		c.JSON(400, gin.H{"error" : "date parameter is required and must be in YYYY-MM-DD"})
		return 
	}

	var check bool = false 
	for _, room := range models.Rooms {
		if room.Id == roomId {
			check = true 
			break
		}
	}
	if !check {
		c.JSON(200, []models.Slot{})
		return
	}

	result := []models.Slot{}

	for _, slot := range models.Slots {
		if slot.Start.Format("2006-01-02") == date.Format("2006-01-02") &&
		slot.RoomId == roomId { //слот подошел
			var check bool = true
			for _, booking := range models.Bookings {
				if booking.SlotId == slot.Id  && booking.Status != "cancelled"{ //проверили что его не забронировали
					check = false
					break
				}
			}
			if check {
				result = append(result, slot)
			}
		}
	}
	c.JSON(200, result)
}