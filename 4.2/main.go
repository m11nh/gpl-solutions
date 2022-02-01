// Converts and prints SHA values of input from stdin
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var shaType int 

// Setting up the flag from arguments
func init() {
	flag.IntVar(&shaType, "type", 256, "Provides the hashing algorithm used. E.g. 256 uses sha256. Default is sha256")
	flag.Parse()
	if !validShaType(shaType) {
		fmt.Fprintf(os.Stderr, "%d is an invalid sha type\n", shaType)
		os.Exit(1)
	}
}

func validShaType(a int) bool {
	if a != 256 && a != 384 && a != 512 {
		return false
	}
	return true
}

// Gets user input, prints out the hash of it
func main() {
	var i string 
	fmt.Scanf("%s", &i)
	f := typeToSHA(shaType)
	fmt.Printf("the sha%d of %s is: %x\n", shaType, i, f([]byte(i)))
}

// provided the sha type, returns the relevant sha function
func typeToSHA(t int) func(b []byte) []byte {
	switch t {
	case 256: 
		return func(b []byte) []byte {
			x := sha256.Sum256(b)
			return x[:]
		}
	case 384: 
		return func(b []byte) []byte {
			x := sha512.Sum384(b)
			return x[:]
		}
	default:
		// 512
		return func(b []byte) []byte {
			x := sha512.Sum512(b)
			return x[:]
		}
	}
}
