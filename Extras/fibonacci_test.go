package Extras

import (
	"reflect"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8}

	t.Run("Basic test", func(t *testing.T) {
		if got := fibonacci(); !reflect.DeepEqual(got, want) {
			t.Errorf("fibonacci() = %v, want %v", got, want)
		}
	})
}
