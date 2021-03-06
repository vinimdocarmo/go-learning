package drum

import (
	"bytes"
	"path"
	"testing"
)

func TestEncodeFile(t *testing.T) {
	tData := []struct {
		path   string
		output []byte
	}{
		{
			"formatted_1.txt",
			[]byte{
				0x53, 0x50, 0x4c, 0x49, 0x43, 0x45, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc5, 0x30, 0x2e,
				0x38, 0x30, 0x38, 0x2d, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0xf0, 0x42, 0x00, 0x00, 0x00, 0x00, 0x04, 0x6b, 0x69, 0x63, 0x6b, 0x01, 0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05,
				0x73, 0x6e, 0x61, 0x72, 0x65, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x04, 0x63, 0x6c, 0x61, 0x70, 0x00, 0x00,
				0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00,
				0x00, 0x00, 0x07, 0x68, 0x68, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00, 0x08, 0x68,
				0x68, 0x2d, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01, 0x05, 0x00, 0x00, 0x00, 0x07, 0x63, 0x6f, 0x77, 0x62,
				0x65, 0x6c, 0x6c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
				0x00, 0x00, 0x00,
			},
		},
	}

	for _, exp := range tData {
		encoded, err := EncodeFile(path.Join("fixtures", exp.path))
		if err != nil {
			t.Fatalf("something went wrong encoding %s - %v", exp.path, err)
		}
		if ok := bytes.Equal(encoded, exp.output); !ok {
			t.Logf("encoded:\n%#v\n", encoded)
			t.Logf("expected:\n%#v\n", exp.output)
			t.Fatalf("%s wasn't encoded as expect.\nGot:\n%s\nExpected:\n%s",
				exp.path, encoded, exp.output)
		}
	}
}
