package main

// PalindromePermutation : Given a string, write a function to check if it is a
// permutation of a palindrome. A palindrome is a word or phrase that is the
// same forwards and backwards. A permutation is a rearrangement of letters.The
// palindrome does not need to be limited to just dictionary words. EXAMPLE
// Input: Tact Coa Output: True (permutations: "taco cat". "atco cta". etc.)
// Hints: #106, #121, #134, #136
func PalindromePermutation(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Create a map from rune->int (counter)

	// Iterate all runes on string and count them
	// Iterate the map with counters and remove all elements divisible by 2
	// If the len(counter) == 0 or len(counter) == 1, return true, else false

	counters := make(map[rune]uint)
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			counters[r]++
		}
	}

	for k, v := range counters {
		if v%2 == 0 {
			delete(counters, k)
		}
	}

	return len(counters) == 1 || len(counters) == 0
}
