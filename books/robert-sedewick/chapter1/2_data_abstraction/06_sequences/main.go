package main

// A string s is a circular rotation of a string t if it matches when the
// characters are circularly shifted by any number of positions; e.g., ACTGACG is
// a circular shift of TGACGAC, and vice versa. Detecting this condition is
// important in the study of genomic sequences. Write a program that checks
// whether two given strings s and t are circular shifts of one another.
//
// Hint : The solution is a one-liner with indexOf(), length(), and string
// concatenation.

func IsCircular(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1 == (s2[i:len(s1)] + s2[0:i]) {
			return true
		}
	}

	return false
}
