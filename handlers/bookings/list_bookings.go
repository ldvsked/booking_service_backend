package bookings

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/internships-backend/test-backend-ldvsked/models"
)

func GetBookingsList(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(403, gin.H{"error" : "call for admin, baby"})
		return
	}
	page, err1 := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, err2 := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{"error" : "page and pageSize must be integer"})
		return
	}
	if page < 1 || pageSize < 1 || pageSize > 100 {
		c.JSON(400, gin.H{"error" : "page and pageSize min - 1, pageSaze max - 100"})
		return
	}
	start := pageSize * (page - 1)
	if start > len(models.Bookings) {
		c.JSON(404, gin.H{"error" : "page is too big"})
		return
	}
	end := start + pageSize
	if end > len(models.Bookings) {
		end = len(models.Bookings)
	}
	c.JSON(200, gin.H{"bookings" : models.Bookings[start:end]})
}