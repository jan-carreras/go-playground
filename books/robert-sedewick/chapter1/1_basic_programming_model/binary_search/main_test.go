package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestSearch(t *testing.T) {
	require.Equal(t, -1, Search([]int{}, 1))
	require.Equal(t, 0, Search([]int{1}, 1))
	require.Equal(t, -1, Search([]int{1}, 2))
	require.Equal(t, 2, Search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
	require.Equal(t, 9, Search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	require.Equal(t, 0, Search([]int{-5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -5))
}

func Test_MinIntCannotBeAbsed(t *testing.T) {
	n := math.MinInt
	if n < 0 {
		n = -n
	}

	require.Equal(t, math.MinInt, n)
}

func Test_Operations(t *testing.T) {
	fmt.Println(2.0e-6 * 100000000.1)
	fmt.Println(0.000002 * 100000000.1)
}

func Test_Unsigned(t *testing.T) {
	var i uint
	fmt.Println(i - 1)
}
