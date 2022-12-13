package isbn

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}

	var sum int

	for index, c := range isbn {
		var number int
		if index == 9 && (c == 'X' || c == 'x') {
			number = 10
		} else if c >= '0' && c <= '9' {
			number = int(c - '0')
		} else {
			return false // Any other letter
		}

		sum += (10 - index) * number
	}

	return sum%11 == 0
}
