package prime

func Factors(n int64) (rsp []int64) {
	for i := int64(2); i <= n; {
		if n%i == 0 {
			rsp = append(rsp, i)
			n /= i
			continue
		}
		i++
	}

	return rsp
}
