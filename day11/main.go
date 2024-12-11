package main

import (
	"fmt"
	"maps"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/mathx"
)

func main() {
	raw := strings.TrimSpace(input.New().Slurp())
	nums := make(map[int]int)

	for _, n := range strings.Fields(raw) {
		nums[conv.Atoi(n)]++
	}

	numSteps := 75

	for i := range numSteps {
		tmp := make(map[int]int, len(nums)*2)

		for n, count := range nums {
			s := fmt.Sprint(n)

			switch {
			case n == 0:
				tmp[1] += count

			case len(s)%2 == 0:
				mid := len(s) / 2
				tmp[conv.Atoi(s[0:mid])] += count
				tmp[conv.Atoi(s[mid:])] += count

			default:
				tmp[n*2024] += count
			}
		}

		nums = tmp

		if i == 24 {
			fmt.Println("part 1:", mathx.Sum(maps.Values(nums)))
		}
	}

	fmt.Println("part 2:", mathx.Sum(maps.Values(nums)))
}
