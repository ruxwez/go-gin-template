package v1Router

import (
	"github.com/gin-gonic/gin"
	v1Interfaces "go-gin-template/internal/v1/interfaces"
)

func Routes(v1 *gin.RouterGroup) {
	v1.GET("/test", v1Interfaces.TestInterface)
}
