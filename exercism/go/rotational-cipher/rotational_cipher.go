package rotationalcipher

import (
	"bytes"
)

// RotationalCipher we're assuming ASCII
func RotationalCipher(plain string, shiftKey int) string {
	shiftKey %= 26

	out := bytes.Buffer{}

	for _, c := range plain {
		if c >= 'a' && c <= 'z' {
			out.WriteRune(convert(c, 'a', shiftKey))
		} else if c >= 'A' && c <= 'Z' {
			out.WriteRune(convert(c, 'A', shiftKey))
		} else {
			out.WriteRune(c)
		}
	}

	return out.String()
}

func convert(c, start rune, shiftKey int) rune {
	return (((c - start) + rune(shiftKey)) % 26) + start
}
