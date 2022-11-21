package __1_22_traces

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRank(t *testing.T) {
	require.Equal(t, 1, Rank([]int{1, 2, 3, 4, 5, 6}, 1))
	require.Equal(t, 5, Rank([]int{1, 2, 3, 4, 5, 6}, 5))
	require.Equal(t, -1, Rank([]int{1, 2, 3, 4, 5, 6}, 7))

	/**
	STDOUT GENERATED:

	lo=0 hi=5
		⮑ lo=0 hi=1
	lo=0 hi=5
		⮑ lo=3 hi=5
	lo=0 hi=5
		⮑ lo=3 hi=5
			⮑ lo=5 hi=5
				⮑ lo=6 hi=5
	*/

}
