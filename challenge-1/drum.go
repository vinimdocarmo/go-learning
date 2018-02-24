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

// Track is the high level representation of a sound track of an instrument
type Track struct {
	ID         int8
	NameLength int32
	Instrument string
	Quarters   [4][4]byte
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
type Pattern struct {
	Header Header
	Tracks []Track
}

// Format formats the splice file pattern to meanful string
func (p Pattern) String() string {
	return fmt.Sprintf("Saved with HW Version: %v\nTempo: %v", p.Header.Version, p.Header.Tempo)
}
