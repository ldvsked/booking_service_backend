package bookings

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/internships-backend/test-backend-ldvsked/models"
)

func CancelBooking(c *gin.Context) {
	role := c.GetString("role")
	if role != "user" {
		c.JSON(403, gin.H{"error" : "only for users"})
		return
	}
	userId, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error" : "invalid userId"})
		return
	}

	bookingId, err := uuid.Parse(c.Param("bookingId"))
	if err != nil {
		c.JSON(400, gin.H{"error" : "invalid bookingId"})
		return
	}

	var index int = -1
	for indexBooking, booking := range models.Bookings {
		if booking.Id == bookingId {
			if booking.UserId != userId {
				c.JSON(403, gin.H{"error" : "you can cancel only your booking"})
				return
			}
			index = indexBooking
			break
		}
	}
	if index == -1 {
		c.JSON(404, gin.H{"error" : "not found this booking"})
		return
	}
	models.Bookings[index].Status = "cancelled"
	c.JSON(200, gin.H{"booking" : models.Bookings[index]})
}