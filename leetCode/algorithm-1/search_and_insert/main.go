package search_and_insert

/** Given a sorted array of distinct integers and a target value, return the
index if the target is found. If not, return the index where it would be if it
were inserted in order. You must write an algorithm with O(log n) runtime
complexity.
*/

func SearchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	mid := 0
	for low <= high {
		mid = (high-low)/2 + low
		switch {
		case nums[mid] > target:
			high = mid - 1
		case nums[mid] < target:
			low = mid + 1
		default:
			return mid // Number found
		}
	}

	return low
}
