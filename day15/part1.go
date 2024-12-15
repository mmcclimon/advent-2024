package main

import (
	"fmt"
	"maps"
	"strings"
)

type Map map[xy]MapObject

type MapObject interface {
	fmt.Stringer
	Move(Map, rune) bool
	IsBox() bool
}

type Movable struct {
	x, y int
	char rune
}

// Wall doesn't need to store its own coordinates because it doesn't ever move.
type Wall struct{}

func part1(hunks [][]string) {
	grid := make(Map)
	var robot *Movable

	for y, line := range hunks[0] {
		for x, char := range line {
			var obj MapObject
			switch char {
			case '#':
				obj = &Wall{}
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
func (m *Movable) IsBox() bool    { return m.char == 'O' }

func (m *Movable) Move(grid Map, dir rune) bool {
	next := xyForDir(m.x, m.y, dir)
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

func (_ *Wall) String() string               { return "#" }
func (_ *Wall) IsBox() bool                  { return false }
func (_ *Wall) Move(grid Map, dir rune) bool { return false }

//nolint:unused
func printGrid(grid Map) {
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
				fmt.Print(obj)
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("")
}
