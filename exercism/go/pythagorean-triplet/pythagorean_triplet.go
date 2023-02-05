package pythagorean

import (
	"math"
)

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) (rsp []Triplet) {
	N := float64(max)
	for a := float64(min); a <= N; a++ {
		for b := a + 1; b <= N; b++ {
			k := a*a + b*b
			c := math.Sqrt(k)
			if c != math.Trunc(c) {
				continue
			} else if c > N {
				continue
			}

			rsp = append(rsp, Triplet{int(a), int(b), int(c)})
		}
	}
	return rsp
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) (rsp []Triplet) {
	N := float64(p)
	for a := 1.0; a < N; a++ {
		for b := a + 1; a+b < N; b++ {
			k := a*a + b*b
			c := math.Sqrt(k)
			if c != math.Trunc(c) {
				continue
			}

			if a+b+c == N {
				rsp = append(rsp, Triplet{int(a), int(b), int(c)})
			}
		}
	}
	return rsp
}
