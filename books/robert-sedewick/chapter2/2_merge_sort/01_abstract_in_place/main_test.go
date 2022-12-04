package abstract_in_place_test

import (
	"bytes"
	"fmt"
	abstract_in_place "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter2/2_merge_sort/01_abstract_in_place"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestInPlace(t *testing.T) {
	s := new(abstract_in_place.Sort[string])

	input := strings.Split("A B", " ")
	s.InPlace(input)
	require.Equal(t, "AB", strings.Join(input, ""))

	input = strings.Split("B A", " ")
	s.InPlace(input)
	require.Equal(t, "AB", strings.Join(input, ""))

	input = strings.Split("A B C D", " ")
	s.InPlace(input)
	require.Equal(t, "ABCD", strings.Join(input, ""))
	input = strings.Split("D C B A", " ")
	s.InPlace(input)
	require.Equal(t, "BADC", strings.Join(input, ""))

	input = strings.Split("E E G M R A C E R T", " ")
	s.InPlace(input)
	require.Equal(t, "ACEEEGMRRT", strings.Join(input, ""))
}

func TestInPlace2(t *testing.T) {
	b := new(bytes.Buffer)
	s := new(abstract_in_place.Sort[string]).WithDebug(b)

	input := strings.Split("E E G M R A C E R T", " ")
	s.InPlace(input)
	require.Equal(t, "ACEEEGMRRT", strings.Join(input, ""))

	fmt.Println(b.String())
}
