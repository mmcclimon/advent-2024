package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
)

var (
	initialA     = 47719761
	instructions = []int{2, 4, 1, 5, 7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0}
)

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

type CPU struct {
	a, b, c int
	ip      int
	ops     []int
	output  []int
}

func main() {
	cpu := &CPU{
		a:   initialA,
		b:   0,
		c:   0,
		ops: instructions,
	}

	cpu.Run()
	fmt.Println(cpu.OutputString())

	part2()
}

/*
This program is:

while a != 0:
    b = (a % 8) ^ 5
    c = a // (2**b)
    a = a // 8
    b = b ^ c  ^ 6
    out.append(b % 8)
*/

func part2() {
	start := int(math.Pow(8, float64(len(instructions)-1)))
	start = 1

	for i := start; true; i++ {
		out := runOptimized(i)
		fmt.Printf("%5d  %010b\n", i, out)

		if i > 128 {
			break
		}

		/*

			cpu := &CPU{
				a:   i,
				b:   0,
				c:   0,
				ops: instructions,
			}

			cpu.Run()
			if slices.Equal(instructions, cpu.output) {
				fmt.Println("\npart 2:", i)
				break
			}

			if i%1000 == 0 {
				fmt.Printf("\r%d (%d)", i, len(cpu.output))
			}
		*/
	}
}

func runOptimized(start int) int64 {
	a := start
	var b, c int
	var out []int

	for a != 0 {
		b = (a % 8) ^ 5
		c = a / int(math.Pow(2, float64(b)))
		a = a / 8
		b = b ^ c ^ 6
		out = append(out, b%8)
	}

	strs := make([]string, len(out))
	for i, o := range out {
		strs[i] = fmt.Sprint(o)
	}

	n, err := strconv.ParseInt(strings.Join(strs, ""), 8, 0)
	assert.Nil(err)

	return n
}

func (cpu *CPU) Run() {
	for cpu.ip < len(cpu.ops) {
		cpu.processOp()
	}
}

func (cpu *CPU) processOp() {
	var f func(int)

	op := cpu.ops[cpu.ip]
	arg := cpu.ops[cpu.ip+1]

	cpu.ip += 2

	switch op {
	case adv:
		f = cpu.adv
	case bxl:
		f = cpu.bxl
	case bst:
		f = cpu.bst
	case jnz:
		f = cpu.jnz
	case bxc:
		f = cpu.bxc
	case out:
		f = cpu.out
	case bdv:
		f = cpu.bdv
	case cdv:
		f = cpu.cdv
	}

	f(arg)
}

func (cpu *CPU) combo(arg int) int {
	switch arg {
	case 0, 1, 2, 3:
		return arg
	case 4:
		return cpu.a
	case 5:
		return cpu.b
	case 6:
		return cpu.c
	default:
		panic("invalid arg")
	}
}

func (cp *CPU) cstr(arg int) string {
	switch arg {
	case 0, 1, 2, 3:
		return fmt.Sprint(arg)
	case 4:
		return "a"
	case 5:
		return "b"
	case 6:
		return "c"
	default:
		panic("invalid arg")
	}

}

func (cpu *CPU) div(arg int) int {
	num := cpu.a
	denom := int(math.Pow(float64(2), float64(cpu.combo(arg))))
	return num / denom
}

func (cpu *CPU) adv(arg int) {
	cpu.a = cpu.div(arg)
}

func (cpu *CPU) bxl(arg int) {
	cpu.b = cpu.b ^ arg
}

func (cpu *CPU) bst(arg int) {
	cpu.b = cpu.combo(arg) & 0b111
}

func (cpu *CPU) jnz(arg int) {
	if cpu.a == 0 {
		return
	}

	cpu.ip = arg
}

func (cpu *CPU) bxc(_ int) {
	cpu.b = cpu.b ^ cpu.c
}

func (cpu *CPU) out(arg int) {
	out := cpu.combo(arg) & 0b111
	cpu.output = append(cpu.output, out)
}

func (cpu *CPU) bdv(arg int) {
	cpu.b = cpu.div(arg)
}

func (cpu *CPU) cdv(arg int) {
	cpu.c = cpu.div(arg)
}

func (cpu *CPU) OutputString() string {
	strs := make([]string, len(cpu.output))
	for i, o := range cpu.output {
		strs[i] = fmt.Sprint(o)
	}

	return strings.Join(strs, ",")
}

/*

	b = a % 8  	  # 2 4
	b = b ^ 5		  # 1 5
	c = a / b     # 7 5
	a = a / 8		  # 0 3
	b = b ^ c		  # 4 1
	b = b ^ 6		  # 1 6
	output b % 8  # 5 5
	goto start

while a != 0:
	b = a % 8
	b = b ^ 5
	c = a / b
	a = a / 8
	b = b ^ c
	b = b ^ 6
	output b


a = 47719761

while a != 0:
	b = a % 8		b = 1
	b = b ^ 5		b = 4
	c = a / b		c = 9543952
	a = a / 8		a = 5964970
	b = b ^ c		b =
	b = b ^ 6		b = 9543954
	output b

*/
