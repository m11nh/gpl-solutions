// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//!+array
	a := []byte("f的jk的   salfj")
	reverse(a)
	fmt.Println(string(a))
}

func rev(s []byte) {
	N := len(s)
	for i := 0; i < i/2; i++ {
		end := N - 1 - i 
		s[i], s[end] = s[end], s[i]
	}
}

// reverse reverses a slice of bytes in place.
func reverse(s []byte) {
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		rev(s[i:i + size])
		i += size
	}
	rev(s[:])
}
