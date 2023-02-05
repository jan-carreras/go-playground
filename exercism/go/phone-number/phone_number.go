package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidLength = errors.New("invalid length")
var ErrInvalidCharacters = errors.New("invalid characters")
var ErrInvalidAreaCode = errors.New("invalid area code")
var ErrInvalidExchangeCode = errors.New("invalid exchange code")

const (
	numberExpectedLength = 10
)

func Number(phoneNumber string) (string, error) {
	return numbers(phoneNumber)
}

func AreaCode(phoneNumber string) (string, error) {
	n, err := numbers(phoneNumber)
	if err != nil {
		return "", err
	}
	return n[:3], err
}

func Format(phoneNumber string) (string, error) {
	n, err := numbers(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}

func numbers(number string) (string, error) {
	number, err := clean(number)
	if err != nil {
		return "", err
	}

	// Remove country code '1' if needed
	if len(number) == 11 && number[0] == '1' {
		number = number[1:]
	}

	// Sanity checks
	if len(number) != numberExpectedLength {
		return "", ErrInvalidLength
	} else if n := number[0]; n == '0' || n == '1' { // Area code cannot start with 0 or 1
		return "", ErrInvalidAreaCode
	} else if n := number[3]; n == '0' || n == '1' { // Exchange code cannot start with 0 or 1
		return "", ErrInvalidExchangeCode
	}

	return number, nil
}

func clean(number string) (string, error) {
	b := &strings.Builder{}
	b.Grow(len(number))
	for _, c := range number {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			b.WriteRune(c)
		case '-', ' ', '(', ')', '.', '+': // Ignore those chars
		default:
			return "", ErrInvalidCharacters
		}
	}
	return b.String(), nil
}
