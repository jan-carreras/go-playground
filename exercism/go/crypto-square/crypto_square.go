package cryptosquare

import (
	"fmt"
	"strings"
)

func Encode(str string) string {
	buf := strings.Builder{}
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			buf.WriteRune(c)
		} else if c >= 'A' && c <= 'Z' {
			buf.WriteByte(byte(c - 'A' + 'a'))
		} else if c >= '1' && c <= '9' {
			buf.WriteRune(c)
		}
	}

	str = buf.String()
	buf.Reset()

	r, c := 1, 1
	for r*c < len(str) {
		if r == c {
			c++
		} else {
			r++
		}
	}

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			idx := j*c + i
			if idx >= len(str) {
				break
			}
			buf.WriteByte(str[idx])
		}
	}

	str = buf.String()
	buf.Reset()

	fmt.Println(str)

	padding := c - (c*r - len(str))

	// "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn  sseoau "

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			idx := r*i + j
			if idx >= len(str) {
				break
			}

			if i >= padding && j == r-1 {
				buf.WriteString("  ")
				buf.WriteByte(str[idx])
			} else {
				buf.WriteByte(str[idx])
				if idx == len(str)-1 {
					break
				}
				if j == r-1 || idx == len(str)-1 {
					buf.WriteByte(' ')
				}
			}
		}

	}
	str = buf.String()
	fmt.Printf("%q\n", str)
	buf.Reset()

	return str
}
