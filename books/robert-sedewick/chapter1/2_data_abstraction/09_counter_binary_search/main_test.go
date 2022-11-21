package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRank(t *testing.T) {
	var c Counter
	require.Equal(t, 1, Rank(&c, []int{1, 2, 3, 4}, 2))
	require.Equal(t, 2, c.Value())
}

func TestCounter(t *testing.T) {
	var c Counter
	c.Increment()
	c.Increment()
	require.Equal(t, 2, c.Value())

	c.IncrementN(2)
	require.Equal(t, 4, c.Value())
}
