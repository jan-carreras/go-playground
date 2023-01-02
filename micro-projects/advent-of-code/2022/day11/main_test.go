package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test(t *testing.T) {

	f, err := os.Open("input.txt")
	require.NoError(t, err)

	err = run(f)
	require.NoError(t, err)

	// First try:

	// That's not the right answer; your answer is too low. If you're stuck, make
	// sure you're using the full input data; there are also some general tips on the
	// about page, or you can ask for hints on the subreddit. Please wait one minute
	// before trying again. (You guessed 108216.) [Return to Day 11]

	// Second try:

	// PART 1. What is the level of monkey business after 20 rounds of stuff-slinging simian shenanigans? 110220

	// PART 2. What is the level of monkey business after 20 rounds of stuff-slinging simian shenanigans? 19457438264

	/**
	I've cheated on part 2. No idea how to approach it, and I still don't grasp
	the mathematical underlying logic :/
	*/

}
