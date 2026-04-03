package handlers 

import (
	"github.com/gin-gonic/gin"
	"github.com/internships-backend/test-backend-ldvsked/models"
)

func GetRooms(c *gin.Context) {
	c.JSON(200, map[string]any{"rooms":models.Rooms})
}

