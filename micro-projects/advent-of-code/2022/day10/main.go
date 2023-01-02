package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if _, err := run(os.Stdin); err != nil {
		log.Fatalln(err)
	}
}

func run(input io.Reader) (int, error) {
	buf := bufio.NewScanner(input)

	cycles := []int{1}

	for buf.Scan() {
		line := buf.Text()
		if strings.HasPrefix(line, "noop") {
			cycles = append(cycles, cycles[len(cycles)-1])
		} else if strings.HasPrefix(line, "addx") {
			cycles = append(cycles, cycles[len(cycles)-1])
			var number int
			_, err := fmt.Fscanf(strings.NewReader(line), "addx %d", &number)
			if err != nil {
				return 0, err
			}
			cycles = append(cycles, cycles[len(cycles)-1]+number)
		} else {
			return 0, fmt.Errorf("unknown operation: %s", line)
		}
	}

	sum := 0
	for _, i := range []int{20, 60, 100, 140, 180, 220} {
		fmt.Println(i, signalStrength(cycles, i))
		sum += signalStrength(cycles, i)
	}

	fmt.Printf("The sum of these signal strengths is: %d\n", sum)

	return sum, nil
}
func signalStrength(cycles []int, index int) int {
	return index * cycles[index-1]
}
