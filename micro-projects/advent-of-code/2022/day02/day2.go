package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	score, err := part1(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("[first part] my total score would be: %d\n", score)

	score, err = part2(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("[second part] my total score would be: %d\n", score)
}

func part1(input io.Reader) (totalScore int, err error) {
	// Read the input

	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors -> (1 for Rock, 2 for Paper, and 3 for Scissors)
	// Outcome: (0 if you lost, 3 if the round was a draw, and 6 if you won).
	scores := map[string]int{
		"AX": 1 + 3,
		"AY": 2 + 6,
		"AZ": 3 + 0,

		"BX": 1 + 0,
		"BY": 2 + 3,
		"BZ": 3 + 6,

		"CX": 1 + 6,
		"CY": 2 + 0,
		"CZ": 3 + 3,
	}

	s := bufio.NewScanner(input)
	for s.Scan() {
		var they, me string
		if _, err := fmt.Sscanf(s.Text(), "%s %s", &they, &me); err != nil {
			return 0, fmt.Errorf("unable to parse hands: %w", err)
		}

		score, ok := scores[they+me]
		if !ok {
			return 0, fmt.Errorf("unknown hand combination: %score Vs %score", they, me)
		}

		totalScore += score // Accumulate totalScore
	}

	return totalScore, nil
}

func part2(input io.Reader) (totalScore int, err error) {
	// Read the input

	// A for Rock, B for Paper, and C for Scissors
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
	// X for Rock, Y for Paper, and Z for Scissors -> (1 for Rock, 2 for Paper, and 3 for Scissors)
	// Outcome: (0 if you lost, 3 if the round was a draw, and 6 if you won).
	scores := map[string]int{
		// A for Rock
		"AX": 3 + 0,
		"AY": 1 + 3,
		"AZ": 2 + 6,

		// B for Paper
		"BX": 1 + 0,
		"BY": 2 + 3,
		"BZ": 3 + 6,

		// C for Scissors
		"CX": 2 + 0,
		"CY": 3 + 3,
		"CZ": 1 + 6,
	}

	s := bufio.NewScanner(input)
	for s.Scan() {
		var they, me string
		if _, err := fmt.Sscanf(s.Text(), "%s %s", &they, &me); err != nil {
			return 0, fmt.Errorf("unable to parse hands: %w", err)
		}

		score, ok := scores[they+me]
		if !ok {
			return 0, fmt.Errorf("unknown hand combination: %score Vs %score", they, me)
		}

		totalScore += score // Accumulate totalScore
	}

	return totalScore, nil
}
