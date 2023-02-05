package armstrong

func IsNumber(n int) bool {
	exponent := digits(n)

	sum := 0
	for nc := n; nc > 0; nc /= 10 {
		sum += pow(nc%10, exponent)
	}

	return sum == n
}

func digits(n int) (digits int) {
	for ; n > 0; n /= 10 {
		digits++
	}
	return digits
}

func pow(base, pow int) int {
	result := 1
	for j := 0; j < pow; j++ {
		result *= base
	}
	return result
}
