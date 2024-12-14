package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/mathx"
	"github.com/mmcclimon/advent-2024/advent/operator"
)

const (
	Width  = 101
	Height = 103
)

type Quadrant int

const (
	Mid Quadrant = iota
	NW
	NE
	SE
	SW
)

type Robot struct {
	x, y   int
	xv, yv int
}

func main() {
	robots1 := make([]*Robot, 0)
	robots2 := make([]*Robot, 0)

	for line := range input.New().Lines() {
		robots1 = append(robots1, makeRobot(line))
		robots2 = append(robots2, makeRobot(line))
	}

	part1(robots1)
	part2(robots2)
}

func part1(robots []*Robot) {
	for range 100 {
		for _, r := range robots {
			r.Tick()
		}
	}

	quads := make(map[Quadrant]int)
	for _, r := range robots {
		quads[r.Quadrant()]++
	}

	delete(quads, Mid)

	total := 1
	for _, v := range quads {
		total *= v
	}

	fmt.Println("part 1:", total)
}

func part2(robots []*Robot) {
	// The first version of this program printed out the grid after every tick.
	// On doing so, I noticed that every 101 ticks they started to coalesce, so
	// then I started checking those.

	// Specific to my input,
	nextStop := 29
	cycleLen := 101

	for i := 0; true; i++ {
		if i == nextStop {
			nextStop += cycleLen

			printGrid(robots)
			if wait(i) {
				fmt.Println("part 2:", i)
				break
			}
		}

		for _, r := range robots {
			r.Tick()
		}
	}
}

var extract = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func makeRobot(line string) *Robot {
	m := extract.FindStringSubmatch(line)
	return &Robot{
		x:  conv.Atoi(m[1]),
		y:  conv.Atoi(m[2]),
		xv: conv.Atoi(m[3]),
		yv: conv.Atoi(m[4]),
	}
}

func (r *Robot) Tick() {
	r.x = mathx.Mod(r.x+r.xv, Width)
	r.y = mathx.Mod(r.y+r.yv, Height)
}

func (r *Robot) Quadrant() Quadrant {
	xMid := Width / 2
	yMid := Height / 2

	switch {
	case r.x < xMid && r.y < yMid:
		return NW
	case r.x > xMid && r.y < yMid:
		return NE
	case r.x < xMid && r.y > yMid:
		return SW
	case r.x > xMid && r.y > yMid:
		return SE
	default:
		return Mid
	}
}

func printGrid(robots []*Robot) {
	type xy struct{ x, y int }

	grid := make(map[xy]int, len(robots))
	for _, r := range robots {
		grid[xy{r.x, r.y}]++
	}

	for y := range Height {
		for x := range Width {
			_, ok := grid[xy{x, y}]
			fmt.Print(operator.CrummyTernary(ok, "*", " "))
		}
		fmt.Print("\n")
	}
}

func wait(step int) bool {
	fmt.Printf("[step %d] done? [yn] ", step)

	reader := bufio.NewReader(os.Stdin)
	resp, err := reader.ReadString('\n')
	if err != nil {
		return true
	}

	return strings.TrimSpace(resp) == "y"
}
