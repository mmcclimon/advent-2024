package main

import (
	"fmt"
	"slices"

	"github.com/mmcclimon/advent-2024/advent/input"
)

func main() {
	var locks, keys [][]int

	for hunk := range input.New().Hunks() {
		heights := slices.Repeat([]int{-1}, 5)
		for _, line := range hunk {
			for i, char := range line {
				if char == '#' {
					heights[i]++
				}
			}
		}

		if hunk[0][0] == '#' {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	total := 0

	for _, lock := range locks {

	key:
		for _, key := range keys {
			for i := range 5 {
				if lock[i]+key[i] > 5 {
					continue key
				}
			}

			total++
		}
	}

	fmt.Println("part 1:", total)
}
