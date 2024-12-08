package main

import (
	"fmt"
	"regexp"

	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type rc struct {
	r, c int
}

func main() {
	grid := make(map[rc]string)
	ants := make(map[string][]rc)
	isAntenna := regexp.MustCompile(`[a-zA-Z0-9]`)

	for r, line := range input.New().EnumerateLines() {
		for c, letter := range line {
			l := string(letter)
			grid[rc{r, c}] = l

			if isAntenna.MatchString(string(letter)) {
				ants[l] = append(ants[l], rc{r, c})
			}
		}
	}

	part1(ants, grid)
	part2(ants, grid)
}

func part1(ants map[string][]rc, grid map[rc]string) {
	antinodes := collections.NewSet[rc]()

	for _, list := range ants {
		if len(list) < 2 {
			continue
		}

		for i := range list {
			for j := range list {
				if i == j {
					continue
				}

				a, b := list[i], list[j]

				rd := b.r - a.r
				cd := b.c - a.c

				antinode := rc{b.r + rd, b.c + cd}
				if _, ok := grid[antinode]; ok {
					antinodes.Add(antinode)
				}
			}
		}
	}

	fmt.Println("part 1:", len(antinodes))
}

func part2(ants map[string][]rc, grid map[rc]string) {
	antinodes := collections.NewSet[rc]()

	for _, list := range ants {
		if len(list) < 2 {
			continue
		}

		for i := range list {
			for j := range list {
				if i == j {
					continue
				}

				a, b := list[i], list[j]

				rd := b.r - a.r
				cd := b.c - a.c

				// antinodes.Add(rc{b.r, b.c})

				for scale := 0; true; scale++ {
					antinode := rc{b.r + (scale * rd), b.c + (scale * cd)}
					_, ok := grid[antinode]
					if !ok {
						break
					}

					antinodes.Add(antinode)
				}
			}
		}
	}

	fmt.Println("part 2:", len(antinodes))
}
