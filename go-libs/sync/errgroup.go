package errgroup

import (
	"context"
	"sync"
)

type token struct{}

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid, has no limit on the number of active goroutines,
// and does not cancel on error.
type Group struct {
	cancellation context.CancelFunc

	wg sync.WaitGroup

	tokens chan token

	err     error
	errOnce sync.Once
}

func (g *Group) done() {
	if g.tokens != nil {
		<-g.tokens
	}

	g.wg.Done()
}

func (g *Group) cancel() {
	if g.cancellation != nil {
		g.cancellation()
	}
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first. 🟩
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancellation := context.WithCancel(ctx)
	return &Group{cancellation: cancellation}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	defer g.cancel()
	g.wg.Wait()
	return g.err
}

// Go calls the given function in a new goroutine.
// It blocks until the new goroutine can be added without the number of
// active goroutines in the group exceeding the configured limit.
//
// The first call to return a non-nil error cancels the group's context, if the
// group was created by calling WithContext. The error will be returned by Wait.
func (g *Group) Go(f func() error) {

	// DO: Acquire lock
	if g.tokens != nil {
		g.tokens <- token{}
	}

	g.wg.Add(1)
	go g.run(f)

}

func (g *Group) run(f func() error) {
	defer g.done()

	err := f()
	if err != nil {
		g.errOnce.Do(func() {
			g.err = err
			g.cancel()
		})
	}
}

// TryGo calls the given function in a new goroutine only if the number of
// active goroutines in the group is currently below the configured limit.
//
// The return value reports whether the goroutine was started.
func (g *Group) TryGo(f func() error) bool {

	if g.tokens != nil {
		select {
		case g.tokens <- token{}:
			// We acquired a token, we can run the function
		default:
			return false
		}
	}

	g.wg.Add(1)
	go g.run(f)

	return true
}

// SetLimit limits the number of active goroutines in this group to at most n.
// A negative value indicates no limit.
//
// Any subsequent call to the Go method will block until it can add an active
// goroutine without exceeding the configured limit.
//
// The limit must not be modified while any goroutines in the group are active.
func (g *Group) SetLimit(n int) {
	if n < 0 {
		g.tokens = nil
		return
	}

	if len(g.tokens) != 0 {
		panic("limit must not be modified while any goroutines in the group are active")
	}

	g.tokens = make(chan token, n)
}
