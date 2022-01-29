// does conversions from kg to pounds vice versa
package weightconv

import "fmt"

type Kilogram float64
type Pound float64 

// converts input in Kilograms to Pounds
func Ktop(i Kilogram) Pound { return Pound(i * 2.205) }

// converts input in Pounds to Kilograms
func PtoK(i Pound) Kilogram { return Kilogram(i / 2.205) }


func (i Kilogram) String() string { return fmt.Sprintf("%gKg", i) }
func (i Pound) String() string { return fmt.Sprintf("%glbs", i) }

