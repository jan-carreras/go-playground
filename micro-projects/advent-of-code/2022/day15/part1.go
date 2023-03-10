package day15

import (
	"fmt"
	"io"
	"strings"
)

type point struct {
	x, y int
}

func (p point) Equals(other point) bool {
	return p.x == other.x && p.y == other.y
}

func (p point) Distance(other point) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (p point) String() string {
	return fmt.Sprintf("x=%d y=%d", p.x, p.y)
}

type sensorPair struct {
	sensor point
	beacon point
}

func (s *sensorPair) Radius() int {
	return s.sensor.Distance(s.beacon)
}

func (s *sensorPair) InsideRange(other point) bool {
	return s.sensor.Distance(other) <= s.Radius()
}

// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
func part1(input io.Reader) error {
	raw, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	sensorPairs := make([]sensorPair, 0)
	for _, line := range strings.Split(string(raw), "\n") {
		pair := sensorPair{}
		_, err := fmt.Sscanf(
			line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&pair.sensor.x, &pair.sensor.y, &pair.beacon.x, &pair.beacon.y,
		)
		if err != nil {
			return err
		}
		sensorPairs = append(sensorPairs, pair)
	}

	min, max := sensorPairs[0], sensorPairs[0]
	for _, pair := range sensorPairs {
		if pair.sensor.x-pair.Radius() < min.sensor.x-max.Radius() {
			min = pair
		}
		if pair.sensor.x+pair.Radius() > max.sensor.x+max.Radius() {
			max = pair
		}
	}

	limit := 4000000

END:
	for y := 0; y <= limit; y++ {
		/*line := make(map[int]bool, limit)
		for _, pair := range sensorPairs {
			pair.
		}*/

		for x := 0; x <= limit; x++ {
			lineScanner := point{x: x, y: y}
			inRange := false
			for _, pair := range sensorPairs {
				if pair.InsideRange(lineScanner) {
					inRange = true
					break
				}
			}
			if !inRange {
				fmt.Println("Found the answer", limit*x+y, x, y)
				break END
			}
		}
	}

	// 4343576 too low // 4453516 too low // Part 1: Correct 5073496
	//fmt.Println("In the row where y=2000000, how many positions cannot contain a beacon?", count-1)

	return nil
}
