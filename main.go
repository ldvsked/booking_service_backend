package main

import (
	"log"

	"github.com/gin-contrib/cors" // приклеивать разрешающие заголовки, чтобы браузер не ругался, сервер разрешает использовать его данные
	"github.com/gin-gonic/gin"
	"github.com/internships-backend/test-backend-ldvsked/handlers"
	//"github.com/internships-backend/test-backend-ldvsked/handlers/bookings"
	"github.com/internships-backend/test-backend-ldvsked/handlers/rooms"

	"github.com/jmoiron/sqlx"
	_"github.com/lib/pq" //использую не функции, а побочные эффекты, Go сам сходит к этому драйверу когда надо будет sqlx
)


func main() {

	var dsn string = "host=db port=5432 user=postgres password=password dbname=booking_db sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Критическая ошибка: не удалось подключиться в БД: %s", err)
	}

	log.Println("Связь с базой установлена!")
	defer db.Close()

	var r *gin.Engine = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, 
		AllowCredentials: true,
	}))

	r.GET("/_info", handlers.Info)
	r.POST("/dummyLogin", handlers.DummyLogin)

	var handler rooms.Handler = rooms.Handler{DB: db}

	r.GET("/rooms/list", handlers.AuthMiddleware, handler.GetRooms)
	// r.POST("/rooms/create", handlers.AuthMiddleware, rooms.CreateRoom)
	// r.GET("/rooms/:roomId/slots/list", handlers.AuthMiddleware, handlers.GetSlots)

	// r.POST("/rooms/:roomId/schedule/create", handlers.AuthMiddleware, handlers.CreateShedule)

	// r.POST("/bookings/create", handlers.AuthMiddleware, bookings.CreateBooking)
	// r.GET("/bookings/list", handlers.AuthMiddleware, bookings.GetBookingsList)
	// r.GET("/bookings/my", handlers.AuthMiddleware, bookings.GetMyBookings)
	// r.POST("/bookings/:bookingId/cancel", handlers.AuthMiddleware, bookings.CancelBooking)

	r.Run(":8080")

}