package main

import (
	"bufio"
	"errors"
	"io"
)

func Part1(input io.Reader) (int, error) {
	sum := 0
	s := bufio.NewScanner(input)
	for s.Scan() {
		rucksack := s.Text()

		compartment1, compartment2 := makeSet(rucksack[:len(rucksack)/2]), makeSet(rucksack[len(rucksack)/2:])

		item, err := getRandomElement(unions(compartment1, compartment2))
		if err != nil {
			return 0, err
		}
		p, err := priority(item)
		if err != nil {
			return 0, err
		}
		sum += p
	}

	return sum, s.Err()
}

func Part2(input io.Reader) (int, error) {
	sum := 0
	s := bufio.NewScanner(input)

	elfGroup := make([]map[rune]bool, 3)
	for i := 0; s.Scan(); i = (i + 1) % 3 {
		rucksack := makeSet(s.Text())
		elfGroup[i] = rucksack

		if i == 2 {
			commonItem, err := getRandomElement(unions(elfGroup...))
			if err != nil {
				return 0, err
			}

			p, err := priority(commonItem)
			if err != nil {
				return 0, err
			}

			sum += p
		}
	}

	return sum, s.Err()
}

// priority returns the priority of each item
func priority(r rune) (int, error) {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1), nil
	}

	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27), nil
	}

	return 0, errors.New("unknown character " + string(r))
}

// makeSet generates a set from each character on a string
func makeSet(s string) map[rune]bool {
	set := make(map[rune]bool)
	for _, item := range s {
		set[item] = true
	}
	return set
}

// getRandomElement returns a random element of a set
func getRandomElement(set map[rune]bool) (rune, error) {
	for r := range set {
		return r, nil
	}
	return 0, errors.New("set is empty, thus cannot get a random element")
}

// unions returns a new union containing elements present in all sets
func unions(sets ...map[rune]bool) map[rune]bool {
	if len(sets) == 0 {
		return make(map[rune]bool)
	} else if len(sets) == 1 {
		return sets[0]
	}

	u := union(sets[0], sets[1])
	for i := 2; i < len(sets); i++ {
		u = union(u, sets[i])
	}

	return u
}

// union returns a new set union of s1 and s2 (elements contained in both sets)
func union(s1, s2 map[rune]bool) map[rune]bool {
	u := make(map[rune]bool)
	for k := range s2 {
		if _, found := s1[k]; found {
			u[k] = true
		}
	}

	return u
}
