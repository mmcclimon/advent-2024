package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type Report []int

func main() {
	safe1, safe2 := 0, 0

	for line := range input.New().Lines() {
		r := NewReport(line)

		if r.IsSafe() {
			safe1++
		}

		if r.IsSafe2() {
			safe2++
		}
	}

	fmt.Println("part 1:", safe1)
	fmt.Println("part 2:", safe2)
}

func NewReport(text string) Report {
	levels := conv.ToInts(strings.Fields(text))
	return Report(levels)
}

func (r Report) IsSafe() bool {
	var hasPos, hasNeg bool

	for i := range len(r) - 1 {
		delta := r[i+1] - r[i]

		switch {
		case delta == 0,
			delta < -3,
			delta > 3:
			return false
		case delta > 0:
			hasPos = true
		case delta < 0:
			hasNeg = true
		}

		if hasPos && hasNeg {
			return false
		}
	}

	return true
}

func (r Report) IsSafe2() bool {
	if r.IsSafe() {
		return true
	}

	// stupid and inefficient
	for i := range r {
		l2 := slices.Delete(slices.Clone(r), i, i+1)

		if Report(l2).IsSafe() {
			return true
		}
	}

	return false
}
