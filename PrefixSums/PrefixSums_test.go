package PrefixSums

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_calculatePrefixSums(t *testing.T) {
	tests := []struct {
		A    []int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{0, 1, 3, 6, 10, 15},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.A, tt.want), func(t *testing.T) {
			if got := calculatePrefixSums(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculatePrefixSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSuffixSums(t *testing.T) {
	tests := []struct {
		A    []int
		want []int
	}{
		{
			[]int{3, 5, 1, 6, 6},
			[]int{21, 18, 13, 12, 6},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.A, tt.want), func(t *testing.T) {
			if got := calculateSuffixSums(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateSuffixSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
