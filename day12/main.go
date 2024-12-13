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
	grid := make(map[rc]string) // pos => letter
	plots := make(map[rc]int)   // pos => plot id

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

		// find all connected plots
		for item := range search(grid, pos) {
			plots[item] = id
		}

		id++
	}

	part1(grid, plots)
	part2(plots)
}

func part1(grid map[rc]string, plots map[rc]int) {
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

func part2(plots map[rc]int) {
	// First, key by id.
	byID := make(map[int][]rc)

	for pos, id := range plots {
		byID[id] = append(byID[id], pos)
	}

	total := 0

	// SURELY there is a better way to do this, but.
	// For every distinct group...
	for id, positions := range byID {
		// Make a set of the corners, keyed by upper-left coordinate
		corners := collections.NewSet[rc]()

		for _, pos := range positions {
			var (
				up       = rc{pos.r - 1, pos.c}
				down     = rc{pos.r + 1, pos.c}
				left     = rc{pos.r, pos.c - 1}
				right    = rc{pos.r, pos.c + 1}
				hasUp    = plots[up] == id
				hasDown  = plots[down] == id
				hasLeft  = plots[left] == id
				hasRight = plots[right] == id
			)

			if !hasLeft && !hasUp {
				corners.Add(rc{pos.r, pos.c}) // outside corner, top left
			}

			if !hasUp && !hasRight {
				corners.Add(rc{pos.r, pos.c + 1}) // outside corner, top right
			}

			if !hasRight && !hasDown {
				corners.Add(rc{pos.r + 1, pos.c + 1}) // outside corner, bottom right
			}

			if !hasDown && !hasLeft {
				corners.Add(rc{pos.r + 1, pos.c}) // outside corner, bottom left
			}

			// Now, we also need to check for inside corners. We do this by finding
			// all the points outside this polygon, and checking if there's an
			// inside corner.
			for _, neighb := range []rc{up, down, left, right} {
				if plots[neighb] == id {
					continue // this is part of the shape, we don't care
				}

				var (
					u = plots[rc{neighb.r - 1, neighb.c}] == id
					d = plots[rc{neighb.r + 1, neighb.c}] == id
					l = plots[rc{neighb.r, neighb.c - 1}] == id
					r = plots[rc{neighb.r, neighb.c + 1}] == id
				)

				if u && l {
					corners.Add(rc{neighb.r, neighb.c}) // inside corner, top left
				}

				if u && r {
					corners.Add(rc{neighb.r, neighb.c + 1}) // inside corner, top right
				}

				if r && d {
					corners.Add(rc{neighb.r + 1, neighb.c + 1}) // inside corner, bottom right
				}

				if d && l {
					corners.Add(rc{neighb.r + 1, neighb.c}) // inside corner, bottom left
				}

			}
		}

		numCorners := len(corners)

		// Now, we have to double count inside corners, maybe.
		for corner := range corners.Iter() {
			tr := plots[rc{corner.r - 1, corner.c}]
			tl := plots[rc{corner.r - 1, corner.c - 1}]
			br := plots[rc{corner.r, corner.c}]
			bl := plots[rc{corner.r, corner.c - 1}]

			// This is if (ne/sw diagonal || nw/se diagonal)
			if (tr == id && tr == bl && tr != tl && tr != br) ||
				(tl == id && tl == br && tl != tr && tl != bl) {
				numCorners++
			}
		}

		total += len(positions) * numCorners
	}

	fmt.Println("part 2:", total)
}

// This is just a breadth-first search; returns a set of all the coordinates
// in the starting plot (including the start)
func search(grid map[rc]string, start rc) collections.Set[rc] {
	q := collections.NewDeque[rc]()
	seen := collections.NewSet[rc]()

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

	return seen
}

// Returns a slice of connected neighbors (same letter in the grid)
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
