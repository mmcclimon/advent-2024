package main

import (
	"fmt"
	"maps"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/operator"
)

type Map2 map[xy]MapObject2

type MapObject2 interface {
	Move(Map2, rune)
	CanMove(Map2, rune) bool
	CharFor(x, y int) string
}

type Wall2 struct {
	// leftmost x position
	x, y int
}

type Box2 struct {
	// leftmost x position
	x, y int
}

type Robot2 struct {
	x, y int
}

func part2(hunks [][]string) {
	directions := strings.Join(hunks[1], "")

	grid := make(Map2)
	var robot *Robot2

	for y, line := range hunks[0] {
		for x, char := range line {
			x *= 2

			switch char {
			case '#':
				wall := &Wall2{x, y}
				grid[xy{x, y}] = wall
				grid[xy{x + 1, y}] = wall

			case '@':
				robot = &Robot2{x, y}
				grid[xy{x, y}] = robot
				grid[xy{x + 1, y}] = nil

			case 'O':
				box := &Box2{x, y}
				grid[xy{x, y}] = box
				grid[xy{x + 1, y}] = box
			}
		}
	}

	// printGrid2(grid)

	for _, dir := range directions {
		robot.Move(grid, dir)
	}

	total := 0

	for obj := range maps.Values(grid) {
		box, ok := obj.(*Box2)
		if !ok {
			continue
		}

		total += 100*box.y + box.x
	}

	// we always double-count boxes in the above
	fmt.Println("part 2:", total/2)
}

func (_ *Robot2) CharFor(_, _ int) string { return "@" }
func (b *Box2) CharFor(x, _ int) string {
	return operator.CrummyTernary(x == b.x, "[", "]")
}

func (r *Robot2) CanMove(grid Map2, dir rune) bool {
	var next xy
	switch dir {
	case '^':
		next = xy{r.x, r.y - 1}
	case 'v':
		next = xy{r.x, r.y + 1}
	case '<':
		next = xy{r.x - 1, r.y}
	case '>':
		next = xy{r.x + 1, r.y}
	}

	obj := grid[next]

	return obj == nil || obj.CanMove(grid, dir)
}

func (b *Box2) CanMove(grid Map2, dir rune) (result bool) {
	switch dir {
	case '<':
		next := grid[xy{b.x - 1, b.y}]
		return next == nil || next.CanMove(grid, dir)

	case '>':
		next := grid[xy{b.x + 2, b.y}]
		return next == nil || next.CanMove(grid, dir)

	case '^':
		ul := grid[xy{b.x, b.y - 1}]
		ur := grid[xy{b.x + 1, b.y - 1}]
		return (ul == nil || ul.CanMove(grid, dir)) &&
			(ur == nil || ur.CanMove(grid, dir))

	case 'v':
		dl := grid[xy{b.x, b.y + 1}]
		dr := grid[xy{b.x + 1, b.y + 1}]
		return (dl == nil || dl.CanMove(grid, dir)) &&
			(dr == nil || dr.CanMove(grid, dir))

	default:
		panic("unreachable")
	}
}

func (r *Robot2) Move(grid Map2, dir rune) {
	if !r.CanMove(grid, dir) {
		return
	}

	var next xy
	switch dir {
	case '^':
		next = xy{r.x, r.y - 1}
	case 'v':
		next = xy{r.x, r.y + 1}
	case '<':
		next = xy{r.x - 1, r.y}
	case '>':
		next = xy{r.x + 1, r.y}
	default:
		panic("bad direction")
	}

	if obj := grid[next]; obj != nil {
		obj.Move(grid, dir)
	}

	grid[xy{r.x, r.y}] = nil
	grid[next] = r
	r.x = next.x
	r.y = next.y
}

func (b *Box2) Move(grid Map2, dir rune) {
	if !b.CanMove(grid, dir) {
		fmt.Printf("cannot move box at %+v to %c\n", b, dir)
		return
	}

	switch dir {
	case '<':
		next := xy{b.x - 1, b.y}
		if obj := grid[next]; obj != nil {
			obj.Move(grid, dir)
		}

		grid[xy{b.x + 1, b.y}] = nil
		grid[next] = b
		b.x--

	case '>':
		next := xy{b.x + 2, b.y}
		if obj := grid[next]; obj != nil {
			obj.Move(grid, dir)
		}

		grid[xy{b.x, b.y}] = nil
		grid[next] = b
		b.x++

	case '^':
		// We know, now, that we _can_ move anything in the way.
		l := xy{b.x, b.y - 1}
		if obj := grid[l]; obj != nil {
			obj.Move(grid, dir)
		}

		r := xy{b.x + 1, b.y - 1}
		if obj := grid[r]; obj != nil {
			obj.Move(grid, dir)
		}

		grid[xy{b.x, b.y}] = nil
		grid[xy{b.x + 1, b.y}] = nil
		grid[xy{b.x, b.y - 1}] = b
		grid[xy{b.x + 1, b.y - 1}] = b
		b.y--

	case 'v':
		// We know, now, that we _can_ move anything in the way.
		l := xy{b.x, b.y + 1}
		if obj := grid[l]; obj != nil {
			obj.Move(grid, dir)
		}

		r := xy{b.x + 1, b.y + 1}
		if obj := grid[r]; obj != nil {
			obj.Move(grid, dir)
		}

		grid[xy{b.x, b.y}] = nil
		grid[xy{b.x + 1, b.y}] = nil
		grid[xy{b.x, b.y + 1}] = b
		grid[xy{b.x + 1, b.y + 1}] = b
		b.y++

	default:
		panic("unimplemented")

	}

}

func (_ *Wall2) CharFor(_, _ int) string  { return "#" }
func (_ *Wall2) CanMove(Map2, rune) bool  { return false }
func (_ *Wall2) Move(grid Map2, dir rune) {}

//nolint:unused
func printGrid2(grid Map2) {
	maxX, maxY := -1, -1
	for k := range grid {
		maxX = max(maxX, k.x)
		maxY = max(maxY, k.y)
	}

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			obj := grid[xy{x, y}]
			if obj == nil {
				fmt.Print(".")
			} else {
				fmt.Print(obj.CharFor(x, y))
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("")
}
