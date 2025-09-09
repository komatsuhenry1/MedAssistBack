package utils

import (
	"strings"
	"unicode"
)

func CapitalizeFirstWord(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}
	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
