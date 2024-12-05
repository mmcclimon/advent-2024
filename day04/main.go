package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2024/advent/input"
)

type rc struct {
	r, c int
}

func main() {
	grid := make(map[rc]rune)

	for r, line := range input.New().EnumerateLines() {
		for c, letter := range line {
			grid[rc{r, c}] = letter
		}
	}

	found1, found2 := 0, 0

	for k, l := range grid {
		if l == 'X' {
			found1 += check1(grid, k)
		}

		if l == 'A' {
			found2 += check2(grid, k)
		}
	}

	fmt.Println("part 1:", found1)
	fmt.Println("part 2:", found2)
}

func check1(grid map[rc]rune, start rc) int {
	found := 0
	r, c := start.r, start.c

	chars := make([]rune, 4)

	for rd := -1; rd <= 1; rd++ {
		for cd := -1; cd <= 1; cd++ {
			for i := range 4 {
				chars[i] = grid[rc{r + (i * rd), c + (i * cd)}]
			}

			if string(chars) == "XMAS" {
				found++
			}
		}
	}

	return found
}

func check2(grid map[rc]rune, start rc) int {
	r, c := start.r, start.c

	se := string([]rune{grid[rc{r - 1, c - 1}], grid[start], grid[rc{r + 1, c + 1}]})
	sw := string([]rune{grid[rc{r - 1, c + 1}], grid[start], grid[rc{r + 1, c - 1}]})

	isMas := func(s string) bool { return s == "MAS" || s == "SAM" }

	if isMas(se) && isMas(sw) {
		return 1
	}

	return 0
}
