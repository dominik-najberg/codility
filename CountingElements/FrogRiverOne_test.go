package CountingElements

import (
	"fmt"
	"testing"
)

func TestFrogRiverOneSolution(t *testing.T) {
	tests := []struct {
		X    int
		A    []int
		want int
	}{
		{
			5,
			[]int{1, 3, 1, 4, 2, 3, 5, 4},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v -> %v", tt.X, tt.A, tt.want), func(t *testing.T) {
			if got := FrogRiverOneSolution(tt.X, tt.A); got != tt.want {
				t.Errorf("FrogRiverSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
