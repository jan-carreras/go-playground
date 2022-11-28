package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/5_union_find/union_find/union_find"
	"io"
	"log"
	"os"
)

type pair struct {
	a, b int
}

func main() {
	var algorithm string
	flag.StringVar(&algorithm, "algorithm", "", "algorithm to be used: quickFind,quickUnion,weightedQuickUnion")
	flag.Parsed()

	if algorithm == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := Run(algorithm, os.Stdin, os.Stdout); err != nil {
		log.Fatalln(err)
	}
}

func Run(algorithm string, input io.Reader, output io.Writer) error {
	s := bufio.NewScanner(input)

	sites, err := readSites(s)
	if err != nil {
		return err
	}

	unionFind, err := union_find.New(union_find.Algo(algorithm), sites)
	if err != nil {
		return err
	}

	pairs, err := readPairs(s)
	if err != nil {
		return err
	}

	for _, p := range pairs {
		if unionFind.Connected(p.a, p.b) {
			continue
		}
		unionFind.Union(p.a, p.b)
		_, err := fmt.Fprintf(output, "%d %d\n", p.a, p.b)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(output, "%d components", unionFind.Count())
	if err != nil {
		return err
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
