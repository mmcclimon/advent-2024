package main

import (
	"fmt"
	"slices"

	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/operator"
)

type xy struct {
	x, y int
}

type Map map[xy]MapObject

type MapObject interface {
	fmt.Stringer
	Move(Map, rune) bool
	IsBox() bool
}

func main() {
	hunks := slices.Collect(input.New().Hunks())

	part1(hunks)
	part2(hunks)
}

//nolint:unused
func printGrid(grid Map, wide bool) {
	maxX, maxY := -1, -1
	for k := range grid {
		maxX = max(maxX, k.x)
		maxY = max(maxY, k.y)
	}

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			obj := grid[xy{x, y}]
			if obj == nil {
				fmt.Print(operator.CrummyTernary(wide, "..", "."))
			} else {
				fmt.Print(obj)
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("")
}
