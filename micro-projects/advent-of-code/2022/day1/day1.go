package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	maxCalories, err := part1(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("The elf with most calories is: %d\n", maxCalories)
}

func part1(input io.Reader) (maxCalories int, err error) {
	s := bufio.NewScanner(input)

	lastElfCalories := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			maxCalories = max(maxCalories, lastElfCalories)
			lastElfCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %w", err)
		}

		lastElfCalories += calories
	}

	return max(maxCalories, lastElfCalories), nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
