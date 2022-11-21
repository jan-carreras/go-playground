package main

import (
	"fmt"
	"log"
	"os"
)

// 1.2.2 Write an Interval1D client that takes an int value N as command-line
// argument, reads N intervals (each defined by a pair of double values) from
// standard input, and prints all pairs that intersect.

/**
INPUT
5
1.0 2.0
2.0 3.0
1.0 3.0
1.0 5.0
4.0 5.0
*/

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

type Range struct {
	start, end float64
}

func (r Range) Overlaps(otherRange Range) bool {
	return r.end >= otherRange.start && r.start <= otherRange.end
}

type OverlappingRange struct {
	r1 Range
	r2 Range
}

func run() error {
	var linesOfPairs int
	_, err := fmt.Fscanf(os.Stdin, "%d", &linesOfPairs)
	if err != nil {
		return err
	}

	ranges := make([]Range, 0, linesOfPairs)
	for i := 0; i < linesOfPairs; i++ {
		var start, end float64
		_, err := fmt.Fscanf(os.Stdin, "%f %f", &start, &end)
		if err != nil {
			return err
		}

		ranges = append(ranges, Range{start, end})
	}

	for _, overlappingRange := range Intersecting(ranges) {
		fmt.Println("Overlapping range:", overlappingRange)
	}

	return nil
}

func Intersecting(ranges []Range) []OverlappingRange {
	overlaps := make([]OverlappingRange, 0)
	for i := 0; i < len(ranges); i++ {
		rangeI := ranges[i]
		for j := i + 1; j < len(ranges); j++ {
			rangeJ := ranges[j]
			if rangeI.Overlaps(rangeJ) {
				overlaps = append(overlaps, OverlappingRange{rangeI, rangeJ})
			}
		}

	}

	return overlaps
}
