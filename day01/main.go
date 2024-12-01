package main

import (
	"fmt"
	"slices"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/mathx"
)

type Line struct {
	Left  int
	Right int
}

func main() {
	var left, right []int
	rmap := make(map[int]int)

	s := input.NewStrummer()

	for s.HasLines() {
		var line Line
		err := s.Decode(&line)
		assert.Nil(err)

		left = append(left, line.Left)
		right = append(right, line.Right)
		rmap[line.Right]++
	}

	slices.Sort(left)
	slices.Sort(right)

	dist, sim := 0, 0
	for i := range left {
		l, r := left[i], right[i]
		dist += mathx.Abs(l - r)
		sim += l * rmap[l]
	}

	fmt.Println("part 1:", dist)
	fmt.Println("part 2:", sim)
}
