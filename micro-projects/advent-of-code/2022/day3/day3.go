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
		rucksack := s.Text()

		compartment1, compartment2 := makeSet(rucksack[:len(rucksack)/2]), makeSet(rucksack[len(rucksack)/2:])

		item := firstUnion(unions(compartment1, compartment2))
		sum += priority(item)
	}

	return sum, s.Err()
}

func part2(input io.Reader) (int, error) {
	sum := 0
	s := bufio.NewScanner(input)

	elfGroup := make([]map[rune]bool, 3)
	for i := 0; s.Scan(); i = (i + 1) % 3 {
		rucksack := makeSet(s.Text())
		elfGroup[i] = rucksack

		if i == 2 {
			commonItem := firstUnion(unions(elfGroup...))
			sum += priority(commonItem)
		}
	}

	return sum, s.Err()
}

func makeSet(s string) map[rune]bool {
	set := make(map[rune]bool)
	for _, item := range s {
		set[item] = true
	}
	return set
}

func firstUnion(set map[rune]bool) rune {
	for r := range set {
		return r
	}
	panic("the set is empty")
}

func unions(maps ...map[rune]bool) map[rune]bool {
	if len(maps) == 0 {
		return make(map[rune]bool)
	} else if len(maps) == 1 {
		return maps[0]
	}

	u := union(maps[0], maps[1])
	for i := 2; i < len(maps); i++ {
		u = unions(u, maps[i])
	}

	return u
}

func union(m1, m2 map[rune]bool) map[rune]bool {
	u := make(map[rune]bool)
	for k := range m2 {
		if _, found := m1[k]; found {
			u[k] = true
		}
	}

	return u
}
