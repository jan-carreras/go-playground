package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	b := strings.Builder{}

	if number%3 == 0 {
		b.WriteString("Pling")
	}
	if number%5 == 0 {
		b.WriteString("Plang")
	}
	if number%7 == 0 {
		b.WriteString("Plong")
	}

	if b.Len() == 0 {
		b.WriteString(strconv.Itoa(number))
	}

	return b.String()
}
