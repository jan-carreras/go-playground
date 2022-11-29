package union_find

import "fmt"

type Algo string

const QuickFindAlgo Algo = "quickFind"
const QuickUnionAlgo Algo = "quickUnion"
const WeightedQuickUnionAlgo Algo = "weightedQuickUnion"
const WeightedQuickUnionPathCompressionAlgo Algo = "weightedQuickUnionPathCompression"

type UnionFind interface {
	// Union adds a connection between p and q
	Union(p, q int)
	// Find returns the componentID of p
	Find(p int) (componentID int)
	// Connected returns if p and q are in the same component
	Connected(p, q int) (connected bool)
	// Count returns the number of components
	Count() int
}

func New(algorithm Algo, sites int) (UnionFind, error) {
	switch algorithm {
	case QuickFindAlgo:
		return NewQuickFind(sites), nil
	case QuickUnionAlgo:
		return NewQuickUnion(sites), nil
	case WeightedQuickUnionAlgo:
		return NewWeightedQuickUnion(sites), nil
	case WeightedQuickUnionPathCompressionAlgo:
		return NewWeightedQuickUnionPathCompression(sites), nil
	default:
		return nil, fmt.Errorf("unknown algorithm: %s", algorithm)
	}
}
