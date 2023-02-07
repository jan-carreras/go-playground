package atbash

import "strings"

func Atbash(str string) string {
	b := strings.Builder{}
	b.Grow(len(str) + len(str)/5) // Same number of characters + spaces needed

	for _, s := range str {
		if s >= 'a' && s <= 'z' {
			b.WriteByte(byte('z' - (s - 'a')))
		} else if s >= 'A' && s <= 'Z' {
			b.WriteByte(byte('z' - (s - 'A')))
		} else if s >= '0' && s <= '9' {
			b.WriteRune(s)
		}

		if b.Len()%6 == 5 {
			b.WriteRune(' ')
		}
	}

	s := b.String()
	if len(s) > 1 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}

	return s
}
