package grades

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRageGrades(t *testing.T) {
	grade, err := rateGrades([]string{"A+", "C", "F"})
	require.NoError(t, err)
	require.EqualValues(t, 2.11, grade)

	grade, err = rateGrades([]string{"A+", "A+"})
	require.NoError(t, err)
	require.EqualValues(t, 4.33, grade)

	grade, err = rateGrades([]string{"A+++++"})
	require.Error(t, err)
}
