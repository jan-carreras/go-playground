package romannumerals

import (
	"errors"
	"fmt"
	"strings"
)

const (
	maxRomanNumber = 3999
)

var conversions = []struct {
	value int
	digit string
}{
	{1000, "M"}, {900, "CM"},
	{500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"},
	{5, "V"}, {4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 {
		return "", errors.New("zero or negative numbers not supported")
	} else if input > maxRomanNumber {
		return "", fmt.Errorf("max roman number is %d: %d cannot be converted", maxRomanNumber, input)
	}

	b := strings.Builder{}
	for _, conversion := range conversions {
		for input >= conversion.value {
			b.WriteString(conversion.digit)
			input -= conversion.value
		}
	}

	return b.String(), nil
}
