package main

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
	list "leetcode/problems/lists"
	"testing"
)

func TestRemove(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		l := list.NewList([]int{1, 2, 3, 4, 5})
		l = removeNthFromEnd(l, 2)
		assert.NoError(t, cupaloy.SnapshotMulti("remove 2", l.String()))
	})

	t.Run("remove 1", func(t *testing.T) {
		l := list.NewList([]int{1, 2})
		l = removeNthFromEnd(l, 1)
		assert.NoError(t, cupaloy.SnapshotMulti("remove 1 - short list", l.String()))
	})

	t.Run("0 index is invalid - do nothing", func(t *testing.T) {
		l := list.NewList([]int{1, 2, 3, 4, 5})
		l = removeNthFromEnd(l, 0)
		assert.NoError(t, cupaloy.SnapshotMulti("remove 0", l.String()))
	})

	t.Run("remove last", func(t *testing.T) {
		l := list.NewList([]int{1, 2, 3, 4, 5})
		l = removeNthFromEnd(l, 1)
		assert.NoError(t, cupaloy.SnapshotMulti("remove 1", l.String()))
	})

	t.Run("remove first", func(t *testing.T) {
		// Remove first
		l := list.NewList([]int{1, 2, 3, 4, 5})
		l = removeNthFromEnd(l, 5)
		assert.NoError(t, cupaloy.SnapshotMulti("remove 5", l.String()))
	})

}
