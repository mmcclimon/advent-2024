package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type Gate struct {
	in1, in2 string
	op       string
	out      string

	iv1, iv2 *int
	done     bool
	val      int
}

func main() {
	hunks := slices.Collect(input.New().Hunks())

	re := regexp.MustCompile(`(\w+) (AND|OR|XOR) (\w+) -> (\w+)`)

	lookup := make(map[string][]*Gate)

	for _, line := range hunks[1] {
		m := re.FindStringSubmatch(line)
		in1, in2, op, out := m[1], m[3], m[2], m[4]
		gate := &Gate{in1, in2, op, out, nil, nil, false, 0}
		lookup[in1] = append(lookup[in1], gate)
		lookup[in2] = append(lookup[in2], gate)
	}

	q := collections.NewDeque[*Gate]()

	for _, line := range hunks[0] {
		bits := strings.Split(line, ": ")
		wire := bits[0]
		n := conv.Atoi(bits[1])

		for _, gate := range lookup[wire] {
			assert.True(gate != nil, wire)
			gate.Input(wire, n)

			// fmt.Println("INPUT", gate, wire, n)

			if gate.IsReady() {
				// fmt.Println("gate is ready", gate)
				q.Append(gate)
			}
		}
	}

	assert.True(q.Len() > 0, "at least one gate is ready")

	for q.Len() > 0 {
		gate, err := q.PopLeft()
		assert.Nil(err)

		// fmt.Println("process", gate)

		out := gate.out

		for _, g2 := range lookup[out] {

			// if !ok {
			// 	fmt.Println("cannot find output gate", out, gate)
			// 	continue
			// }

			assert.True(g2 != nil, fmt.Sprintf("got gate for %s", out))
			g2.Input(out, gate.Output())

			// fmt.Println("input", g2, out)

			if g2.IsReady() {
				q.Append(g2)
			}
		}
	}

	zs := make(map[string]int)

	for _, gates := range lookup {
		for _, gate := range gates {
			if !strings.HasPrefix(gate.out, "z") {
				continue
			}

			zs[gate.out] = gate.val
		}
	}

	result := 0

	for i := 0; true; i++ {
		bit, ok := zs[fmt.Sprintf("z%02d", i)]
		if !ok {
			break
		}

		n := bit << i

		result |= n

	}

	fmt.Println(result)

}

func (g *Gate) String() string {
	return fmt.Sprintf("{%s %s %s -> %s}", g.in1, g.op, g.in2, g.out)
}

func (g *Gate) Input(which string, val int) {
	switch which {
	case g.in1:
		g.iv1 = &val
	case g.in2:
		g.iv2 = &val
	default:
		panic("bad gate")
	}

	// fmt.Printf("  after input on %s: %v, %v\n", which, g.iv1, g.iv2)

	if g.iv1 == nil || g.iv2 == nil {
		return
	}

	switch g.op {
	case "AND":
		g.val = *g.iv1 & *g.iv2
	case "OR":
		g.val = *g.iv1 | *g.iv2
	case "XOR":
		g.val = *g.iv1 ^ *g.iv2
	default:
		panic("bad op")
	}
}

func (g *Gate) IsReady() bool {
	return g.iv1 != nil && g.iv2 != nil && !g.done
}

func (g *Gate) Output() int {
	if g.iv1 == nil || g.iv2 == nil {
		panic(fmt.Sprintf("cannot output a non-ready gate: %s", g))
	}

	g.done = true
	return g.val
}
