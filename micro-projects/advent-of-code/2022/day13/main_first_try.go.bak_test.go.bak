package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input-tiny.txt")
	require.NoError(t, err)

	pairs, err := readPairs(f)
	require.NoError(t, err)

	expected := []bool{true, true, false, true, false, true, false, false}
	for i, pair := range pairs {
		isValid := pair.IsValid()
		println("pair", i, isValid)
		require.Equal(t, expected[i], isValid, i)
	}

}

func TestRunPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	err = runPart1(f)
	require.NoError(t, err)

	// First try: What is the sum of the indices of those pairs? 5137

	// That's not the right answer; your answer is too low. If you're stuck, make
	// sure you're using the full input data; there are also some general tips on the
	// about page, or you can ask for hints on the subreddit. Please wait one minute
	// before trying again. (You guessed 5137.)
}
