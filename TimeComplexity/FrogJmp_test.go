package TimeComplexity

import (
	"fmt"
	"testing"
)

func TestFrogJmpSolution(t *testing.T) {

	tests := []struct {
		X    int
		Y    int
		D    int
		want int
	}{
		{
			10,
			85,
			30,
			3,
		},
		{
			85,
			85,
			30,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v, %v -> %v", tt.X, tt.Y, tt.D, tt.want), func(t *testing.T) {
			if got := FrogJmpSolution(tt.X, tt.Y, tt.D); got != tt.want {
				t.Errorf("FrogJmpSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
