package TimeComplexity

import (
	"fmt"
	"testing"
)

func TestTapeEquilibriumSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		A    []int
		want int
	}{
		{
			[]int{3, 1, 2, 4, 3},
			1,
		},
		{
			[]int{3, 1},
			2,
		},
		{
			[]int{3},
			0,
		},
		{
			[]int{},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.A, tt.want), func(t *testing.T) {
			if got := TapeEquilibriumSolution(tt.A); got != tt.want {
				t.Errorf("TapeEquilibriumSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
