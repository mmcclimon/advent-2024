package main

import (
	"fmt"
	"regexp"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/mathx"
)

type xy struct {
	x, y int64
}

func main() {
	var sum1, sum2 int64

	for hunk := range input.New().Hunks() {
		p1, p2 := processHunk(hunk)
		sum1 += p1
		sum2 += p2
	}

	fmt.Println("part 1:", sum1)
	fmt.Println("part 2:", sum2)
}

func processHunk(hunk []string) (int64, int64) {
	a := processLine(hunk[0])
	b := processLine(hunk[1])
	target := processLine(hunk[2])

	part1 := scoreForHunk(a, b, target)

	target.x += 10_000_000_000_000
	target.y += 10_000_000_000_000
	part2 := scoreForHunk(a, b, target)

	return part1, part2
}

// This is just: solve a system of two equations:
// aPresses * a.x + bPresses * b.x = target.x
// aPresses * a.y + bPresses * b.y = target.y
func scoreForHunk(a, b, target xy) int64 {
	lcm := mathx.LCM(a.x, a.y)

	xs := lcm / a.x
	ys := lcm / a.y

	// Multiply both sides by the scaling factors and solve for b, then a.
	bPresses := ((target.x * xs) - (target.y * ys)) / ((b.x * xs) - (b.y * ys))
	aPresses := (target.x - (b.x * bPresses)) / a.x

	// Double-check here to make sure the math works out because we're doing
	// integer division abote.
	if aPresses*a.x+bPresses*b.x != target.x ||
		aPresses*a.y+bPresses*b.y != target.y {
		return 0
	}

	return aPresses*3 + bPresses
}

var extracter = regexp.MustCompile(`X.(\d+), Y.(\d+)$`)

func processLine(line string) xy {
	m := extracter.FindStringSubmatch(line)
	return xy{
		x: int64(conv.Atoi(m[1])),
		y: int64(conv.Atoi(m[2])),
	}
}
