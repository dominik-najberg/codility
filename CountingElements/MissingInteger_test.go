package CountingElements

import (
	"fmt"
	"testing"
)

func TestMissingIntegerSolution(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{
			[]int{1, 3, 6, 4, 1, 2},
			5,
		},
		{
			[]int{1, 2, 3},
			4,
		},
		{
			[]int{-1, -3},
			1,
		},
		{
			[]int{1},
			2,
		},
		{
			[]int{100},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.A, tt.want), func(t *testing.T) {
			if got := MissingIntegerSolution(tt.A); got != tt.want {
				t.Errorf("MissingIntegerSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
