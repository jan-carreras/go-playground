package farthest_pair

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFarthestPair(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 5.3, 1.1}
	require.InDelta(t, 4.3, FarthestPair(input), 0.01)
}
