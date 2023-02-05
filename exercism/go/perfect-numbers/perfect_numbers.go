package perfect

import (
	"errors"
)

type Classification string

const (
	ClassificationAbundant  = "abundant"
	ClassificationDeficient = "deficient"
	ClassificationPerfect   = "perfect"
)

var ErrOnlyPositive = errors.New("ErrOnlyPositive")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	} else if n == 1 {
		return ClassificationDeficient, nil
	}

	sum := int64(1)
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if n/i != i {
				sum += n / i
			}
		}
		if sum > n {
			return ClassificationAbundant, nil
		}
	}

	if sum < n {
		return ClassificationDeficient, nil
	}

	return ClassificationPerfect, nil
}
