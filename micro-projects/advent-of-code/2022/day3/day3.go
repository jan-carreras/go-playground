package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sum, err := part1(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("The sum of priorities is: %d\n", sum)
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}

	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	}

	panic("unknown character " + string(r))
}

func part1(input io.Reader) (int, error) {
	sum := 0
	s := bufio.NewScanner(input)
	for s.Scan() {
		// Read the input line by line
		rucksack := s.Text()

		// Split the line into two sections
		compartment1, compartment2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]

		set := make(map[rune]bool)
		for _, item := range compartment1 {
			set[item] = true
		}

		// Iterate thru both sections to search for a common element
		for _, item := range compartment2 {
			if _, duplicate := set[item]; duplicate {
				//fmt.Println(string(item), priority(item), compartment1, compartment2)
				// Punctuate the element + sum all the elements
				sum += priority(item)
				break
			}
		}
	}

	return sum, s.Err()
}
