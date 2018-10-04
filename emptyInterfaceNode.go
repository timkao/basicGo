package main

import (
	"fmt"
)

type Node struct {
	le   *Node
	ri   *Node
	data interface{}
}

func NewNode(left, right *Node) *Node {
	// or return &Node{left, right}
	result := new(Node)
	result.le = left
	result.ri = right
	return result
}

func (p *Node) SetData(data interface{}) {
	p.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("Hello World!")
	leftNode := NewNode(nil, nil)
	leftNode.SetData("I am on the left")
	rightNode := NewNode(nil, nil)
	rightNode.SetData(100)
	root.le = leftNode
	root.ri = rightNode

	fmt.Println(root.data)
	fmt.Println(root.ri.data)
	fmt.Println(root.le.data)
}
