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
	r := sha(shaType, []byte(i))
	fmt.Printf("the sha%d of %s is: %x\n", shaType, i, r)
}

// Gets the sha type, then returns the result of applying the relevant sha function 
func sha(i int, s []byte) interface{} {
	v := []byte(s)
	switch i {
	case 256:
		return sha256.Sum256(v)
	case 384:
		return sha512.Sum384(v)
	default:
		return sha512.Sum512(v)
	}
}
