package main

import (
	"flag"
	"fmt"
	"github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/5_union_find/union_find/union_find"
	"math/rand"
	"os"
	"time"
)

// 1.5.17 Random connections. Develop a UF client ErdosRenyi that takes an
// integer value N from the command line, generates random pairs of integers
// between 0 and N-1, calling connected() to determine if they are connected and
// then union() if not (as in our development client), looping until all sites
// are connected, and printing the number of connections generated. Package your
// program as a static method count() that takes N as argument and returns the
// number of connections and a main() that takes N from the command line, calls
// count(), and prints the returned value.

func main() {
	var randomPairs int
	flag.IntVar(&randomPairs, "r", 0, "random pairs to generate")
	flag.Parse()

	if randomPairs <= 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixMilli())

	connectionsGenerated := count(randomPairs)
	fmt.Println("connectionsGenerated=", connectionsGenerated)

}

func count(randomPairs int) int {
	uf := union_find.NewWeightedQuickUnionPathCompression(randomPairs)

	connectionsGenerated := 0

	for uf.Count() != 1 {
		p, q := rand.Intn(randomPairs), rand.Intn(randomPairs)
		if uf.Connected(p, q) {
			continue
		}
		uf.Union(p, q)
		connectionsGenerated++
	}

	return connectionsGenerated
}
