package node

import (
	"fmt"
	"reflect"
)

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

func NewNode() *Node {
	return &Node{
		Data: nil,
		Next: nil,
		Prev: nil,
	}
}

func (node *Node) GetData() interface{} {
	return node.Data
}

func (node *Node) GetNext() *Node {
	return node.Next
}

func (node *Node) GetPrev() *Node {
	return node.Prev
}

func (node *Node) GetDataType() string {
	return fmt.Sprint(reflect.TypeOf(node.Data))
}
