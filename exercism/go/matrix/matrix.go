package matrix

import (
	"errors"
	"strconv"
	"strings"
)

var errRowsNotSameSize = errors.New("rows have not the same size")
var errInvalidInput = errors.New("invalid input. The matrix cannot be empty or be malformatted")

type Matrix struct {
	m [][]int
}

func isValidInput(m string) bool {
	return m == "" || strings.HasPrefix(m, "\n") || strings.HasSuffix(m, "\n")
}

func readLines(m string) []string {
	return strings.Split(strings.TrimSpace(m), "\n")
}

func readCol(line string) []string {
	return strings.Split(strings.TrimSpace(line), " ")
}

func New(m string) (*Matrix, error) {
	if isValidInput(m ){
		return nil, errInvalidInput
	}

	matrix := Matrix{}
	lines := readLines(m)
	matrix.m = make([][]int, len(lines))
	for i, line := range lines {
		cols := readCol(line)
		if i > 0 && len(cols) != len(matrix.m[i-1]) {
			return nil, errRowsNotSameSize
		}
		matrix.m[i] = make([]int, len(cols))
		for j, num := range cols {
			if num == "" {
				continue
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			matrix.m[i][j] = n
		}
	}

	return &matrix, nil
}

func (m *Matrix) Rows() [][]int {
	newM := make([][]int, len(m.m))
	for i := range m.m {
		newM[i] = make([]int, len(m.m[i]))
		for j := range m.m[i] {
			newM[i][j] = m.m[i][j]
		}

	}
	return newM
}

func (m *Matrix) Cols() [][]int {
	newM := make([][]int, 0)
	for j:=0; j<len(m.m[0]); j++ {
		col := make([]int, 0)
		for i:=0; i<len(m.m); i++ {
			col = append(col, m.m[i][j])
		}
		newM = append(newM, col)
	}
	return newM
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.m) {
		return false
	}
	if col < 0 || col >= len(m.m[0]) {
		return false
	}
	m.m[row][col] = val
	return true
}