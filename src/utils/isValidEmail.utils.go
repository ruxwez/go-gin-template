package utils

import "regexp"

// IsValidEmail checks if a given string is a valid email
func IsValidEmail(email string) bool {
	// The regular expression that we will use to verify the syntax of the email
	// is based on the e-mail format specification (RFC 5322)
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
