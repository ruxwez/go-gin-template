package main

import (
	"go-gin-template/src"
	"go-gin-template/src/utils/logs"
	"go-gin-template/src/vars"

	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load()
	if err != nil {
		logs.Send(vars.Logs.Error, "Could not load .env file")
	}

	src.Start()
}
