package lsproduct

import (
	"errors"
)

var ErrMustBePositive = errors.New("span must be positive")
var ErrNotSmallerThanLength = errors.New("span must be smaller than string length")
var ErrInvalidNumber = errors.New("invalid number")

func LargestSeriesProduct(digits string, span int) (result int64, err error) {
	// Sanity checks
	if span < 0 {
		return 0, ErrMustBePositive
	} else if span > len(digits) {
		return 0, ErrNotSmallerThanLength
	} else if digits == "" {
		return 1, nil
	}

	for i := 0; i < len(digits)-span+1; i++ {
		mult := int64(1)
		for j := i; j < i+span; j++ {
			if !isDigit(digits[j]) {
				return 0, ErrInvalidNumber
			}

			mult *= int64(digits[j] - '0')
		}
		result = max(result, mult)
	}

	return result, nil
}

func isDigit(r uint8) bool {
	return r >= '0' && r <= '9'
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
