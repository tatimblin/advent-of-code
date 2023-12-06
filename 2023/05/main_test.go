package main

import (
	"testing"
)

func TestCalulateSoil(t *testing.T) {
	tests := []struct {
		Name    string
		Seed    int
		Almanac []Conversion
		Output  int
	}{
		{
			Name: "example from question",
			Seed: 79,
			Almanac: []Conversion{
				{Ranges: [][]int{{98, 99}, {50, 98}}, Offset: []int{-48, 2}},
				{Ranges: [][]int{{15, 51}, {52, 3}, {0, 14}}, Offset: []int{-15, -15, 39}},
				{Ranges: [][]int{{53, 60}, {11, 52}, {0, 6}, {7, 10}}, Offset: []int{-4, -11, 42, 50}},
				{Ranges: [][]int{{18, 24}, {25, 94}}, Offset: []int{70, -7}},
				{Ranges: [][]int{{77, 99}, {45, 63}, {64, 76}}, Offset: []int{-32, 36, 4}},
				{Ranges: [][]int{{69, 69}, {0, 68}}, Offset: []int{-69, 1}},
				{Ranges: [][]int{{56, 92}, {93, 95}}, Offset: []int{4, -37}},
			},
			Output: 82,
		},
		{
			Name: "calculates points for a complete win",
			Seed: 3,
			Almanac: []Conversion{
				{Ranges: [][]int{{2, 4}}, Offset: []int{4}},
				{Ranges: [][]int{{7, 9}}, Offset: []int{-2}},
			},
			Output: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := CalculateSoil(tt.Seed, tt.Almanac)
			if got != tt.Output {
				t.Errorf("got %d, wanted %d", got, tt.Output)
			}
		})
	}
}
