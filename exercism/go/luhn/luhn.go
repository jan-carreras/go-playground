package luhn

func Valid(id string) bool {
	sum, digits := 0, 0
	for i := len(id) - 1; i >= 0; i-- {
		if id[i] == ' ' {
			continue // Ignore spaces
		}

		digit := 0
		switch c := id[i]; {
		case c >= '0' && c <= '9':
			digit = int(c - '0')
		default:
			return false // Non-integer character, which makes the input invalid
		}

		mustDouble := digits%2 == 1
		if mustDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		digits++
	}

	return digits >= 2 && sum%10 == 0
}
