package bookings

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"

	"github.com/internships-backend/test-backend-ldvsked/models"
)

func GetMyBookings(c *gin.Context) {
	role := c.GetString("role")

	if role == "admin" {
		c.JSON(403, gin.H{"error" : "only user can book"})
		return
	}
	userId, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error" : "invlaid userId"})
		return
	}
	var result = []models.Booking{}
	for _, booking := range models.Bookings {
		if booking.UserId == userId {
			var slotStart time.Time
			for _, slot := range models.Slots {
				if slot.Id == booking.SlotId {
					slotStart = slot.Start
					break
				}
			}
			if !slotStart.Before(time.Now()) && booking.Status == "active"{
				result = append(result, booking)
			}
		}
			
	}
	c.JSON(200, gin.H{"bookings" : result})
}