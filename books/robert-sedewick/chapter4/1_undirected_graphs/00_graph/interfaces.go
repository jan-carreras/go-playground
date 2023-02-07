package graphs

import "errors"

var ErrInvalidVertex = errors.New("invalid vertex")

type Grapher interface {
	// Vertices return the number of vertices
	Vertices() int
	// Edges return the number of edges
	Edges() int
	// AddEdge adds v-w edge to this graph. 0 <= v|w < Vertices
	// Errors returned can be: ErrInvalidVertex if v or w is out of bounds
	AddEdge(v, w int) error
	// AdjacentVertices returns vertices want to v
	AdjacentVertices(v int) []int
	String() string
}
