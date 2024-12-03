package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	in := input.New().Slurp()
	fmt.Println("part 1:", part1(in))
	fmt.Println("part 2:", part2(in))
}

func part1(line string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	sum := 0

	for _, m := range re.FindAllStringSubmatch(line, -1) {
		sum += mul(m[1], m[2])
	}

	return sum
}

func part2(line string) int {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	sum := 0
	do := true

	for _, m := range re.FindAllStringSubmatch(line, -1) {
		switch {
		case strings.HasPrefix(m[0], "don't"):
			do = false
		case strings.HasPrefix(m[0], "do"):
			do = true
		case strings.HasPrefix(m[0], "mul"):
			if do {
				sum += mul(m[1], m[2])
			}
		default:
			panic(m)
		}
	}

	return sum
}

func mul(m, n string) int {
	a, err := strconv.Atoi(m)
	assert.Nil(err)

	b, err := strconv.Atoi(n)
	assert.Nil(err)

	return a * b
}
