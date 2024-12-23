package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type Graph = map[string][]string

func main() {
	graph := make(Graph)

	for line := range input.New().Lines() {
		bits := strings.Split(line, "-")
		l, r := bits[0], bits[1]
		graph[l] = append(graph[l], r)
		graph[r] = append(graph[r], l)
	}

	part1(graph)
	part2(graph)
}

func part1(graph Graph) {
	trios := collections.NewSet[Trio]()

	for k := range graph {
		// consider only Ts
		if k[0] != 't' {
			continue
		}

		trios.Extend(findTrios(graph, k))
	}

	fmt.Println("part 1:", len(trios))
}

type Trio struct{ a, b, c string }

func newTrio(a, b, c string) Trio {
	s := []string{a, b, c}
	slices.Sort(s)
	return Trio{s[0], s[1], s[2]}
}

// This is stupid but totally works.
func findTrios(graph Graph, start string) collections.Set[Trio] {
	seen := collections.NewSet[Trio]()

	for _, first := range graph[start] {
		for _, second := range graph[first] {
			if second == start {
				continue
			}

			for _, third := range graph[second] {
				if third == start {
					seen.Add(newTrio(first, second, third))
				}
			}
		}
	}

	return seen
}

func part2(graph Graph) {
	p := collections.NewSetFromIter(maps.Keys(graph))

	found := BK(graph, collections.NewSet[string](), p, collections.NewSet[string]())

	var longest []string

	for _, clique := range found {
		s := clique.ToSlice()
		if len(s) > len(longest) {
			longest = s
		}
	}

	slices.Sort(longest)
	fmt.Println("part 2:", strings.Join(longest, ","))

}

// This is the Bron-Kerbosch algorithm.
func BK(graph Graph, r, p, x collections.Set[string]) []collections.Set[string] {
	ret := make([]collections.Set[string], 0)

	if len(p) == 0 && len(x) == 0 {
		ret = append(ret, r)
	}

	for _, v := range slices.Collect(p.Iter()) {
		neighbors := collections.NewSet(graph[v]...)

		r2 := r.Clone()
		r2.Add(v)

		p2 := p.Intersection(neighbors)
		x2 := x.Intersection(neighbors)

		// recurse, recurse!
		ret = append(ret, BK(graph, r2, p2, x2)...)

		p.Delete(v)
		x.Add(v)
	}

	return ret
}
