package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	hunks := slices.Collect(input.New().Hunks())

	towels := strings.Split(hunks[0][0], ", ")
	re := regexp.MustCompile("^(?:" + strings.Join(towels, "|") + ")+$")

	p1, p2 := 0, 0
	for _, line := range hunks[1] {
		if !re.MatchString(line) {
			continue
		}

		p1++
		p2 += try(towels, line)
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
