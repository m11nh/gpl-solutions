package removeDuplicates

// removes adjacent duplicates in a []string slice
func RemoveDuplicates(x []string) []string {
	if len(x) == 0 {
		return x
	}
	l := 1
	for c := 1; c < len(x); c++ {
		if x[c] != x[c-1] {
			x[l] = x[c]
			l++
		}
	}
	return x[:l]
}
