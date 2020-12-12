package main

import (
	"fmt"
	"testing"
)

func TestCountUniqueChars(t *testing.T) {
	t.Run("abcdef", testCountUniqueChars("abcdef", 6))
	t.Run("aabcccdef", testCountUniqueChars("aabcccdef", 6))
}

func testCountUniqueChars(body string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := countUniqueChars(body)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected the unique characters of %v to be %d but instead got %d!", body, expected, actual))
		}
	}
}
