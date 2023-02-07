package sieve

func Sieve(limit int) (out []int) {
	out = make([]int, 0, 32)
	s := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		if s[i] == false {
			out = append(out, i)
			for j := i; j <= limit; j += i {
				s[j] = true
			}
		}
	}

	return out
}
