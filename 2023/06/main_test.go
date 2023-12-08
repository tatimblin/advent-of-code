package main

import (
	"testing"
)

func TestCalcDistance(t *testing.T) {

	tests := []struct {
		Time     int
		Hold     int
		Distance int
	}{
		{Time: 7, Hold: 0, Distance: 0},
		{Time: 7, Hold: 1, Distance: 6},
		{Time: 7, Hold: 2, Distance: 10},
		{Time: 7, Hold: 3, Distance: 12},
		{Time: 7, Hold: 4, Distance: 12},
		{Time: 7, Hold: 5, Distance: 10},
		{Time: 7, Hold: 6, Distance: 6},
		{Time: 7, Hold: 7, Distance: 0},
	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			distance := CalcDistance(tt.Time, tt.Hold)
			if distance != tt.Distance {
				t.Errorf("got %d, wanted %d", distance, tt.Distance)
			}
		})
	}
}
