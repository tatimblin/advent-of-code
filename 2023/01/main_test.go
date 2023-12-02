package main

import (
	"testing"
)

func TestParseString(t *testing.T) {
	tests := []struct {
		Name string
		Input string
		Output int
	}{
		{
			Name: "can parse a typical input",
			Input: "jhsd3jhsd44hj8d",
			Output: 38,
		},
		{
			Name: "can parse a single digit input",
			Input: "5",
			Output: 55,
		},
		{
			Name: "can parse a single alpha input",
			Input: "a",
			Output: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := ParseString(tt.Input); got != tt.Output {
				t.Errorf("got %d, wanted %d", got, tt.Output)
			}
		})
	}
} 

func TestIsByteNumber(t *testing.T) {

	got := IsByteNumber(byte('7'))
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = IsByteNumber(byte('a'))
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
