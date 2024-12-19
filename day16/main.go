package main

import (
	"cmp"
	"fmt"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type Direction string

const (
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
	North Direction = "N"
)

type rc struct {
	r, c int
}

func (pos rc) String() string {
	return fmt.Sprintf("%d,%d", pos.r, pos.c)
}

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

	findShortest(grid, start)
}

type node struct {
	pos rc
	dir Direction
}

func (n node) String() string {
	return fmt.Sprintf("%s/%s", n.pos, n.dir)
}

func findShortest(grid map[rc]rune, start rc) {
	dist := make(map[node]int, len(grid))
	prev := make(map[node][]node, len(grid))

	q := collections.NewMinQueue(func(a, b node) int {
		return cmp.Compare(dist[a], dist[b])
	})

	first := node{start, East}
	dist[first] = 0
	q.Insert(first)

	for q.Len() > 0 {
		cur := q.ExtractMin()

		if grid[cur.pos] == 'E' {
			fmt.Println("part 1:", dist[cur])
			part2(prev, first, cur)
			return
		}

		facing := cur.dir

		for _, v := range neighbors(grid, cur.pos, facing) {
			thisDir := dirFor(cur.pos, v)
			next := node{v, thisDir}

			thisDist := 1
			if thisDir != facing {
				thisDist += 1000
			}

			alt := dist[cur] + thisDist
			existingDist, ok := dist[next]
			if ok && alt > existingDist {
				continue
			}

			// If this is equal to the one we already know about, add it to the
			// list; if it's better (or we don't have one at all), make a new list.
			if alt == existingDist {
				prev[next] = append(prev[next], cur)
			} else {
				prev[next] = []node{cur}
			}

			dist[next] = alt
			q.Insert(next)
		}
	}
}

func neighbors(grid map[rc]rune, node rc, facing Direction) []rc {
	var ret []rc
	for _, pos := range []rc{
		{node.r + 1, node.c},
		{node.r - 1, node.c},
		{node.r, node.c + 1},
		{node.r, node.c - 1},
	} {
		char, ok := grid[pos]
		if char == '#' || !ok {
			continue
		}

		dir := dirFor(node, pos)
		if dir == facing.opposite() {
			continue
		}
		ret = append(ret, pos)
	}

	return ret
}

func (d Direction) opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic("unreachable")
	}
}

func dirFor(from, to rc) Direction {
	switch {
	case from == rc{0, 0}:
		// stupid special case
		return East
	case to == rc{from.r, from.c + 1}:
		return East
	case to == rc{from.r, from.c - 1}:
		return West
	case to == rc{from.r - 1, from.c}:
		return North
	case to == rc{from.r + 1, from.c}:
		return South
	default:
		fmt.Println("bad direction", from, to)
		panic("bad direction")
	}
}

func part2(prev map[node][]node, start, end node) {
	// This is a depth-first search on the previous graph, starting from the end
	// and working back to the beginning.
	s := collections.NewDeque[node]()
	seen := collections.NewSet[node]()

	s.Append(end)

	numPaths := 0

	for s.Len() > 0 {
		pos, err := s.Pop()
		assert.Nil(err)

		if pos == start {
			numPaths++
			seen.Add(pos)
			continue
		}

		if !seen.Contains(pos) {
			seen.Add(pos)
			for _, coords := range prev[pos] {
				s.Append(coords)
			}
		}
	}

	// We did the DFS on nodes, but we actually need to calculate the distinct
	// tiles.
	tiles := collections.NewSet[rc]()
	for node := range seen.Iter() {
		tiles.Add(node.pos)
	}

	fmt.Println("part 2:", len(tiles))
}
