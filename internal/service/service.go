package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(data string) bool {
	morseChars := "-. "
	for _, ch := range data {
		if !strings.ContainsRune(morseChars, ch) {
			return false
		}
	}
	return true
}

func ConvertText(data string) string {
	if isMorse(data) {
		return morse.ToText(data)
	}
	return morse.ToMorse(data)
}
