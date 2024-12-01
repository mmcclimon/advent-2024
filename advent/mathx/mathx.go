package mathx

import "math"

type Numeric interface {
	int | uint | int64 | uint64 |
		float32 | float64
}

func Abs[T Numeric](n T) T {
	return T(math.Abs(float64(n)))
}
