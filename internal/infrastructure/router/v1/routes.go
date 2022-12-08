package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/internal/interfaces"
)

func Routes(v1 *gin.RouterGroup) {
	v1.GET("/test", interfaces.TestInterface)
}
