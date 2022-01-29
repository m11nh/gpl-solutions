package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	for i := 0; i < 10000000; i++ {
		PopCount(1099511627776)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
//!-

/*
performance analysis
- better than 2.4 if we use popcount a lot of times, because for each popcount, we only do 8 operations, not go through all 64 bits. 

*/
