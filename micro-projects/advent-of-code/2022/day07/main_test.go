package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSimpleInput(t *testing.T) {
	f, err := os.Open("input-short.txt")
	require.NoError(t, err)

	fs, err := parseInput(f)
	require.NoError(t, err)

	output := make(map[string]int, 0)
	findSizeLimit(fs.root, 100000, output)
	sum := 0
	for _, size := range output {
		sum += size
	}
	fmt.Println(output)
	require.Equal(t, 95437, sum) // 1575920 is wrong // That's not the right answer; your answer is too low.
}

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	fs, err := parseInput(f)
	require.NoError(t, err)

	output := make(map[string]int, 0)
	findSizeLimit(fs.root, 100000, output)
	sum := 0
	for _, size := range output {
		sum += size
	}
	fmt.Println(output)
	require.Equal(t, 1778099, sum) // 1575920 is wrong // That's not the right answer; your answer is too low.
}

func TestPart2(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	fs, err := parseInput(f)
	require.NoError(t, err)

	usedSpace := usedSize(fs.root)
	freeSpace := DiskSize - usedSpace
	needToFree := UpdateSize - freeSpace

	fmt.Println("Used Space", usedSpace)
	fmt.Println("Free Space", DiskSize-usedSpace)
	fmt.Println("Needed space", UpdateSize)
	fmt.Println("Needed to free", UpdateSize-freeSpace)
	fmt.Println()

	min := DiskSize
	WalkDirectories(fs.root, func(f *File, size int) {
		if size < needToFree {
			return // Deleting this file does not help us at all
		}

		if size < min {
			min = size
			fmt.Println(f.FullPath(), min, size, size-needToFree)
		}
	})

	fmt.Println("The file with the minimum size has", min)
	require.Equal(t, 1623571, min)
}
