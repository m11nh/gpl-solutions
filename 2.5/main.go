package main

func main() {
	for i := 0; i < 10000000; i++ {
		PopCount(1099511627776)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	for ;x > 0; count++ {
		x&=x-1
	}
	return count
}
/*
	Performance analysis
	- Better than 2.4, 2.3 if the number in bit format is sparse in zero's, in which we can skip a lot of unecessary looping
*/
