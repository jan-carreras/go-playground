package __1_14_lg

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestRandom(t *testing.T) {
	require.Equal(t, int(math.Log2(10)), Lg(10))
	require.Equal(t, int(math.Log2(100)), Lg(100))
	require.Equal(t, int(math.Log2(8)), Lg(8))
}
