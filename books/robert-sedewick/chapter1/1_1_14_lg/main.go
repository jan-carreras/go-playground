package __1_14_lg

// Write a static method lg() that takes an int value N as argument and returns
// the largest int not larger than the base-2 logarithm of N. Do not use Math.

// n=10,  result=3
// n=100, result=6

func Lg(n int) int {
	x := 0
	for n-(1<<x) >= 0 {
		x++
	}

	return x - 1
}
