package main

import (
	"bytes"
	"math"
	"math/rand"
)

// 1.2.1 Write a Point2D client that takes an integer value N from the command
// line, generates N random points in the unit square, and computes the distance
// separating the closest pair of points.

type RandomTable [][]string

// NewRandomTable create a RandomTable
func NewRandomTable(n int) RandomTable {
	// Create the table
	t := make(RandomTable, n)
	for y := range t {
		t[y] = make([]string, n)
		for x := range t[y] {
			t[y][x] = "·"
		}
	}

	// Add N random points to the table
	for i := 0; i < n; i++ {
		y, x := rand.Intn(n), rand.Intn(n)
		if t[y][x] == "X" {
			i-- // Point already taken... trying a new one
			continue
		}
		t[y][x] = "X"
	}

	return t
}

func (t RandomTable) ShortestDistance() (closestDistance float64) {
	points := t.getPoints()

	minDistance := math.MaxFloat64
	for i := 0; i < len(points); i++ {
		pointA := points[i]
		for j := i + 1; j < len(points); j++ {
			pointB := points[j]
			distance := pointA.Distance(pointB)
			if distance < minDistance {
				minDistance = distance
			}
		}

	}

	return minDistance
}

func (t RandomTable) getPoints() []Point {
	points := make([]Point, 0)
	for y := range t {
		for x := range t[y] {
			if t[y][x] == "X" {
				points = append(points, Point{x, y})
			}
		}
	}

	return points
}

type Point struct {
	x, y int
}

func (p Point) Distance(otherP Point) float64 {
	b := p.x - otherP.x
	c := p.y - otherP.y

	b *= b
	c *= c

	return math.Sqrt(float64(b + c))
}

// String prints the table
func (t RandomTable) String() string {
	buf := bytes.Buffer{}

	for y := range t {
		for x := range t[y] {
			buf.WriteString(t[y][x] + " ")
		}
		buf.WriteString("\n")
	}

	return buf.String()
}
