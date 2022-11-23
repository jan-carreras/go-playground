package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJosephus(t *testing.T) {
	deaths, err := Josephus(2, 7)
	require.NoError(t, err)
	require.Equal(t, []int{1, 3, 5, 0, 4, 2, 6}, deaths)

	deaths, err = Josephus(1, 7)
	require.NoError(t, err)
	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6}, deaths) // the 6th again, you lucky bastard.
}
