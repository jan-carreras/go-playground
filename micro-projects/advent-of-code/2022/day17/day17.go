package day17

import (
	"fmt"
	"strings"
)

/**

####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##
*/

const (
	jetStreams = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
)

type rock struct {
	shape        [][]string
	positionLeft int
}

func (r rock) Heigth() int {
	return len(r.shape)
}

func (r rock) Width() int {
	return len(r.shape[0])
}

func makeLine() rock {
	return rock{
		shape: [][]string{
			{"#", "#", "#", "#"},
		},
		positionLeft: 2,
	}
}

type chamber struct {
	shape [][]string
}

/*func (c chamber) IsRockValid(r rock) bool {
	rockHeigth := len(c.shape) - r.Heigth()
	for i := 0; i < len(r.shape); i++ {
		for j := 0; j < len(r.shape[0]); j++ {
			x := j+r.positionLeft
			y := i+rockHeigth

		}

	}

}*/

func (c chamber) String() string {
	b := strings.Builder{}
	for i := len(c.shape) - 1; i >= 0; i-- {
		b.WriteString(strings.Join(c.shape[i], " "))
		b.WriteString("\n")
	}
	return b.String()
}

func newChamber() chamber {
	height := 7

	shape := make([][]string, height)
	for i := range shape {
		shape[i] = make([]string, 9)
		for j := range shape[i] {
			shape[i][j] = "."
		}
		shape[i][0] = "#"
		shape[i][8] = "#"
	}
	for j := range shape[0] {
		shape[0][j] = "#"
	}

	return chamber{
		shape: shape,
	}
}

var jetStreamIndex = 0

func part1() error {
	// We're only going to keep the "skyline" that represent the height of each column.
	// Every time a rock falls on top, we update the specific heights.
	// We need to know exactly where the rock is failing, tho.

	c := newChamber()

	line := makeLine()

	fmt.Println(c)
	return nil
}
