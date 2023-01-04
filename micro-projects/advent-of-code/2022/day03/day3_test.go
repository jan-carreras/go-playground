package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	sum, err := Part1(f)
	require.NoError(t, err)

	require.Equal(t, 7785, sum)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	sum, err := Part2(f)
	require.NoError(t, err)

	require.Equal(t, 2633, sum)

}

func TestPriority(t *testing.T) {
	priority := func(r rune) int {
		p, err := priority(r)
		require.NoError(t, err)
		return p
	}

	require.Equal(t, 16, priority('p'))
	require.Equal(t, 38, priority('L'))
	require.Equal(t, 42, priority('P'))
	require.Equal(t, 22, priority('v'))
	require.Equal(t, 20, priority('t'))
	require.Equal(t, 19, priority('s'))
	require.Equal(t, 18, priority('r'))
}

func TestUnions(t *testing.T) {
	m1 := map[rune]bool{'a': true, 'b': true, 'c': true}
	m2 := map[rune]bool{'a': true, 'b': true, 'd': true}
	m3 := map[rune]bool{'a': true, 'e': true, 'd': true}

	require.Equal(t, map[rune]bool{'a': true}, unions(m1, m2, m3))
}

func TestUnion(t *testing.T) {
	m1 := map[rune]bool{'a': true, 'b': true, 'c': true}
	m2 := map[rune]bool{'a': true, 'b': true, 'd': true}
	m3 := map[rune]bool{'a': true, 'e': true, 'd': true}

	// Empty input
	require.Equal(t, map[rune]bool{}, unions())

	// Only one element
	require.Equal(t, m1, unions(m1))

	// Two elements
	require.Equal(t, map[rune]bool{'a': true, 'b': true}, unions(m1, m2))

	require.Equal(t, map[rune]bool{'a': true}, unions(m1, m2, m3))
}
