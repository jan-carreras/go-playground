package bottom_up_mergesort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBottomUpMergeSortWithQueues(t *testing.T) {
	input := []int{5, 4, 8, 1, 0}
	BottomUpMergeSortWithQueues[int](input)
	require.Equal(t, []int{0, 1, 4, 5, 8}, input)
}
