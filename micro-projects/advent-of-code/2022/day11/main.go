package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const N = 20

func main() {
	if err := run(os.Stdin); err != nil {
		log.Fatalln(err)
	}
}

type Monkey struct {
	monkeyNumber  int
	items         []int // Starting items: 85, 94, 97, 92, 56
	itemsCount    int
	operationOp   string // + or *
	operationTerm string // Operation: new = old + 2 ->> 2
	testDivisor   int    // Test: divisible by 19
	testTrue      *Monkey
	testFalse     *Monkey
}

func run(input io.Reader) error {
	monkeys, err := parseFuckingMonkeys(input)
	if err != nil {
		return err
	}

	for i := 0; i < N; i++ {
		if err := computeRound(monkeys); err != nil {
			return err
		}
		//	printRound(monkeys)
	}

	printInspectedItemsPerMonkey(monkeys)

	return nil
}

func printInspectedItemsPerMonkey(monkeys []Monkey) {
	items := make([]int, 0)
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", monkey.monkeyNumber, monkey.itemsCount)
		items = append(items, monkey.itemsCount)
	}

	sort.Ints(items)
	biggest, secondBiggest := items[len(items)-1], items[len(items)-2]
	fmt.Println("PART 1. What is the level of monkey business after 20 rounds of stuff-slinging simian shenanigans?", biggest*secondBiggest)

}

func printRound(monkeys []Monkey) {
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d: %v\n", monkey.monkeyNumber, monkey.items)
	}
}

func computeRound(monkeys []Monkey) error {
	for i, m := range monkeys {
		debug("Monkey %d:\n", i)
		for _, item := range m.items {
			debug("  Monkey inspects an item with a worry level of %d.\n", item)

			n := item
			if m.operationTerm != "old" {
				num, err := strconv.Atoi(m.operationTerm)
				if err != nil {
					return err
				}
				n = num
			}

			worryLevel := 0
			switch m.operationOp {
			case "*":
				worryLevel = n * item
				//worryLevel = new(big.Int).Mul(&n, &item)
				debug("   Worry level is multiplied by %d to %d.\n", n, worryLevel)
			case "+":
				worryLevel = n + item
				//worryLevel = new(big.Int).Add(&n, &item)
				debug("   Worry level increases by %d to %d.\n", n, worryLevel)
			default:
				return fmt.Errorf("unknown operation: %s", m.operationOp)
			}

			/*
				worryLevel /= 1
				debug("   Monkey gets bored with item. Worry level is divided by %d to %d.\n", 3, worryLevel)
			*/
			if worryLevel%m.testDivisor == 0 {
				debug("   Current worry level is divisible by %s.\n", m.testDivisor)
				debug("   Item with worry level %d is thrown to monkey %d.\n", worryLevel, m.testTrue.monkeyNumber)
				m.testTrue.items = append(m.testTrue.items, worryLevel)
			} else {
				debug("   Current worry level is not divisible by %s.\n", m.testDivisor)
				debug("   Item with worry level %d is thrown to monkey %d.\n", worryLevel, m.testFalse.monkeyNumber)
				m.testFalse.items = append(m.testFalse.items, worryLevel)
			}
		}
		monkeys[i].itemsCount += len(monkeys[i].items)
		monkeys[i].items = nil
	}
	return nil
}

func debug(format string, a ...any) {
	//fmt.Printf(format, a...)
}

func parseFuckingMonkeys(input io.Reader) ([]Monkey, error) {
	in, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	monkeys := make([]Monkey, bytes.Count(in, []byte("Monkey ")))
	monkeysRaw := strings.Split(string(in), "\n\n")
	for i, monkeyRaw := range monkeysRaw {
		m := &monkeys[i]
		m.monkeyNumber = i
		for _, line := range strings.Split(monkeyRaw, "\n") {
			line = strings.Trim(line, " ")
			switch {
			case strings.HasPrefix(line, "Monkey"):
				continue
			case strings.HasPrefix(line, "Starting items: "):
				items := strings.Split(line[len("Starting items: "):], ",")
				for _, i := range items {
					n, err := strconv.Atoi(strings.Trim(i, " "))
					if err != nil {
						return nil, err
					}
					m.items = append(m.items, n)
				}
			case strings.HasPrefix(line, "Operation: "):
				operation := line[len("Operation: "):]
				tokens := strings.Split(operation, " ")
				m.operationOp = tokens[3]
				m.operationTerm = tokens[4]
			case strings.HasPrefix(line, "Test: divisible by "):
				operation := line[len("Test: divisible by "):]
				n, err := strconv.Atoi(operation)
				if err != nil {
					return nil, err
				}
				m.testDivisor = n
			case strings.HasPrefix(line, "If true: "):
				to := line[len("If true: throw to monkey "):]
				n, err := strconv.Atoi(to)
				if err != nil {
					return nil, err
				}
				m.testTrue = &monkeys[n]
			case strings.HasPrefix(line, "If false: "):
				to := line[len("If false: throw to monkey "):]
				n, err := strconv.Atoi(to)
				if err != nil {
					return nil, err
				}
				m.testFalse = &monkeys[n]
			}
		}
	}

	return monkeys, nil
}
