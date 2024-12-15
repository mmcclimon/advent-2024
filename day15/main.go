package main

import (
	"fmt"
	"slices"

	"github.com/mmcclimon/advent-2024/advent/input"
)

type Map map[xy]MapObject

type xy struct {
	x, y int
}

func main() {
	hunks := slices.Collect(input.New().Hunks())

	part1(hunks)
}

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
