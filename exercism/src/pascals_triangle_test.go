package src

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	for _, line := range Triangle(10) {
		t.Log(line)
	}
}
