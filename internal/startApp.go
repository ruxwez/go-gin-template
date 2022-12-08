package internal

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/internal/infrastructure/database"
	"go-gin-template/internal/infrastructure/mail"
	"go-gin-template/internal/infrastructure/router"
	"go-gin-template/internal/infrastructure/utils/logs"
	"go-gin-template/internal/infrastructure/utils/vars"
)

func Run() {
	mail.Connection(vars.MailConnection)

	// Connection database
	database.Connection()

	// Config API
	if !vars.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	router.Routes(r)

	// Run API
	err := r.Run("0.0.0.0" + vars.PortAPI)
	if err != nil {
		logs.Send(vars.TypeLogs.Error, err.Error())
		return
	}
}
