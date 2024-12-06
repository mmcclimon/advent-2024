package main

import (
	"fmt"
	"regexp"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	in := input.New().Slurp()

	re := regexp.MustCompile(`(do|don't|mul)\((?:(\d+),(\d+))?\)`)
	do := true
	sum1, sum2 := 0, 0

	for _, m := range re.FindAllStringSubmatch(in, -1) {
		switch m[1] {
		case "do", "don't":
			do = m[1] == "do"

		case "mul":
			res := conv.Atoi(m[2]) * conv.Atoi(m[3])
			sum1 += res
			if do {
				sum2 += res
			}

		default:
			panic(m)
		}
	}

	fmt.Println("part 1:", sum1)
	fmt.Println("part 2:", sum2)
}
