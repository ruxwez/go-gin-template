package main

import (
	env "github.com/joho/godotenv"
	"go-gin-template/internal"
	"go-gin-template/internal/utils/logs"
	"go-gin-template/internal/utils/vars"
)

func main() {
	err := env.Load()
	if err != nil {
		logs.Send(vars.TypeLogs.Error, "Could not load .env file")
	}

	internal.Run()
}
