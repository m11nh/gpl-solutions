// calculates number of bits in a number that are 1
package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func main() {
	fmt.Println(PopCount(1))
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	for ; x != 0; x >>= 1 {
		if x & 1 == 1 {
			count += 1
		}
	}
	return count
}
