package main

import (
	"fmt"
	"slices"
)

var (
	initialA     = 47719761
	instructions = []int{2, 4, 1, 5, 7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0}
)

/*

3   0		6		2		3		4		6		3		1		3		2		0		4		2		3		3
011 000 110 010 011 100 110 011 001 010 010 000 100 010 011 011

*/

func main() {
	out := runOptimized(initialA)
	fmt.Println(out)

	a := 109020013201563
	quine := runOptimized(a)
	fmt.Println(quine)
	fmt.Printf("%0b\n", a)
	fmt.Printf("%0o\n", a)
	part2()
}

//nolint:unused
func part2() {
	//start := int(math.Pow(8, float64(len(instructions)-1)))

	start := 35184375376282
	start = 0b11000110010011100110011001010010000100010011011
	inc := 98304
	prev := 0
	// start = 1

	//  36601184135578
	// 109020013201563

	for a := start; true; a += inc {
		out := runOptimized(a)
		fmt.Printf("\r%d", a)

		if slices.Equal(out[:7], instructions[:7]) {
			fmt.Print("\n")
			fmt.Println(a, len(out), a-prev, out)
			prev = a
		}

		if slices.Equal(out, instructions) {
			fmt.Println("part 2:", a)
			return
		}
	}
}

func runOptimized(start int) []int {
	a := start
	var out []int

	for a != 0 {
		b := (a & 0b111) ^ 5
		c := a >> b
		a >>= 3
		b = b ^ 6 ^ (c & 0b111)
		out = append(out, b)

		// fmt.Printf("start=%010b, a=%010b  b=%010b  c=%010b, out=%010b\n", start, a, b, c, b%8)
	}

	return out

}
