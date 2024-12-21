package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strconv"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
)

type Keypad struct {
	coords    map[rune]xy
	coordsRev collections.Set[xy]
	memo      map[memoKey][]string
}

type xy struct{ x, y int }

type memoKey struct{ from, to rune }

var (
	test = []string{"029A", "980A", "179A", "456A", "379A"}
	//nolint:unused
	actual = []string{"169A", "279A", "540A", "869A", "789A"}
)

func main() {
	strs := test
	total := 0

	for _, want := range strs {
		dirs := inputString(want)

		r1 := findDirections(dirs)
		r2 := findDirections(r1)

		slices.SortFunc(r2, func(a, b string) int {
			return cmp.Compare(len(a), len(b))
		})

		n, err := strconv.ParseInt(want[:len(want)-1], 10, 64)
		assert.Nil(err)

		total += int(n) * len(r2[0])
	}

	fmt.Println(total)
}

func inputString(want string) []string {
	var all []string
	runes := []rune("A" + want)

	for i := range len(runes) - 1 {
		allDirs := numeric.directionsFor(runes[i], runes[i+1])

		if len(all) == 0 {
			all = allDirs
			continue
		}

		var tmp []string
		for _, dirs := range allDirs {
			for _, existing := range all {
				tmp = append(tmp, existing+dirs)
			}
		}

		all = tmp
	}

	return all
}

func findDirections(want []string) []string {
	var all []string
	for _, dirs := range want {
		r1 := inputDirections(dirs)
		all = append(all, r1...)
	}

	return all
}

func inputDirections(want string) []string {
	var all []string
	runes := []rune("A" + want)

	for i := range len(runes) - 1 {
		allDirs := directional.directionsFor(runes[i], runes[i+1])

		if len(all) == 0 {
			all = allDirs
			continue
		}

		var tmp []string
		for _, dirs := range allDirs {
			for _, existing := range all {
				tmp = append(tmp, existing+dirs)
			}
		}

		all = tmp
	}

	return all
}

/*
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+
*/

var numeric = NewKeypad(map[rune]xy{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'0': {1, 3},
	'A': {2, 3},
})

/*
+---+---+---+
|   | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/
var directional = NewKeypad(map[rune]xy{
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
})

func dirFor(from, to xy) string {
	switch {
	case to == xy{from.x - 1, from.y}:
		return "<"
	case to == xy{from.x + 1, from.y}:
		return ">"
	case to == xy{from.x, from.y - 1}:
		return "^"
	case to == xy{from.x, from.y + 1}:
		return "v"
	default:
		panic("unreachable")
	}
}

func NewKeypad(coords map[rune]xy) *Keypad {
	return &Keypad{
		coords:    coords,
		coordsRev: collections.NewSet(slices.Collect(maps.Values(coords))...),
		memo:      make(map[memoKey][]string),
	}
}

func (k *Keypad) directionsFor(from, to rune) []string {
	key := memoKey{from, to}
	if cached, ok := k.memo[key]; ok {
		return cached
	}

	// another day, another dijkstra
	dist := make(map[xy]int, len(k.coords))
	prev := make(map[xy][]xy, len(k.coords))

	q := collections.NewMinQueue(func(a, b xy) int {
		return cmp.Compare(dist[a], dist[b])
	})

	start := k.coords[from]

	dist[start] = 0
	q.Insert(start)

	for q.Len() > 0 {
		cur := q.ExtractMin()

		if k.coords[to] == cur {
			// fmt.Println("found", string(to), dist[cur])
			break
		}

		for _, v := range k.neighbors(cur) {
			alt := dist[cur] + 1
			existingDist, ok := dist[v]

			if ok && alt > existingDist {
				continue
			}

			dist[v] = alt
			q.Insert(v)

			// If this is equal to the one we already know about, add it to the
			// list; if it's better (or we don't have one at all), make a new list.
			if alt == existingDist {
				prev[v] = append(prev[v], cur)
			} else {
				prev[v] = []xy{cur}
			}
		}
	}

	paths := findPaths(prev, k.coords[from], k.coords[to])
	all := collections.NewSet[string]()

	for _, path := range paths {
		var dirs string
		for i := range len(path) - 1 {
			dirs += dirFor(path[i], path[i+1])
		}

		all.Add(dirs + "A")
	}

	ret := slices.Collect(all.Iter())
	k.memo[key] = ret
	return ret
}

func (k *Keypad) neighbors(pos xy) []xy {
	ret := make([]xy, 0, 4)

	for _, candidate := range []xy{
		{pos.x - 1, pos.y},
		{pos.x + 1, pos.y},
		{pos.x, pos.y - 1},
		{pos.x, pos.y + 1},
	} {
		if k.coordsRev.Contains(candidate) {
			ret = append(ret, candidate)
		}
	}

	return ret
}

func findPaths(prev map[xy][]xy, start, end xy) [][]xy {
	// bfs
	s := collections.NewDeque[[]xy]()
	var paths [][]xy

	s.Append([]xy{end})

	for s.Len() > 0 {
		path, err := s.PopLeft()
		assert.Nil(err)

		pos := path[len(path)-1]

		if pos == start {
			slices.Reverse(path)
			paths = append(paths, path)
			continue
		}

		for _, coords := range prev[pos] {
			next := make([]xy, len(path)+1)
			copy(next, path)
			next[len(path)] = coords
			s.Append(next)
		}
	}

	return paths
}
