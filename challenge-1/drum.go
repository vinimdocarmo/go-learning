// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

// Header represents the head of the file
type Header struct {
	Splice [6]byte
	_      [8]byte
}

// Track is the high level representation of a sound track of an instrument
type Track struct {
	ID         [1]byte
	Instrument [10]byte
	Quarters   [4][4]byte
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
type Pattern struct {
	Header  Header
	Version [11]byte
	_       [25]byte //tempo ta por aqui
	Tracks  [4]Track
}

// Format formats the splice file pattern to meanful string
func (p Pattern) String() string {
	s := "Saved with HW Version: " + string(p.Version[:])
	return s
}
