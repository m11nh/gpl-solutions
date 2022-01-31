// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	b := bytes.Buffer{}
	N := len(s)
	start := 0

	// Handle sign 
	if N == 0 { return s }
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		start = 1
	}

	// Find position of '.' if any 
	end := strings.Index(s, ".")
	if end == -1 { end = N }

	// Replace commas up to position end
	inc := 1
	for i := start; i < end; i++ {
		if inc % 3 == 0 {
			b.WriteByte(',')
		}
		b.WriteByte(s[i])
		inc++
	}

	// insert any remaining values
	b.WriteString(s[end:N])
	return b.String()
}

