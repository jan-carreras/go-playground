package encode

import (
	"bytes"
	"fmt"
	"strconv"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	writeToBuff := func(buf bytes.Buffer, count int, seen rune) bytes.Buffer {
		// Do not put a number if the only one character has been seen
		if count != 1 {
			buf.WriteString(fmt.Sprintf("%d", count))
		}
		buf.WriteString(string(seen))
		return buf
	}

	buf := bytes.Buffer{}
	var seen rune
	var count int
	for _, c := range input {
		if seen == 0 {
			// First character — store it and continue processing
			seen = c
			count++
		} else if seen != c {
			// The character has changed — write to buffer and record new char
			buf = writeToBuff(buf, count, seen)
			seen, count = c, 1
		} else {
			count++
		}
	}

	// Process the remaining characters
	buf = writeToBuff(buf, count, seen)

	return buf.String()
}

func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return ""
	}

	buf := bytes.Buffer{}

	var num string
	for _, c := range input {
		// Search for a number specifier
		if c >= '0' && c <= '9' {
			num += strconv.Itoa(int(c - '0'))
			continue
		}

		// If no number precedes, means that we only need one occurrence
		if num == "" {
			buf.WriteRune(c)
			continue
		}

		// Parse the number of repetitions needed and write the corresponding characters
		n, _ := strconv.Atoi(num)
		for i := 0; i < n; i++ {
			buf.WriteRune(c)
		}
		// Reset the parsed number to empty
		num = ""
	}

	return buf.String()
}
