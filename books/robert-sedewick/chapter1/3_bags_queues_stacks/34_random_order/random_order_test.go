package random_order

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestRandomBag_Add(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())

	generate := func() *RandomBag[string] {
		r := new(RandomBag[string])
		r.Add("hello")
		r.Add("beautiful")
		r.Add("world")
		return r
	}

	for i := 0; i < 100; i++ {
		if generate().String() != generate().String() {
			require.True(t, true, "we have generated a different output")
			return
		}
	}

	t.Fatal("we're generating the same string")
}
