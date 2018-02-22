package drum

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

// DecoderSplice represents the structure that is going to decode a .splice file
type DecoderSplice struct {
	r io.Reader
}

func (d DecoderSplice) decode(p *Pattern) error {
	err := binary.Read(d.r, binary.LittleEndian, p)

	if err != nil {
		return err
	}

	return nil
}

func newDecoderSplice(filename string) (DecoderSplice, error) {
	r, err := os.OpenFile(filename, os.O_RDONLY, 0666)

	var ds DecoderSplice

	if err != nil {
		return ds, err
	}

	ds = DecoderSplice{r: bufio.NewReader(r)}

	return ds, nil
}

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
func DecodeFile(path string) (*Pattern, error) {
	p := &Pattern{}

	ds, err := newDecoderSplice(path)

	if err != nil {
		return p, err
	}

	ds.decode(p)

	return p, nil
}
