package main

import (
	"bufio"
	"errors"
	stack "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_generic_stack"
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
	s := new(stack.Stack[string])

	if err := readLines(input, s); err != nil {
		return fmt.Errorf("readLines: %v", err)
	}

	if k >= s.Length() {
		return errors.New("not enough lines")
	}

	// Discard k-1 elements
	for i := 0; i < k; i++ {
		s.Pop()
	}

	_, err := fmt.Fprintf(output, "the %dth line is: %s\n", k, s.Pop())
	if err != nil {
		return err
	}

	return nil
}

func readLines(input io.Reader, s *stack.Stack[string]) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		s.Push(scanner.Text())
	}
	return scanner.Err()
}
