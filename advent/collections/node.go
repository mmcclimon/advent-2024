package collections

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

func makeNode[T any](val T) *Node[T] {
	return &Node[T]{data: val}
}

func (n Node[T]) String() string {
	return fmt.Sprintf("<Node data=%#v>", n.data)
}
