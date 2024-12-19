package collections

import (
	"cmp"
	"fmt"
	"testing"
)

func TestPQ(t *testing.T) {
	pq := NewMinQueue(cmp.Compare[int])
	pq.Insert(1)

	fmt.Println(pq.data)
	pq.Insert(2)
	fmt.Println(pq.data)

	pq.Insert(17)
	pq.Insert(15)
	pq.Insert(97)
	pq.Insert(-64)

	fmt.Println(pq.data)

	fmt.Println(pq.ExtractMin())
	fmt.Println(pq.data)
}
