package logs

func Send(_type string, message string) {
	addLogFile(_type, message)
}
