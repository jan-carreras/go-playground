package close_pair

import (
	"math"
	"sort"
)

// 1.4.16 Closest pair (in one dimension). Write a program that, given an array
// a[] of N double values, finds the closest pair : two values whose difference is
// no greater than the difference of any other pair (in absolute value). The
// running time of your program should be linearithmic in the worst case.

func ClosePair(input []float64) float64 {
	sort.Float64s(input)

	min := math.MaxFloat64
	for i := 0; i < len(input)-1; i++ {
		v := input[i+1] - input[i]
		if v < min {
			min = v
		}
	}
	return min
}
