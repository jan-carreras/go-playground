package shell_trace

import (
	"bytes"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSort_Ints(t *testing.T) {
	input := []int{1, 5, 4, 3, 8, 1}
	s := new(ShellSort[int])
	s.Sort(input)
	require.Equal(t, []int{1, 1, 3, 4, 5, 8}, input)
}

func TestSort_Strings(t *testing.T) {
	input := strings.Split("E A S Y S H E L L S O R T Q U E S T I O N", " ")
	buf := &bytes.Buffer{}
	s := new(ShellSort[string]).WithWriter(buf)

	s.Sort(input)

	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata"))
	require.NoError(t, snapshotter.SnapshotMulti("sorted", input))
	require.NoError(t, snapshotter.SnapshotMulti("output", buf.String()))
}
