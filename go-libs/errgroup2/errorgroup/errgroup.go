// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errgroup provides synchronization, error propagation, and Context
// cancelation for groups of goroutines working on subtasks of a common task.
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
	cancel context.CancelFunc

	wg sync.WaitGroup

	firstError     error
	firstErrorOnce sync.Once

	tokens chan token
}

func (g *Group) done() {
	if g.cancel != nil {
		g.cancel()
	}
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancellation := context.WithCancel(ctx)

	return &Group{
		cancel: cancellation,
		wg:     sync.WaitGroup{},
		tokens: make(chan token, 0),
	}, ctx
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	defer func() {
		// fmt.Println("Cancelling context on Wait")
		g.done()
	}()

	// fmt.Println("Waiting for goroutines to stop")
	g.wg.Wait()
	// fmt.Println("Reporting error, if any")
	return g.firstError
}

// Go calls the given function in a new goroutine.
// It blocks until the new goroutine can be added without the number of
// active goroutines in the group exceeding the configured limit.
//
// The first call to return a non-nil error cancels the group's context, if the
// group was created by calling WithContext. The error will be returned by Wait.
func (g *Group) Go(f func() error) {
	// fmt.Println("Go: new job to be done")

	releaseToken := func() {}

	if cap(g.tokens) > 0 {
		// fmt.Println("Tokens are limited")
		// Try to acquire the token
		// fmt.Println("Trying to acquire the token")
		g.tokens <- token{}
		// fmt.Println("Token acquired")
		releaseToken = func() {
			// Release the token
			<-g.tokens
			// fmt.Println("Token released")
		}
	}

	// fmt.Println("Add goroutine to wait group")
	g.wg.Add(1)
	// fmt.Println("Create goroutine")
	go g.runAndReport(f, releaseToken)
}

func (g *Group) runAndReport(f func() error, releaseToken func()) {
	defer func() {
		// fmt.Println("wg.Done()")
		g.wg.Done()
	}()
	defer releaseToken()

	// Perform the task
	// fmt.Println("Running the given task")
	err := f()

	if err != nil {
		// Making sure that we assign the first error found and don't change it afterwards

		// fmt.Println("Reporting error once")
		g.firstErrorOnce.Do(func() {
			// fmt.Println("Error reported")
			g.firstError = err
		})
		// fmt.Println("Cancelling the context")
		g.done()
	}
}

// TryGo calls the given function in a new goroutine only if the number of
// active goroutines in the group is currently below the configured limit.
//
// The return value reports whether the goroutine was started.
func (g *Group) TryGo(f func() error) bool {
	releaseToken := func() {}

	// Check if there is a limit, otherwise just run the code
	if cap(g.tokens) < 0 {
		// fmt.Println("TryGo: Unlimited tokens. We can run the task")
		g.wg.Add(1)
		go g.runAndReport(f, releaseToken)

		return true
	}

	// If there is a limit, do a select to try to acquire a token
	select {
	case g.tokens <- token{}:
		// The token release must happen inside runAndReport
		// If we can, run the stuff and return true
		releaseToken = func() {
			<-g.tokens // Release the token!
		}
		g.wg.Add(1)
		go g.runAndReport(f, releaseToken)
		return true

	default:
		// If we cannot, return false
		return false
	}

}

// SetLimit limits the number of active goroutines in this group to at most n.
// A negative value indicates no limit.
//
// Any subsequent call to the Go method will block until it can add an active
// goroutine without exceeding the configured limit.
//
// The limit must not be modified while any goroutines in the group are active.
func (g *Group) SetLimit(n int) {
	// fmt.Println("Setting limit to", n)
	// TODO: We could add validation that there are not running goroutines
	g.tokens = make(chan token, n)
}
