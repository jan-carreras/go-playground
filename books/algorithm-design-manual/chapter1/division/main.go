package division

import (
	"fmt"
	"math"
)

func Div(a, b int) int {
	i := 0
	for rem := a; rem >= b; i++ {
		rem = rem - b
	}

	return i
}

func FastDiv(a, b int) (result int) {
	dividend := a // 10

	for dividend >= b {
		divisor := b // 3

		var i int
		for i = 1; (divisor + divisor) <= dividend; i++ {
			divisor += divisor // divisor *= 2
		}
		result += i
		fmt.Println("[", a, b, "]", i, dividend, divisor, result)

		dividend = dividend - divisor // last iteration's divisor
	}

	return result
}

func FasterDiv2(a, b int) (result int) {
	// Check limits
	// abs(a), abs(b)
	for a-b >= 0 {
		i := 0
		for ; a-(b<<1<<i) >= 0; i++ {
		}
		result += 1 << i
		//fmt.Printf("[%d, %d] i=%d res=%d\n", a, b, i, result)
		a -= b << i
		//fmt.Println(a, b<<i, i)
	}

	// Return correct sign

	return result
}

func FasterDiv(a, b int) (result int) {
	if math.MinInt32 == a && b == -1 {
		return math.MaxInt32
	}

	// Check limits
	dividend, divisor := a, b
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for a-b >= 0 {
		x := 0
		for a-(b<<1<<x) >= 0 {
			x++
		}

		result += 1 << x // Misconception! We need the MULTIPLIER, not the INDEX POSITION!!
		a -= b << x
	}

	// Return correct sign
	if dividend >= 0 && divisor >= 0 {
		return result
	}

	if dividend <= 0 && divisor <= 0 {
		return result
	}

	return -result
}

func FD(dividend, divisor int) (groupsOfDivisors int) {
	// Overflow checks

	positiveResult := (dividend > 0) == (divisor > 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	// Algorithm: How many groups of DIVISOR can we make out of
	// the dividend

	for dividend >= divisor {
		x := 0
		for dividend-(divisor<<1<<x) >= 0 {
			x++
		}

		groupsOfDivisors += 1 << x
		dividend -= divisor << x
	}

	if positiveResult {
		if groupsOfDivisors > math.MaxInt32 {
			return math.MaxInt32
		}
		return groupsOfDivisors
	}

	if -groupsOfDivisors < math.MinInt32 {
		return math.MinInt32
	}

	return -groupsOfDivisors
}
