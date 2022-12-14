package insertion_trace

import (
	"bytes"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSort_Ints(t *testing.T) {
	input := []int{1, 5, 4, 3, 8, 1}
	s := new(InsertSort[int])
	s.Sort(input)
	require.Equal(t, []int{1, 1, 3, 4, 5, 8}, input)
}

func TestSort_Strings(t *testing.T) {
	input := []string{"E", "A", "S", "Y", "Q", "U", "E", "S", "T", "I", "O", "N"}
	buf := &bytes.Buffer{}
	s := new(InsertSort[string]).WithWriter(buf)

	s.Sort(input)

	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("testdata"))
	require.NoError(t, snapshotter.SnapshotMulti("sorted", input))
	require.NoError(t, snapshotter.SnapshotMulti("output", buf.String()))
}
