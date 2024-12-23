package main

import (
	"fmt"
	"sync"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/mmcclimon/advent-2024/advent/collections"
	"github.com/mmcclimon/advent-2024/advent/conv"
	"github.com/mmcclimon/advent-2024/advent/input"
	"golang.org/x/sync/errgroup"
)

type PRNG struct {
	secret   int
	orig     int
	rands    []int
	deltas   []int
	deltaIdx map[Seq]int
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

	var eg errgroup.Group
	eg.SetLimit(32)

	var counts sync.Map

	for seq := range allDeltas {
		eg.Go(func() error {
			bananas := 0
			for _, buyer := range buyers {
				bananas += buyer.BuyAt(seq)
			}

			counts.Store(seq, bananas)
			return nil
		})
	}

	assert.Nil(eg.Wait())

	counts.Range(func(k, v any) bool {
		p2 = max(p2, v.(int))
		return true
	})

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)

}

func NewRNG(secret int) *PRNG {
	return &PRNG{
		secret:   secret,
		orig:     secret,
		rands:    make([]int, 0, 2001),
		deltas:   make([]int, 0, 2000),
		deltaIdx: make(map[Seq]int),
	}
}

func (r *PRNG) Prefill() {
	r.rands = append(r.rands, r.secret)

	for i := range 2000 {
		r.rands = append(r.rands, r.Rand())
		r.deltas = append(r.deltas, (r.rands[i+1]%10)-(r.rands[i]%10))

		// at i=3, 0:4
		if i > 3 {
			start := i - 3
			seq := seqFromSlice(r.deltas[start : i+1])
			_, ok := r.deltaIdx[seq]
			if !ok {
				r.deltaIdx[seq] = start
			}
		}
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

	for seq := range r.deltaIdx {
		distinct.Add(seq)
	}

	return distinct
}

func (r *PRNG) BuyAt(seq Seq) int {
	start, ok := r.deltaIdx[seq]
	if !ok {
		return 0
	}

	return r.rands[start+4] % 10
}

func seqFromSlice(s []int) Seq {
	return Seq{s[0], s[1], s[2], s[3]}
}
