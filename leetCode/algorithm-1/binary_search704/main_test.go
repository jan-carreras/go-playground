package binary_search704

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch(t *testing.T) {
	require.Equal(t, 4, Search([]int{-1, 0, 3, 5, 9, 12}, 9))
	require.Equal(t, -1, Search([]int{-1, 0, 3, 5, 9, 12}, 2))
	require.Equal(t, -1, Search([]int{}, 2))
	require.Equal(t, 0, Search([]int{2}, 2))
	require.Equal(t, 0, Search([]int{2, 3}, 2))
	require.Equal(t, 1, Search([]int{2, 3}, 3))
	require.Equal(t, 1, Search([]int{2, 3, 4}, 3))
}
