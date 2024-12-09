package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
	"github.com/mmcclimon/advent-2024/advent/operator"
)

type Disk struct {
	s             []int
	maxID         int
	firstEmptyIdx int
	lastFullIdx   int
	// only used in part 2
	ids        map[int]blockData
	freeBlocks []*blockData
}

type blockData struct{ idx, length int }

const Empty = -1

func main() {
	in := strings.TrimSpace(input.New().Slurp())

	part1(NewDisk(in))
	part2(NewDisk(in))
}

func NewDisk(line string) *Disk {
	s := make([]int, 0, len(line)*5)
	isFile := true

	var id, lfi int
	var freeBlocks []*blockData
	ids := make(map[int]blockData, len(line))

	for _, char := range line {
		toAppend := operator.CrummyTernary(isFile, id, Empty)
		startIdx := len(s)
		n := conv.Atoi(string(char))
		block := blockData{startIdx, n}

		for range n {
			s = append(s, toAppend)
		}

		if isFile {
			lfi = len(s) - 1
			ids[id] = block
			id++
		} else {
			freeBlocks = append(freeBlocks, &block)
		}

		isFile = !isFile
	}

	return &Disk{
		s:             s,
		maxID:         id - 1,
		firstEmptyIdx: slices.Index(s, Empty), // will be fast: always <= 10
		lastFullIdx:   lfi,
		ids:           ids,
		freeBlocks:    freeBlocks,
	}
}

func part1(disk *Disk) {
	lp := disk.firstEmptyIdx
	rp := disk.lastFullIdx

	for lp < rp {
		disk.swap(lp, rp)

		lp = disk.firstEmptyIdx
		rp = disk.lastFullIdx
	}

	fmt.Println("part 1:", disk.Checksum())
}

func (d *Disk) swap(left, right int) {
	d.s[left] = d.s[right]
	d.s[right] = Empty

	// recompute indexes
	for i := d.lastFullIdx; i >= 0; i-- {
		if d.s[i] != Empty {
			d.lastFullIdx = i
			break
		}
	}

	for i := d.firstEmptyIdx; i < len(d.s); i++ {
		if d.s[i] == Empty {
			d.firstEmptyIdx = i
			break
		}
	}
}

func (d *Disk) swapBlocks(free *blockData, id blockData, freeIdx int) {
	for i := range id.length {
		d.swap(free.idx+i, id.idx+i)
	}

	if free.length == id.length {
		d.freeBlocks = slices.Delete(d.freeBlocks, freeIdx, freeIdx+1)
		return
	}

	free.length -= id.length
	free.idx += id.length
}

func (d *Disk) Checksum() int {
	checksum := 0
	for i, n := range d.s {
		if n == Empty {
			continue
		}

		checksum += i * n
	}

	return checksum
}

func part2(disk *Disk) {
	for id := disk.maxID; id >= 0; id-- {
		data := disk.dataForID(id)
		i, free := disk.findFreeBlock(data.length)

		if free == nil || free.idx >= data.idx {
			continue
		}

		disk.swapBlocks(free, data, i)
	}

	fmt.Println("part 2:", disk.Checksum())
}

func (d *Disk) dataForID(id int) blockData {
	return d.ids[id]
}

func (d *Disk) findFreeBlock(length int) (int, *blockData) {
	for i, block := range d.freeBlocks {
		if block.length >= length {
			return i, block
		}
	}

	return -1, nil
}
