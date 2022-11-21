package first_bad_version

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	expected = 4
	require.Equal(t, expected, FirstBadVersion(100))

	expected = 2
	require.Equal(t, expected, FirstBadVersion(3))

	expected = 3
	require.Equal(t, expected, FirstBadVersion(3))

	expected = 1
	require.Equal(t, expected, FirstBadVersion(1))

	expected = 1
	require.Equal(t, expected, FirstBadVersion(3))

	expected = 1
	require.Equal(t, expected, FirstBadVersion(100))

	expected = 100
	require.Equal(t, expected, FirstBadVersion(100))

	expected = 99
	require.Equal(t, expected, FirstBadVersion(100))
}
