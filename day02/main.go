package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type Report struct {
	Levels []int
}

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
	var levels []int
	for _, data := range strings.Fields(text) {
		n, err := strconv.Atoi(data)
		assert.Nil(err)
		levels = append(levels, n)
	}

	return Report{levels}
}

func (r Report) IsSafe() bool {
	var hasPos, hasNeg bool

	for i := range len(r.Levels) - 1 {
		delta := r.Levels[i+1] - r.Levels[i]
		if delta < -3 || delta > 3 {
			return false
		}

		switch {
		case delta == 0:
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
	for i := range len(r.Levels) {
		l2 := make([]int, 0, len(r.Levels))

		for j := range r.Levels {
			if i == j {
				continue
			}

			l2 = append(l2, r.Levels[j])
		}

		if (Report{l2}).IsSafe() {
			return true
		}
	}

	return false
}
