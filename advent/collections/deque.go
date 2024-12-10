package collections

import (
	"errors"
	"iter"
)

var (
	ErrEmptyPop  = errors.New("cannot pop an empty deque")
	ErrEmptyPeek = errors.New("cannot peek an empty deque")
)

type Deque[T any] struct {
	length int
	left   *Node[T]
	right  *Node[T]
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (d *Deque[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for cur := d.left; cur != nil; cur = cur.next {
			yield(cur.data)
		}
	}
}

func (d *Deque[T]) Len() int {
	return d.length
}

func (d *Deque[T]) ToSlice() []T {
	ret := make([]T, 0, d.length)
	for val := range d.Iter() {
		ret = append(ret, val)
	}

	return ret
}

func (d *Deque[T]) Append(val T) {
	node := makeNode(val)
	d.length++

	if d.right == nil {
		d.left = node
		d.right = node
		return
	}

	d.right.next = node
	node.prev = d.right
	d.right = node
}

func (d *Deque[T]) AppendLeft(val T) {
	node := makeNode(val)
	d.length++

	if d.left == nil {
		d.left = node
		d.right = node
		return
	}

	d.left.prev = node
	node.next = d.left
	d.left = node
}

func (d *Deque[T]) Pop() (T, error) {
	if d.length == 0 {
		return *new(T), ErrEmptyPop
	}

	ret := d.right.data
	d.right = d.right.prev
	d.length--

	// If we popped the last element, also delete the left one.
	if d.right == nil {
		d.left = nil
	} else {
		d.right.next = nil
	}

	return ret, nil
}

func (d *Deque[T]) PopLeft() (T, error) {
	if d.length == 0 {
		return *new(T), ErrEmptyPop
	}

	ret := d.left.data
	d.left = d.left.next
	d.length--

	// If we popped the last element, also delete the right one.
	if d.left == nil {
		d.right = nil
	} else {
		d.left.prev = nil
	}

	return ret, nil
}

func (d *Deque[T]) Peek() (T, error) {
	if d.length == 0 {
		return *new(T), ErrEmptyPeek
	}

	return d.right.data, nil
}

func (d *Deque[T]) PeekLeft() (T, error) {
	if d.length == 0 {
		return *new(T), ErrEmptyPeek
	}

	return d.left.data, nil
}
