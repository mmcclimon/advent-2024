package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/mmcclimon/advent-2024/advent/assert"
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
			res := mul(m[2], m[3])
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

func mul(m, n string) int {
	a, err := strconv.Atoi(m)
	assert.Nil(err)

	b, err := strconv.Atoi(n)
	assert.Nil(err)

	return a * b
}
