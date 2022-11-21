package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsCircular(t *testing.T) {
	require.True(t, IsCircular("ACTGACG", "ACTGACG"))
	require.True(t, IsCircular("ACTGACG", "TGACGAC"))
	require.False(t, IsCircular("ACTGACT", "TGACGAC"))
	require.False(t, IsCircular("ACTGACTTT", "TGACGAC"))
	require.False(t, IsCircular("ACTGACTTT", ""))
}
