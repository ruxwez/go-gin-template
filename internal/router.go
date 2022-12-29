package internal

import (
	"github.com/gin-gonic/gin"
	v1Router "go-gin-template/internal/v1/infrastructure/router"
)

func routes(r *gin.Engine) {
	versionOne := r.Group("/v1")
	{
		v1Router.Routes(versionOne)
	}
}
