package main

import (
	"bufio"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	input := bufio.NewScanner(f)

	stacks, err := readStacks(input)
	require.NoError(t, err)

	require.NoError(t, cupaloy.SnapshotMulti("readStacks", strings.Join(stacks, ", ")))

	instructions, err := readInstructions(input)
	require.NoError(t, err)

	require.NoError(t, cupaloy.SnapshotMulti("readInstructions", instruction{}))

	processInstructions(stacks, instructions)

	require.NoError(t, cupaloy.SnapshotMulti("processInstructions", stacks))

	topCrates := getTopCrates(stacks)
	require.NoError(t, cupaloy.SnapshotMulti("solution", topCrates))
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	input := bufio.NewScanner(f)

	stacks, err := readStacks(input)
	require.NoError(t, err)

	require.NoError(t, cupaloy.SnapshotMulti("readStacks", strings.Join(stacks, ", ")))

	instructions, err := readInstructions(input)
	require.NoError(t, err)

	require.NoError(t, cupaloy.SnapshotMulti("readInstructions", instruction{}))

	processInstructionsV2(stacks, instructions)

	require.NoError(t, cupaloy.SnapshotMulti("processInstructions", stacks))

	topCrates := getTopCrates(stacks)
	require.NoError(t, cupaloy.SnapshotMulti("solution", topCrates))
}
