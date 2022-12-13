package quicksort

import "math/rand"

func Sort(input []int) {
	rand.Shuffle(len(input), func(i, j int) {
		input[i], input[j] = input[j], input[i]
	})
	sort(input, 0, len(input)-1)
}

func sort(arr []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(arr, low, high)
	sort(arr, low, p-1)
	sort(arr, p+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
