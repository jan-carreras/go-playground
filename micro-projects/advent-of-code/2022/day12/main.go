package main

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
)

func main() {
	err := run(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
}

type point struct {
	x, y   int
	parent *point
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p point) Equal(other point) bool {
	return p.x == other.x && p.y == other.y
}

func isWalkable(board [][]string, p1, p2 point) bool {
	// Ignore off-limit points
	if p2.x < 0 || p2.x >= len(board[0]) {
		return false
	}
	if p2.y < 0 || p2.y >= len(board) {
		return false
	}

	getNormalisedValue := func(p point) rune {
		value := board[p.y][p.x]
		switch value {
		case "E":
			return 'z'
		case "S":
			return 'a'
		}
		return rune(value[0])
	}

	// Normalise S and E to their corresponding values
	a := getNormalisedValue(p1)
	b := getNormalisedValue(p2)

	// Compute the difference between the two points
	diff := abs(a - b)
	return diff <= 1
}

func abs(a rune) rune {
	if a < 0 {
		return -a
	}
	return a
}

func run(input io.Reader) error {
	board, err := readBoard(input)
	if err != nil {
		return err
	}

	start, end := startEndPositions(board)

	set := make(map[string]bool)

	l := list.New()
	l.PushFront(start)

	for l.Len() != 0 {
		n := l.Remove(l.Front()).(point)
		if _, visited := set[n.String()]; visited {
			continue
		}
		set[n.String()] = true

		if n.Equal(end) {
			printResult(n, board)
			return nil
		}

		for _, p := range possibleNextSteps(board, n) {
			m := n
			p.parent = &m // To be able to trace back the path
			l.PushBack(p)
		}
	}

	return fmt.Errorf("we haven't found a solution :(")
}

// runPart2 : same as part1 but in reverse order. We walk from the End to the first
// point with value 'a'
func runPart2(input io.Reader) error {
	board, err := readBoard(input)
	if err != nil {
		return err
	}

	_, end := startEndPositions(board)
	start := end // We start walking from the top of the hill, finding the lowest point

	set := make(map[string]bool)

	l := list.New()
	l.PushFront(start)

	isStartPoint := func(p point) bool {
		value := board[p.y][p.x]
		switch value {
		case "S", "a":
			return true
		}
		return false
	}

	for l.Len() != 0 {
		n := l.Remove(l.Front()).(point)
		if _, visited := set[n.String()]; visited {
			continue
		}
		set[n.String()] = true

		if isStartPoint(n) { // We have found the lowest point
			printResult(n, board)
			return nil
		}

		for _, p := range possibleNextSteps(board, n) {
			m := n
			p.parent = &m // To be able to trace back the path
			l.PushBack(p)
		}
	}

	return fmt.Errorf("we haven't found a solution :(")
}

func printResult(n point, board [][]string) {
	fmt.Println("We have found a route!!")
	steps := 0
	node := &n
	for node != nil {
		node = node.parent
		steps++
	}
	fmt.Println("Number of steps:", steps-1) // We don't count the starting point

	printTrail(n, board)
}

func printTrail(po point, board [][]string) {
	path := make(map[string]bool)
	n := &po
	for n != nil {
		path[n.String()] = true
		n = n.parent
	}

	white := color.New(color.FgGreen).Add(color.Bold)

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			p := point{x: x, y: y}
			if _, ok := path[p.String()]; ok {
				white.Printf("%s", board[y][x])
			} else {
				fmt.Printf("%s", board[y][x])
			}
		}
		fmt.Println()
	}
}

func possibleNextSteps(board [][]string, p point) (next []point) {
	directions := []point{
		{x: p.x - 1, y: p.y}, // left
		{x: p.x + 1, y: p.y}, // right
		{x: p.x, y: p.y + 1}, // up
		{x: p.x, y: p.y - 1}, // down
	}

	for _, direction := range directions {
		if isWalkable(board, p, direction) {
			next = append(next, direction)
		}
	}

	return next
}

func startEndPositions(board [][]string) (start point, end point) {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == "S" {
				start = point{x: x, y: y}
			}
			if board[y][x] == "E" {
				end = point{x: x, y: y}
			}
		}
	}

	return start, end
}

func readBoard(input io.Reader) (board [][]string, err error) {
	b := bufio.NewScanner(input)
	for b.Scan() {
		l := make([]string, 0)
		for _, c := range b.Text() {
			l = append(l, string(c))
		}
		board = append(board, l)
	}

	return board, b.Err()
}
