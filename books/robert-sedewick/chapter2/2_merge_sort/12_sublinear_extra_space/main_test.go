package sublinear_extra_space

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMerge(t *testing.T) {
	input := []int{7, 2, 9, 4, 5, 6, 1, 8, 3}
	MergeSort(input)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, input)
}

func TestInsertSortEachBlock(t *testing.T) {
	input := []int{3, 2, 1, 6, 5, 4, 9, 8, 7}
	insertSortEachBlock(input)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, input)
}

func TestSortByBlock(t *testing.T) {
	input := []int{4, 0, 0, 1, 0, 0, 2, 0, 0, 3, 0, 0}
	sortByBlock(input)

	require.Equal(t, []int{1, 0, 0, 2, 0, 0, 3, 0, 0, 4, 0, 0}, input)
}
