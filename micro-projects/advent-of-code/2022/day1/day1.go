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
	maxCalories, err := run(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("The elf with most calories is: %d\n", maxCalories)
}

func run(input io.Reader) (maxCalories int, err error) {
	s := bufio.NewScanner(input)

	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	lastElfCalories := 0
	for _, line := range lines {
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
	maxCalories = max(maxCalories, lastElfCalories)

	return maxCalories, nil
}
