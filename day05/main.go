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
		n := numForLine(line, rules)
		sum1 += n

		if n == 0 {
			sum2 += reorder(line, rules)
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func numForLine(line string, rules map[int]collections.Set[int]) int {
	seen := collections.NewSet[int]()

	nums := strings.Split(line, ",")

	for _, num := range nums {
		n, err := strconv.Atoi(num)
		assert.Nil(err)

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

	middle := nums[len(nums)/2]
	n, _ := strconv.Atoi(middle)
	return n
}

func reorder(line string, rules map[int]collections.Set[int]) int {
	var nums []int
	for _, num := range strings.Split(line, ",") {
		n, err := strconv.Atoi(num)
		assert.Nil(err)
		nums = append(nums, n)
	}

	slices.SortFunc(nums, func(a, b int) int {
		aRules, ok := rules[a]
		if !ok {
			// if no a, then b should come first
			return 1
		}

		bRules, ok := rules[b]
		if !ok {
			// if no b, then a should come first
			return -1
		}

		switch {
		case aRules.Contains(b):
			return -1
		case bRules.Contains(a):
			return 1
		default:
			return 0
		}
	})

	return nums[len(nums)/2]
}
