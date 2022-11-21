package main

/*
*

One Away: There are three types of edits that can be performed on strings:
insert a character, remove a character, or replace a character. Given two
strings, write a function to check if they are one edit (or zero edits) away.
*/

func Distance(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	lenDiff := len(s1) - len(s2)
	if lenDiff > 1 || lenDiff < -1 {
		return false
	}

	changeBudged := 1
	switch {
	case len(s1) == len(s2):
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				changeBudged--
			}
		}
	case len(s1) > len(s2):
		removeBudget := 1
		i1, i2 := 0, 0
		for i2 < len(s2) {
			if s1[i1] != s2[i2] {
				i2++
				removeBudget-- // Removing from the "long one"
			}
			i1++
			i2++
		}

		for i := 0; i < len(s2); i++ {

		}
		// We need to either add or remove a character
		panic("implement me")

	default:
		panic("unknown case")

	}

	return changeBudged >= 0
}
