package __1_11_box

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	b := make(Box, 10)
	for i := range b {
		b[i] = make([]bool, 10)
		for j := range b[i] {
			if rand.Intn(10) == 0 {
				b[i][j] = true
			}
		}
	}

	fmt.Println(b)

}
