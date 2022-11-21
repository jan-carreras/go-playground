package insert_sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []int{1, 5, 7, 2, 3, 0}
	InsertionSort2(input)
	require.Equal(t, []int{0, 1, 2, 3, 5, 7}, input)
}
