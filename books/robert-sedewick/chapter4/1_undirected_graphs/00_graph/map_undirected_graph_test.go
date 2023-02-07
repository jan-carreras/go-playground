package graphs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMapUndirectedGraph(t *testing.T) {
	graph := readMapGraph(t, "testdata/tinyG.txt")

	cases := []Case{
		{vertex: 0, want: []int{5, 1, 2, 6}},
		{vertex: 1, want: []int{0}},
		{vertex: 4, want: []int{6, 3, 5}},
	}

	for _, c := range cases {
		if got := graph.AdjacentVertices(c.vertex); !equalVertices(got, c.want) {
			t.Errorf("invalid vertices: %v want %v", got, c.want)
		}
	}

	if got, want := graph.Edges(), 13; got != want {
		t.Errorf("invalid number of edges: %d want %d", got, want)
	}

	if got, want := graph.Vertices(), 13; got != want {
		t.Errorf("invalid number of vertices: %d want %d", got, want)
	}
}

// Reads a file with the following format:
// 13    # <--- V
// 13    # <--- E
// 0 5   # <--- v-w pair
// 4 3   # <--- v-w pair
func readMapGraph(t *testing.T, file string) *MapUndirectedGraph {
	f, err := os.Open(file)
	if err != nil {
		t.Errorf("unable to open file: %v", err)
	}
	t.Cleanup(func() { _ = f.Close() })

	buf := bufio.NewScanner(f)
	buf.Scan()
	vertexCount, err := strconv.Atoi(buf.Text())
	if err != nil {
		t.Errorf("invalid vertex count: %v", err)
	}

	graph := NewMapUndirectedGraph(vertexCount)

	buf.Scan() // Ignore second line — it describes number of edges
	for buf.Scan() {
		var v, e int
		_, err := fmt.Fscanf(strings.NewReader(buf.Text()), "%d %d", &v, &e)
		if err != nil {
			t.Errorf("error reading vertex: %v", err)
		}
		graph.AddEdge(v, e)
	}

	if buf.Err() != nil {
		t.Errorf("expecting no error when reading input: %v", buf.Err())
	}
	return graph
}
