package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"

	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/mathx"
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
	part1(grid, dist)
	part2(grid, dist)
}

func makeDist(grid map[rc]rune, start rc) map[rc]int {
	dist := make(map[rc]int, len(grid))
	dist[start] = 0

	pos := start

	for grid[pos] != 'E' {
		for _, candidate := range neighbors(grid, pos) {
			next := grid[candidate]
			_, ok := dist[candidate]
			if next == '#' || ok {
				continue
			}

			dist[candidate] = dist[pos] + 1
			pos = candidate
			break
		}
	}

	return dist
}

func part1(grid map[rc]rune, dist map[rc]int) {
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

func part2(grid map[rc]rune, path map[rc]int) {
	shortcuts := make(map[int]int)

	for pos, thisDist := range path {
		if grid[pos] == 'E' {
			continue
		}

		dist := make(map[rc]int, len(grid))

		q := collections.NewMinQueue(func(a, b rc) int {
			return cmp.Compare(dist[a], dist[b])
		})

		dist[pos] = 0
		q.Insert(pos)

		for q.Len() > 0 {
			cur := q.ExtractMin()

			if grid[cur] == 'E' {
				continue
			}

			// We do 25 here because the dijkstra distance isn't exactly the thing
			// we want (apparently).
			for _, v := range neighbors(grid, cur) {
				alt := dist[cur] + 1
				_, haveDist := dist[v]
				if alt <= 25 && (!haveDist || alt < dist[v]) {
					dist[v] = alt
					q.Insert(v)
				}
			}
		}

		for finish := range dist {
			haveDist, onPath := path[finish]
			if !onPath {
				continue
			}

			// I don't know why the taxicab distance is sometimes not the same as
			// the dijkstra gistance, but the taxicab math makes the example work
			// out correctly.
			newDist := mathx.Abs(pos.r-finish.r) + mathx.Abs(pos.c-finish.c)
			if newDist > 20 {
				continue
			}

			alt := thisDist + newDist

			if alt < haveDist {
				shortcuts[haveDist-alt]++
			}
		}
	}

	total := 0
	for _, k := range slices.Sorted(maps.Keys(shortcuts)) {
		if k < 100 {
			continue
		}

		total += shortcuts[k]
	}

	fmt.Println("part 2:", total)
}

func neighbors(grid map[rc]rune, pos rc) []rc {
	ret := make([]rc, 0, 4)

	for _, candidate := range []rc{
		{pos.r - 1, pos.c},
		{pos.r + 1, pos.c},
		{pos.r, pos.c - 1},
		{pos.r, pos.c + 1},
	} {
		_, ok := grid[candidate]
		if ok {
			ret = append(ret, candidate)
		}
	}

	return ret
}
