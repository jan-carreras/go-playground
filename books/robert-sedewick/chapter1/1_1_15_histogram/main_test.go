package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHistogram(t *testing.T) {
	require.Equal(t, []int{0, 2, 1}, Histogram([]int{1, 2, 1, 2, 3}, 3))
	require.Equal(t, []int{0, 3, 0}, Histogram([]int{1, 1, 1}, 3))
	require.Equal(t, []int{0, 0, 0}, Histogram([]int{-1, 5, 5}, 3))
}
