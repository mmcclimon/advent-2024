package mathx

import (
	"iter"
	"math"
)

type Numeric interface {
	int | uint | int64 | uint64 |
		float32 | float64
}

func Abs[T Numeric](n T) T {
	return T(math.Abs(float64(n)))
}

func Sum[T Numeric](seq iter.Seq[T]) T {
	sum := T(0)
	for n := range seq {
		sum += n
	}

	return sum
}
