package encoding

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// SHA256Hash returns a string representation of the SHA256 hash
func SHA256Hash(r io.Reader) (string, error) {
	hashFunc := sha256.New()
	if _, err := io.Copy(hashFunc, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(hashFunc.Sum(nil)), nil
}
