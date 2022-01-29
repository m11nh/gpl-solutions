// package lengthconv converts length in feet and meters
package lengthconv

import (
	"fmt"
)

type Feet float64
type Meters float64 

// converts feet to meters
func Ftom(i Feet) Meters { return Meters(i / 3.281) }

// converts meters to feet
func MtoF(i Meters) Feet { return Feet(i * 3.281) }

func (i Meters) String() string { return fmt.Sprintf("%g Meters", i) }
func (i Feet) String() string { return fmt.Sprintf("%g Feet", i) }


