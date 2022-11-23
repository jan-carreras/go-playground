package ring_buffer_test

import (
	"context"
	ring_buffer "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/39_ring_buffer"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestNewWithStorage_BelowLimit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 500*time.Millisecond)
	defer cancel()

	ring := ring_buffer.NewWithStorage[string](ctx, 2)
	require.NoError(t, ring.Enqueue("hello"))
	require.NoError(t, ring.Enqueue("world"))

	time.Sleep(time.Millisecond)

	require.EqualValues(t, "[hello world]", ring.Data())

	noErr := noErrorsPlease(t)

	require.Equal(t, "hello", noErr(ring.Dequeue()))
	require.Equal(t, "world", noErr(ring.Dequeue()))

	_, err := ring.Dequeue() // We have not inserted much more
	require.ErrorContains(t, err, "context deadline exceeded")
}

func TestRingBuffer_Concurrency(t *testing.T) {
	wg := sync.WaitGroup{}
	ioOperations := 10000

	ctx, cancel := context.WithTimeout(context.TODO(), 500*time.Millisecond)
	defer cancel()

	ring := ring_buffer.NewWithStorage[int](ctx, 10)

	go func() {
		defer wg.Done()
		for i := 0; i < ioOperations; i++ {
			require.NoError(t, ring.Enqueue(i))
		}

	}()

	go func() {
		defer wg.Done()

		for i := 0; i < ioOperations; i++ {
			v, err := ring.Dequeue()
			require.NoError(t, err)
			require.Equal(t, i, v)
		}
	}()

	wg.Wait()
}

func TestRingBuffer_SlowReader(t *testing.T) {
	wg := sync.WaitGroup{}
	ioOperations := 100

	ctx, cancel := context.WithTimeout(context.TODO(), 500*time.Millisecond)
	defer cancel()

	ring := ring_buffer.NewWithStorage[int](ctx, 10)

	go func() {
		defer wg.Done()
		for i := 0; i < ioOperations; i++ {
			require.NoError(t, ring.Enqueue(i)) // Patient writer
		}

	}()

	go func() {
		defer wg.Done()

		for i := 0; i < ioOperations; i++ {
			time.Sleep(time.Millisecond) // Slow reader
			v, err := ring.Dequeue()
			require.NoError(t, err)
			require.Equal(t, i, v)
		}
	}()

	wg.Wait()
}

func noErrorsPlease(t *testing.T) func(a any, err error) any {
	return func(a any, err error) any {
		require.NoError(t, err)
		return a
	}
}
