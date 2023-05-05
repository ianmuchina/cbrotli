package cbrotli

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"testing"
)

//go:embed tests/file1.txt.sbr
var encoded_br []byte

//go:embed tests/file1.txt.sbr
var encoded_sbr []byte

//go:embed tests/file1.txt
var content []byte

//go:embed tests/dict.txt
var dict_content []byte

func TestDecodeCustomDict(t *testing.T) {
	// Decode with shared Dictionary
	decoded2, err2 := DecodeWithCustomDictionary(encoded_sbr, dict_content)
	if err2 != nil {
		fmt.Printf("err a, %v\n", err2)
	}
	// Decode without shared dictionary. Should have error
	_, err1 := Decode(encoded_sbr)

	// Dictionary error
	if err1 == fmt.Errorf("cbrotli: DICTIONARY") {
		fmt.Printf("yes, err b ;;%v;;\n", err1)
	} else if err1 != nil {
		// some other error
		log.Fatal(err1)
	}

	// fmt.Printf("%s", err1)
	// fmt.Println("===")
	// fmt.Printf("decoded1: %x\n", sha256.Sum256(decoded1))
	// fmt.Printf("decoded2: %x\n", sha256.Sum256(decoded2))
	// fmt.Printf("original: %x\n", sha256.Sum256(content))
	// fmt.Println("===")

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
