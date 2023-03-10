package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func parse(line string) (parsed any) {
	json.Unmarshal([]byte(line), &parsed)
	return parsed
}

func marshal(line any) (encoded string) {
	b, _ := json.Marshal(line)
	return string(b)
}

func cmp(left, right any) int {
	l, lok := left.([]any)
	r, rok := right.([]any)

	switch {
	case !lok && !rok:
		return int(left.(float64) - right.(float64))
	case !lok:
		l = []any{left}
	case !rok:
		r = []any{right}
	}

	for i := 0; i < len(l) && i < len(r); i++ {
		if res := cmp(l[i], r[i]); res != 0 {
			return res
		}
	}

	return len(l) - len(r)
}

func listsAreEqual(left, right string) bool {
	l, r := parse(left), parse(right)

	return cmp(l, r) <= 0
}

func part1(input io.Reader) error {
	raw, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	sum := 0
	lines := strings.Split(string(raw), "\n")
	for i := 0; i < len(lines); i += 3 {
		left, right := lines[i], lines[i+1]
		if listsAreEqual(left, right) {
			// /3 because each group is three lines.
			// +1 because indexes are zero-based and we want the first pair to be 1, not zero.
			sum += (i / 3) + 1
		}
	}

	fmt.Println("What is the sum of the indices of those pairs?", sum)

	return nil
}
