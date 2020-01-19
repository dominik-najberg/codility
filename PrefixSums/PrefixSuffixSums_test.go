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
		{
			[]int{2, 2, 1, 4, 1},
			[]int{0, 2, 4, 5, 9, 10},
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
		{
			[]int{6, 2, 3, 6, 1},
			[]int{18, 12, 10, 7, 1},
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

func Test_countTotalOfSliceOfSums(t *testing.T) {
	tests := []struct {
		A    []int
		x    int
		y    int
		want int
	}{
		{
			[]int{0, 1, 3, 6, 10, 15}, // prefixSums for the following slice: []int{1, 2, 3, 4, 5}
			0,
			2,
			6,
		},
		{
			[]int{21, 18, 13, 12, 6}, // prefixSums for the following slice: []int{3, 5, 1, 6, 6}
			0,
			2,
			9,
		},
		{
			[]int{0, 2, 4, 5, 9, 10}, // prefixSums for the following slice: []int{2, 2, 1, 4, 1}
			0,
			2,
			5,
		},
		{
			[]int{18, 12, 10, 7, 1}, // prefixSums for the following slice: []int{6, 2, 3, 6, 1}
			0,
			2,
			11,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, x: %v, y: %v -> %v", tt.A, tt.x, tt.y, tt.want), func(t *testing.T) {
			if got := countTotalOfSliceOfSums(tt.A, tt.x, tt.y); got != tt.want {
				t.Errorf("countTotalOfSliceOfSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	tests := []struct {
		vs   []int
		want int
	}{
		{
			[]int{10, 1, 29},
			1,
		},
		{
			[]int{100, -1, 901, -200},
			-200,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.vs, tt.want), func(t *testing.T) {
			if got := min(tt.vs...); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	tests := []struct {
		vs   []int
		want int
	}{
		{
			[]int{10, 1, 29},
			29,
		},
		{
			[]int{100, -1, 901, -200},
			901,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v", tt.vs, tt.want), func(t *testing.T) {
			if got := max(tt.vs...); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mushrooms(t *testing.T) {
	tests := []struct {
		A        []int
		startPos int
		steps    int
		want     int
	}{
		{
			A:        []int{2, 3, 7, 5, 1, 3, 9},
			startPos: 4,
			steps:    6,
			want:     25,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, start: %v, steps: %v -> %v", tt.A, tt.startPos, tt.steps, tt.want), func(t *testing.T) {
			if got := mushrooms(tt.A, tt.startPos, tt.steps); got != tt.want {
				t.Errorf("mushrooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
