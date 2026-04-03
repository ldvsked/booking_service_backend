package main

import (
	"github.com/gin-contrib/cors" // приклеивать разрешающие заголовки, чтобы браузер не ругался, сервер разрешает использовать его данные
	"github.com/gin-gonic/gin"
	"github.com/internships-backend/test-backend-ldvsked/handlers/rooms"
	"github.com/internships-backend/test-backend-ldvsked/handlers"
	
)


func main() {
	var r *gin.Engine = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, 
		AllowCredentials: true,
	}))

	r.GET("/rooms/list", rooms.GetRooms)
	r.POST("/dummyLogin", handlers.DummyLogin)
	r.GET("/_info", handlers.Info)
	r.POST("/rooms/create", handlers.AuthMiddleware, rooms.CreateRoom)

	r.Run("localhost:8080")

}