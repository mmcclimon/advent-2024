package main

import (
	"fmt"
	"iter"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	next, stop := iter.Pull(input.New().Lines())
	defer stop()

	line, ok := next()
	assert.True(ok, "pulled line")

	towels := strings.Split(line, ", ")

	p1, p2 := 0, 0
	for line, ok := next(); ok; line, ok = next() {
		n := try(towels, line)
		p2 += n

		if n > 0 {
			p1++
		}
	}

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)
}

func try(towels []string, line string) int {
	res := []int{1}

	var waysHere int
	for i := 1; i <= len(line); i++ {
		waysHere = 0

		for _, towel := range towels {
			// If we're looking at "bbrbw", and towel "bw", and there are three
			// different ways to get to "bbr", then there are also 3 different ways
			// to get to "bbrbw".
			if strings.HasSuffix(line[:i], towel) {
				waysHere += res[i-len(towel)]
			}
		}

		res = append(res, waysHere)
	}

	return waysHere
}
