package main

import (
	"strings"
	"unicode"
)

func Trim(text string) string {
	return strings.TrimSpace(text)
}

func CapitalizeUppercase(text string) string {
	if strings.ToUpper(text) == text {
		words := strings.Split(strings.ToLower(text), " ")

		for i, word := range words {
			r := []rune(word)
			if len(r) > 0 {
				r[0] = unicode.ToUpper(r[0])
				words[i] = string(r)
			}
		}

		return strings.Join(words, " ")
	}
	return text
}
