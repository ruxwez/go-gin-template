package logs

func Send(_type string, message string) {
	// Agregar condicinal si quiere discord
	SendDiscord(_type, message)
	AddLogFile(message, _type)
}
