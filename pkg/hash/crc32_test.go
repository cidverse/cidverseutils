package hash

import (
	"bytes"
	"testing"
)

func TestCRC32HashEmpty(t *testing.T) {
	emptyInput := bytes.NewReader([]byte{})
	hash, err := CRC32Hash(emptyInput)
	if err != nil {
		t.Errorf("CRC32Hash returned an error: %s", err)
	}
	expectedHash := "00000000"
	if hash != expectedHash {
		t.Errorf("Expected empty hash, but got %s", hash)
	}
}

func TestCRC32HashWithContent(t *testing.T) {
	input := bytes.NewReader([]byte("sample data"))
	hash, err := CRC32Hash(input)
	if err != nil {
		t.Errorf("CRC32Hash returned an error: %s", err)
	}
	expectedHash := "6a47a630"
	if hash != expectedHash {
		t.Errorf("Expected hash '%s', but got %s", expectedHash, hash)
	}
}
