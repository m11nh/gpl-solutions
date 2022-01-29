// calculates number of bits in a number that are 1
package main

func main() {
	for i := 0; i < 10000000; i++ {
		PopCount(1099511627776)
	}
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


/*
	Better than 2.3 if we only use PopCount once, as we only go through 64 operations, instead of the init 256 operations needed in 2.3
*/
