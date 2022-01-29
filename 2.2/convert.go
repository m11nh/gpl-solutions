// Converts input into length, weight and temperature measurements
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/miniman1/gobooksolutions/2.2/lengthconv"
	"github.com/miniman1/gobooksolutions/2.2/tempconv"
	"github.com/miniman1/gobooksolutions/2.2/weightconv"
)

// scan args or user input, prints out conversions
func main() {
	args := os.Args[1:]
	for _, v := range args {
		input, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input: %s\n", v)
			os.Exit(1)
		}
		convertAndDisplay(input)
	}

	if len(args) == 0 {
		var input string
		fmt.Scanln(&input)
		for {
			v, err2 := strconv.ParseFloat(input, 64)
			if err2 != nil {
				fmt.Fprintf(os.Stderr, "invalid input: %s\n", input)
				os.Exit(1)
			}
			convertAndDisplay(v)
			fmt.Scanln(&input)
		}
	}
}

// converts and displays input into length, weight and temperature measurements
func convertAndDisplay(i float64) {
	c := tempconv.Celsius(i)
	f := tempconv.Fahrenheit(i)
	fmt.Printf("temperature conversions: %s = %s, %s = %s\n", c, tempconv.CToF(c), f, tempconv.FToC(f))

	m := lengthconv.Meters(i)
	ft := lengthconv.Feet(i)
	fmt.Printf("length conversions: %s = %s, %s = %s\n", m, lengthconv.MtoF(m), ft, lengthconv.Ftom(ft))

	kg := weightconv.Kilogram(i)
	lbs := weightconv.Pound(i)
	fmt.Printf("weight conversions: %s = %s, %s = %s\n", kg, weightconv.Ktop(kg), lbs, weightconv.PtoK(lbs))
}
