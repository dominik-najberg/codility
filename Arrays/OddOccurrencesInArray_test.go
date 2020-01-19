package Arrays

import (
	"fmt"
	"testing"
)

func TestOddOccurrencesInArraySolution(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{
			A:    []int{9, 3, 9, 3, 9, 7, 9},
			want: 7,
		},
		{
			A:    []int{9, 3, 9, 3, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.A, tt.want), func(t *testing.T) {
			if got := OddOccurrencesInArraySolution(tt.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
