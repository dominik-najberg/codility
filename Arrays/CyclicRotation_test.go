package Arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCyclicRotationSolution(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want []int
	}{
		{
			A:    []int{3, 8, 9, 7, 6},
			K:    3,
			want: []int{9, 7, 6, 3, 8},
		},
		{
			A:    []int{3, 8, 9, 7, 6},
			K:    1,
			want: []int{6, 3, 8, 9, 7},
		},
		{
			A:    []int{0, 0, 0},
			K:    3,
			want: []int{0, 0, 0},
		},
		{
			A:    []int{0},
			K:    3,
			want: []int{0},
		},
		{
			A:    []int{},
			K:    3,
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.A, tt.K), func(t *testing.T) {
			if got := CyclicRotationSolution(tt.A, tt.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
