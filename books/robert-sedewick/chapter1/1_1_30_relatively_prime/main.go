package main

import (
	"bytes"
	"fmt"
)

// Array exercise. Write a code fragment that creates an N-by-N boolean array
// a[][] such that a[i][j] is true if i and j are relatively prime (have no
// common factors), and false otherwise.

type Table [][]bool

func NewTable(n int) Table {
	table := make(Table, n)
	for y := 0; y < n; y++ {
		table[y] = make([]bool, n)
		for x := range table[y] {
			table[y][x] = areRelativePrime(y, x)
		}
	}
	return table
}

func (t Table) String() string {
	buf := bytes.Buffer{}
	for y := range t {
		for x := range t[y] {
			if t[y][x] {
				buf.WriteString(fmt.Sprint("T "))
			} else {
				buf.WriteString(fmt.Sprint("F "))
			}
		}
		buf.WriteString(fmt.Sprint("\n"))
	}

	return buf.String()
}

func areRelativePrime(x, y int) bool {
	if x == 0 || y == 0 {
		return false
	}

	div := float64(x) / float64(y)
	if div == float64(int(div)) {
		return true
	}

	div = float64(y) / float64(x)
	return div == float64(int(div))
}
