package main

import (
	"github.com/gin-contrib/cors" // приклеивать разрешающие заголовки, чтобы браузер не ругался, сервер разрешает использовать его данные
	"github.com/gin-gonic/gin"
	"github.com/internships-backend/test-backend-ldvsked/handlers"
	"github.com/internships-backend/test-backend-ldvsked/handlers/bookings"
	"github.com/internships-backend/test-backend-ldvsked/handlers/rooms"
)


func main() {
	var r *gin.Engine = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, 
		AllowCredentials: true,
	}))

	r.GET("/_info", handlers.Info)
	r.POST("/dummyLogin", handlers.DummyLogin)

	r.GET("/rooms/list", rooms.GetRooms)
	r.POST("/rooms/create", handlers.AuthMiddleware, rooms.CreateRoom)
	r.GET("/rooms/:roomId/slots/list", handlers.AuthMiddleware, handlers.GetSlots)

	r.POST("/rooms/:roomId/schedule/create", handlers.AuthMiddleware, handlers.CreateShedule)

	r.POST("/bookings/create", handlers.AuthMiddleware, bookings.CreateBooking)
	r.GET("/bookings/list", handlers.AuthMiddleware, bookings.GetBookingsList)
	r.GET("/bookings/my", handlers.AuthMiddleware, bookings.GetMyBookings)
	r.POST("/bookings/:bookingId/cancel", handlers.AuthMiddleware, bookings.CancelBooking)

	r.Run(":8080")

}