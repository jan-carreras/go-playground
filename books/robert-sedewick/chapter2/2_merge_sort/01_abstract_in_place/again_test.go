package abstract_in_place_test

import (
	abstract_in_place "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter2/2_merge_sort/01_abstract_in_place"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestAgain(t *testing.T) {
	s := new(abstract_in_place.Sort2[string])
	input := []string{"A", "B"}
	s.Sort(input)
	require.Equal(t, "AB", strings.Join(input, ""))

	input = []string{"B", "A"}
	s.Sort(input)
	require.Equal(t, "AB", strings.Join(input, ""))

	input = []string{"A", "B", "C", "D"}
	s.Sort(input)
	require.Equal(t, "ABCD", strings.Join(input, ""))

	input = []string{"C", "B", "A", "D"}
	s.Sort(input)
	require.Equal(t, "ACBD", strings.Join(input, ""))

	input = []string{"A", "B", "C", "D", "E"}
	s.Sort(input)
	require.Equal(t, "ABCDE", strings.Join(input, ""))

	input = []string{"B", "E", "A", "C", "D"}
	s.Sort(input)
	require.Equal(t, "ABCDE", strings.Join(input, ""))

	input = strings.Split("B A", " ")
	s.Sort(input)
	require.Equal(t, "AB", strings.Join(input, ""))

	input = strings.Split("A B C D", " ")
	s.Sort(input)
	require.Equal(t, "ABCD", strings.Join(input, ""))
	input = strings.Split("D C B A", " ")
	s.Sort(input)
	require.Equal(t, "BADC", strings.Join(input, ""))

	input = strings.Split("E E G M R A C E R T", " ")
	s.Sort(input)
	require.Equal(t, "ACEEEGMRRT", strings.Join(input, ""))
}
