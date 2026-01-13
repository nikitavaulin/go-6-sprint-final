package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// isMorse проверяет написан ли текст на азбуке Морзе или нет
func isMorse(data string) bool {
	morseChars := "-. "
	for _, ch := range data {
		if !strings.ContainsRune(morseChars, ch) {
			return false
		}
	}
	return true
}

// ConvertText конвертирует текст исходя из алфавита его символов
func ConvertText(data string) string {
	if isMorse(data) {
		return morse.ToText(data)
	}
	return morse.ToMorse(data)
}
