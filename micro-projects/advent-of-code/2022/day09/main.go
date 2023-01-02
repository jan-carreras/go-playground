package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// I can record the X,Y positions on a hashmap where T has been
// function T needs to move? If false, we continue

/**+

Store head last position. If we're moving vertically or horizontally,
T must go on the last position of H


If moving diagonally, we have to move on the last position of H

*/

func main() {
	err := run(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
}

func run(input io.Reader) error {
	instructions, err := readInstructions(input)
	if err != nil {
		return fmt.Errorf("unable to read instructions: %w", err)
	}

	processInstructions(instructions)

	return nil
}

type Point struct {
	y, x int
}

type PositionCounter map[string]int

func (pc PositionCounter) Count(p Point) {
	key := fmt.Sprintf("%d-%d", p.x, p.y)
	pc[key]++
}

func processInstructions(instructions []string) int {
	head, tail, pc := Point{}, Point{}, PositionCounter{}

	pc.Count(tail)
	for _, instruction := range instructions {
		direction, times := parseInstruction(instruction)
		for i := 0; i < times; i++ {
			oldHead := head
			switch direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			default:
				panic("unknown direction")
			}

			// We don't need to update tails, right?
			if !isTailOk(head, tail) {
				switch direction {
				case "R", "L":
					tail = oldHead
				case "U", "D":
					tail = oldHead
				}
				pc.Count(tail)
			}

			//fmt.Println(direction, "head:", head, "tail:", tail)

		}
	}

	//printBoard(pc)
	return len(pc)
}

func isTailOk(head, tail Point) bool {
	if abs(head.x-tail.x) > 1 {
		return false
	}

	if abs(head.y-tail.y) > 1 {
		return false
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func printBoard(pc PositionCounter) {
	board := make([][]string, 10)
	for i := range board {
		board[i] = make([]string, 10)
	}

	x, y := 0, 0
	for k := range pc {
		fmt.Sscanf(k, "%d-%d", &x, &y) // TODO: Pretend that works
		board[y][x] = "X"
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board); x++ {
			fmt.Print(board[y][x], "")
		}
		fmt.Println("")

	}
}

func parseInstruction(instruction string) (direction string, times int) {
	parts := strings.Split(instruction, " ")
	direction, timesS := parts[0], parts[1]

	times, _ = strconv.Atoi(timesS) // TODO: Error ignored! Expecting valid input!
	return direction, times
}

func readInstructions(input io.Reader) ([]string, error) {
	in, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("unable to read input: %w", err)
	}

	return strings.Split(string(in), "\n"), nil
}
