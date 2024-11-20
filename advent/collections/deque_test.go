package collections

import (
	"slices"
	"testing"
)

func TestDeque(t *testing.T) {
	d := NewDeque[int]()

	sliceOk(t, d, []int{})

	_, err := d.Pop()
	if err == nil {
		t.Error("did not get error on empty Pop")
	}

	d.Append(42)
	sliceOk(t, d, []int{42})

	d.AppendLeft(17)
	sliceOk(t, d, []int{17, 42})

	d.Append(99)
	sliceOk(t, d, []int{17, 42, 99})

	el, _ := d.Pop()
	elOk(t, el, 99)
	sliceOk(t, d, []int{17, 42})

	el, _ = d.PopLeft()
	elOk(t, el, 17)
	sliceOk(t, d, []int{42})

	d.Pop()
	sliceOk(t, d, []int{})
}

func sliceOk[T comparable](t *testing.T, d *Deque[T], expect []T) {
	got := d.ToSlice()
	if !slices.Equal(got, expect) {
		t.Errorf("got: %+v, expect: %+v", got, expect)
	}
}

func elOk[T comparable](t *testing.T, got T, expect T) {
	if got != expect {
		t.Errorf("got: %+v, expect: %+v", got, expect)
	}
}
