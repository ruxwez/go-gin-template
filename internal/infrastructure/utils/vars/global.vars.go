package vars

import "os"

var (
	AppName               string = "Go Fiber Template V1"
	AppIconNotTransparent string = ""
	PortAPI               string = ":8080"
	Debug                 bool   = false

	// MAIL
	EmailHost     string = os.Getenv("EMAIL_HOST")
	EmailPort     int    = 22
	EmailMail     string = os.Getenv("EMAIL_MAIL")
	EmailPassword string = os.Getenv("EMAIL_PASSWORD")

	// LOGS
	DiscordHook string = os.Getenv("DISCORD_HOOK_LOG")
)
