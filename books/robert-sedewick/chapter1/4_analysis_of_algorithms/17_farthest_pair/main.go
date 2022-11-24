package farthest_pair

import "sort"

// 1.4.17 Farthest pair (in one dimension)
//
// Write a program that, given an array a[] of N double values, finds a farthest
// pair : two values whose difference is no smaller than the the difference of
// any other pair (in absolute value). The running time of your program should be
// linear in the worst case.

func FarthestPair(input []float64) float64 {
	sort.Float64s(input)
	return input[len(input)-1] - input[0]
}
