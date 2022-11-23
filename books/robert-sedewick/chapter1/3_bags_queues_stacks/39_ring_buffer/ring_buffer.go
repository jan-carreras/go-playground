package ring_buffer

import (
	"context"
)

// Ring buffer. A ring buffer, or circular queue, is a FIFO data structure of a
// fixed capacity N. It is useful for transferring data between asynchronous
// processes or for storing log files.
//
// When the buffer is empty, the consumer waits until data is deposited; when the
// buffer is full, the producer waits to deposit data. Develop an API for a
// RingBuffer and an implementation that uses an array representation (with
// circular wrap-around).

// 🤔 In Go this is basically a buffered channel

type RingBuffer[T any] struct {
	// Can be cancelled to signal to writers/readers that they don't have to listen anymore
	ctx context.Context

	// Buffered channel
	data chan T
}

func New[T any](ctx context.Context, capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		ctx:  ctx,
		data: make(chan T, capacity),
	}
}

func (r *RingBuffer[T]) Enqueue(v T) error {
	select {
	case r.data <- v:
		return nil
	case <-r.ctx.Done():
		// TODO: Close r.data channel
		return r.ctx.Err()
	}
}

func (r *RingBuffer[T]) Dequeue() (T, error) {
	select {
	case rsp := <-r.data:
		return rsp, nil
	case <-r.ctx.Done():
		// TODO: Close r.data channel
		return *new(T), r.ctx.Err()
	}
}
