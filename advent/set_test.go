package advent

import (
	"testing"
)

func TestIter(t *testing.T) {
	s := NewSet(42, 17, 29)

	expect := map[int]bool{
		42: true,
		17: true,
		29: true,
	}

	called := 0

	for el := range s.Iter() {
		delete(expect, el)
		called++
	}

	if len(expect) != 0 || called != 3 {
		t.Error("iter did something weird")
		t.Log(expect)
	}
}
