package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("./input.txt")
	require.NoError(t, err)

	maxCalories, err := part1(f)
	require.NoError(t, err)

	require.Equal(t, 68467, maxCalories)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("./input.txt")
	require.NoError(t, err)

	maxCalories, err := part2(f)
	require.NoError(t, err)

	require.Equal(t, 203420, maxCalories)
}
