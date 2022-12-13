package tree

import "fmt"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

/**
{5 1}
{3 2}
{2 0}
{4 1}
{1 0}
{0 0}
{6 2}
*/
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	fmt.Println()

	return nil, nil
}
