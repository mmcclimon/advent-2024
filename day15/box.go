package main

import "github.com/mmcclimon/advent-2024/advent/operator"

type Box struct {
	// leftmost x position
	x, y int
	wide bool
}

func (b *Box) xy() xy { return xy{b.x, b.y} }

func (b *Box) CharFor(x int) string {
	if !b.wide {
		return "O"
	}

	return operator.CrummyTernary(x == b.x, "[", "]")
}

func (b *Box) CanMove(grid Map, dir rune) (result bool) {
	if !b.wide {
		return canMoveSingle(grid, b, dir)
	}

	switch dir {
	case '<':
		return canMoveSingle(grid, b, dir)

	case '>':
		next := grid[xyForDir(b.x+1, b.y, dir)]
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

func (b *Box) Move(grid Map, dir rune) {
	if !b.CanMove(grid, dir) {
		panic("must not call Move() without calling CanMove() first")
	}

	if !b.wide {
		next := moveSingle(grid, b, dir)
		b.x = next.x
		b.y = next.y
		return
	}

	var nextY int

	switch dir {
	case '<':
		next := xy{b.x - 1, b.y}
		if obj := grid[next]; obj != nil {
			obj.Move(grid, dir)
		}

		grid[xy{b.x + 1, b.y}] = nil
		grid[next] = b
		b.x--
		return

	case '>':
		next := xy{b.x + 2, b.y}
		if obj := grid[next]; obj != nil {
			obj.Move(grid, dir)
		}

		grid[xy{b.x, b.y}] = nil
		grid[next] = b
		b.x++
		return

	case '^':
		nextY = b.y - 1

	case 'v':
		nextY = b.y + 1

	default:
		panic("bad direction")
	}

	l := xy{b.x, nextY}
	if obj := grid[l]; obj != nil {
		obj.Move(grid, dir)
	}

	r := xy{b.x + 1, nextY}
	if obj := grid[r]; obj != nil {
		obj.Move(grid, dir)
	}

	grid[xy{b.x, b.y}] = nil
	grid[xy{b.x + 1, b.y}] = nil
	grid[xy{b.x, nextY}] = b
	grid[xy{b.x + 1, nextY}] = b
	b.y = nextY
}
