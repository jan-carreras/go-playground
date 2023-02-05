package pangram

const (
	azChars = 25
)

func IsPangram(input string) bool {
	if len(input) < azChars {
		return false
	}

	seen, totalSeen := make([]bool, azChars+1), 0
	for _, char := range input {
		if char >= 'A' && char <= 'Z' { // To lowercase
			char = char - 'A'
		} else if char >= 'a' && char <= 'z' {
			char = char - 'a'
		} else {
			continue // Ignore non-ascii [a-z] chars (fuck unicode, BTW)
		}

		if !seen[int(char)] {
			seen[int(char)] = true
			totalSeen++
		}

		if totalSeen > azChars {
			return true
		}
	}

	return false
}
