package collections

import (
	"fmt"
	"iter"
)

type LinkedList[T any] struct {
	root *Node[T]
}

func NewList[T any](elems ...T) *LinkedList[T] {
	ll := LinkedList[T]{}
	if len(elems) == 0 {
		return &ll
	}

	ll.root = makeNode(elems[0])
	prev := ll.root

	for _, item := range elems[1:] {
		node := makeNode(item)
		prev.next = node
		node.prev = prev
		prev = node
	}

	return &ll
}

func (ll *LinkedList[T]) String() string {
	return fmt.Sprintf("<LinkedList root=%s>", ll.root)
}

func (ll *LinkedList[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for cur := ll.root; cur != nil; cur = cur.next {
			yield(cur.data)
		}
	}
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.root == nil
}

func (ll *LinkedList[T]) Cons(value T) {
	node := makeNode(value)
	if ll.IsEmpty() {
		ll.root = node
		return
	}

	node.next = ll.root
	ll.root.prev = node
	ll.root = node
}
