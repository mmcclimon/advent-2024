package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type rc struct {
	r, c int
}

func main() {
	grid := make(map[rc]int)

	for r, line := range input.New().EnumerateLines() {
		for c, num := range line {
			if !('0' <= num && num <= '9') {
				continue
			}

			n := conv.Atoi(string(num))
			grid[rc{r, c}] = n
		}
	}

	part1(grid)
	part2(grid)
}

func part1(grid map[rc]int) {
	sum := 0

	for k, v := range grid {
		if v != 0 {
			continue
		}

		sum += scoreFor(grid, k, false)
	}

	fmt.Println("part 1:", sum)
}

func part2(grid map[rc]int) {
	sum := 0

	for k, v := range grid {
		if v != 0 {
			continue
		}

		sum += scoreFor(grid, k, true)
	}

	fmt.Println("part 2:", sum)
}

func scoreFor(grid map[rc]int, start rc, part2 bool) int {
	// bfs
	q := collections.NewDeque[rc]()
	seen := collections.NewSet[rc]()

	q.Append(start)

	score := 0

	for q.Len() > 0 {
		pos, err := q.PopLeft()
		assert.Nil(err)

		n := grid[pos]
		if n == 9 {
			score++
			continue
		}

		for _, coords := range neighbors(grid, pos) {
			if grid[coords] != n+1 || seen.Contains(coords) {
				continue
			}

			// fmt.Printf("looking at %+v, num=%d\n", coords, grid[coords])

			seen.Add(coords)
			q.Append(coords)
		}
	}

	return score
}

func neighbors(grid map[rc]int, start rc) []rc {
	ret := make([]rc, 0, 4)

	for _, candidate := range []rc{
		{start.r - 1, start.c},
		{start.r + 1, start.c},
		{start.r, start.c - 1},
		{start.r, start.c + 1},
	} {
		_, ok := grid[candidate]
		if ok {
			ret = append(ret, candidate)
		}
	}

	return ret
}
