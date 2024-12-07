package main

import (
	"fmt"
	"maps"

	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

type rc struct {
	r, c int
}

type rcd struct {
	r, c int
	d    Direction
}

type Guard struct {
	r, c int
	dir  Direction
	seen collections.Set[rc]
}

func main() {
	grid := make(map[rc]rune)
	var g1, g2 Guard

	for r, line := range input.New().EnumerateLines() {
		for c, letter := range line {
			if letter == '^' {
				g1 = NewGuard(r, c)
				g2 = NewGuard(r, c)
				letter = '.'
			}

			grid[rc{r, c}] = letter
		}
	}

	run(grid, g1)
	run2(grid, g2, g1.seen)
}

func run(grid map[rc]rune, guard Guard) {
	for guard.InMap(grid) {
		if !guard.Forward(grid) {
			guard.TurnRight()
		}
	}

	// subtract 1 because this includes the final position outside the map
	fmt.Println("part 1:", len(guard.seen)-1)
}

// This is stupid slow.
func run2(grid map[rc]rune, guard Guard, candidates collections.Set[rc]) {
	total := 0

COORD:
	for coord := range candidates.Iter() {
		g := NewGuard(guard.r, guard.c)
		dupe := maps.Clone(grid)
		dupe[coord] = '#'

		seen := collections.NewSet[rcd]()

		for g.InMap(dupe) {
			if !g.Forward(dupe) {
				g.TurnRight()
			}

			if seen.Contains(g.rcd()) {
				total++
				continue COORD
			}

			seen.Add(g.rcd())
		}
	}

	fmt.Println("part 2:", total-1)
}

func NewGuard(r, c int) Guard {
	g := Guard{r, c, N, collections.NewSet[rc]()}
	g.seen.Add(g.rc())
	return g
}

func (g *Guard) InMap(grid map[rc]rune) bool {
	_, ok := grid[g.rc()]
	return ok
}

func (g *Guard) rc() rc {
	return rc{g.r, g.c}
}

func (g *Guard) rcd() rcd {
	return rcd{g.r, g.c, g.dir}
}

func (g *Guard) Forward(grid map[rc]rune) bool {
	var pos rc
	switch g.dir {
	case N:
		pos = rc{g.r - 1, g.c}
	case S:
		pos = rc{g.r + 1, g.c}
	case E:
		pos = rc{g.r, g.c + 1}
	case W:
		pos = rc{g.r, g.c - 1}
	}

	if grid[pos] != '#' {
		g.r = pos.r
		g.c = pos.c
		g.seen.Add(g.rc())
		return true
	}

	return false
}

func (g *Guard) TurnRight() {
	switch g.dir {
	case N:
		g.dir = E
	case E:
		g.dir = S
	case S:
		g.dir = W
	case W:
		g.dir = N
	}
}

func (g Guard) WouldMakeLoop(grid map[rc]rune) bool {
	var vec rc
	switch g.dir {
	case N:
		vec = rc{0, 1} // s
	case S:
		vec = rc{0, -1} // n
	case E:
		vec = rc{-1, 0} // w
	case W:
		vec = rc{1, 0} // e
	}

	r := g.r + vec.r
	c := g.c + vec.c

	for char, ok := grid[rc{r, c}]; ok; char, ok = grid[rc{r, c}] {
		if char == '#' {
			// fmt.Println(r, c, g.seen.Contains(rc{r, c}))
			return true
		}

		r += vec.r
		c += vec.c
	}

	return false
}
