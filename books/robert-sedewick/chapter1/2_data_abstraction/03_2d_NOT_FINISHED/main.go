package main

import (
	"flag"
	"fmt"
	"github.com/fogleman/gg"
	"log"
	"math/rand"
	"time"
)

/**
1.2.3

✅ Write an Interval2D client that takes command line arguments N,min,and max

✅ and generates N random 2D intervals whose width and height are uniformly
distributed between min and max in the unit square.

✅ Draw them on StdDraw and

❗️ print the number of pairs of intervals that intersect

❗️ and the number of intervals that are contained in one another.
*/

const boardSize = 1000

func main() {
	rand.Seed(time.Now().UnixMilli())

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	// Read arguments
	var n int
	var min, max float64
	if err := parseArgs(&n, &min, &max); err != nil {
		return err
	}

	points := createImage(min, max, n)

	if err := drawPoints(points); err != nil {
		return err
	}

	return nil
}

func drawPoints(intervals []Interval) error {
	dc := gg.NewContext(boardSize, boardSize)

	for _, interval := range intervals {
		dc.DrawRectangle(interval.Point.X, interval.Point.Y, interval.Width, interval.Height)
		dc.SetRGB(rand.Float64()*255, 0, 0)
		dc.Fill()
	}

	// TODO: Save it in the exercise directory, instead of the root path
	if err := dc.SavePNG("out.png"); err != nil {
		return err
	}

	return nil
}

// Interval describes an interval in 2 dimensions. Quite honestly I don't know why the fuck I haven't named this
// as box :facepalm:
type Interval struct {
	Point  gg.Point
	Width  float64
	Height float64
}

func (i Interval) Intersects(otherInterval Interval) bool {

	return true
}

func createImage(min float64, max float64, n int) []Interval {
	// Make intervals
	randomFloat := func() float64 { return rand.Float64() * (boardSize - max) }
	randomBoundedFloat := func() float64 { return min + rand.Float64()*(max-min) }

	intervals := make([]Interval, 0, n)
	for i := 0; i < n; i++ {
		intervals = append(intervals, Interval{
			Point:  gg.Point{X: randomFloat(), Y: randomFloat()},
			Width:  randomBoundedFloat(),
			Height: randomBoundedFloat(),
		})
	}
	return intervals
}

func parseArgs(n *int, min, max *float64) error {
	flag.IntVar(n, "N", 0, "an int")
	flag.Float64Var(min, "min", 0, "an int")
	flag.Float64Var(max, "max", 0, "an int")
	flag.Parse()

	if *max <= *min {
		return fmt.Errorf("max must be greater than min: max=%f min=%f", *max, *min)
	}

	if *n <= 0 {
		return fmt.Errorf("N must be greater than 0: N=%d", n)
	}

	return nil
}
