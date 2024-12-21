package main

import (
	"cmp"
	"fmt"
	"sync/atomic"

	"golang.org/x/sync/errgroup"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
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
	fmt.Println("part 1:", findShortcuts(grid, dist, 2))
	fmt.Println("part 2:", findShortcuts(grid, dist, 20))
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

func findShortcuts(grid map[rc]rune, path map[rc]int, shortcutLen int) int64 {
	var eg errgroup.Group
	eg.SetLimit(32)

	var total atomic.Int64

	for start, thisDist := range path {
		if grid[start] == 'E' {
			continue
		}

		eg.Go(func() error {
			// another day, another dijkstra
			dist := make(map[rc]int, shortcutLen*shortcutLen)
			q := collections.NewMinQueue(func(a, b rc) int {
				return cmp.Compare(dist[a], dist[b])
			})

			dist[start] = 0
			q.Insert(start)

			for q.Len() > 0 {
				cur := q.ExtractMin()

				for _, v := range neighbors(grid, cur) {
					alt := dist[cur] + 1
					_, haveDist := dist[v]

					if alt <= shortcutLen && (!haveDist || alt < dist[v]) {
						dist[v] = alt
						q.Insert(v)
					}
				}
			}

			// now, check all the new distances and sum up the ones where the shortcut
			// is at least 100.
			for finish, newDist := range dist {
				haveDist := path[finish]
				alt := thisDist + newDist

				if alt < haveDist && haveDist-alt >= 100 {
					total.Add(1)
				}
			}

			return nil
		})
	}

	err := eg.Wait()
	assert.Nil(err)

	return total.Load()
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
