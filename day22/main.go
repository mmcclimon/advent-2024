package main

import (
	"fmt"

	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
)

type PRNG struct {
	secret int
	orig   int
	rands  []int
	deltas []int
}

type Seq struct {
	a, b, c, d int
}

func main() {
	p1, p2 := 0, 0

	var buyers []*PRNG
	allDeltas := collections.NewSet[Seq]()

	for buyer := range input.New().Lines() {
		n := conv.Atoi(buyer)
		rng := NewRNG(n)
		buyers = append(buyers, rng)

		rng.Prefill()
		allDeltas.Extend(rng.Deltas())
		p1 += rng.Nth(2000)
	}

	done := 0
	for seq := range allDeltas {
		if done%100 == 0 {
			fmt.Print("\r", done)
		}

		done++

		bananas := 0
		for _, buyer := range buyers {
			bananas += buyer.BuyAt(seq)
		}

		p2 = max(p2, bananas)
	}

	fmt.Print("\n")

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)

}

func NewRNG(secret int) *PRNG {
	return &PRNG{
		secret: secret,
		orig:   secret,
		rands:  make([]int, 0, 2001),
		deltas: make([]int, 0, 2000),
	}
}

func (r *PRNG) Prefill() {
	r.rands = append(r.rands, r.secret)

	for i := range 2000 {
		r.rands = append(r.rands, r.Rand())
		r.deltas = append(r.deltas, (r.rands[i+1]%10)-(r.rands[i]%10))
	}
}

func (r *PRNG) Nth(n int) int {
	return r.rands[n]
}

func (r *PRNG) Rand() int {
	r.mix(r.secret * 64)
	r.mix(r.secret / 32)
	r.mix(r.secret * 2048)
	return r.secret
}

func (r *PRNG) mix(n int) {
	r.secret = (r.secret ^ n) % 16777216
}

func (r *PRNG) Deltas() collections.Set[Seq] {
	distinct := collections.NewSet[Seq]()

	for i := range 2000 - 3 {
		distinct.Add(seqFromSlice(r.deltas[i : i+4]))
	}

	return distinct
}

func (r *PRNG) BuyAt(seq Seq) int {
	for i := range 2000 - 3 {
		s := seqFromSlice(r.deltas[i : i+4])
		if s == seq {
			return r.rands[i+4] % 10
		}
	}

	return 0
}

func seqFromSlice(s []int) Seq {
	return Seq{s[0], s[1], s[2], s[3]}
}
