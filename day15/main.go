package main

import (
	"slices"

	"github.com/mmcclimon/advent-2024/advent/input"
)

type xy struct {
	x, y int
}

func main() {
	hunks := slices.Collect(input.New().Hunks())

	part1(hunks)
	part2(hunks)
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
