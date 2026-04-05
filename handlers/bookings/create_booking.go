package bookings

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/internships-backend/test-backend-ldvsked/models"
)

type BookingRequest struct {
	SlotId uuid.UUID `json:"slotId"`
	CreateConferenceLink bool `json:"createConferenceLink"`
}

func CreateBooking(c *gin.Context) {
	role := c.GetString("role")
	if role != "user" {
		c.JSON(403, map[string]any{"error" : "you can't make bookings"})
		return
	}
	userId, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error" : "invalid user_id"})
		return
	} 

	var bookingRequest BookingRequest
	if err := c.ShouldBindJSON(&bookingRequest); err != nil {
		c.JSON(400, gin.H{"error" : "invalid slotId or need of conference link"})
		return
	}

	var check bool = false 
	for _, slot := range models.Slots {
		if slot.Id == bookingRequest.SlotId {
			check = true 
			if slot.Start.Before(time.Now()) {
				c.JSON(400, gin.H{"error" : "slot is in the past"})
				return
			}
			break
		}
	}
	if !check {
		c.JSON(404, gin.H{"error" : "we don't have this slot"})
		return
	}
	for _, booking := range models.Bookings {
		if booking.SlotId == bookingRequest.SlotId && booking.Status == "active" {
			c.JSON(400, gin.H{"error" : "this slot is owned by someone else"})
			return
		}
	}

	var conferenceLink string = ""
	if bookingRequest.CreateConferenceLink {
		conferenceLink = "https://meet.example.com/" + uuid.New().String()
	}

	var newBooking models.Booking = models.Booking{Id: uuid.New(), 
		SlotId: bookingRequest.SlotId, 
		UserId: userId,
		Status: "active", 
		ConferenceLink: conferenceLink,
		CreatedAt: time.Now(),
	}
	models.Bookings = append(models.Bookings, newBooking)
	c.JSON(201, gin.H{"booking" : newBooking})
}