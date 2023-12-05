package main

import (
	"testing"
)

func TestParseScratcher(t *testing.T) {

	tests := []struct {
		Name   string
		Input  string
		Output Scratcher
	}{
		{
			Name:  "calculates points for a complete win",
			Input: "Card 1: 1 2 3 | 1 2 3",
			Output: Scratcher{
				ID:       1,
				Winning:  []int{1, 2, 3},
				Matching: []int{1, 2, 3},
				Points:   4,
			},
		},
		{
			Name:  "calculates points for a complete loss",
			Input: "Card 1: 1 2 3 | 4 5 6",
			Output: Scratcher{
				ID:       1,
				Winning:  []int{1, 2, 3},
				Matching: []int{1, 2, 3},
				Points:   0,
			},
		},
		{
			Name:  "calculates points for a half win",
			Input: "Card 1: 1 2 3 | 1 2 3 4 5 6",
			Output: Scratcher{
				ID:       1,
				Winning:  []int{1, 2, 3},
				Matching: []int{1, 2, 3},
				Points:   4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			scratcher := ParseScratcher(tt.Input)
			if scratcher.Points != tt.Output.Points {
				t.Errorf("got %d, wanted %d", scratcher.Points, tt.Output.Points)
			}
		})
	}
}
