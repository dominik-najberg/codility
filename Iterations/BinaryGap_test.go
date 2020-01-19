package Iterations

import (
	"fmt"
	"testing"
)

func TestBinaryGapSolution(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			9,
			2,
		},
		{
			1041,
			5,
		},
		{
			104120,
			2,
		},
		{
			32,
			0,
		},
		{
			66561,
			9,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing %v", tt.input), func(t *testing.T) {
			if got := BinaryGapSolution(tt.input); got != tt.want {
				t.Errorf("BinaryGapSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
