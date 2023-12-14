package module

import (
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		Name    string
		In      *Universe
		ExpandX []int
		ExpandY []int
	}{
		{
			Name: "Small sample",
			In: &Universe{
				Matrix: [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}},
			},
			ExpandX: []int{0, 1, 2},
			ExpandY: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.In.Expand()
			for i, value := range tt.In.ExpandX {
				if value != tt.ExpandX[i] {
					t.Errorf("got %d, wanted %d", value, tt.ExpandX[i])
				}
			}
			for i, value := range tt.In.ExpandY {
				if value != tt.ExpandY[i] {
					t.Errorf("got %d, wanted %d", value, tt.ExpandY[i])
				}
			}
		})
	}
}
