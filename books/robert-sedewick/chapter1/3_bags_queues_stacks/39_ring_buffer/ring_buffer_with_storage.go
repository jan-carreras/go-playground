package ring_buffer

import (
	"context"
	"fmt"
	"time"
)

type RingBufferStorage[T any] struct {
	// Can be cancelled to signal to writers/readers that they don't have to listen anymore
	ctx    context.Context
	cancel context.CancelFunc

	// Buffered channel
	data        []T
	length      int
	read, write int

	incoming chan T
	outgoing chan T
}

func NewWithStorage[T any](ctx context.Context, capacity int) *RingBufferStorage[T] {
	ctx, cancel := context.WithCancel(ctx)

	rb := &RingBufferStorage[T]{
		ctx:    ctx,
		cancel: cancel,

		data: make([]T, capacity),

		incoming: make(chan T, 0), // Unbuffered
		outgoing: make(chan T, 0), // Unbuffered
	}
	go rb.in()
	go rb.out()

	return rb
}

func (r *RingBufferStorage[T]) closeAll() {
	//close(r.incoming)
	//close(r.outgoing)
	r.cancel()
	r.data = nil
}

func (r *RingBufferStorage[T]) in() {
	for {
		for r.length == cap(r.data) {
			// TODO: How can we prevent this? 🤔More channels? :facepalm?
			time.Sleep(1 * time.Millisecond) // Waiting for someone to read...
		}

		select {
		case v := <-r.incoming:
			r.data[r.write] = v
			r.write = (r.write + 1) % cap(r.data)
			r.length++
		case <-r.ctx.Done():
			r.closeAll()
			return
		}

	}
}

func (r *RingBufferStorage[T]) out() {
	for {
		for r.length == 0 {
			time.Sleep(1 * time.Millisecond)
		}

		v := r.data[r.read]
        r.read = (r.read + 1) % cap(r.data)
        r.length--

		select {
		case r.outgoing <- v: // Blocking until someone reads
		case <-r.ctx.Done():
			r.closeAll()
			return
		}
	}
}

func (r *RingBufferStorage[T]) Enqueue(v T) error {
	select {
	case r.incoming <- v:
		return nil
	case <-r.ctx.Done():
		return r.ctx.Err()
	}
}

func (r *RingBufferStorage[T]) Dequeue() (T, error) {
	select {
	case rsp := <-r.outgoing:
		return rsp, nil
	case <-r.ctx.Done():
		return *new(T), r.ctx.Err()
	}
}

// For testing purposes only
func (r *RingBufferStorage[T]) Data() string {
	return fmt.Sprintf("%v", r.data)
}
