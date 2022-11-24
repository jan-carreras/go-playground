package main

import (
	"bufio"
	"errors"
	adt "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_adt"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var k int
	flag.IntVar(&k, "k", 0, "kth element to print")
	flag.Parse()

	if err := Tail(os.Stdin, os.Stdout, k); err != nil {
		log.Fatalln(err)
	}
}

func Tail(input io.Reader, output io.Writer, k int) error {
	s := adt.NewTypeStack[string]()

	if err := readLines(input, s); err != nil {
		return fmt.Errorf("readLines: %v", err)
	}

	if k >= s.Len() {
		return errors.New("not enough lines")
	}

	// Discard k-1 elements
	for i := 0; i < k; i++ {
		s.SPop()
	}

	_, err := fmt.Fprintf(output, "the %dth line is: %s\n", k, s.SPop())
	if err != nil {
		return err
	}

	return nil
}

func readLines(input io.Reader, s *adt.TypeStack[string]) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		s.Push(scanner.Text())
	}
	return scanner.Err()
}
