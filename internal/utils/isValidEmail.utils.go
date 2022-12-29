package utils

import "regexp"

// IsValidEmail verifica si una cadena dada es un correo electrónico válido
func IsValidEmail(email string) bool {
	// La expresión regular que utilizaremos para verificar la sintaxis del correo electrónico
	// está basada en la especificación del formato de correo electrónico (RFC 5322)
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
