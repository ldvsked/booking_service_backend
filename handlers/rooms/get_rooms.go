package rooms

import (
	"github.com/gin-gonic/gin"
	"github.com/internships-backend/test-backend-ldvsked/models"

	"github.com/jmoiron/sqlx"
)

type Handler struct{
	DB *sqlx.DB //чтобы в мейне положить
}

func (h *Handler) GetRooms(c *gin.Context) {
	var roomList []models.Room = []models.Room{}

	err := h.DB.Select(&roomList, "Select * from room")
	if err != nil {
		c.JSON(500, gin.H{"error" : "Can't get rooms"})
		return
	}

	c.JSON(200, gin.H{"rooms" : roomList})
}

