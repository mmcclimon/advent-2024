package mathx

import (
	"iter"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

func Abs[T Numeric](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

func Sum[T Numeric](seq iter.Seq[T]) T {
	sum := T(0)
	for n := range seq {
		sum += n
	}

	return sum
}

func GCD[T constraints.Integer](m, n T) T {
	if n == 0 {
		return m
	}

	return GCD(n, m%n)
}

func LCM[T constraints.Integer](m, n T) T {
	return Abs(m*n) / GCD(m, n)
}
