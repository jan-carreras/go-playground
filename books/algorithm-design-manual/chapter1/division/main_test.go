package division

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestDiv(t *testing.T) {
	require.Equal(t, 10/3, Div(10, 3))
	require.Equal(t, 10/2, Div(10, 2))
	require.Equal(t, 100/2, Div(100, 2))
	require.Equal(t, 100/3, Div(100, 3))
}

func TestFastDiv(t *testing.T) {
	require.Equal(t, 10/3, FasterDiv(10, 3))
	require.Equal(t, 10/2, FasterDiv(10, 2))
	require.Equal(t, 100/2, FasterDiv(100, 2))
	require.Equal(t, 100/3, FasterDiv(100, 3))
	require.Equal(t, -10/3, FasterDiv(-10, 3))
	require.Equal(t, -1/-1, FasterDiv(-1, -1))
}

func TestFD(t *testing.T) {
	require.Equal(t, 10/3, FD(10, 3))
	require.Equal(t, 10/2, FD(10, 2))
	require.Equal(t, 100/2, FD(100, 2))
	require.Equal(t, 100/3, FD(100, 3))
	require.Equal(t, -10/3, FD(-10, 3))
	require.Equal(t, -1/-1, FD(-1, -1))
	require.Equal(t, math.MaxInt32, FD(math.MaxInt32+10, 1))
	require.Equal(t, math.MinInt32, FD(math.MinInt32-10, 1))
}
