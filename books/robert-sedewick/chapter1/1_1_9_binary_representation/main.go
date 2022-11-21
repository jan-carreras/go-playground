package __1_9_binary_representation

import (
	"bytes"
)

// Write a code fragment that puts the binary representation of a positive integer N
// into a String s.

// Binary
// Warning: this solution works but it's very ineffective since strings are copied every time
func Binary(input int) (output string) {
	if input <= 0 {
		panic("invalid input")
	}

	for ; input > 0; input /= 2 {
		output = string(byte('0'+input%2)) + output
	}

	return output
}

func BinaryFast(input int) string {
	if input <= 0 {
		panic("invalid input")
	}

	buf := bytes.Buffer{}

	for ; input > 0; input /= 2 {
		buf.WriteByte('0' + byte(1&input))
	}

	revResult := buf.Bytes()

	l := len(revResult)
	for i := 0; i < l/2; i++ {
		revResult[i], revResult[l-1-i] = revResult[l-1-i], revResult[i]
	}

	return string(revResult)
}

func BinaryAgain(input int) (output string) {
	for ; input > 0; input /= 2 {
		output = string(byte('0'+(input%2))) + output
	}
	return output
}
