package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/input"
)

type Map map[xy]MapObject

type xy struct {
	x, y int
}

type Box struct {
	x, y int
}

type Robot struct {
	x, y int
}

type Wall struct {
	x, y int
}

type MapObject interface {
	fmt.Stringer
	Move(Map, rune) bool
}

func main() {
	hunks := slices.Collect(input.New().Hunks())
	grid := make(Map)
	var robot *Robot

	for y, line := range hunks[0] {
		for x, char := range line {
			var obj MapObject
			switch char {
			case '#':
				obj = &Wall{x, y}
			case '@':
				robot = &Robot{x, y}
				obj = robot
			case 'O':
				obj = &Box{x, y}
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

	printGrid(grid)

	total := 0

	for obj := range maps.Values(grid) {
		box, ok := obj.(*Box)
		if !ok {
			continue
		}

		total += 100*box.y + box.x
	}

	fmt.Println(total)

}

func (_ *Box) String() string   { return "O" }
func (_ *Robot) String() string { return "@" }
func (_ *Wall) String() string  { return "#" }

func (r *Robot) Move(grid Map, dir rune) bool {
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

	obj := grid[next]

	canMove := obj == nil || obj.Move(grid, dir)
	if !canMove {
		return false
	}

	grid[xy{r.x, r.y}] = nil
	grid[next] = r
	r.x = next.x
	r.y = next.y
	return true
}

func (b *Box) Move(grid Map, dir rune) bool {
	// TODO reduce copypasta
	var next xy
	switch dir {
	case '^':
		next = xy{b.x, b.y - 1}
	case 'v':
		next = xy{b.x, b.y + 1}
	case '<':
		next = xy{b.x - 1, b.y}
	case '>':
		next = xy{b.x + 1, b.y}
	default:
		panic("bad direction")
	}

	obj := grid[next]

	canMove := obj == nil || obj.Move(grid, dir)
	if !canMove {
		return false
	}

	grid[xy{b.x, b.y}] = nil
	grid[next] = b
	b.x = next.x
	b.y = next.y
	return true
}

func (_ *Wall) Move(grid Map, dir rune) bool {
	return false
}

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
