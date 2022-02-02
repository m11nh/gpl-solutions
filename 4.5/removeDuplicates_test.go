package removeDuplicates_test

import (
	"reflect"
	"testing"

	"github.com/miniman1/gobooksolutions/4.5"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		x, want []string
	}{
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"a"},
			[]string{"a"},
		},
		{
			[]string{"a", "b"},
			[]string{"a", "b"},
		},
		{
			[]string{"a", "a"},
			[]string{"a"},
		},
		{
			[]string{"a", "b", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"a", "b", "c", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"a", "b", "b", "b", "c", "c"},
			[]string{"a", "b", "c"},
		},
	}
	for _, tc := range tests {
		copyx := make([]string, len(tc.x))
		copy(copyx, tc.x)
		result := removeDuplicates.RemoveDuplicates(tc.x)
		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("Attempted removeDuplicates(%v) - Got %v, expected %v\n", copyx, result, tc.want)
		}
	}
}
