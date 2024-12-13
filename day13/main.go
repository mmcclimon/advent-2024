package main

import (
	"fmt"
	"regexp"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type xy struct {
	x, y uint64
}

func main() {
	var sum1, sum2 uint64

	for hunk := range input.New().Hunks() {
		p1, p2 := processHunk(hunk)
		sum1 += p1
		sum2 += p2
	}

	fmt.Println("part 1:", sum1)
	fmt.Println("part 2:", sum2)
	fmt.Println("num computations", done)
}

var done = 0

func processHunk(hunk []string) (uint64, uint64) {
	a := processLine(hunk[0])
	b := processLine(hunk[1])
	target := processLine(hunk[2])

	part1 := scoreForHunk(a, b, target)

	target.x += 10_000_000_000_000
	target.y += 10_000_000_000_000
	// part2 := scoreForHunk(a, b, target)

	return part1, 0
}

func scoreForHunk(a, b, target xy) uint64 {
	var best uint64

	// This is the largest number of steps we'd ever have to take there using
	// just A or B presses.
	aMax := min(target.x/a.x, target.y/a.y)

	// obviously stupid
	for aTry := range aMax + 1 {
		ax := aTry * a.x
		ay := aTry * a.y

		// Determine the point to try B
		bStart := min(
			(target.x-ax)/b.x,
			(target.y-ay)/b.y,
		)

		for bTry := bStart; true; bTry++ {
			x := aTry*a.x + bTry*b.x
			y := aTry*a.y + bTry*b.y

			// fmt.Println(aTry, bTry)
			done++

			if x > target.x || y > target.y {
				break
			}

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
		x: uint64(conv.Atoi(m[1])),
		y: uint64(conv.Atoi(m[2])),
	}
}
