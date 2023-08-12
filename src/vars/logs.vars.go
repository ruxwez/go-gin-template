package vars

var (
	DiscordLogsIcon = ""
)

type typeLogs struct {
	Error   string
	Warning string
	Info    string
}

var Logs = typeLogs{
	Error:   "error",
	Warning: "warning",
	Info:    "info",
}
