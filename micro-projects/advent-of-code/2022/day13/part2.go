package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func part2(input io.Reader) error {
	raw, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	ordered := []any{parse("[[2]]"), parse("[[6]]")}

	lines := strings.Split(string(raw), "\n")
	for i := 0; i < len(lines); i += 3 {
		left, right := lines[i], lines[i+1]
		ordered = append(ordered, parse(left), parse(right))
	}

	sort.Slice(ordered, func(i, j int) bool {
		return cmp(ordered[i], ordered[j]) <= 0
	})

	result := 1
	for index, a := range ordered {
		if marshal(a) == "[[2]]" || marshal(a) == "[[6]]" {
			result *= index + 1
		}
	}

	fmt.Println("What is the decoder key for the distress signal?", result)

	return nil
}
