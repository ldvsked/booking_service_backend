package main

import (
	"github.com/gin-contrib/cors" // приклеивать разрешающие заголовки, чтобы браузер не ругался, сервер разрешает использовать его данные
	"github.com/gin-gonic/gin"
	"github.com/internships-backend/test-backend-ldvsked/handlers"
)

func SayHi(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"message": "Hi, girls!"})
}


func main() {
	var r *gin.Engine = gin.Default()

	r.Use(cors.Default())

	r.GET("/say_hi", SayHi)
	r.GET("/rooms/list", handlers.GetRooms)
	r.POST("/dummyLogin", handlers.DummyLogin)
	r.GET("/_info", handlers.Info)

	r.Run("localhost:8080")

}