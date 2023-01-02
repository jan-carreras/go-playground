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

	fmt.Println()
	fmt.Printf("Part 1: The sum of these signal strengths is: %d\n", sum)
	fmt.Println()

	fmt.Println("Part 2")
	fmt.Println("SCREEN")
	cycle := 1
	for j := 0; j < 6; j++ {
		for i := 0; i < 40; i++ {
			value := getSpriteChar(cycles[cycle-1], i)

			fmt.Print(value)
			cycle++
		}
		fmt.Println("")
	}
	fmt.Println("END SCREEN")

	return sum, nil
}

func getSpriteChar(center int, index int) string {
	sprite := ""
	for i := 0; i < 40; i++ {
		if i == center || i == center-1 || i == center+1 {
			sprite += "#"
		} else {
			sprite += " "
		}
	}

	return string(sprite[index])

}

func signalStrength(cycles []int, index int) int {
	return index * cycles[index-1]
}
