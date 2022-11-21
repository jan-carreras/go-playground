package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrintGrades(t *testing.T) {
	gs := make(Grades, 0)
	gs = append(gs, Grade{name: "john", a: 10, b: 2})
	gs = append(gs, Grade{name: "bob", a: 10, b: 3})
	gs = append(gs, Grade{name: "alice", a: 10, b: 4})

	expectedOutput := `           name	      sumScores	    scoresCount	            avg
****************************************************************
           john	             10	              2	          5.000
            bob	             10	              3	          3.333
          alice	             10	              4	          2.500
`

	require.Equal(t, expectedOutput, gs.String())
}
