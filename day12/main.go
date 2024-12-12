package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type rc struct {
	r, c int
}

type stats struct {
	area, perimeter int
}

func main() {
	grid := make(map[rc]string)
	plots := make(map[rc]int)

	for r, line := range input.New().EnumerateLines() {
		for c, char := range line {
			grid[rc{r, c}] = string(char)
		}
	}

	id := 1

	for pos := range grid {
		if plots[pos] != 0 {
			continue // already seen this one
		}

		for item := range search(grid, pos) {
			plots[item] = id
		}

		id++
	}

	calc := make(map[int]stats)

	for pos, id := range plots {
		st := calc[id]

		st.area++
		st.perimeter += 4 - len(neighbors(grid, pos))

		calc[id] = st
	}

	total := 0
	for _, st := range calc {
		total += st.area * st.perimeter
	}

	fmt.Println("part 1:", total)
}

func search(grid map[rc]string, start rc) collections.Set[rc] {
	// dfs
	q := collections.NewDeque[rc]()
	seen := collections.NewSet[rc]()

	// fmt.Println("looking at", grid[start], start)

	q.Append(start)
	seen.Add(start)

	for q.Len() > 0 {
		pos, err := q.PopLeft()
		assert.Nil(err)

		for _, neighbor := range neighbors(grid, pos) {
			if seen.Contains(neighbor) {
				continue
			}

			seen.Add(neighbor)
			q.Append(neighbor)
		}
	}

	// fmt.Println("found", slices.Collect(maps.Keys(seen)))

	return seen
}

func neighbors(grid map[rc]string, start rc) []rc {
	ret := make([]rc, 0, 4)

	for _, candidate := range []rc{
		{start.r - 1, start.c},
		{start.r + 1, start.c},
		{start.r, start.c - 1},
		{start.r, start.c + 1},
	} {
		char, ok := grid[candidate]

		if ok && char == grid[start] {
			ret = append(ret, candidate)
		}
	}

	return ret
}
