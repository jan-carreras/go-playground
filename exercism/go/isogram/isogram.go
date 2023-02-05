package isogram

import "strings"

func IsIsogram(word string) bool {
	set := make(map[rune]bool)
	for _, c := range strings.ToLower(word) {
		switch c {
		case ' ', '-':
			continue // Ignore character
		}

		if _, found := set[c]; found {
			return false
		}
		set[c] = true
	}

	return true
}
