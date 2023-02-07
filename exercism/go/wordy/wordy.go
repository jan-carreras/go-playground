package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(?m)(-?\d+)`)

func Answer(question string) (int, bool) {

	question = strings.ToLower(question)

	switch {

	case strings.Contains(question, "plus"):
		a, b, ok := readTwo(question)
		if !ok {
			return 0, false
		}
		return a + b, true
	case strings.Contains(question, "minus"):
		a, b, ok := readTwo(question)
		if !ok {
			return 0, false
		}
		return a - b, true
	case strings.Contains(question, "what is"):
		return readOne(question)
	}

	return 0, true
}

func readTwo(text string) (int, int, bool) {
	ints, ok := readNumbers(text, 2)
	if !ok {
		return 0, 0, false
	}

	return ints[0], ints[1], true
}

func readOne(text string) (int, bool) {
	ints, ok := readNumbers(text, 1)
	if !ok {
		return 0, false
	}

	return ints[0], true
}

func readNumbers(text string, count int) ([]int, bool) {
	g := re.FindAllString(text, -1)
	if len(g) != count {
		return nil, false
	}

	gg := make([]int, 0, len(g))
	for _, number := range g {
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			return nil, false
		}
		gg = append(gg, numberInt)
	}

	return gg, true

}
