package CountingElements

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxCountersSolution(t *testing.T) {
	tests := []struct {
		N    int
		A    []int
		want []int
	}{
		{
			5,
			[]int{3, 4, 4, 6, 1, 4, 4},
			[]int{3, 2, 2, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v -> %v", tt.N, tt.A, tt.want), func(t *testing.T) {
			if got := MaxCountersSolution(tt.N, tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxCountersSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
