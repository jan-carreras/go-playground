package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIntersecting(t *testing.T) {
	r1 := Range{start: 1, end: 2}
	r2 := Range{start: 2, end: 3}
	r3 := Range{start: 4, end: 5}

	ranges := []Range{r1, r2, r3}

	require.Equal(t, []OverlappingRange{{r1, r2}}, Intersecting(ranges))
}
