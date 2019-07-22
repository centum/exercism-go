// Package triangle implement function KindFromSides
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = iota // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides return kind triangle
func KindFromSides(a, b, c float64) Kind {
	for _, v := range []float64{a, b, c} {
		if v == 0 || math.IsNaN(v) || math.IsInf(v, 1) || math.IsInf(v, -1) {
			return NaT
		}
	}

	if a + b < c || a + c < b || c + b < a {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
