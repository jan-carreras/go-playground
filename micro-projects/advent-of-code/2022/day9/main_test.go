package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	instructions, err := readInstructions(f)
	require.NoError(t, err)

	processed := processInstructions(instructions)
	require.Equal(t, 5981, processed)
}
