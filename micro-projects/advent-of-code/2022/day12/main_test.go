package main

import (
	"github.com/fatih/color"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test(t *testing.T) {
	color.NoColor = false
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	err = run(f)
	require.NoError(t, err)

	// Part 1 solution
	/**
	We have found a route!!
	Number of steps: 423
	*/
}
