package main

import "sync"

type Redis struct {
	storeMux sync.Mutex
	store    map[string]int
}

func (r *Redis) foo() {
	// TODO: Finish it
}
