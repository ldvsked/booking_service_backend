package rooms

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/internships-backend/test-backend-ldvsked/models"
	"github.com/google/uuid"
)

type CreateRoomRequest struct {
	Name string `json:"name"`
	Capacity int `json:"capacity"`
	Description string `json:"description"`
}

func CreateRoom(c *gin.Context) {
	role := c.GetString("role")

	if role != "admin" {
		c.JSON(403, map[string]any{"error":"forbidden"})
		return
	}

	var createRoomRequest CreateRoomRequest
	//ошибка в синтаксисе, невалидные типы
	if err := c.ShouldBindJSON(&createRoomRequest); err != nil || createRoomRequest.Name == ""{
		c.JSON(400, map[string]any{"error" : "invalid request"})
		return
	}
	newRoom := models.Room{Id: uuid.New(), 
		Name: createRoomRequest.Name,
		Description: createRoomRequest.Description,
		Capacity: createRoomRequest.Capacity,
		CreatedAt: time.Now(),
	}
	models.Rooms = append(models.Rooms, newRoom)
	c.JSON(201, map[string]any{"room": newRoom})
}
