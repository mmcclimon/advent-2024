package main

import (
	"cmp"
	"fmt"
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

		if done < nItems {
			continue
		}

		dist := shortestPath(grid, size)

		if done == nItems {
			fmt.Println("part 1:", dist)
		}

		if dist == -1 {
			fmt.Println("part 2:", line)
			break
		}
	}
}

func shortestPath(grid collections.Set[xy], size int) int {
	// dijkstra, much more straightforward than day 16
	dist := make(map[xy]int, size*size)
	prev := make(map[xy]xy, size*size)

	q := collections.NewMinQueue(func(a, b xy) int {
		return cmp.Compare(dist[a], dist[b])
	})

	dist[xy{0, 0}] = 0
	q.Insert(xy{0, 0})

	for q.Len() > 0 {
		cur := q.ExtractMin()
		// fmt.Println("looking at", cur, dist[cur])

		if cur == (xy{size, size}) {
			return dist[cur]
		}

		// fmt.Println(dist[cur])
		for _, v := range neighbors(grid, size, cur) {
			alt := dist[cur] + 1
			if dist[v] == 0 || alt < dist[v] {
				dist[v] = alt
				q.Insert(v)
				prev[v] = cur
			}
		}
	}

	return -1
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
