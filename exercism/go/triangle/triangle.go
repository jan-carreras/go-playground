// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a+b <= c || a+c <= b || b+c <= a { // Is a triangle?
		return NaT
	} else if a == b && a == c {
		return Equ
	} else if a != b && a != c && b != c {
		return Sca
	} else {
		return Iso
	}
}
