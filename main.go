package main 

import (
	"github.com/gin-gonic/gin"
)

func SayHi(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"message": "Hi, girls!"})
}

func main() {
	var r *gin.Engine = gin.Default()

	r.GET("/say_hi", SayHi)

	r.Run("127.0.0.1:8080")

}