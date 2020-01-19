package CountingElements

import (
	"fmt"
	"testing"
)

func TestPermutationDetectorSolution(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{
			[]int{4, 1, 3, 2},
			1,
		},
		{
			[]int{4, 1, 3},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{2},
			0,
		},
		{
			[]int{1000000000},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.A, tt.want), func(t *testing.T) {
			if got := PermutationDetectorSolution(tt.A); got != tt.want {
				t.Errorf("PermutationDetectorSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
