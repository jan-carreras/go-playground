package tree

import (
	"fmt"
	"sort"
)

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

func NewNode(ID int) *Node {
	return &Node{ID: ID, Children: make([]*Node, 0)}
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	nodes := make(map[int]*Node)
	for i, record := range records {
		if i != record.ID { // Sanity check
			return nil, fmt.Errorf("non-continuous ID: ids are expected to be sequential")
		} else if _, duplicate := nodes[record.ID]; duplicate { // Sanity check
			return nil, fmt.Errorf("record already imported: duplicates are not allowed")
		} else if record.ID == 0 && record.Parent != 0 { // Sanity check
			return nil, fmt.Errorf("root node (id=0) should have parent=0")
		}

		if record.ID == 0 { // Root node
			nodes[record.ID] = NewNode(record.ID)
			continue
		}

		parent, found := nodes[record.Parent]
		if !found { // Sanity check
			return nil, fmt.Errorf("parent for node %d not found: %v", record.ID, record)
		}

		nodes[record.ID] = NewNode(record.ID)
		parent.Children = append(parent.Children, nodes[record.ID])
	}

	root, found := nodes[0]
	if !found {
		return nil, fmt.Errorf("node 0  (root) not found!")
	}

	return root, nil
}
