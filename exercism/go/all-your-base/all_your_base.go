package allyourbase

import "errors"

var ErrInvalidDigit = errors.New("all digits must satisfy 0 <= d < input base")

// ConvertToBase : Convert a number, represented as a sequence of digits in one base, to any other base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	} else if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	if inputBase == outputBase {
		return inputDigits, nil
	}

	number := 0
	for i, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, ErrInvalidDigit
		}

		exp := len(inputDigits) - i - 1
		number += d * pow(inputBase, exp)
	}

	if number == 0 {
		return []int{0}, nil
	}

	output := make([]int, 0)
	for number > 0 {
		output = append(output, number%outputBase)
		number /= outputBase
	}

	for i := 0; i < len(output)/2; i++ {
		output[i], output[len(output)-i-1] = output[len(output)-i-1], output[i]
	}

	return output, nil
}

func pow(base, exp int) int {
	result := 1
	for i := 1; i <= exp; i++ {
		result *= base
	}

	return result
}
