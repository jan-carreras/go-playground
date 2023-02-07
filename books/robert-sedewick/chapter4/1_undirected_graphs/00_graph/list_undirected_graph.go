package graphs

import (
	"fmt"
	"strings"
)

var _ Grapher = &MapUndirectedGraph{}

type ListUndirectedGraph struct {
	edges int
	data  [][]int
}

// NewListUndirectedGraph creates a V-vertex graph with no edges
func NewListUndirectedGraph(vertices int) *ListUndirectedGraph {
	return &ListUndirectedGraph{
		data: make([][]int, vertices),
	}
}

func (g *ListUndirectedGraph) Vertices() int {
	return len(g.data)
}

func (g *ListUndirectedGraph) Edges() int {
	return g.edges
}

func (g *ListUndirectedGraph) AddEdge(v, w int) error {
	for _, i := range g.data[v] {
		if i == w {
			break // Already related, no-op.
		}
	}
	g.edges++
	g.data[v] = append(g.data[v], w)
	g.data[w] = append(g.data[w], v)

	return nil
}

func (g *ListUndirectedGraph) AdjacentVertices(v int) []int {
	return g.data[v]
}

func (g *ListUndirectedGraph) String() string {
	b := strings.Builder{}
	for index, data := range g.data {
		if len(data) == 0 {
			continue
		}
		b.WriteString(fmt.Sprintf("%d: %v\n", index, data))
	}

	return b.String()
}
