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

	x, y := getXY(hunks[0])
	lookup := makeGates(hunks[1])
	result := run(lookup, x, y)
	fmt.Println("part 1:", result)

	// This is just a loop to print out where things look awry, then I actually
	// solved by squinting at graphviz output.
	for i := range 45 {
		x := 1 << i
		y := 1

		lookup := makeGates(hunks[1])
		run(lookup, x, y)

		result := collectResult(lookup)

		//nolint:staticcheck
		if result != x+y {
			fmt.Printf("BAD %d + %d != %d (%d) \n", x, y, result, i)
		} else {
			// fmt.Printf("OK  %d + %d = %d\n", x, y, result)
		}
	}
}

var gateRE = regexp.MustCompile(`(\w+) (AND|OR|XOR) (\w+) -> (\w+)`)

func makeGates(lines []string) map[string][]*Gate {
	lookup := make(map[string][]*Gate)

	for _, line := range lines {
		m := gateRE.FindStringSubmatch(line)
		in1, in2, op, out := m[1], m[3], m[2], m[4]
		gate := &Gate{in1, in2, op, out, nil, nil, false, 0}
		lookup[in1] = append(lookup[in1], gate)
		lookup[in2] = append(lookup[in2], gate)
	}

	return lookup
}

func getXY(lines []string) (int, int) {
	var x, y int

	for _, line := range lines {
		bits := strings.Split(line, ": ")
		wire := bits[0]

		which := wire[0]
		shift := conv.Atoi(wire[1:])

		n := conv.Atoi(bits[1]) << shift

		switch which {
		case 'x':
			x |= n
		case 'y':
			y |= n
		}

		/*

		 */
	}

	return x, y
}

func run(lookup map[string][]*Gate, x, y int) int {
	q := collections.NewDeque[*Gate]()

	for i := range 45 {
		xBit := (x & (1 << i)) >> i
		xWire := fmt.Sprintf("x%02d", i)
		for _, gate := range lookup[xWire] {
			assert.True(gate != nil, xWire)
			gate.Input(xWire, xBit)

			if gate.IsReady() {
				q.Append(gate)
			}
		}

		yBit := (y & (1 << i)) >> i
		yWire := fmt.Sprintf("y%02d", i)
		for _, gate := range lookup[yWire] {
			assert.True(gate != nil, yWire)
			gate.Input(yWire, yBit)

			if gate.IsReady() {
				q.Append(gate)
			}
		}
	}

	assert.True(q.Len() > 0, "at least one gate is ready")

	for q.Len() > 0 {
		gate, err := q.PopLeft()
		assert.Nil(err)

		out := gate.out

		for _, g2 := range lookup[out] {
			assert.True(g2 != nil, fmt.Sprintf("got gate for %s", out))
			g2.Input(out, gate.Output())

			if g2.IsReady() {
				q.Append(g2)
			}
		}
	}

	return collectResult(lookup)
}

func collectResult(lookup map[string][]*Gate) int {
	// find the output
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

		result |= bit << i
	}

	return result
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

/*

Manual inspection

jgt, mht
z05, hdt
z09, gbf
z30,nbf

*/
