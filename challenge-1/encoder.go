package drum

import (
	"io"
	"os"
)

// EncoderSplice represents the structure that is going to encode
// a splice formmatted text to a .splice file
type EncoderSplice struct {
	w    io.Writer
	file *os.File
}
