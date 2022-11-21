// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errgroup provides synchronization, error propagation, and Context
// cancelation for groups of goroutines working on subtasks of a common task. 🟩
package errgroup

import (
	"context"
	"sync"
)

type token struct{}

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task. 🟩
//
// A zero Group is valid, has no limit on the number of active goroutines,
// and does not cancel on error. 🟨
type Group struct {
	cancel context.CancelFunc

	wg sync.WaitGroup

	tokens chan token

	firstError     error
	firstErrorOnce sync.Once
}

func (g *Group) done() {
	if g.cancel != nil {
		g.cancel()
	}
}

// WithContext returns a new Group and an associated Context derived from ctx. 🟩
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first. 🟩
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancellation := context.WithCancel(ctx)

	return &Group{
		cancel: cancellation,
	}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them. 🟩
func (g *Group) Wait() error {
	// Cancel the Context when we return 🟩
	defer g.done()

	// We need to block until all goroutines have returned 🟩
	g.wg.Wait()

	// return the first non-nil error 🟩
	return g.firstError
}

func (g *Group) tokenLimit() bool {
	return cap(g.tokens) > 0
}

// Go calls the given function in a new goroutine. 🟩
// It blocks until the new goroutine can be added without the number of
// active goroutines in the group exceeding the configured limit. 🟩
//
// The first call to return a non-nil error cancels the group's context, if the 🟩
// group was created by calling WithContext. The error will be returned by Wait. 🟩
func (g *Group) Go(f func() error) {
	// It blocks until the new goroutine can be added without the number of
	// active goroutines in the group exceeding the configured limit. 🟩
	// Check that we're below the limit, or if there is a limit at all

	if g.tokenLimit() {
		// Acquire the token. Blocks if no tokens are available
		g.tokens <- token{}
	}

	// calls the given function in a new goroutine. 🟩
	g.wg.Add(1)
	go g.run(f)

}

func (g *Group) run(f func() error) {
	defer g.wg.Done()
	defer func() {
		if g.tokenLimit() {
			<-g.tokens // Release our token
		}
	}()

	// When the function returns non-nil, we need to cancel the context 🟩
	err := f()
	if err != nil {
		// Way to share errors with Wait is needed 🟩
		g.firstErrorOnce.Do(func() {
			g.firstError = err
		})
		g.done()
	}
}

// TryGo calls the given function in a new goroutine only if the number of
// active goroutines in the group is currently below the configured limit. 🟩
//
// The return value reports whether the goroutine was started. 🟩
func (g *Group) TryGo(f func() error) bool {
	// If we don't have a limit, we can always execute 🟩
	if cap(g.tokens) == 0 {
		return false
	}

	if !g.tokenLimit() {
		g.wg.Add(1)
		go g.run(f)
		return true
	}

	// if the number of active goroutines in the group is currently below the configured limit 🟩
	select {
	case g.tokens <- token{}:
		// If we do have the token, we can process it
		g.wg.Add(1)
		go g.run(f)

		return true
	default:
		return false
	}
}

// SetLimit limits the number of active goroutines in this group to at most n.
// A negative value indicates no limit. 🟩
//
// Any subsequent call to the Go method will block until it can add an active
// goroutine without exceeding the configured limit. 🟩
//
// The limit must not be modified while any goroutines in the group are active. 🟩
func (g *Group) SetLimit(n int) {
	// Set limit of goroutines 🟩
	g.tokens = make(chan token, n)
}
