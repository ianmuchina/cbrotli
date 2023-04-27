package cbrotli

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"fmt"
	"io"
	"testing"
)

//go:embed cdnjs.cloudflare.com/ajax/libs/react/18.2.0/umd/react.production.min.js.br
var encoded_br []byte

//go:embed cdnjs.cloudflare.com/ajax/libs/react/18.2.0/umd/react.production.min.js.sbr
var encoded_sbr []byte

//go:embed cdnjs.cloudflare.com/ajax/libs/react/18.2.0/umd/react.production.min.js
var content []byte

//go:embed cdnjs.cloudflare.com/ajax/libs/react/17.0.2/umd/react.production.min.js
var dict_content []byte

func TestDecodeCustomDict(t *testing.T) {
	decoded1, err := Decode(encoded_sbr)

	if err != nil {
		fmt.Println("err")
	}
	decoded2, err := DecodeWithCustomDictionary(encoded_sbr, dict_content)
	fmt.Println("===")
	fmt.Printf("decoded1: %x\n", sha256.Sum256(decoded1))
	fmt.Printf("decoded2: %x\n", sha256.Sum256(decoded2))
	fmt.Printf("original: %x\n", sha256.Sum256(content))
	fmt.Println("===")
	// TODO: Error checks
	if err != nil {
		t.Errorf("Decode: %v", err)
	}
	if !bytes.Equal(decoded2, content) {
		t.Errorf(""+
			"Decode content:\n"+
			"%q\n"+
			"want:\n"+
			"<%d bytes>",
			decoded2, len(content))
	}
}

func TestEncoderLargeInput2(t *testing.T) {
	input := content
	out := bytes.Buffer{}
	e := NewWriter(&out, WriterOptions{
		Quality:               5,
		UsePreparedDictionary: true,
		PreparedDictionary:    dict_content,
	})

	in := bytes.NewReader(input)
	n, err := io.Copy(e, in)
	if err != nil {
		t.Errorf("Copy Error: %v", err)
	}
	if int(n) != len(input) {
		t.Errorf("Copy() n=%v, want %v", n, len(input))
	}
	if err := e.Close(); err != nil {
		t.Errorf("Close Error after copied %d bytes: %v", n, err)
	}
	if err := checkCompressedData2(out.Bytes(), input, dict_content); err != nil {
		t.Error(err)
	}
}
