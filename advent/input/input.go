package input

import (
	"bufio"
	"io"
	"iter"
	"os"
	"strconv"

	"github.com/mmcclimon/advent-2024/advent/assert"
	"github.com/xdg-go/strum"
)

type Input struct {
	r io.Reader
}

// Get an io.Reader for the first command-line arg; defaulting to stdin.
func New() *Input {
	if len(os.Args) == 1 {
		return &Input{r: os.Stdin}
	}

	f, err := os.Open(os.Args[1])
	assert.Nil(err)

	return &Input{r: f}
}

func (i *Input) Strummer() *Strummer {
	return &Strummer{st: strum.NewDecoder(i.r)}
}

// NB throws away errors.
func (i *Input) Lines() iter.Seq[string] {
	scanner := bufio.NewScanner(i.r)
	return func(yield func(string) bool) {
		for scanner.Scan() {
			yield(scanner.Text())
		}
	}
}

func (i *Input) Ints() iter.Seq[int] {
	return func(yield func(int) bool) {
		for line := range i.Lines() {
			n, _ := strconv.Atoi(line)
			yield(n)
		}
	}
}

func (i *Input) Hunks() iter.Seq[[]string] {
	var buf []string

	return func(yield func([]string) bool) {
		for line := range i.Lines() {
			if line == "" {
				yield(buf)
				buf = nil
				continue
			}

			buf = append(buf, line)
		}

		if len(buf) > 0 {
			yield(buf)
		}
	}
}
