package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestDay2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	score, err := run(f)
	require.NoError(t, err)

	require.Equal(t, 13052, score)
}
