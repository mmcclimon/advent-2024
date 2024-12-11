package main

import (
	"fmt"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	raw := strings.TrimSpace(input.New().Slurp())

	nums := conv.ToInts(strings.Fields(raw))

	numSteps := 25

	for i := range numSteps {
		fmt.Println(i)
		tmp := make([]int, 0, len(nums)*2)

		for _, n := range nums {
			s := fmt.Sprint(n)

			switch {
			case n == 0:
				tmp = append(tmp, 1)

			case len(s)%2 == 0:
				mid := len(s) / 2
				tmp = append(tmp, conv.Atoi(s[0:mid]), conv.Atoi(s[mid:]))

			default:
				tmp = append(tmp, n*2024)
			}
		}

		nums = tmp
		// fmt.Println(nums)
	}

	fmt.Println(len(nums))
}
