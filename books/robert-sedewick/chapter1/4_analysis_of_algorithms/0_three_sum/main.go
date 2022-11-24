package three_sum

func ThreeSum(lst []int) (count int) {
	l := len(lst)

	cache := make(map[int]int)
	for i, v := range lst {
		cache[v] = i
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			complementary := -(lst[i] + lst[j])
			if position, ok := cache[complementary]; ok && position > j {
				count++
			}
		}
	}
	return count
}
