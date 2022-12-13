package summultiples

func SumMultiples(limit int, divisors ...int) int {
	set := make(map[int]bool, 0)
	for _, divisor := range divisors {
		accumulator := 0
		for accumulator < limit && divisor != 0 {
			set[accumulator] = true
			accumulator += divisor
		}
	}

	sum := 0
	for k := range set {
		sum += k
	}

	return sum
}
