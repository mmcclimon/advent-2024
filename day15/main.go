package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	hunks := slices.Collect(input.New().Hunks())

	fmt.Println("part 1:", run(hunks, false))
	fmt.Println("part 2:", run(hunks, true))
}

func run(hunks [][]string, wide bool) int {
	grid := MakeGrid(hunks[0], wide)

	// Run the thing.
	directions := strings.Join(hunks[1], "")
	for _, dir := range directions {
		grid.robot.Move(grid.m, dir)
	}

	// grid.Print()

	return grid.Total()
}
