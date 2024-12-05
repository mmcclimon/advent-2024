package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	hunks := slices.Collect(input.New().Hunks())

	rules := make(map[int]collections.Set[int])

	for _, line := range hunks[0] {
		bits := strings.Split(line, "|")
		l, err := strconv.Atoi(bits[0])
		assert.Nil(err)

		r, err := strconv.Atoi(bits[1])
		assert.Nil(err)

		if _, ok := rules[l]; !ok {
			rules[l] = collections.NewSet[int]()
		}

		rules[l].Add(r)
	}

	sum1, sum2 := 0, 0

	for _, line := range hunks[1] {
		n := numForLine(toInts(line), rules)
		sum1 += n

		if n == 0 {
			sum2 += reorder(toInts(line), rules)
		}
	}

	fmt.Println("part 1:", sum1)
	fmt.Println("part 2:", sum2)
}

func toInts(line string) []int {
	var nums []int

	for _, num := range strings.Split(line, ",") {
		n, err := strconv.Atoi(num)
		assert.Nil(err)
		nums = append(nums, n)
	}

	return nums
}

func numForLine(nums []int, rules map[int]collections.Set[int]) int {
	seen := collections.NewSet[int]()

	for _, n := range nums {
		seen.Add(n)

		forbidden, ok := rules[n]
		if !ok {
			continue
		}

		for rule := range forbidden.Iter() {
			if seen.Contains(rule) {
				return 0
			}
		}
	}

	return nums[len(nums)/2]
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

	return nums[len(nums)/2]
}
