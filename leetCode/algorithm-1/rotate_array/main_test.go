package rotate_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRotate(t *testing.T) {
	var input []int

	input = []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(input, 0)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, input)

	input = []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(input, 7)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, input)

	input = []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(input, 1)
	require.Equal(t, []int{7, 1, 2, 3, 4, 5, 6}, input)

	input = []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(input, 2)
	require.Equal(t, []int{6, 7, 1, 2, 3, 4, 5}, input)

	input = []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(input, 3)
	require.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, input)

	input = []int{-1, -100, 3, 99}
	Rotate(input, 2)
	require.Equal(t, []int{3, 99, -1, -100}, input)
}
