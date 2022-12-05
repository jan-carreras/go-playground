package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	sum, err := part1(f)
	require.NoError(t, err)

	require.Equal(t, 7785, sum) // This is incorrect, tho
}

func TestPriority(t *testing.T) {
	require.Equal(t, 16, priority('p'))
	require.Equal(t, 38, priority('L'))
	require.Equal(t, 42, priority('P'))
	require.Equal(t, 22, priority('v'))
	require.Equal(t, 20, priority('t'))
	require.Equal(t, 19, priority('s'))
}
