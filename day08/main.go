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

	findAntinodes(ants, grid)
}

func findAntinodes(ants map[string][]rc, grid map[rc]string) {
	p1 := collections.NewSet[rc]()
	p2 := collections.NewSet[rc]()

	for _, list := range ants {
		for i := range list {
			for j := range list {
				if i == j {
					continue
				}

				a, b := list[i], list[j]

				rd := b.r - a.r
				cd := b.c - a.c

				for scale := 0; true; scale++ {
					node := rc{b.r + (scale * rd), b.c + (scale * cd)}
					_, ok := grid[node]
					if !ok {
						break
					}

					if scale == 1 {
						p1.Add(node)
					}

					p2.Add(node)
				}
			}
		}
	}

	fmt.Println("part 1:", len(p1))
	fmt.Println("part 2:", len(p2))
}
