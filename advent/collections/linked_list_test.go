package collections

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := NewList(4, 9, 16, 25)

	// for n := range ll.Iter() {
	// 	fmt.Println(n)
	// }

	ll.Cons(1)

	for n := range ll.Iter() {
		fmt.Println(n)
	}
}
