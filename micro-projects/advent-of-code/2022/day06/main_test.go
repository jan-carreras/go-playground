package day6

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestExamples(t *testing.T) {
	require.Equal(t, -1, startOfTransmission("abc"))
	require.Equal(t, 4, startOfTransmission("abcd"))
	require.Equal(t, 5, startOfTransmission("aabcd"))
	require.Equal(t, 6, startOfTransmission("aaabcd"))
	require.Equal(t, 5, startOfTransmission("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	require.Equal(t, 6, startOfTransmission("nppdvjthqldpwncqszvftbrmjlhg"))
	require.Equal(t, 10, startOfTransmission("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	require.Equal(t, 11, startOfTransmission("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestExamplesPart2(t *testing.T) {
	require.Equal(t, 19, startMessageMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	require.Equal(t, 23, startMessageMarker("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	require.Equal(t, 23, startMessageMarker("nppdvjthqldpwncqszvftbrmjlhg"))
	require.Equal(t, 29, startMessageMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	require.Equal(t, 26, startMessageMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestPart1And2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	transmission, err := io.ReadAll(f)
	require.NoError(t, err)

	// Part 1
	require.Equal(t, 1100, startOfTransmission(string(transmission)))

	// Part 2
	require.Equal(t, 2421, startMessageMarker(string(transmission)))
}
