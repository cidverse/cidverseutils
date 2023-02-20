package encoding

import (
	"bytes"
	"testing"
)

func TestSHA256HashEmpty(t *testing.T) {
	emptyInput := bytes.NewReader([]byte{})
	hash, err := SHA256Hash(emptyInput)
	if err != nil {
		t.Errorf("FileSHA256Hash returned an error: %s", err)
	}
	if hash != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
		t.Errorf("Expected empty hash, but got %s", hash)
	}
}

func TestSHA256HashWithContent(t *testing.T) {
	input := bytes.NewReader([]byte("sample data"))
	hash, err := SHA256Hash(input)
	if err != nil {
		t.Errorf("FileSHA256Hash returned an error: %s", err)
	}
	if hash != "f107aac59dff1d49ebfedb7f03877eaa0297f9a7d3cff26edfc75406f222256d" {
		t.Errorf("Expected hash 'f107aac59dff1d49ebfedb7f03877eaa0297f9a7d3cff26edfc75406f222256d', but got %s", hash)
	}
}
