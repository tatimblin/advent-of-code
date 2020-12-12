package main

import (
	"strings"
)

func countUniqueChars(body string) int {
	chars := strings.Split(body, "")
	charMap := make(map[string]bool)

	for _, char := range chars {
		charMap[char] = true
	}

	return len(charMap)
}
