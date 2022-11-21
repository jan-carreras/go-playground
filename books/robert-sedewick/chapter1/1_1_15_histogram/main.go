package main

func Histogram(lst []int, m int) []int {
	output := make([]int, m)

	for i := 0; i < m; i++ {
		value := lst[i]
		if value < 0 || value >= m {
			continue // Ignore that character
		}
		output[value] += 1
	}

	return output
}
