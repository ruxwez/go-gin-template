package vars

type typeLogs struct {
	Error   string
	Warning string
}

var TypeLogs = typeLogs{
	Error:   "errors",
	Warning: "warnings",
}
