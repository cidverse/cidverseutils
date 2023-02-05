package encoding

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
)

// GZIPBase64EncodeBytes returns a string representation of the input byte slice after compressing and encoding it using GZIP and base64
func GZIPBase64EncodeBytes(input []byte) (string, error) {
	var sarifGzipped bytes.Buffer
	gz := gzip.NewWriter(&sarifGzipped)
	if _, err := gz.Write(input); err != nil {
		return "", fmt.Errorf("failed to gzip input")
	}
	if err := gz.Close(); err != nil {
		return "", fmt.Errorf("failed to gzip input")
	}
	return base64.URLEncoding.EncodeToString(sarifGzipped.Bytes()), nil
}
