package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-gin-template/internal/infrastructure/router/v1"
)

func Routes(r *gin.Engine) {
	versionOne := r.Group("/v1")
	{
		v1.Routes(versionOne)
	}
}
