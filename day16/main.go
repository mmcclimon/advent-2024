package main

import (
	"fmt"
	"math"

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
	prev := make(map[rc]node, len(grid))

	q := collections.NewSet[node]()

	for pos, char := range grid {
		if char == '#' {
			continue
		}

		for _, dir := range []Direction{North, South, East, West} {
			n := node{pos, dir}
			q.Add(n)
			dist[n] = math.MaxInt
		}
	}

	dist[node{start, East}] = 0

	for len(q) > 0 {
		fmt.Println(len(q))
		var u node
		for pos := range q.Iter() {
			if u == (node{}) || dist[pos] < dist[u] {
				u = pos
			}
		}

		q.Delete(u)

		if grid[u.pos] == 'E' {
			fmt.Println("found it!")
			fmt.Println(dist[u])
			// spew.Dump(dist)
			// spew.Dump(prev)
			return
		}

		facing := u.dir

		for _, v := range neighbors(grid, u.pos) {
			thisDir := dirFor(u.pos, v)
			next := node{v, thisDir}

			if !q.Contains(next) {
				continue
			}

			thisDist := 1
			if thisDir != facing {
				thisDist += 1000

				// stupid
				if thisDir == North && facing == South ||
					thisDir == South && facing == North ||
					thisDir == East && facing == West ||
					thisDir == West && facing == East {
					thisDist += 1000
				}
			}

			// fmt.Printf("  facing %s, want to face %s to get to %+v\n", facing, thisDir, v)
			// fmt.Printf("  dist so far = %d, thisDist = %d\n", dist[u], thisDist)
			alt := dist[u] + thisDist
			if dist[next] == 0 || alt < dist[next] {
				dist[next] = alt
				prev[v] = u
			}
		}
	}
}

func neighbors(grid map[rc]rune, node rc) []rc {
	var ret []rc
	for _, pos := range []rc{
		{node.r + 1, node.c},
		{node.r - 1, node.c},
		{node.r, node.c + 1},
		{node.r, node.c - 1},
	} {
		char := grid[pos]
		if char == '.' || char == 'E' || char == 'S' {
			ret = append(ret, pos)
		}
	}

	return ret
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
