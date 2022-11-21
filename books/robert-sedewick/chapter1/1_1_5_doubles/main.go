package __1_5_doubles

func BetweenRange(a float64) bool {
	return a >= 0 && a <= 1
}

func Check(a, b float64) bool {
	return BetweenRange(a) && BetweenRange(b)
}

func CheckN(lst ...float64) bool {
	for _, f := range lst {
		if !BetweenRange(f) {
			return false
		}
	}

	return true
}
