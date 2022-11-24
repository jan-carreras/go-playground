package close_pair

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClosePair(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 5.3, 1.1}
	require.InDelta(t, 0.1, ClosePair(input), 0.01)
}
