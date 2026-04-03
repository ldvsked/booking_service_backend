package main

import (
	"github.com/gin-contrib/cors" // приклеивать разрешающие заголовки, чтобы браузер не ругался, сервер разрешает использовать его данные
	"github.com/gin-gonic/gin"
	"github.com/internships-backend/test-backend-ldvsked/rooms"
	"github.com/internships-backend/test-backend-ldvsked/handlers"
	
)


func main() {
	var r *gin.Engine = gin.Default()

	r.Use(cors.Default())

	r.GET("/rooms/list", handlers.GetRooms)
	r.POST("/dummyLogin", handlers.DummyLogin)
	r.GET("/_info", handlers.Info)
	r.POST("/rooms/create", handlers.AuthMiddleware, handlers.CreateRoom)

	r.Run("localhost:8080")

}