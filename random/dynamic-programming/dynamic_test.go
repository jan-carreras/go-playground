package dynamic_programming

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDynamic(t *testing.T) {
	items := []Item{
		{
			name:   "Water",
			weight: 3,
			score:  10,
		},
		{
			name:   "Book",
			weight: 1,
			score:  3,
		},
		{
			name:   "Food",
			weight: 2,
			score:  9,
		},
		{
			name:   "Jacket",
			weight: 2,
			score:  5,
		},
		{
			name:   "Camera",
			weight: 1,
			score:  6,
		},
	}

	result := itemsToTake(items, 6)

	require.Equal(t, 20, result.score)
	require.Equal(t, "FJC", result.name)

}
