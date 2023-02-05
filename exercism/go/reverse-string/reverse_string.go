package reverse

func Reverse(input string) string {
	s := []rune(input)

	for i, l := 0, len(s); i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}

	return string(s)
}
