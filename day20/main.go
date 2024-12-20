package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2024/advent/input"
)

type rc struct{ r, c int }

func main() {
	grid := make(map[rc]rune)
	var start rc

	for r, line := range input.New().EnumerateLines() {
		for c, char := range line {
			grid[rc{r, c}] = char
			if char == 'S' {
				start = rc{r, c}
			}
		}
	}

	dist := makeDist(grid, start)
	findShortcuts(grid, dist)
}

func makeDist(grid map[rc]rune, start rc) map[rc]int {
	dist := make(map[rc]int, len(grid))
	dist[start] = 0

	pos := start

	for grid[pos] != 'E' {
		for _, candidate := range []rc{
			{pos.r - 1, pos.c},
			{pos.r + 1, pos.c},
			{pos.r, pos.c - 1},
			{pos.r, pos.c + 1},
		} {

			next, ok := grid[candidate]
			if !ok || next == '#' {
				continue
			}

			if dist[candidate] > 0 {
				continue
			}

			dist[candidate] = dist[pos] + 1
			pos = candidate
			break
		}
	}

	return dist
}

func findShortcuts(grid map[rc]rune, dist map[rc]int) {
	shortcuts := make(map[int]int)

	for pos, thisDist := range dist {
		if grid[pos] == 'E' {
			continue
		}

		for _, vec := range []rc{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			candidate := rc{pos.r + vec.r, pos.c + vec.c}
			if grid[candidate] != '#' {
				continue
			}

			finish := dist[rc{pos.r + 2*vec.r, pos.c + 2*vec.c}]
			if finish == 0 || finish < thisDist {
				continue
			}

			sl := finish - thisDist - 2
			shortcuts[sl]++
		}
	}

	total := 0
	for k, v := range shortcuts {
		if k >= 100 {
			total += v
		}
	}

	fmt.Println("part 1:", total)
}
