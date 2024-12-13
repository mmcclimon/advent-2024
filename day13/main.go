package main

import (
	"fmt"
	"regexp"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type xy struct {
	x, y int
}

func main() {
	sum := 0

	for hunk := range input.New().Hunks() {
		sum += processHunk(hunk)
	}

	fmt.Println("part 1:", sum)
}

func processHunk(hunk []string) int {
	a := processLine(hunk[0])
	b := processLine(hunk[1])
	target := processLine(hunk[2])

	best := 0

	// obviously stupid
	for aTry := range 101 {
		ax := aTry * a.x
		ay := aTry * a.y
		if ax > target.x || ay > target.y {
			break
		}

		for bTry := range 101 {
			x := aTry*a.x + bTry*b.x
			y := aTry*a.y + bTry*b.y

			if x == target.x && y == target.y {
				cost := aTry*3 + bTry
				if best == 0 || cost < best {
					best = cost
				}
			}
		}
	}

	return best
}

var extracter = regexp.MustCompile(`X.(\d+), Y.(\d+)$`)

func processLine(line string) xy {
	m := extracter.FindStringSubmatch(line)
	return xy{
		x: conv.Atoi(m[1]),
		y: conv.Atoi(m[2]),
	}
}
