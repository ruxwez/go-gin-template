package src

import (
	v1Router "go-gin-template/src/core/v1/router"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	// Create a v1 group for all v1 routes
	v1 := r.Group("/v1")
	{
		v1Router.Routes(v1)
	}
}
