package handlers 

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

var (
	AdminUUID uuid.UUID = uuid.MustParse("00000000-0000-0000-0000-000000000001")
	UserUUID uuid.UUID = uuid.MustParse("00000000-0000-0000-0000-000000000002")
)


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