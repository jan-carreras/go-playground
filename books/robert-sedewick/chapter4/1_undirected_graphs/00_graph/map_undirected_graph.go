package graphs

import (
	"fmt"
	"strings"
)

var _ Grapher = &MapUndirectedGraph{}

type MapUndirectedGraph struct {
	edges int
	data  []map[int]bool
}

// NewMapUndirectedGraph creates a V-vertex graph with no edges
func NewMapUndirectedGraph(vertices int) *MapUndirectedGraph {
	data := make([]map[int]bool, vertices)
	for index := range data {
		data[index] = make(map[int]bool)
	}
	return &MapUndirectedGraph{
		data: data,
	}
}

func (g *MapUndirectedGraph) Vertices() int {
	return len(g.data)
}

func (g *MapUndirectedGraph) Edges() int {
	return g.edges
}

func (g *MapUndirectedGraph) AddEdge(v, w int) error {
	if err := g.areVertexInBounds(v, w); err != nil {
		return err
	}

	if _, ok := g.data[v][w]; ok {
		return nil
	}

	g.edges++
	// Will panic if the map is not initialized
	g.data[v][w] = true
	g.data[w][v] = true

	return nil
}

func (g *MapUndirectedGraph) AdjacentVertices(v int) []int {
	r := make([]int, 0, len(g.data[v]))
	for v := range g.data[v] {
		r = append(r, v)
	}
	return r
}

func (g *MapUndirectedGraph) String() string {
	b := strings.Builder{}
	for index, data := range g.data {
		if len(data) == 0 {
			continue
		}
		b.WriteString(fmt.Sprintf("%d: %v\n", index, data))
	}

	return b.String()
}

// areVertexInBounds errors when one vertex is 0 > v[n] >= g.data.len()
func (g *MapUndirectedGraph) areVertexInBounds(vs ...int) error {
	for _, v := range vs {
		if v < 0 || v >= len(g.data) {
			return fmt.Errorf("vertex %v expected to be between 0 and %d: %w", v, len(g.data), ErrInvalidVertex)
		}
	}
	return nil
}
