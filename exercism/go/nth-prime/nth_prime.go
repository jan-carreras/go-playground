package prime

import (
	"errors"
	"math"
)

var ErrInvalidInput = errors.New("ErrInvalidInput")

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, ErrInvalidInput
	}

	number := 1
	for n > 0 {
		number++
		if isPrime(number) {
			n--
		}
	}

	return number, nil
}

func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
