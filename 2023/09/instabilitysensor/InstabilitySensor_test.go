package instabilitysensor

import (
	"testing"
)

func TestCalculateDifferences(t *testing.T) {
	tests := []struct {
		Name string
		In   Log
		Size int
	}{
		{
			Name: "Small sample",
			In:   Log{History: [][]int{{1, 2, 3}}},
			Size: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.In.CalculateDifferences()
			if len(tt.In.History) != tt.Size {
				t.Errorf("got %d, wanted %d", len(tt.In.History), tt.Size)
			}
		})
	}
}

func TestAddEstimates(t *testing.T) {
	tests := []struct {
		Name string
		In   Log
		Out  int
	}{
		{
			Name: "Small sample",
			In:   Log{History: [][]int{{1, 2, 3}, {1, 1}, {0}}},
			Out:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.In.AddEstimates()
			if got != tt.Out {
				t.Errorf("got %d, wanted %d", got, tt.Out)
			}
		})
	}
}
