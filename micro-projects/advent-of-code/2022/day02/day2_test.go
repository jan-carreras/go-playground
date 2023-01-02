package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	score, err := part1(f)
	require.NoError(t, err)

	require.Equal(t, 13052, score)
}

func TestDay2Part2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	score, err := part2(f)
	require.NoError(t, err)

	require.Equal(t, 13693, score)
}
