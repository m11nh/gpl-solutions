// Tests for functions in shadiff.go
package shadiff_test

import (
	"testing"
	"github.com/miniman1/gobooksolutions/4.1"
)

// Testing the Diffsha256 function from shadiff module
func testshadiff(t *testing.T) {
	tests := []struct {
		a, b [32]byte
		want int
	}{
		{[32]byte{0}, [32]byte{6}, 2},
		{[32]byte{1, 2, 3}, [32]byte{4, 5, 6}, 7},	
	}
	for _, tc := range tests {
		r := shadiff.Diffsha256(tc.a, tc.b)
		if  r != tc.want {
			t.Errorf("diffsha256(%v, %v), got %d, want %d\n", tc.a, tc.b, r, tc.want)
		}
	}
}

