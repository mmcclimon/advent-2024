package input

import (
	"bufio"
	"io"
	"iter"
	"strconv"
)

// NB throws away errors.
func Lines(r io.Reader) iter.Seq[string] {
	scanner := bufio.NewScanner(r)
	return func(yield func(string) bool) {
		for scanner.Scan() {
			yield(scanner.Text())
		}
	}
}

func Ints(r io.Reader) iter.Seq[int] {
	return func(yield func(int) bool) {
		for line := range Lines(r) {
			n, _ := strconv.Atoi(line)
			yield(n)
		}
	}
}

func Hunks(r io.Reader) iter.Seq[[]string] {
	var buf []string

	return func(yield func([]string) bool) {
		for line := range Lines(r) {
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
