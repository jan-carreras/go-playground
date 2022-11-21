package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	b := makeBoard()
	fmt.Println(b)
}

func makeBoard() Board {
	counter := 0
	y := make(Board, 2)
	for i := range y {
		y[i] = make([]int, 3)
		for x := range y[i] {
			y[i][x] = counter
			counter++
		}
	}

	return y
}
