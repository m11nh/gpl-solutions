package anagram_test

import (
	"testing"
	"github.com/miniman1/gobooksolutions/3.12/anagram"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		same bool
	}{
		{"cat", "tac", true},
		{"cat", "tic", false},
		{"", "", true},
		{"b", "b", true},
		{"batta", "batta", true},
	}
	for _, test := range tests {
		r := anagram.IsAnagram(test.a, test.b)
		if r != test.same {
			t.Errorf("expected IsAnagram(%q, %q) to be %v, instead we got %v\n )", test.a, test.b, test.same, r)
		}
	}
}
