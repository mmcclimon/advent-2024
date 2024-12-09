package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

const Empty = -1

func main() {
	in := strings.TrimSpace(input.New().Slurp())

	disk1 := make([]int, 0, len(in)*5)
	isFile := true
	id := 0

	for _, char := range in {
		n := conv.Atoi(string(char))
		toAppend := Empty
		if isFile {
			toAppend = id
		}

		for range n {
			disk1 = append(disk1, toAppend)
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
	rp := lastFullIndex(disk)

	for lp < rp {
		disk[lp] = disk[rp]
		disk[rp] = Empty

		lp = slices.Index(disk, Empty)
		rp = lastFullIndex(disk)
	}

	fmt.Println("part 1:", checksum(disk))
}

func lastFullIndex(disk []int) int {
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
		start, length := startAndLengthForID(disk, id)
		idx := findFreeBlock(disk, length)

		if idx == -1 || idx >= start {
			continue
		}

		for i := range length {
			disk[idx+i] = id
			disk[start+i] = Empty
		}

	}

	fmt.Println("part 2:", checksum(disk))
}

func startAndLengthForID(disk []int, id int) (int, int) {
	start := slices.Index(disk, id)

	length := 0
	for i := start; i < len(disk) && disk[i] == id; i++ {
		length++
	}

	return start, length
}

func findFreeBlock(disk []int, length int) int {
	for i := 0; i < len(disk); i++ {
		if disk[i] != Empty {
			continue
		}

		start := i

		for j := i; j < len(disk) && disk[j] == Empty; j++ {
			if 1+j-start >= length {
				return start
			}

			i++
		}
	}

	return -1
}
