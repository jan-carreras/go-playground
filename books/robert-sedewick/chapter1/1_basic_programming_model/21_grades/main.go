package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 1.1.21 Write a program that reads in lines from standard input with each line
// containing a name and two integers and then uses printf() to print a table
// with a column of the names, the integers, and the result of dividing the first
// by the second, accurate to three decimal places. You could use a program like
// this to tabulate batting averages for baseball players or grades for students.

/**

INPUT

John 10 4
Bob 3 7
Alice 2 9
*/

type Grade struct {
	name string
	a, b int
}

type Grades []Grade

func (grades Grades) String() string {
	buf := bytes.Buffer{}

	header := fmt.Sprintf("%15s\t%15s\t%15s\t%15s\n", "name", "sumScores", "scoresCount", "avg")
	buf.WriteString(header)

	buf.WriteString(strings.Repeat("*", len(header)))
	buf.WriteString("\n")

	for _, grade := range grades {
		line := fmt.Sprintf("%15s\t%15d\t%15d\t%15.3f\n", grade.name, grade.a, grade.b, float64(grade.a)/float64(grade.b))
		buf.WriteString(line)
	}

	return buf.String()
}
