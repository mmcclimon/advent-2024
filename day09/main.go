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

	disk1 := make([]int, 0, len(in)*5)
	isFile := true
	id := 0

	for _, char := range in {
		if char == '\n' {
			continue
		}

		n := conv.Atoi(string(char))
		for range n {
			if isFile {
				disk1 = append(disk1, id)
			} else {
				disk1 = append(disk1, Empty)
			}
		}

		if isFile {
			id++
		}
		isFile = !isFile
	}

	disk2 := make([]int, len(disk1))
	copy(disk2, disk1)

	part1(disk1)
	part2(disk2, id-1)
}

func part1(disk []int) {
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

	fmt.Println("part 1:", checksum(disk))
}

func lastFull(disk []int) int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != Empty {
			return i
		}
	}

	return -1
}

func checksum(disk []int) int {
	checksum := 0
	for i, n := range disk {
		if n == Empty {
			continue
		}

		checksum += i * n
	}

	return checksum
}

func part2(disk []int, startID int) {
	for id := startID; id >= 0; id-- {
		// loop
		start, length := startEndForFull(disk, id)
		free := indexForFree(disk, length)

		if free == -1 || free >= start {
			continue
		}

		// fmt.Printf("moving id %d (len=%d) from %d to %d\n", id, length, start, free)

		for i := range length {
			disk[free+i] = id
			disk[start+i] = Empty
		}

	}

	fmt.Println("part 2:", checksum(disk))
}

func startEndForFull(disk []int, id int) (int, int) {
	start := slices.Index(disk, id)
	length := 0
	for i := start; i < len(disk) && disk[i] == id; i++ {
		length++
	}

	return start, length
	// fmt.Println(start, length)
}

func indexForFree(disk []int, length int) int {
	if length >= 10 {
		panic("length too big?")
	}

	for i := 0; i < len(disk); i++ {
		if disk[i] != Empty {
			continue
		}

		start := i
		// fmt.Println("looking for space of", length, "starting at", i)

		for j := i; j < len(disk) && disk[j] == Empty; j++ {
			// fmt.Println("checking", j, "at i", i)
			if 1+j-start >= length {
				return start
			}

			i++
		}

		// i = start
	}

	// panic("no free space found")
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
