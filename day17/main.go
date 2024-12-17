package main

import (
	"fmt"
	"math"
	"strings"
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
		a:   47719761,
		b:   0,
		c:   0,
		ops: []int{2, 4, 1, 5, 7, 5, 0, 3, 4, 1, 1, 6, 5, 5, 3, 0},
	}

	cpu.Run()

	fmt.Println(cpu.OutputString())
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

func (cpu *CPU) adv(arg int) {
	num := cpu.a
	denom := int(math.Pow(float64(2), float64(cpu.combo(arg))))
	cpu.a = num / denom
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
	num := cpu.a
	denom := int(math.Pow(float64(2), float64(cpu.combo(arg))))
	cpu.b = num / denom
}

func (cpu *CPU) cdv(arg int) {
	num := cpu.a
	denom := int(math.Pow(float64(2), float64(cpu.combo(arg))))
	cpu.c = num / denom
}

func (cpu *CPU) OutputString() string {
	strs := make([]string, len(cpu.output))
	for i, o := range cpu.output {
		strs[i] = fmt.Sprint(o)
	}

	return strings.Join(strs, ",")
}
