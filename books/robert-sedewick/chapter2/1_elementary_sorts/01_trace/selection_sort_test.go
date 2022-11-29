package selection_sort

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSort_Ints(t *testing.T) {
	input := []int{1, 5, 4, 3, 8, 1}
	s := SelectionSort[int]{}
	s.Sort(input)
	require.Equal(t, []int{1, 1, 3, 4, 5, 8}, input)
}

func TestSort_Strings(t *testing.T) {
	input := []string{"E", "A", "S", "Y", "Q", "U", "E", "S", "T", "I", "O", "N"}
	buf := &bytes.Buffer{}
	s := SelectionSort[string]{
		writer: buf,
	}
	s.Sort(input)
	require.Equal(t, "AENSOQEISTUY", strings.Join(input, ""))

	require.Equal(t, `i= 0 min= 1 [E A S Y Q U E S T I O N]
i= 1 min= 1 [A E S Y Q U E S T I O N]
i= 2 min=11 [A E S Y Q U E S T I O N]
i= 3 min=11 [A E N Y Q U E S T I O S]
i= 4 min=10 [A E N S Q U E S T I O Y]
i= 5 min=10 [A E N S O U E S T I Q Y]
i= 6 min= 6 [A E N S O Q E S T I U Y]
i= 7 min= 9 [A E N S O Q E S T I U Y]
i= 8 min= 9 [A E N S O Q E I T S U Y]
i= 9 min= 9 [A E N S O Q E I S T U Y]
i=10 min=10 [A E N S O Q E I S T U Y]
i=11 min=11 [A E N S O Q E I S T U Y]
`, buf.String())
}
