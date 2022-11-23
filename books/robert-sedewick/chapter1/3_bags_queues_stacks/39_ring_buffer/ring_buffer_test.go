package ring_buffer

import (
	"context"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestRingBuffer_Enqueue(t *testing.T) {
	ctx, cancellation := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancellation()

	r := New[string](ctx, 3)

	require.NoError(t, r.Enqueue("hello"))
	require.NoError(t, r.Enqueue("beautiful"))
	require.NoError(t, r.Enqueue("world"))

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		require.NoError(t, r.Enqueue("lorem"))
		require.NoError(t, r.Enqueue("ipsum"))
		require.NoError(t, r.Enqueue("amet"))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		check := func(v string, err error) string {
			require.NoError(t, err)
			return v
		}
		require.Equal(t, "hello", check(r.Dequeue()))
		require.Equal(t, "beautiful", check(r.Dequeue()))
		require.Equal(t, "world", check(r.Dequeue()))

		require.Equal(t, "lorem", check(r.Dequeue()))
		require.Equal(t, "ipsum", check(r.Dequeue()))
		require.Equal(t, "amet", check(r.Dequeue()))
	}()

	wg.Wait()
}
