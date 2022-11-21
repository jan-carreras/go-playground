package __1_28_remove_duplicated

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch(t *testing.T) {

	bs := NewBinarySearch([]int{1, 1, 2, 3, 4, 5, 5, 6, 6})
	fmt.Println(bs)

	require.Equal(t, 2, bs.Search(3))
	require.Equal(t, -1, bs.Search(7))
}
