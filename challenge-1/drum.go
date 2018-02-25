// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"fmt"
)

// Header represents the head of the file
// Header.Version has 32 bytes
type Header struct {
	Splice     [6]byte
	PatterSize int64
	Version    string
	Tempo      float32
}

// Quarter representes one quarter of 4 steps
type Quarter [4]bool

// Quarters represents an sclice of 4 quarters. 16 steps total
type Quarters [4]Quarter

// Track is the high level representation of a sound track of an Name
type Track struct {
	ID         uint8
	NameLength uint32
	Name       string
	Quarters   Quarters
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
type Pattern struct {
	Header Header
	Tracks []Track
}

func (p Pattern) String() string {
	s := fmt.Sprintf("Saved with HW Version: %v\nTempo: %v\n", p.Header.Version, p.Header.Tempo)

	for _, t := range p.Tracks {
		s += fmt.Sprintf("(%v) %v\t%v\n", t.ID, t.Name, t.Quarters)
	}

	return s
}

func (q Quarters) String() string {
	var s string

	for _, quarter := range q {
		s += fmt.Sprintf("|%v", quarter)
	}

	s += "|"

	return s
}

func (q Quarter) String() string {
	var s string

	for _, step := range q {
		if step == true {
			s += "x"
		} else if step == false {
			s += "-"
		}
	}

	return s
}
