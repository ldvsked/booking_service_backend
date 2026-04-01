package main

import (
	"time"

	"github.com/gin-contrib/cors" // приклеивать разрешающие заголовки, чтобы браузер не ругался, сервер разрешает использовать его данные
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	AdminUUID uuid.UUID = uuid.MustParse("00000000-0000-0000-0000-000000000001")
	UserUUID uuid.UUID = uuid.MustParse("00000000-0000-0000-0000-000000000002")
)



func SayHi(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"message": "Hi, girls!"})
}


func main() {
	var r *gin.Engine = gin.Default()

	r.Use(cors.Default())

	r.GET("/say_hi", SayHi)
	r.GET("/rooms/list", GetRooms)
	r.POST("/dummyLogin", DummyLogin)

	r.Run("localhost:8080")

}

func GetRooms(c *gin.Context) {
	c.JSON(200, map[string]any{"rooms":Rooms})
}

type Role struct {
	Role string `json:"role"`
}

func DummyLogin(c *gin.Context) {
	var role Role 
	if err := c.ShouldBindJSON(&role); err != nil || role.Role != "admin" && role.Role != "user" {
		c.JSON(400, map[string]interface{}{"error" : "Невалидный запрос"})
		return
	}

	var userId = UserUUID

	if role.Role == "admin" {
		userId = AdminUUID
	}

	claims := jwt.MapClaims{
		"user_id" : userId.String(),
		"role" : role.Role,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte("i_want_to_sleep")
	tokenString, err := token.SignedString(secret)
	if err != nil{
		c.JSON(500, map[string]any{"error": "На сервере проблема"})
		return 
	}
	c.JSON(200, map[string]any {"token" : tokenString})

}
