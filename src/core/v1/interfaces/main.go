package interfaces

import "github.com/gin-gonic/gin"

func PingInterface(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
