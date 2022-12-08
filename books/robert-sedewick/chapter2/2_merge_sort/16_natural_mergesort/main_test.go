package natural_mergesort

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	input := []int{100, 200, 300, 400, 110, 120, 130, 140, 310, 320, 330, 20, 500}
	Sort(input)
	require.Equal(t, []int{20, 100, 110, 120, 130, 140, 200, 300, 310, 320, 330, 400, 500}, input)

	input = []int{100}
	Sort(input)
	require.Equal(t, []int{100}, input)

	input = []int{200, 100}
	Sort(input)
	require.Equal(t, []int{100, 200}, input)

	input = make([]int, rand.Intn(1000)+100)
	rand.Seed(time.Now().UnixMilli())
	for i := range input {
		input[i] = rand.Intn(10000)
	}

	Sort(input)
	require.True(t, isSorted(input))
}

func isSorted(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] > input[i+1] {
			return false
		}
	}
	return true
}
