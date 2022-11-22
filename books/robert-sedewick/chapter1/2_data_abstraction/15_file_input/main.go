package main

import (
	"bytes"
	"io"
	"strconv"
)

func ReadInts(input io.Reader) ([]int, error) {
	in, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	if len(in) == 0 {
		return nil, nil
	}

	rawInts := bytes.Split(in, []byte(" "))
	ints := make([]int, 0, len(rawInts))
	for _, rawInt := range rawInts {
		i, err := strconv.Atoi(string(rawInt))
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}

	return ints, nil
}
