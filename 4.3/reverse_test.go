package reverse_test

import (
	"testing"

	"github.com/miniman1/gobooksolutions/4.3"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		i, want [5]int
	}{
		{
			[5]int{1, 2, 3, 4, 5}, 
			[5]int{5, 4, 3, 2, 1}, 
		}, 
		{
			[5]int{}, 
			[5]int{}, 
		}, 
	}
	for _, tc := range tests {
		if *reverse.ReverseInt(&tc.i) !=  tc.want {
			t.Errorf("ReverseInt yielded %v, instead of %v\n", tc.i, tc.want)
		}
	}
}
