package src

import (
	"go-gin-template/src/database"
	"go-gin-template/src/utils/logs"
	"go-gin-template/src/vars"

	"github.com/gin-gonic/gin"
)

func Start() {

	// Database connection
	database.Init()

	// Set Debug
	if !vars.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())

	routes(r)

	// Run API
	err := r.Run("0.0.0.0" + vars.PortAPI)
	if err != nil {
		logs.Send(vars.Logs.Error, err.Error())
		return
	}

}
