package main

import (
	"fmt"
	"maps"

	"github.com/mmcclimon/advent-2024/advent/operator"
)

type Map map[xy]MapObject

type Grid struct {
	m     Map
	robot *Robot
	wide  bool
}

type MapObject interface {
	Move(Map, rune)
	CanMove(Map, rune) bool
	CharFor(x int) string
}

type xy struct {
	x, y int
}

func xyForDir(x, y int, dir rune) xy {
	switch dir {
	case '^':
		return xy{x, y - 1}
	case 'v':
		return xy{x, y + 1}
	case '<':
		return xy{x - 1, y}
	case '>':
		return xy{x + 1, y}
	default:
		panic("unreachable")
	}
}

func MakeGrid(in []string, wide bool) *Grid {
	grid := make(Map)
	var robot *Robot

	// Make the grid
	for y, line := range in {
		for x, char := range line {
			if wide {
				x *= 2
			}

			var obj MapObject

			switch char {
			case '#':
				obj = &Wall{}

			case '@':
				robot = &Robot{x, y}
				obj = robot

			case 'O':
				obj = &Box{x, y, wide}
			}

			grid[xy{x, y}] = obj

			if wide {
				right := operator.CrummyTernary(obj == robot, nil, obj)
				grid[xy{x + 1, y}] = right
			}
		}
	}

	return &Grid{grid, robot, wide}
}

func (g *Grid) Total() int {
	total := 0

	for obj := range maps.Values(g.m) {
		if box, ok := obj.(*Box); ok {
			total += 100*box.y + box.x
		}
	}

	if g.wide {
		// We always double-count boxes above
		total /= 2
	}

	return total
}

func (g *Grid) Print() {
	maxX, maxY := -1, -1
	for k := range g.m {
		maxX = max(maxX, k.x)
		maxY = max(maxY, k.y)
	}

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			obj := g.m[xy{x, y}]
			if obj == nil {
				fmt.Print(".")
			} else {
				fmt.Print(obj.CharFor(x))
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("")
}

type movable interface {
	MapObject
	xy() xy
}

func moveSingle(grid Map, m movable, dir rune) xy {
	pos := m.xy()
	next := xyForDir(pos.x, pos.y, dir)
	if obj := grid[next]; obj != nil {
		obj.Move(grid, dir)
	}

	grid[pos] = nil
	grid[next] = m
	return next
}

func canMoveSingle(grid Map, m movable, dir rune) bool {
	pos := m.xy()
	next := grid[xyForDir(pos.x, pos.y, dir)]
	return next == nil || next.CanMove(grid, dir)
}
