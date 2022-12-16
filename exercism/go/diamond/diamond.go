package diamond

import (
	"errors"
	"strings"
)

func printLine(c byte, max int) string {
	str := make([]byte, max)
	for i := 0; i < max; i++ {
		str[i] = '-'
	}

	diff := c - 'A'
	str[(max/2)-int(diff)] = c
	str[(max/2)+int(diff)] = c

	return string(str)
}

func Gen(c byte) (string, error) {
	// Know how many lines do I have to print
	// Know which character do I have to print on each line
	// function that prints the extra chars
	if c < 'A' || c > 'Z' {
		return "", errors.New("out of bounds")
	}

	max := (int(c-'A') * 2) + 1

	letter := byte('A')
	lines := make([]string, max)
	for i := 0; i < max; i++ {
		lines[i] = printLine(letter, max)
		if i < max/2 {
			letter += 1
		} else {
			letter -= 1
		}

	}
	return strings.Join(lines, "\n"), nil
}
