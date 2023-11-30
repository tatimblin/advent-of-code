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

func countMatchingChars(answers []string) int {
	charMap := make(map[string]int)
	matches := 0

	for _, submission := range answers {
		chars := strings.Split(submission, "")
		for _, char := range chars {
			charMap[char]++
		}
	}

	for _, count := range charMap {
		if count == len(answers) {
			matches++
		}
	}

	return matches
}
