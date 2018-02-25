package drum

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"os"
)

// DecoderSplice represents the structure that is going to decode a .splice file
type DecoderSplice struct {
	r    io.Reader
	file *os.File
}

func (d DecoderSplice) decode(p *Pattern) error {
	err := d.decodeHeader(&p.Header)

	if err != nil {
		return err
	}

	p.Tracks = []Track{}

	for {
		track := Track{}

		err = d.decodeTrack(&track)

		// If there is no byte left for read
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		p.Tracks = append(p.Tracks, track)
	}
}

func (d DecoderSplice) decodeHeader(h *Header) error {
	var header struct {
		Splice  [6]byte
		Size    int64
		Version [32]byte
	}

	err := binary.Read(d.r, binary.BigEndian, &header)

	if err != nil {
		return err
	}

	var tempo float32

	err = binary.Read(d.r, binary.LittleEndian, &tempo)

	if err != nil {
		return err
	}

	h.PatterSize = header.Size
	h.Splice = header.Splice
	h.Version = string(bytes.TrimRight(header.Version[:], string(0)))
	h.Tempo = tempo

	return nil
}

func (d DecoderSplice) decodeTrack(t *Track) error {
	var trackHeader struct {
		ID      uint8
		NameLen uint32
	}

	err := binary.Read(d.r, binary.BigEndian, &trackHeader)

	if err != nil {
		return err
	}

	t.ID = trackHeader.ID
	t.NameLength = trackHeader.NameLen

	name := make([]byte, trackHeader.NameLen)

	err = binary.Read(d.r, binary.BigEndian, name)

	if err != nil {
		return err
	}

	t.Name = string(name)

	err = binary.Read(d.r, binary.BigEndian, &t.Quarters)

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

	ds = DecoderSplice{r: bufio.NewReader(r), file: r}

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

	defer ds.file.Close()

	ds.decode(p)

	return p, nil
}
