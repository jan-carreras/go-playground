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
	err := run(os.Stdin, 2)
	if err != nil {
		log.Fatalln(err)
	}
}

func run(input io.Reader, version int) error {
	s := bufio.NewScanner(input)

	stacks, err := readStacks(s)
	if err != nil {
		return fmt.Errorf("unable to read stacks: %w", err)
	}

	instructions, err := readInstructions(s)
	if err != nil {
		return fmt.Errorf("unable to read the instructions: %w", err)
	}

	if version == 1 {
		processInstructions(stacks, instructions)
	} else if version == 2 {

		processInstructionsV2(stacks, instructions)
	} else {
		return fmt.Errorf("unknown version %d", version)
	}

	topCrates := getTopCrates(stacks)
	fmt.Println("After the rearrangement procedure completes, what crate ends up on top of each stack?", topCrates)

	return nil
}

func getTopCrates(stacks []string) string {
	output := ""
	for _, stack := range stacks {
		if stack == "" {
			continue

		}
		output += string(stack[len(stack)-1])

	}
	return output
}

func processInstructions(stacks []string, instructions []instruction) {
	for _, instruction := range instructions {
		from := stacks[instruction.from-1]
		// Move the crates in reverse order to the destination
		stacks[instruction.to-1] = stacks[instruction.to-1] + reverseString(from[len(from)-instruction.count:])
		// Remove crates from the origin
		stacks[instruction.from-1] = from[:(len(from) - instruction.count)]
	}
}

func processInstructionsV2(stacks []string, instructions []instruction) {
	for _, instruction := range instructions {
		from := stacks[instruction.from-1]
		// Move the crates in reverse order to the destination
		stacks[instruction.to-1] = stacks[instruction.to-1] + from[len(from)-instruction.count:]
		// Remove crates from the origin
		stacks[instruction.from-1] = from[:(len(from) - instruction.count)]
	}
}

type instruction struct {
	count int
	from  int
	to    int
}

func readInstructions(input *bufio.Scanner) ([]instruction, error) {
	instructions := make([]instruction, 0)
	for input.Scan() {
		t := input.Text()
		inst := instruction{}
		_, err := fmt.Sscanf(t, "move %d from %d to %d", &inst.count, &inst.from, &inst.to)
		if err != nil {
			fmt.Println(instructions)
			return nil, fmt.Errorf("invalid line %q: %w", t, err)
		}
		instructions = append(instructions, inst)

	}

	return instructions, input.Err()
}

func readStacks(input *bufio.Scanner) ([]string, error) {
	stacks := make([]string, 9) // TODO: I can make this dynamic.
	for input.Scan() {
		t := input.Text()
		if t == "" {
			break // End of the block describing the stacks
		}

		for i := 1; i < len(input.Text()); i += 4 {
			crate := t[i]
			if crate == ' ' {
				continue // Empty crate, we don't care about it
			}
			if crate >= '0' && crate <= '9' {
				break // We ignore the line of numbers
			}

			idx := (i - 1) / 4
			stacks[idx] = stacks[idx] + string(crate)
		}
	}

	for i := range stacks {
		stacks[i] = reverseString(stacks[i])
	}

	return stacks, input.Err()
}

func reverseString(input string) string {
	b := bytes.Buffer{}
	for i := len(input) - 1; i >= 0; i-- {
		b.WriteString(string(input[i]))
	}
	return b.String()
}
