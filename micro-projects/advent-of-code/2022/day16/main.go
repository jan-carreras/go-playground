package day16

import (
	"fmt"
	"io"
	"strings"
)

type valve struct {
	name   string
	rate   int
	valves []string
}

func day1(input io.Reader) error {
	raw, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	valves := make([]valve, 0)

	lines := strings.Split(string(raw), "\n")
	for _, line := range lines {
		v := valve{}
		_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &v.name, &v.rate)
		if err != nil {
			return err
		}

		parts := strings.Split(line, ";")
		tunnelsStr := parts[1][len(" tunnel leads to valve ")+1:]
		v.valves = strings.Split(tunnelsStr, ", ")
		valves = append(valves, v)

	}
	fmt.Println(valves)
	return nil
}
