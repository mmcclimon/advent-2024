package collections

import (
	"iter"
)

type Set[T comparable] map[T]struct{}

var unit struct{}

func NewSet[T comparable](elems ...T) Set[T] {
	s := make(map[T]struct{}, len(elems))
	for _, elem := range elems {
		s[elem] = unit
	}

	return Set[T](s)
}

func (s Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		s[elem] = unit
	}
}

func (s Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s {
			if !yield(v) {
				return
			}
		}
	}
}
