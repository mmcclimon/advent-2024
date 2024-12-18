package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type xy struct{ x, y int }

const (
	nItems = 1024
	size   = 70
)

func main() {
	grid := collections.NewSet[xy]()

	done := 0

	for line := range input.New().Lines() {
		pos := conv.ToInts(strings.Split(line, ","))
		grid.Add(xy{pos[0], pos[1]})

		done++
		if done == nItems {
			break
		}
	}

	part1(grid, size)
}

func part1(grid collections.Set[xy], size int) {
	// dijkstra, much more straightforward than day 16
	dist := make(map[xy]int, size*size)
	prev := make(map[xy]xy, size*size)

	q := collections.NewSet[xy]()

	for y := range size + 1 {
		for x := range size + 1 {
			pos := xy{x, y}
			q.Add(pos)
			dist[pos] = math.MaxInt
		}
	}

	dist[xy{0, 0}] = 0

	for len(q) > 0 {
		cur := q.Peek()
		for pos := range q.Iter() {
			if dist[pos] <= dist[cur] {
				cur = pos
			}
		}

		q.Delete(cur)

		if cur == (xy{size, size}) {
			fmt.Println("part 1:", dist[cur])
			return
		}

		for _, v := range neighbors(grid, size, cur) {
			if !q.Contains(v) {
				continue
			}

			alt := dist[cur] + 1
			if dist[v] == 0 || alt < dist[v] {
				dist[v] = alt
				prev[v] = cur
			}
		}
	}

}

func neighbors(grid collections.Set[xy], size int, start xy) []xy {
	var ret []xy
	for _, pos := range []xy{
		{start.x + 1, start.y},
		{start.x - 1, start.y},
		{start.x, start.y + 1},
		{start.x, start.y - 1},
	} {
		if pos.x < 0 ||
			pos.y < 0 ||
			pos.x > size ||
			pos.y > size ||
			grid.Contains(pos) {
			continue
		}

		ret = append(ret, pos)
	}

	return ret
}

func (pos xy) String() string {
	return fmt.Sprintf("%d,%d", pos.x, pos.y)
}
