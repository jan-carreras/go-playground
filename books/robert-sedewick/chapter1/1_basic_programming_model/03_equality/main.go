package main

import (
	"fmt"
	"log"
	"os"
)

// Write a program that takes three integer command-line arguments and prints
// equal if all three are equal, and not equal otherwise

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("invalid input")
	}

	if equal(os.Args[1:]) {
		fmt.Println("equal")
		return
	}

	fmt.Println("not equal")

}

// equal return true if all values of the input are equal
// returns false if the input is empty
func equal(input []string) bool {
	if len(input) == 0 {
		return false
	}

	first := input[0]
	for i := 1; i < len(input); i++ {
		if first != input[i] {
			return false
		}
	}

	return true
}
