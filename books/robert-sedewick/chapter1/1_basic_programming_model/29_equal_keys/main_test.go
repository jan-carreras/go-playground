package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSmallerThan(t *testing.T) {
	bs := NewBinarySearch([]int{1, 1, 2, 3, 3})
	require.Equal(t, 3, bs.SmallerThan(3))
	require.Equal(t, -1, bs.SmallerThan(10))
	require.Equal(t, 0, bs.SmallerThan(1))
	require.Equal(t, 2, bs.SmallerThan(2))

	require.Equal(t, 2, bs.Count(3))
	require.Equal(t, -1, bs.Count(10))
	require.Equal(t, 2, bs.Count(1))
	require.Equal(t, 1, bs.Count(2))

}
