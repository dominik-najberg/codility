package TimeComplexity

import (
	"fmt"
	"testing"
)

func TestPermMissingElemSolution(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{
			[]int{2, 3, 1, 5},
			4,
		},
		{
			[]int{0, 2, 3, 1, 5},
			4,
		},
		{
			[]int{0},
			1,
		},
		{
			[]int{},
			1,
		},
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{2, 3},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.A, tt.want), func(t *testing.T) {
			if got := PermMissingElemSolution(tt.A); got != tt.want {
				t.Errorf("PermMissingElemSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
