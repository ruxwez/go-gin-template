package enums

type typeLogs struct {
	Error   string
	Warning string
	HTTP    string
}

var TypeLogs = typeLogs{
	Error:   "errors",
	Warning: "warnings",
	HTTP:    "http",
}
