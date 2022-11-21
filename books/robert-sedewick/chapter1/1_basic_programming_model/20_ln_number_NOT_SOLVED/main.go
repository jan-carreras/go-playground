package main

// 1.1.20 Write a recursive static method that computes the value of: ln(N!)

// factorial 5: 5 * 4 * 3 * 2 * 1

func lnFactorial(n int) int {
	return 0
}

func factorial(n float64) float64 {
	if n == 1 {
		return n
	}

	return factorial(n-1) * n
}
