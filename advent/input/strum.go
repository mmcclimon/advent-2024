package input

import (
	"io"

	"github.com/xdg-go/strum"
)

type Strummer struct {
	st   *strum.Decoder
	done bool
}

func NewStrummer() *Strummer {
	return &Strummer{st: strum.NewDecoder(New().r)}
}

func (s *Strummer) Next() bool {
	return !s.done
}

func (s *Strummer) Decode(val any) error {
	err := s.st.Decode(val)

	if err == io.EOF {
		s.done = true
		return nil
	}

	return err
}
