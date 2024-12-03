package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/input"
)

var (
	re1 = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	re2 = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
)

func main() {
	sum1, sum2 := 0, 0

	for line := range input.New().Lines() {
		sum1 += part1(line)
		sum2 += part2(line)
	}

	fmt.Println("part 1:", sum1)
	fmt.Println("part 2:", sum2)
}

func part1(line string) int {
	matches := re1.FindAllStringSubmatch(line, -1)
	sum := 0

	for _, m := range matches {
		sum += mul(m[1], m[2])
	}

	return sum
}

var enabled = true

func part2(line string) int {
	matches := re2.FindAllStringSubmatch(line, -1)
	sum := 0

	for _, m := range matches {
		switch {
		case strings.HasPrefix(m[0], "don't"):
			enabled = false
		case strings.HasPrefix(m[0], "do"):
			enabled = true
		case strings.HasPrefix(m[0], "mul"):
			if enabled {
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
