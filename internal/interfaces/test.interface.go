package interfaces

import (
	"github.com/gin-gonic/gin"
)

func TestInterface(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Test",
	})
}
