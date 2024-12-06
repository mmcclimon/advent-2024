package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	hunks := slices.Collect(input.New().Hunks())

	rules := make(map[int]collections.Set[int])

	for _, line := range hunks[0] {
		bits := strings.Split(line, "|")
		l := conv.Atoi(bits[0])
		r := conv.Atoi(bits[1])

		if _, ok := rules[l]; !ok {
			rules[l] = collections.NewSet[int]()
		}

		rules[l].Add(r)
	}

	sum1, sum2 := 0, 0

	for _, line := range hunks[1] {
		nums := conv.ToInts(strings.Split(line, ","))

		mid, isOrdered := middleForLine(nums, rules)
		sum1 += mid

		if !isOrdered {
			sum2 += reorder(nums, rules)
		}
	}

	fmt.Println("part 1:", sum1)
	fmt.Println("part 2:", sum2)
}

func middleForLine(nums []int, rules map[int]collections.Set[int]) (int, bool) {
	seen := collections.NewSet[int]()

	for _, n := range nums {
		seen.Add(n)

		forbidden, ok := rules[n]
		if !ok {
			continue
		}

		for rule := range forbidden.Iter() {
			if seen.Contains(rule) {
				return 0, false
			}
		}
	}

	return nums[len(nums)/2], true
}

func reorder(nums []int, rules map[int]collections.Set[int]) int {
	slices.SortFunc(nums, func(a, b int) int {
		switch {
		case rules[a].Contains(b):
			return -1
		case rules[b].Contains(a):
			return 1
		default:
			return 0
		}
	})

	mid, isOrdered := middleForLine(nums, rules)
	assert.True(isOrdered, "line is ordered")

	return mid
}
