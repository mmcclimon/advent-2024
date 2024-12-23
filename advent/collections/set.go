package collections

import (
	"iter"
	"slices"
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

func NewSetFromIter[T comparable](iterator iter.Seq[T]) Set[T] {
	s := NewSet[T]()
	for elem := range iterator {
		s[elem] = unit
	}

	return s
}

func (s Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		s[elem] = unit
	}
}

func (s Set[T]) Delete(elem T) {
	delete(s, elem)
}

func (s Set[T]) Contains(elem T) bool {
	if s == nil {
		return false
	}

	_, ok := s[elem]
	return ok
}

func (s Set[T]) Peek() T {
	for item := range s.Iter() {
		return item
	}

	panic("cannot peek an empty set")
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

func (s Set[T]) ToSlice() []T {
	return slices.Collect(s.Iter())
}

func (s Set[T]) Extend(other Set[T]) {
	for elem := range other.Iter() {
		s.Add(elem)
	}
}

func (s Set[T]) Clone() Set[T] {
	dupe := make(map[T]struct{}, len(s))
	for k := range s {
		dupe[k] = unit
	}
	return dupe
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	ret := make(map[T]struct{})

	for k := range s {
		if other.Contains(k) {
			ret[k] = unit
		}
	}

	return ret
}
