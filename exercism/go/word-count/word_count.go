package wordcount

import (
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	f := make(Frequency)

	for _, c := range strings.FieldsFunc(strings.ToLower(phrase), split) {
		f[clean(c)]++
	}

	return f
}

func clean(s string) string {
	b := strings.Builder{}

	validApostrophe := func(i int) bool {
		if i == 0 || i == len(s)-1 {
			return false
		}

		return true
	}

	for i, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || (c == '\'' && validApostrophe(i)) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func split(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == ','
}
