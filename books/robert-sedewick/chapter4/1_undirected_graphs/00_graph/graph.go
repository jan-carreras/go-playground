package graphs

func Degree(graph Grapher, v int) int {
	return len(graph.AdjacentVertices(v))
}

func MaxDegree(graph Grapher) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxDegree := 0
	for i := 0; i < graph.Vertices(); i++ {
		maxDegree = max(maxDegree, Degree(graph, i))
	}

	return maxDegree
}

func AverageDegree(graph Grapher) int {
	return 2 * graph.Edges() / graph.Vertices()
}

func NumberOfSelfLoops(graph Grapher) int {
	count := 0
	for v := 0; v < graph.Vertices(); v++ {
		for _, w := range graph.AdjacentVertices(v) {
			if v == w {
				count++
			}
		}
	}
	return count / 2 // each edge counted twice

}
