package main

import (
	"fmt"

	"github.com/neo0801/MyCode/learn/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3, Left: nil, Right: nil}
	root.Left = &tree.Node{Value: 0}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = &tree.Node{Value: 2, Left: nil, Right: nil}
	root.Right.Left.SetValue(4)

	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
