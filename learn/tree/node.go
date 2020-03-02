package tree

import "fmt"

// Node has a value and two children
type Node struct {
	Value       float64
	Left, Right *Node
}

// SetValue sets node's value
func (node *Node) SetValue(v float64) {
	node.Value = v
}

// Print a node
func (node *Node) Print() {
	fmt.Printf("%f\n", node.Value)
}

// Traverse a node
func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// TraverseFunc travels a node and apply a function f
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
