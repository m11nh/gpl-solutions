package squash_test

import (
	"reflect"
	"testing"
	"github.com/miniman1/gobooksolutions/4.6"
)

func TestSquash(t *testing.T) {
	tests := []struct{
		input, want []byte
	}{
		{
			input: []byte("f的jk的   salfj"),
			want: []byte("f的jk的 salfj"),
		},
		{
			input: []byte("f的    jk的   salfj"),
			want: []byte("f的 jk的 salfj"),
		},
		{
			input: []byte("f的 jk的   salfj"),
			want: []byte("f的 jk的 salfj"),
		},
		{
			input: []byte(""),
			want: []byte(""),
		},
		{
			input: []byte("的"),
			want: []byte("的"),
		},
		{
			input: []byte(" "),
			want: []byte(" "),
		},
		{
			input: []byte("    "),
			want: []byte(" "),
		},
		{
			input: []byte("f的 jk的   salfj     "),
			want: []byte("f的 jk的 salfj "),
		},
	}
	for _, tc := range tests {
		inputCopy := make([]byte, len(tc.input))
		copy(inputCopy, tc.input)
		result := squash.Squash(tc.input)
		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("Squash(%q) resulted in %q, but wanted %q\n", inputCopy, result, tc.want)
		}
	}
}
