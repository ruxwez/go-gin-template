package router

import (
	"go-gin-template/src/core/v1/interfaces"

	"github.com/gin-gonic/gin"
)

func Routes(v1 *gin.RouterGroup) {
	v1.GET("/ping", interfaces.PingInterface)
}
