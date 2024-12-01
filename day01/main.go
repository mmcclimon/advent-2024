package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/mathx"
)

func main() {
	var left, right []int
	rmap := make(map[int]int)

	for line := range input.New().Lines() {
		digits := strings.Fields(line)

		l, err := strconv.Atoi(digits[0])
		assert.Nil(err)
		left = append(left, l)

		r, err := strconv.Atoi(digits[1])
		assert.Nil(err)
		right = append(right, r)
		rmap[r]++
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
