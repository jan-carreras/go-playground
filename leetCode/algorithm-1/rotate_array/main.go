package rotate_array

func Rotate(nums []int, k int) {
	l := len(nums)

	k = k % l
	copy(nums, append(nums[l-k:], nums[0:l-k]...))
}
