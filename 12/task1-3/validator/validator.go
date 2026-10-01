package validator

import (
	"strings"
)

func IsValidEmail(email string) bool {
	atIndex := strings.Index(email, "@")

	if atIndex <= 0 || atIndex == len(email)-1 || strings.Count(email, "@") != 1 {
		return false
	}

	domain := email[atIndex+1:]

	lastDotIndex := strings.LastIndex(domain, ".")

	if lastDotIndex <= 0 || lastDotIndex == len(domain)-1 {
		return false
	}

	return true
}
