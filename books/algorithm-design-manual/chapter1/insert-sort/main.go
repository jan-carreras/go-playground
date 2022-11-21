package insert_sort

func InsertionSort(s []int) {
	for i := 0; i < len(s); i++ {
		j := i
		for (j > 0) && s[j] < s[j-1] {
			s[j], s[j-1] = s[j-1], s[j]
			j -= 1
		}
	}
}

func InsertionSort2(s []int) {
	for i := 1; i < len(s); i++ {
		j := i
		for j > 0 && s[j] < s[j-1] {
			s[j], s[j-1] = s[j-1], s[j]
			j--
		}
	}
}
