package bitronic_search

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	input := []int{4, 10, 12, 13, 14, 15, 9, 7, 5, 3, 1}
	require.True(t, BitronicSearch(input, 4))

	input = []int{4, 10, 12, 13, 14, 15, 9, 7, 5, 3, 1}
	require.True(t, BitronicSearch(input, 3))

	// Create a lot of Bitronic slices and check if they work, searching for a random key
	for i := 5; i < 100; i++ {
		createTheDamnBitronicThingie(i)
		rndKey := input[rand.Intn(len(input)-1)]
		require.True(t, BitronicSearch(input, rndKey))
		require.False(t, BitronicSearch(input, -1))  // Too small
		require.False(t, BitronicSearch(input, 1e6)) // Too big
	}
}

func createTheDamnBitronicThingie(N int) []int {
	input := make([]int, 0, N*2)
	for i := 0; i < N; i += 2 {
		input = append(input, i)
	}

	M := N
	if N%2 == 0 {
		M--
	}
	for i := M; i > 0; i -= 2 {
		input = append(input, i)
	}
	return input
}

func TestSorting(t *testing.T) {
	input := []int{4, 10, 12, 13, 14, 15, 9, 7, 5, 3, 1}
	BitronicSearchSorting(input, 9)
}

func TestBitronicSearch(t *testing.T) {
	input := []int{4, 10, 12, 13, 14, 15, 9, 7, 5, 3, 1}
	require.True(t, BitronicSearchBruteforce(input, 9))
	require.False(t, BitronicSearchBruteforce(input, 2))
}
