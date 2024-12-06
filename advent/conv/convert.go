package conv

import (
	"strconv"

	"github.com/mmcclimon/advent-2024/advent/assert"
)

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	assert.Nil(err)
	return n
}

func ToInts(strs []string) []int {
	nums := make([]int, len(strs))

	for i, s := range strs {
		nums[i] = Atoi(s)
	}

	return nums
}
