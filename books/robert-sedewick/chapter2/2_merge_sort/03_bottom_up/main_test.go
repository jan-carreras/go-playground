package bottom_up_test

import (
	"bytes"
	"github.com/bradleyjkemp/cupaloy"
	bottom_up "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter2/2_merge_sort/03_bottom_up"
	"github.com/stretchr/testify/require"
	"sort"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	output := new(bytes.Buffer)
	s := new(bottom_up.BottomUpSort[string]).WithDebug(output)

	input := strings.Split("E A S Y Q U E S T I O N", " ")
	s.Sort(input)

	inputCopy := input
	sort.Strings(inputCopy)

	require.Equal(t, inputCopy, input)

	cupaloy.SnapshotT(t, output.String())
}
