package binary_search704

/** Given an array of integers nums which is sorted in ascending order, and an
integer target, write a function to search target in nums. If target exists,
then return its index. Otherwise, return -1. You must write an algorithm with
O(log n) runtime complexity.
*/

func Search(nums []int, target int) int {
	lower, upper := 0, len(nums)-1
	mid := 0
	for lower <= upper {
		mid = ((upper - lower) / 2) + lower
		switch {
		case nums[mid] < target:
			lower = mid + 1
		case nums[mid] > target:
			upper = mid - 1
		default:
			return mid
		}
	}

	return -1
}
