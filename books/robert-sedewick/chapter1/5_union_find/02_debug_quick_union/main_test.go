package debug_quick_union

import (
	"github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/5_union_find/union_find/union_find"
	"github.com/stretchr/testify/require"
	"testing"
)

// 1.5.2 Do Exercise1.5.1, but use quick-union.In addition,draw the forest of
// trees represented by the id[] array after each input pair is processed.

func Test(t *testing.T) {
	qf := union_find.NewQuickUnion(10)
	// 9-0 3-4 5-8 7-2 2-1 5-7 0-3 4-2
	qf.Union(9, 0)
	qf.Union(3, 4)
	qf.Union(5, 8)
	qf.Union(7, 2)
	qf.Union(2, 1)
	qf.Union(5, 7)
	qf.Union(0, 3)
	qf.Union(4, 2)

	require.Equal(t, []int{4, 1, 1, 4, 1, 8, 6, 2, 1, 0}, qf.Debug().ID)
	require.Equal(t, 10, qf.Debug().IDAccesses)

	// TODO: Draw the forest of trees
}
