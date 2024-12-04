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

	for rd := -1; rd <= 1; rd++ {
		for cd := -1; cd <= 1; cd++ {
			if rd == 0 && cd == 0 {
				continue
			}

			if check_delta(grid, start, rd, cd) {
				found++
			}
		}
	}

	return found
}

func check_delta(grid map[rc]rune, start rc, rd, cd int) bool {
	r, c := start.r, start.c

	chars := make([]rune, 4)

	for i := range 4 {
		chars[i] = grid[rc{r + (i * rd), c + (i * cd)}]
	}

	return string(chars) == "XMAS"
}

func check2(grid map[rc]rune, start rc) int {
	r, c := start.r, start.c

	se := string([]rune{grid[rc{r - 1, c - 1}], grid[start], grid[rc{r + 1, c + 1}]})
	sw := string([]rune{grid[rc{r - 1, c + 1}], grid[start], grid[rc{r + 1, c - 1}]})

	if (se == "MAS" || se == "SAM") && (sw == "MAS" || sw == "SAM") {
		return 1
	}

	return 0
}
