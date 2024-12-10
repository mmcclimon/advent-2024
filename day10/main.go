package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/operator"
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

	p1, p2 := 0, 0

	for k, v := range grid {
		if v != 0 {
			continue
		}

		p1 += scoreFor(grid, k, false)
		p2 += scoreFor(grid, k, true)
	}

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)
}

func scoreFor(grid map[rc]int, start rc, part2 bool) int {
	// dfs
	s := collections.NewDeque[rc]()
	seen := collections.NewSet[rc]()

	s.Append(start)

	numPaths := 0
	ends := collections.NewSet[rc]()

	for s.Len() > 0 {
		pos, err := s.Pop()
		assert.Nil(err)

		if grid[pos] == 9 {
			ends.Add(pos)
			numPaths++
			continue
		}

		if !seen.Contains(pos) {
			seen.Add(pos)
		}

		for _, coords := range neighbors(grid, pos) {
			s.Append(coords)
		}
	}

	return operator.CrummyTernary(part2, numPaths, len(ends))
}

func neighbors(grid map[rc]int, start rc) []rc {
	ret := make([]rc, 0, 4)

	for _, candidate := range []rc{
		{start.r - 1, start.c},
		{start.r + 1, start.c},
		{start.r, start.c - 1},
		{start.r, start.c + 1},
	} {
		n, ok := grid[candidate]
		if ok && n == grid[start]+1 {
			ret = append(ret, candidate)
		}
	}

	return ret
}
