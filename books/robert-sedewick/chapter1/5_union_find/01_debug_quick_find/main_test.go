package main

import (
	"github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/5_union_find/union_find/union_find"
	"github.com/stretchr/testify/require"
	"testing"
)

// 1.5.1 Show the contents of the id[] array and the number of times the array is
// accessed for each input pair when you use quick-find for the sequence 9-0 3-4
// 5-8 7-2 2-1 5-7 0-3 4-2.

func Test(t *testing.T) {
	qf := union_find.NewQuickFind(10)
	// 9-0 3-4 5-8 7-2 2-1 5-7 0-3 4-2
	qf.Union(9, 0)
	qf.Union(3, 4)
	qf.Union(5, 8)
	qf.Union(7, 2)
	qf.Union(2, 1)
	qf.Union(5, 7)
	qf.Union(0, 3)
	qf.Union(4, 2)

	require.Equal(t, []int{1, 1, 1, 1, 1, 1, 6, 1, 1, 1}, qf.Debug().ID)
	require.Equal(t, 96, qf.Debug().IDAccesses)
}
