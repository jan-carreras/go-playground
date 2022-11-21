package __1_13_transposition

import (
	"bytes"
	"fmt"
)

// Write a code fragment to print the transposition (rows and columns changed) of
// a two-dimensional array with M rows and N columns.

type Board [][]int

func (b Board) String() string {
	output := bytes.Buffer{}
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[0]); x++ {
			output.WriteString(fmt.Sprintf("%d ", b[y][x]))
		}
		output.WriteByte('\n')

	}
	output.WriteString("\n\n")

	for x := 0; x < len(b[0]); x++ {
		for y := 0; y < len(b); y++ {
			output.WriteString(fmt.Sprintf("%d ", b[y][x]))
		}
		output.WriteByte('\n')

	}
	return output.String()
}
