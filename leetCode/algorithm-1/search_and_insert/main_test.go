package search_and_insert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	assert.Equal(t, 2, SearchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, SearchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, SearchInsert([]int{1, 3, 5, 6}, 7))
	assert.Equal(t, 1, SearchInsert([]int{1}, 2))
	assert.Equal(t, 0, SearchInsert([]int{1}, 0))
	assert.Equal(t, 0, SearchInsert([]int{}, 0))
}
