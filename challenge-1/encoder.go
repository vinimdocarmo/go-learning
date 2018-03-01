package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

// EncoderSplice represents the structure that is going to encode
// a splice formmatted text to a .splice file
type EncoderSplice struct {
	file *os.File
	buf  *bytes.Buffer
}

func (es EncoderSplice) encode() error {
	err := binary.Write(es.buf, binary.BigEndian, []byte("SPLICE"))

	if err != nil {
		return err
	}

	err = es.writeSize()

	if err != nil {
		return err
	}

	err = es.writeVersion()

	if err != nil {
		return err
	}

	err = es.writeTempo()

	if err != nil {
		return err
	}

	return nil
}

func (es EncoderSplice) writeSize() error {
	fstat, err := es.file.Stat()

	if err != nil {
		return err
	}

	err = binary.Write(es.buf, binary.BigEndian, fstat.Size())

	if err != nil {
		return err
	}

	return nil
}

func (es EncoderSplice) writeVersion() error {
	var version string
	versionSize := 32

	fmt.Fscanf(es.file, "Saved with HW Version: %s", &version)

	err := binary.Write(es.buf, binary.BigEndian, []byte(version))

	if err != nil {
		return err
	}

	remainingSize := versionSize - len(version)
	zeros := bytes.Repeat([]byte{0x00}, remainingSize)

	err = binary.Write(es.buf, binary.BigEndian, zeros)

	if err != nil {
		return err
	}

	return nil
}

func (es EncoderSplice) writeTempo() error {
	var tempo float32

	fmt.Fscanf(es.file, "Tempo: %g", &tempo)

	err := binary.Write(es.buf, binary.LittleEndian, tempo)

	if err != nil {
		return err
	}

	return nil
}

func newEncoderSplice(filename string) (EncoderSplice, error) {
	r, err := os.OpenFile(filename, os.O_RDONLY, 0666)

	var es EncoderSplice

	if err != nil {
		return es, err
	}

	es = EncoderSplice{file: r, buf: new(bytes.Buffer)}

	return es, nil
}

// EncodeFile encode a formatted drum file into a slice of bytes
func EncodeFile(path string) ([]byte, error) {
	es, err := newEncoderSplice(path)

	if err != nil {
		return nil, err
	}

	defer es.file.Close()

	err = es.encode()

	if err != nil {
		return nil, err
	}

	return es.buf.Bytes(), nil
}
