package main

import (
	"container/list"
	"fmt"
)

type UnionFind interface {
	// Union adds a connection between p and q
	Union(p, q int)
	// Find returns the componentID of p
	Find(p int) (componentID string)
	// Connected returns if p and q are in the same component
	Connected(p, q int) (connected bool)
	// Count returns the number of components
	Count() int
}

type UF struct {
	components []*list.List
}

func NewUF(sites int) *UF {
	components := make([]*list.List, sites)
	for i := range components {
		components[i] = &list.List{}
		components[i].PushFront(i)
	}

	return &UF{
		components: components,
	}
}

func (u *UF) componentID(component *list.List) string {
	return fmt.Sprintf("%p", component)
}

func (u *UF) Union(p, q int) {
	if u.Connected(p, q) {
		return // If they are connected already, there is nothing to do here
	}

	compP := u.components[p]
	compP.PushBackList(u.components[q])

	old := u.components[q]
	for i := range u.components {
		if u.components[i] == old {
			u.components[i] = compP
		}

	}
}

func (u *UF) Find(p int) (componentID string) {
	return u.componentID(u.components[p])
}

func (u *UF) Connected(p, q int) (connected bool) {
	return u.Find(p) == u.Find(q)
}

func (u *UF) Count() int {
	set := make(map[string]bool)
	for _, component := range u.components {
		set[u.componentID(component)] = true
	}

	return len(set)
}
