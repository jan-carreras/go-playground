package ast

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSingleArraySymbolTable(t *testing.T) {
	ast := SingleArraySymbolTable[string, float64]{}

	ast.Del("nonexistent key")
	ast.Set("invalid", 0)
	ast.Del("invalid")

	ast.Set("A+", 4.33)
	ast.Set("A", 4)
	ast.Set("A-", 3.67)
	ast.Set("D", 100)
	ast.Set("D", 1) // Overwriting with the correct value
	ast.Set("F", 1)
	ast.Set("F", 0) // Overwriting with the correct value
	ast.Set("B+", 3.33)
	ast.Set("B", 3)
	ast.Set("B-", 2.67)
	ast.Set("C+", 2.33)
	ast.Set("C", 2)
	ast.Set("C-", 1.67)

	ast.Del("A+")
	ast.Del("D")
	ast.Del("F")

	ast.Set("A+", 4.33)
	ast.Set("D", 1)
	ast.Set("F", 0)

	for i := 0; i < ast.size; i++ {
		fmt.Printf("%2s: %3.2f\n", ast.kv[i].key, ast.kv[i].value)
	}

	require.Nil(t, ast.Get("nonexistent key"))
	require.EqualValues(t, 4, *ast.Get("A"))
	require.EqualValues(t, 2, *ast.Get("C"))
}

func TestSingleArrayKeys(t *testing.T) {
	ast := SingleArraySymbolTable[string, float64]{}

	ast.Set("A+", 4.33)
	ast.Set("A", 4)
	ast.Set("A-", 3.67)
	ast.Set("D", 100)
	ast.Set("D", 1) // Overwriting with

	require.Equal(t, []string{"A", "A+", "A-", "D"}, ast.Keys())
}
