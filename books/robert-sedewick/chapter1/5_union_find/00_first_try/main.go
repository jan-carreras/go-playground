package main

import (
	"bufio"
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

type pair struct {
	a, b int
}

func run(input io.Reader) error {
	s := bufio.NewScanner(input)

	sites, err := readSites(s)
	if err != nil {
		return err
	}

	pairs, err := readPairs(s)
	if err != nil {
		return err
	}

	fmt.Println("sites=", sites, "pairs=", pairs)

	uf := NewUF(sites)
	for _, p := range pairs {
		if uf.Connected(p.a, p.b) {
			continue
		}
		uf.Union(p.a, p.b)
		fmt.Println(p.a, p.b)
	}

	return nil
}

func readSites(s *bufio.Scanner) (int, error) {
	s.Scan()

	var sites int
	_, err := fmt.Sscanf(s.Text(), "%d", &sites)
	if err != nil {
		return 0, fmt.Errorf("unable to scan sites: %w", err)
	}
	return sites, nil
}

func readPairs(s *bufio.Scanner) ([]pair, error) {
	pairs := make([]pair, 0)
	for s.Scan() {
		p := pair{}
		_, err := fmt.Sscanf(s.Text(), "%d %d", &p.a, &p.b)
		if err != nil {
			return nil, fmt.Errorf("unable to parse pair: %w", err)
		}

		pairs = append(pairs, p)

		if err := s.Err(); err != nil {
			return nil, err
		}
	}

	return pairs, nil
}
