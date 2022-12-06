package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(os.Stdin); err != nil {
		log.Fatalln(err)
	}

}

func run(input io.Reader) error {
	b, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	total, err := countIsContained(bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	fmt.Printf("[part 1] Ranges that fully contain the other: %d\n", total)

	total, err = countOverlaps(bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	fmt.Printf("[part 2] Ranges that overlap the other: %d\n", total)

	return nil
}

type assignement struct {
	start, end int
}

func countIsContained(input io.Reader) (total int, err error) {
	// read the input from stdin
	s := bufio.NewScanner(input)
	for s.Scan() {
		pair := s.Text()
		// parse the pairs 2-4,6-8
		var a, b assignement
		_, err := fmt.Sscanf(pair, "%d-%d,%d-%d\n", &a.start, &a.end, &b.start, &b.end)
		if err != nil {
			return 0, fmt.Errorf("unable to parse: %q: %w", pair, err)
		}

		// check if a overlaps b, or b overlaps a
		if a.isContained(b) || b.isContained(a) {
			// count them - caveat: 2-4,2-4 how many times need to be counted? probably just one, I would say
			total += 1
		}
	}

	// return result
	return total, s.Err()
}

func countOverlaps(input io.Reader) (total int, err error) {
	// read the input from stdin
	s := bufio.NewScanner(input)
	for s.Scan() {
		pair := s.Text()
		// parse the pairs 2-4,6-8
		var a, b assignement
		_, err := fmt.Sscanf(pair, "%d-%d,%d-%d\n", &a.start, &a.end, &b.start, &b.end)
		if err != nil {
			return 0, fmt.Errorf("unable to parse: %q: %w", pair, err)
		}

		// check if a overlaps b
		if a.overlaps(b) {
			total += 1
		}
	}

	// return result
	return total, s.Err()
}

// isContained: returns true if a is not inside b. Returns true if ranges are equal
func (a assignement) isContained(b assignement) bool {
	return a.start >= b.start && a.start <= b.end &&
		a.end >= b.start && a.end <= b.end
}

// checks if a and b overlaps. a.overlaps(b) == b.overlaps(a)
func (a assignement) overlaps(b assignement) bool {
	return a.start <= b.end && a.end >= b.start
}
