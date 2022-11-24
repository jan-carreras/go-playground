package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	input := "1.3.45_stack_and_other_stuff"
	var chapter, module, exercise int
	var name string

	fmt.Println(fmt.Sscanf(input, "%d.%d.%d_%s", &chapter, &module, &exercise, &name))

	fmt.Println(chapter, module, exercise, name)

}
