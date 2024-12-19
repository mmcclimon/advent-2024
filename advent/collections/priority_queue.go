package collections

import (
	"fmt"
)

type MinQueue[T comparable] struct {
	data     []T
	contains map[T]int
	n        int
	cmp      func(a, b T) int
}

func NewMinQueue[T comparable](compare func(a, b T) int) *MinQueue[T] {
	return &MinQueue[T]{
		data:     []T{*new(T)},
		contains: make(map[T]int),
		n:        0,
		cmp:      compare,
	}
}

func (mq *MinQueue[T]) Len() int {
	return mq.n
}

func (mq *MinQueue[T]) Debug() {
	fmt.Println(mq.data)
}

func (mq *MinQueue[T]) Insert(item T) {
	mq.data = append(mq.data, item)
	mq.contains[item]++
	mq.n++
	mq.swim(mq.n)
}

func (mq *MinQueue[T]) ExtractMin() T {
	ret := mq.data[1]
	mq.exchange(1, mq.n)
	mq.data = mq.data[:mq.n]
	mq.n--
	mq.sink(1)

	mq.contains[ret]--
	if mq.contains[ret] == 0 {
		delete(mq.contains, ret)
	}

	return ret
}

func (mq *MinQueue[T]) Contains(item T) bool {
	return mq.contains[item] > 0
}

func (mq *MinQueue[T]) swim(idx int) {
	for idx > 1 && mq.greater(idx/2, idx) {
		mq.exchange(idx/2, idx)
		idx = idx / 2
	}
}

func (mq *MinQueue[T]) sink(idx int) {
	for 2*idx <= mq.n {
		j := 2 * idx
		if j < mq.n && mq.greater(j, j+1) {
			j++
		}

		if !mq.greater(idx, j) {
			break
		}

		mq.exchange(idx, j)
		idx = j
	}
}

func (mq *MinQueue[T]) greater(a, b int) bool {
	return mq.cmp(mq.data[a], mq.data[b]) > 0
}

func (mq *MinQueue[T]) exchange(a, b int) {
	mq.data[a], mq.data[b] = mq.data[b], mq.data[a]
}
