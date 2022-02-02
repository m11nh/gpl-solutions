package rotate

// Rotates all the elements in a slice left
func Rotate(x []int) []int {
	n := len(x)
	if n <= 1 {
		return x
	}
	f := x[0]
	copy(x[:n-1], x[1:])
	x[n-1] = f
	return x
}
