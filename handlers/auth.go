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

//jwt валидный
func AuthMiddleware(c *gin.Context) {
	//вытащить токен из заголовка 
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, map[string]interface{}{"error": "missing authorization header"})
		return
	}

	//распарсить строку в jwt и подпись проверить
	authHeader = authHeader[7:]
	token, err := jwt.Parse(authHeader, func(token *jwt.Token)(interface{}, error) {
		return []byte("i_want_to_sleep"), nil
	})
	if err != nil {
		c.JSON(401, map[string]any{"error":"faked jwt"})
		return
	}
	
	//вытащить claims и проверить что там есть
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Status(400)
		return
	}
	//проверка и на есть и на приведение к строке
	role, ok1 := claims["role"].(string)
	userId, ok2 := claims["user_id"].(string)
	if !ok1 || !ok2 || userId == "" || !(role == "admin" || role == "user") {
		c.JSON(401, "invalid token")
		return
	}
	
	c.Set("userId", userId)
	c.Set("role", role)
	c.Next()
}

