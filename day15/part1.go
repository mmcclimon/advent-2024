package main

import (
	"fmt"
	"maps"
	"strings"
)

type Movable struct {
	x, y int
	char rune
}

type Wall struct {
	x, y int
}

type MapObject interface {
	fmt.Stringer
	Move(Map, rune) bool
	IsBox() bool
}

func part1(hunks [][]string) {
	grid := make(Map)
	var robot *Movable

	for y, line := range hunks[0] {
		for x, char := range line {
			var obj MapObject
			switch char {
			case '#':
				obj = &Wall{x, y}
			case '@':
				robot = &Movable{x, y, '@'}
				obj = robot
			case 'O':
				obj = &Movable{x, y, 'O'}
			}

			grid[xy{x, y}] = obj
		}
	}

	directions := strings.Join(hunks[1], "")

	done := 0
	for _, dir := range directions {
		robot.Move(grid, dir)
		done++
	}

	total := 0

	for obj := range maps.Values(grid) {
		if obj == nil || !obj.IsBox() {
			continue
		}

		box := obj.(*Movable)
		total += 100*box.y + box.x
	}

	// printGrid(grid)
	fmt.Println("part 1:", total)
}

func (m *Movable) String() string { return string(m.char) }
func (_ *Wall) String() string    { return "#" }

func (m *Movable) IsBox() bool { return m.char == 'O' }
func (_ *Wall) IsBox() bool    { return false }

func (m *Movable) Move(grid Map, dir rune) bool {
	var next xy
	switch dir {
	case '^':
		next = xy{m.x, m.y - 1}
	case 'v':
		next = xy{m.x, m.y + 1}
	case '<':
		next = xy{m.x - 1, m.y}
	case '>':
		next = xy{m.x + 1, m.y}
	default:
		panic("bad direction")
	}

	obj := grid[next]

	canMove := obj == nil || obj.Move(grid, dir)
	if !canMove {
		return false
	}

	grid[xy{m.x, m.y}] = nil
	grid[next] = m
	m.x = next.x
	m.y = next.y
	return true
}

func (_ *Wall) Move(grid Map, dir rune) bool {
	return false
}
