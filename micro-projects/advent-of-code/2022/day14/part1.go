package day14

import (
	"fmt"
	"io"
	"strings"
)

type point struct {
	left, bottom int
}

type rockPath struct {
	points []point
}

type cave [][]string

func (c cave) String() string {
	b := strings.Builder{}
	for _, level := range c {
		b.WriteString(strings.Join(level, " "))
		b.WriteString("\n")
	}
	return b.String()
}

func findHeightWidth(paths []rockPath) (int, int) {
	height, width := 0, 0
	for _, path := range paths {
		for _, p := range path.points {
			if width < p.left {
				width = p.left
			}
			if height < p.bottom {
				height = p.bottom
			}
		}
	}
	return height, width
}

func makeCave(paths []rockPath) cave {
	height, width := findHeightWidth(paths)

	c := make(cave, height+3)
	for i := range c {
		c[i] = make([]string, (width*2)+1)
		for j := range c[i] {
			c[i][j] = "."
		}
	}

	for _, path := range paths {
		for i := 0; i < len(path.points)-1; i++ {
			from, to := path.points[i], path.points[i+1]
			minLeft, maxLeft := min(from.left, to.left), max(from.left, to.left)
			minRight, maxRight := min(from.bottom, to.bottom), max(from.bottom, to.bottom)
			for j := minRight; j <= maxRight; j++ {
				for k := minLeft; k <= maxLeft; k++ {
					c[j][k] = "#"
				}
			}
		}
	}

	last := len(c) - 1
	for i := 0; i < len(c[last]); i++ {
		c[last][i] = "#"
	}

	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (rp *rockPath) normalise(base int) {
	for i, point := range rp.points {
		point.left -= base
		rp.points[i] = point
	}
}

func normaliseRockPaths(rockPaths []rockPath) ([]rockPath, int) {
	min := rockPaths[0].points[0].left
	for _, rp := range rockPaths {
		for _, p := range rp.points {
			if min > p.left {
				min = p.left
			}
		}
	}
	min = 0

	for _, rp := range rockPaths {
		rp.normalise(min)
	}

	return rockPaths, min
}

func part1(input io.Reader) error {
	rockPath, err := parse(input)
	if err != nil {
		return err
	}

	rockPath, _ = normaliseRockPaths(rockPath)

	cave := makeCave(rockPath)

	sandStartsHere := point{left: 500, bottom: 0}
	cave[sandStartsHere.bottom][sandStartsHere.left] = "@"

EXIT:
	for i := 0; true; i++ {
		sandDrips := sandStartsHere
		loop := true
		for loop {

			if (sandDrips.bottom == len(cave)-1) || sandDrips.left == 0 || sandDrips.left == len(cave[0]) {
				fmt.Println("Solution found!", i)
				break EXIT
			}

			if cave[sandDrips.bottom+1][sandDrips.left] == "." {
				sandDrips.bottom++
			} else if cave[sandDrips.bottom+1][sandDrips.left-1] == "." {
				sandDrips.bottom++
				sandDrips.left--
			} else if cave[sandDrips.bottom+1][sandDrips.left+1] == "." {
				sandDrips.bottom++
				sandDrips.left++
			} else {
				loop = false
			}

			if cave[sandDrips.bottom][sandDrips.left] == "@" {
				fmt.Println("Cave is full!", i+1)
				break EXIT
			}

		}
		cave[sandDrips.bottom][sandDrips.left] = "S"
	}

	fmt.Println(cave)

	return nil
}

func parse(input io.Reader) ([]rockPath, error) {
	raw, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	scanTraces := make([]rockPath, 0)
	for _, line := range strings.Split(string(raw), "\n") {
		rp := rockPath{}
		for _, coordinates := range strings.Split(line, " -> ") {
			p := point{}
			_, err := fmt.Sscanf(coordinates, "%d,%d", &p.left, &p.bottom)
			if err != nil {
				return nil, err
			}
			rp.points = append(rp.points, p)
		}
		scanTraces = append(scanTraces, rp)
	}

	return scanTraces, nil
}
