package main

import (
	"bytes"
	"fmt"
)

const (
	trueChar  = '*'
	falseChar = ' '
)

type Box [][]bool

// 1.1.11 Write a code fragment that prints the contents of a two-dimensional
// boolean array, using * to represent true and a space to represent false.
// Include row and column numbers.

func (b Box) String() string {
	/*c := [][]bool(b)
	output += fmt.Sprintf("debug:\n%+v\n\n", c)*/

	output := bytes.Buffer{}

	// Print the header
	header := &bytes.Buffer{}
	header.WriteString("     ")
	for y := 0; y < len(b); y++ {
		header.WriteString(fmt.Sprintf("%d ", y))
	}

	header.WriteString("\n     ")
	for y := 0; y < len(b); y++ {
		header.WriteString("↓ ")
	}

	output.ReadFrom(header)

	// Write the table
	for y := 0; y < len(b); y++ {
		line := &bytes.Buffer{}
		line.WriteString(fmt.Sprintf("%d -> ", y))
		for x := 0; x < len(b[0]); x++ {
			if b[y][x] {
				line.WriteByte(trueChar)
			} else {
				line.WriteByte(falseChar)
			}
			line.WriteByte(' ')
		}
		output.WriteByte('\n')
		output.ReadFrom(line)
	}

	return output.String()
}
