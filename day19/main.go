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

	total := 0
	for _, line := range hunks[1] {
		if re.MatchString(line) {
			total++
		}
	}

	fmt.Println(total)
}
