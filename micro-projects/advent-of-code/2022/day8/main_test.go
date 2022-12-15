package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	forest, err := readForest(f)
	require.NoError(t, err)

	visible := countVisibleTrees(forest)
	require.Equal(t, 1849, visible)
}

func TestPart2Small(t *testing.T) {
	f, err := os.Open("input-small.txt")
	require.NoError(t, err)

	forest, err := readForest(f)
	require.NoError(t, err)

	score := scenicScore(forest, 1, 2)
	require.Equal(t, 4, score)

	score = scenicScore(forest, 3, 2)
	require.Equal(t, 8, score)

	maxScore := maxScenicScore(forest)
	require.Equal(t, 16, maxScore)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	forest, err := readForest(f)
	require.NoError(t, err)

	maxScore := maxScenicScore(forest)
	require.Equal(t, 201600, maxScore)
}
