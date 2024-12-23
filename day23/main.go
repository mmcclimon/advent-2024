package main

import (
	"fmt"
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

	trios := collections.NewSet[Trio]()

	// consider only Ts
	for k := range graph {
		if k[0] != 't' {
			continue
		}

		trios.Extend(find(graph, k))
	}

	fmt.Println(len(trios))
}

type Trio struct{ a, b, c string }

func newTrio(a, b, c string) Trio {
	s := []string{a, b, c}
	slices.Sort(s)
	return Trio{s[0], s[1], s[2]}
}

// This is stupid but totally works.
func find(graph Graph, start string) collections.Set[Trio] {
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
