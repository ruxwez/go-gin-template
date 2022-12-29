package internal

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/internal/database"
	"go-gin-template/internal/mail"
	"go-gin-template/internal/utils/logs"
	vars2 "go-gin-template/internal/utils/vars"
)

func Run() {
	mail.Connection(vars2.MailConnection)

	// Connection database
	database.Connection()

	// Config API
	if !vars2.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	routes(r)

	// Run API
	err := r.Run("0.0.0.0" + vars2.PortAPI)
	if err != nil {
		logs.Send(vars2.TypeLogs.Error, err.Error())
		return
	}
}
