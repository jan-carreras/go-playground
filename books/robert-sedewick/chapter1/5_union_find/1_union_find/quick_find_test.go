package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strings"
	"testing"
)

func TestQuickFind(t *testing.T) {
	tests := []struct {
		algorithm string
		input     string
		output    string
	}{
		{
			algorithm: "quickFind",
			input:     "./testdata/tinyUF.txt",
			output:    "./testdata/tinyUFoutput.txt",
		},
		{
			algorithm: "quickFind",
			input:     "./testdata/mediumUF.txt",
			output:    "./testdata/mediumUFoutput.txt",
		},
		// Skipping ./testdata/largeUF.txt -> Not solvable
		{
			algorithm: "quickUnion",
			input:     "./testdata/tinyUF.txt",
			output:    "./testdata/tinyUFoutput.txt",
		},
		{
			algorithm: "quickUnion",
			input:     "./testdata/mediumUF.txt",
			output:    "./testdata/mediumUFoutput.txt",
		},
		// Skipping ./testdata/largeUF.txt -> Not solvable
		{
			algorithm: "weightedQuickUnion",
			input:     "./testdata/tinyUF.txt",
			output:    "./testdata/tinyUFoutput.txt",
		},
		{
			algorithm: "weightedQuickUnion",
			input:     "./testdata/mediumUF.txt",
			output:    "./testdata/mediumUFoutput.txt",
		},

		{ // Takes 2.56s seconds
			algorithm: "weightedQuickUnion",
			input:     "./testdata/largeUF.txt",
			output:    "./testdata/largeUFoutput.txt",
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s - %s", tt.algorithm, tt.input)
		t.Run(testName, func(t *testing.T) {
			input := strings.NewReader(readFile(t, tt.input))
			output := &bytes.Buffer{}

			err := Run(tt.algorithm, input, output)
			require.NoError(t, err)

			require.Equal(t, readFile(t, tt.output), output.String())
		})
	}
}

func readFile(t *testing.T, filename string) string {
	f, err := os.Open(filename)
	require.NoError(t, err)
	defer f.Close()

	data, err := io.ReadAll(f)
	require.NoError(t, err)

	return string(data)
}
