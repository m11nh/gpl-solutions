// Counts number of differnt bits between two SHA256 hashes
package shadiff

const (
	bitsInByte = 8 
)

// Returns number of different bits between two sha256
func Diffsha256(a [32]byte, b [32]byte) int {
	diff := 0
	for i := 0; i < 32; i++ {
		diff += bitDiff(a[i], b[i])	
	}
	return diff
}

// Counts the difference in bits between two bytes
func bitDiff(a byte, b byte) int {
	count := 0
	for i := 0; i < bitsInByte; i++ {
		if a&1 == b&1 {
			count++
		}
		a >>= 1
		b >>= 1
	}
	return 0
}

