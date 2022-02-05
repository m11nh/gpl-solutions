// squashes adjacent spaces in a byte slice into a single ASCII space
package squash

import (
	"unicode"
)

// squashes each run of adjacent Unicode spaces in a UTF-8 encoded []byte slice into a single ASCII space
func Squash(x []byte) []byte {
	N := len(x)
	marked := false
	markedIndex := 0
	for i := 1; i < N; i++ {
		if marked == false && unicode.IsSpace(rune(x[i])) && unicode.IsSpace(rune(x[i-1])) {
			marked = true 
			markedIndex = i
		} else if marked == true && !unicode.IsSpace(rune(x[i])) {
			leftShift(&x, i, i - markedIndex)	
			N -= i - markedIndex
			marked = false
		}
	}
	if marked == true {
		N = markedIndex
	}
	return x[:N]
}

// shift the bytes from from and at the right of pos left n times
func leftShift(x *[]byte, pos int, n int) {
	N := len(*x)
	for i := pos; i < N; i++ {
		(*x)[i - n] = (*x)[i]
	}
}
	
