package module

import (
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		Name     string
		From     Galaxy
		To       Galaxy
		Distance int
	}{
		{
			Name:     "Diagonal",
			From:     Galaxy{Y: 0, X: 4},
			To:       Galaxy{Y: 10, X: 9},
			Distance: 15,
		},
		{
			Name:     "Diagonal 2",
			From:     Galaxy{Y: 2, X: 0},
			To:       Galaxy{Y: 7, X: 12},
			Distance: 17,
		},
		{
			Name:     "Horizontal",
			From:     Galaxy{Y: 11, X: 0},
			To:       Galaxy{Y: 11, X: 5},
			Distance: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := CalculateDistance(tt.From, tt.To)
			if got != tt.Distance {
				t.Errorf("got %d, wanted %d", got, tt.Distance)
			}
		})
	}
}

func TestCountRange(t *testing.T) {
	tests := []struct {
		Name  string
		Arr   []int
		From  int
		To    int
		Count int
	}{
		{
			Name:  "Inner slice",
			Arr:   []int{22, 54, 56, 64, 85, 92, 128},
			From:  36,
			To:    114,
			Count: 5,
		},
		{
			Name:  "Left slice",
			Arr:   []int{22, 54, 56, 64, 85, 92, 128},
			From:  2,
			To:    60,
			Count: 3,
		},
		{
			Name:  "Right slice",
			Arr:   []int{22, 54, 56, 64, 85, 92, 128},
			From:  65,
			To:    130,
			Count: 3,
		},
		{
			Name:  "Complete slice",
			Arr:   []int{22, 54, 56, 64, 85, 92, 128},
			From:  10,
			To:    200,
			Count: 7,
		},
		{
			Name:  "Empty slice",
			Arr:   []int{1, 2, 4, 5},
			From:  3,
			To:    3,
			Count: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := CountRange(tt.Arr, tt.From, tt.To)
			if got != tt.Count {
				t.Errorf("got %d, wanted %d", got, tt.Count)
			}
		})
	}
}
