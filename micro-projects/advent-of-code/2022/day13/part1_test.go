package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	err = part1(f)
	require.NoError(t, err)
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	err = part2(f)
	require.NoError(t, err)
}
