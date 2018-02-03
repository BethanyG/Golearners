// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides returns whether a triange is Equ, Iso, Sca or NaT.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	for _, s := range sides {
		if math.IsNaN(s) || math.IsInf(s, 0) {
			return NaT
		}
	}

	if a > 0 && b > 0 && c > 0 && a+b >= c && a+c >= b && b+c >= a {
		if a == b && b == c {
			return Equ
		} else if a != b && b != c && a != c {
			return Sca
		}
		return Iso
	}
	return NaT
}
