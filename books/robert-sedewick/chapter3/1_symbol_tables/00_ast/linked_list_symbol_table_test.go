package ast

import (
	"container/list"
	"fmt"
	"testing"
)

func TestOrderedLinkedListSymbolTable(t *testing.T) {
	ast := OrderedLinkedListSymbolTable[string, float64]{}

	ast.Set("jan", 1234)
	ast.Set("bob", 4321)
	ast.Set("alice", 444999)

	ast.Set("A+", 4.33)
	ast.Set("D", 1)
	ast.Set("F", 0)
	ast.Del("D")

	nodeKey, nodeValue := ast.keys.Front(), ast.values.Front()
	for nodeKey != nil {
		if nodeValue == nil {
			printList(ast.keys)
			printList(ast.values)
			panic("fuck, nodeValue is nil when it shouldn't")
		}
		fmt.Printf("%2s: %3.2f\n", nodeKey.Value.(string), nodeValue.Value.(float64))
		nodeKey, nodeValue = nodeKey.Next(), nodeValue.Next()
	}
}

func printList(l list.List) {
	fmt.Println("length=", l.Len())
	node := l.Front()
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next()
	}

}
