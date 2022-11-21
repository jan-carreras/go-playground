package squares_sorted

import "sort"

func SortedSquares(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		nums[i] *= nums[i]
	}

	sort.Ints(nums)

	return nums
}
