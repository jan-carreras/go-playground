package quicksort

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	sort3 "sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	input := []int{6, 2, 5, 7, 1, 3, 90, 8}
	Sort(input)
	require.Equal(t, []int{1, 2, 3, 5, 6, 7, 8, 90}, input)

	input = make([]int, rand.Intn(100000))
	for i := 0; i < len(input); i++ {
		input[i] = rand.Intn(100000)
	}

	inputCopy := input
	Sort(input)

	sort3.Ints(inputCopy)

	require.Equal(t, inputCopy, input)
}

func TestQuickSort2(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	input := []int{6, 2, 5, 7, 1, 3, 90, 8}
	Sort2(input)
	require.Equal(t, []int{1, 2, 3, 5, 6, 7, 8, 90}, input)

	input = make([]int, rand.Intn(100000))
	for i := 0; i < len(input); i++ {
		input[i] = rand.Intn(100000)
	}

	inputCopy := input
	Sort(input)

	sort3.Ints(inputCopy)

	require.Equal(t, inputCopy, input)
}
