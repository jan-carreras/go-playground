package main

import (
	"io"
	"sort"
	"strconv"
	"strings"
)

func part1(input io.Reader) (int, error) {
	elfsCalories, err := readElfCalories(input)
	if err != nil {
		return 0, err
	}

	maxCalories := maxInt(elfsCalories)

	return maxCalories, nil
}

func part2(input io.Reader) (calories int, err error) {
	elfs, err := readElfCalories(input)
	if err != nil {
		return 0, err
	}

	sort.Ints(elfs)

	// Find the top three Elves carrying the most Calories
	return sum(elfs[len(elfs)-3:]...), nil
}

// readElfCalories returns the calories being carried by each elf
func readElfCalories(input io.Reader) (elfsCalories []int, err error) {
	raw, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	// Each elf is differentiated to the next elf by a double \n
	elfs := strings.Split(string(raw), "\n\n")
	for _, elf := range elfs {
		calories, err := parseInts(strings.Split(elf, "\n"))
		if err != nil {
			return nil, err
		}
		elfsCalories = append(elfsCalories, sum(calories...))
	}

	return elfsCalories, nil
}

func parseInts(strings []string) (rsp []int, err error) {
	for _, s := range strings {
		if s == "" { // Ignore empty lines. Might be present at the end of the input
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		rsp = append(rsp, i)
	}

	return rsp, nil
}

func sum(input ...int) (r int) {
	for _, i := range input {
		r += i
	}

	return r
}

func maxInt(ints []int) int {
	m := 0
	for _, i := range ints {
		m = max(m, i)
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
