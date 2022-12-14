package pq_test

import (
	"fmt"
	pq "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter2/4_priority_queues/00_pq"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxPQ(t *testing.T) {
	maxPQ := pq.NewMaxPQ[string](4)

	for i := 0; i < 10; i++ {
		maxPQ.Insert(pq.NewValue(fmt.Sprintf("whatever %d", i), i))
	}

	for !maxPQ.IsEmpty() {
		v, ok := maxPQ.DeleteMax()
		require.True(t, ok)
		require.Equal(t, maxPQ.Size(), v.Priority())
	}
}
