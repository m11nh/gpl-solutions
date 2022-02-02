package rotate_test

import (
	"github.com/miniman1/gobooksolutions/4.4"
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		x, want []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 3, 4, 5, 1},
		},
	}
	for _, tc := range tests {
		result := rotate.Rotate(tc.x)
		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("Wanted %v, instead got %v", tc.want, result)
		}
	}
}
