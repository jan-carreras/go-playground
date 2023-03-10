package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(?m)(\d+)`)

func main() {
	err := runPart1(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
}

type pair struct {
	left, right []string
}

func (p pair) IsValid() bool {
	fmt.Println("****************")
	fmt.Println(p.left, p.right, "?????")

	for i := 0; i < len(p.right); i++ {
		// If the left list runs out of items first, the inputs are in the right order.
		if i == len(p.left) {
			fmt.Println("left list finished")
			return true
		}

		fmt.Println("comparing:", p.left[i], p.right[i])

		// If the left integer is lower than the right integer, the inputs are in the
		// right order
		if p.left[i] < p.right[i] {
			fmt.Println("left side is smaller. Right order")
			return true
		} else if p.left[i] > p.right[i] {
			// If the left integer is higher than the right integer, the inputs are not in
			// the right order.
			fmt.Println("left side is bigger. Wrong order")
			return false
		}

		// Otherwise, the inputs are the same integer; continue checking the next part of the input.

	}

	if len(p.left) != 0 && len(p.left) == len(p.right) {
		fmt.Println("lists are equal and non-empty. life is good")
		return true
	}

	fmt.Println("end of right list")
	return false
}

func runPart1(input io.Reader) error {
	pairs, err := readPairs(input)
	if err != nil {
		return err
	}

	sum := 0
	for i, pair := range pairs {
		if pair.IsValid() {
			sum += i + 1
		}
	}

	fmt.Println("What is the sum of the indices of those pairs?", sum)

	return nil
}

func readPairs(input io.Reader) ([]pair, error) {
	raw, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	pairs := make([]pair, 0)
	for _, group := range strings.Split(string(raw), "\n\n") {
		parts := strings.Split(group, "\n")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid number of parts. expected 2 having %d", len(parts))
		}

		pairs = append(pairs, pair{
			left:  re.FindAllString(parts[0], -1),
			right: re.FindAllString(parts[1], -1),
		})
	}
	return pairs, nil
}
