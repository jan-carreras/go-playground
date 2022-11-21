package squares_sorted

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	require.Equal(t, []int{0, 1, 9, 16, 100}, SortedSquares([]int{-4, -1, 0, 3, 10}))
}
