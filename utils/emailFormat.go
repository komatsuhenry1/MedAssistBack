package utils

import (
	"errors"
	"regexp"
	"strings"
)

var emailPattern = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

func EmailRegex(email string) (string, error) {
	normalized := strings.TrimSpace(strings.ToLower(email))

	if !emailPattern.MatchString(normalized) {
		return "", errors.New("Formato de e-mail inválido")
	}

	return normalized, nil
}
