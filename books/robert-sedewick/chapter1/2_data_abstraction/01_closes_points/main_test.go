package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestTable(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	table := NewRandomTable(10)
	fmt.Println(table.String())
}

func TestShortestDistance(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	table := NewRandomTable(10)
	fmt.Println(table)
	fmt.Println(table.ShortestDistance())

	/**
	OUTPUT Can look like:


	· X · · · · · X · ·
	· · · · X · · · · ·
	· · X · · · · · · ·
	· · · · · · · · · ·
	· · · X · · · · · ·
	· · · · · · · · · ·
	· · · · · · · · · X
	· · · · · · X · · ·
	· X · · · X · · · ·
	· · · · · · · X · ·

	1.4142135623730951

	*/

}

func TestDistance(t *testing.T) {
	a := Point{x: 1, y: 1}
	b := Point{x: 4, y: 3}
	require.InDelta(t, 3.6, a.Distance(b), 0.01)
}
