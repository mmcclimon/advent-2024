package main

import (
	"fmt"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type Operator int

const (
	Add Operator = iota
	Mul
	Concat
)

func main() {
	sum1, sum2 := 0, 0

	for line := range input.New().Lines() {
		fields := strings.Fields(line)
		target := conv.Atoi(fields[0][:len(fields[0])-1])

		nums := conv.ToInts(fields[1:])
		// fmt.Println(target, nums)

		if check(target, nums) {
			// fmt.Println("yay!", target)
			sum1 += target
			// sum2 += target
			// continue
		}

		if check2(target, nums) {
			// fmt.Println("yay!", target)
			sum2 += target
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func check(target int, nums []int) bool {
	allOps := generateOps(len(nums)-1, false)

	for _, opList := range allOps {
		result := nums[0]

		for i, op := range opList {
			n := nums[i+1]
			switch op {
			case Add:
				result += n
			case Mul:
				result *= n
			}
		}

		if target == result {
			return true
		}
	}

	return false
}

func check2(target int, nums []int) bool {
	allOps := generateOps(len(nums)-1, true)

	for _, opList := range allOps {
		result := nums[0]

		for i, op := range opList {
			n := nums[i+1]
			switch op {
			case Add:
				result += n
			case Mul:
				result *= n
			case Concat:
				result = conv.Atoi(fmt.Sprintf("%d%d", result, n))
			}
		}

		if target == result {
			return true
		}
	}

	return false
}

func generateOps(n int, withConcat bool) [][]Operator {
	ops := [][]Operator{nil}

	for range n {
		var tmp [][]Operator

		for _, list := range ops {
			add := make([]Operator, len(list)+1)
			copy(add, list)
			add[len(add)-1] = Add

			mul := make([]Operator, len(list)+1)
			copy(mul, list)
			mul[len(mul)-1] = Mul

			// fmt.Println("HEY", list, nums)

			tmp = append(tmp, add, mul)

			if withConcat {
				concat := make([]Operator, len(list)+1)
				copy(concat, list)
				concat[len(concat)-1] = Concat
				tmp = append(tmp, concat)
			}
		}

		ops = tmp
	}

	return ops
}
