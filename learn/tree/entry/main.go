package main

import (
	"fmt"

	"github.com/neo0801/MyCode/learn/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3.0, Left: nil, Right: nil}
	root.Left = &tree.Node{Value: 0.0}
	root.Right = &tree.Node{Value: 5.0, Left: nil, Right: nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = &tree.Node{Value: 2.0, Left: nil, Right: nil}
	root.Right.Left.SetValue(4.0)

	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0.0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Printf("max node value is:%f\n", maxNode)
}
