package handlers 

import (
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	c.Status(200)
}

