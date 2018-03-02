package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// EncoderSplice represents the structure that is going to encode
// a splice formmatted text to a .splice file
type EncoderSplice struct {
	r           io.Reader
	buf         *bytes.Buffer
	bytesWriten int
}

func (es EncoderSplice) encode() error {
	err := binary.Write(es.buf, binary.BigEndian, []byte("SPLICE"))

	if err != nil {
		return err
	}

	es.bytesWriten += 6 //6 bytes that represents the string "SPLICE"

	err = es.writeSize()

	if err != nil {
		return err
	}

	es.bytesWriten += 8 //8 bytes that represents the size of size itself

	err = es.writeVersion()

	if err != nil {
		return err
	}

	es.bytesWriten += 32 //32 bytes that represents the splice version

	err = es.writeTempo()

	if err != nil {
		return err
	}

	es.bytesWriten += 4 //4 bytes that represents the tempo

	return nil
}

func (es EncoderSplice) writeSize() error {
	return nil
}

// writeVersion writes into the buffer a big endian 32-bytes representing the splice version
func (es EncoderSplice) writeVersion() error {
	var version string
	versionSize := 32

	fmt.Fscanf(es.r, "Saved with HW Version: %s", &version)

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

	fmt.Fscanf(es.r, "Tempo: %g", &tempo)

	err := binary.Write(es.buf, binary.LittleEndian, tempo)

	if err != nil {
		return err
	}

	return nil
}

func newEncoderSplice(filename string) (EncoderSplice, error) {
	f, err := os.Open(filename)

	var es EncoderSplice

	if err != nil {
		return es, err
	}

	defer f.Close()

	br := new([]byte)

	_, err = f.Read(*br)

	if err != nil {
		return es, err
	}

	r := bytes.NewReader(*br)

	es = EncoderSplice{bytesWriten: 0, r: r, buf: new(bytes.Buffer)}

	return es, nil
}

// EncodeFile encode a formatted drum file into a slice of bytes
func EncodeFile(path string) ([]byte, error) {
	es, err := newEncoderSplice(path)

	if err != nil {
		return nil, err
	}

	err = es.encode()

	if err != nil {
		return nil, err
	}

	return es.buf.Bytes(), nil
}
