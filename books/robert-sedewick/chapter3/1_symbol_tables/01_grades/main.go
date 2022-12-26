package grades

import (
	"fmt"
	st "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter3/1_symbol_tables/00_ast"
)

// 3.1.1
//
// Write a client that creates a symbol table mapping letter grades to numerical
// scores, as in the table below, then reads from standard input a list of letter
// grades and computes and prints the GPA (the average of the numbers
// corresponding to the grades).

func rateGrades(grades []string) (score float64, err error) {
	scoresToGrade := scoreMap()

	for _, grade := range grades {
		s := scoresToGrade.Get(grade)
		if s == nil {
			return 0, fmt.Errorf("unknown grade %q", grade)
		}
		score += *s
	}

	return score / float64(len(grades)), nil
}

func scoreMap() st.ArraySymbolTable[string, float64] {
	st := st.ArraySymbolTable[string, float64]{}
	st.Set("A+", 4.33)
	st.Set("A", 4)
	st.Set("A-", 3.67)
	st.Set("D", 100)
	st.Set("D", 1) // Overwriting with the correct value
	st.Set("F", 0)
	st.Set("B+", 3.33)
	st.Set("B", 3)
	st.Set("B-", 2.67)
	st.Set("C+", 2.33)
	st.Set("C", 2)
	st.Set("C-", 1.67)
	return st
}
