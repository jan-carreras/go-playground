package grains

import "errors"

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("number must be [1,64]")
	}
	return 1 << (number - 1), nil
}

func Total() (sum uint64) {
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		sum += s
	}
	return sum
}
