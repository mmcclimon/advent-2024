package main

import (
	"fmt"
	"slices"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

const Empty = -1

func main() {
	in := input.New().Slurp()

	disk := make([]int, 0, len(in)*5)
	isFile := true
	id := 0

	for _, char := range in {
		if char == '\n' {
			continue
		}

		n := conv.Atoi(string(char))
		for range n {
			if isFile {
				disk = append(disk, id)
			} else {
				disk = append(disk, Empty)
			}
		}

		if isFile {
			id++
		}
		isFile = !isFile
	}

	part1(disk)
}

func part1(disk []int) {
	// printDisk(disk)
	lp := slices.Index(disk, Empty)
	rp := lastFull(disk)

	for lp < rp {
		assert.True(disk[lp] == Empty, "left pointer is empty")
		assert.True(lp < rp, "lp < rp")

		disk[lp] = disk[rp]
		disk[rp] = Empty

		lp = slices.Index(disk, Empty)
		rp = lastFull(disk)
	}

	// printDisk(disk)
	// take one more pass: we could compute this as we go, but also meh
	checksum := 0
	for i, n := range disk {
		if n == Empty {
			break
		}
		// fmt.Printf("%d * %d\n", i, n)

		checksum += i * n
	}

	fmt.Println(checksum)
}

func lastFull(disk []int) int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != Empty {
			return i
		}
	}

	return -1
}

func printDisk(disk []int) {
	for _, n := range disk {
		switch n {
		case Empty:
			fmt.Print(".")
		default:
			fmt.Print(n)
		}
	}

	fmt.Print("\n")
}
