package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// SHA256Hash returns a SHA256 hash
func SHA256Hash(r io.Reader) (string, error) {
	hashFunc := sha256.New()
	if _, err := io.Copy(hashFunc, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(hashFunc.Sum(nil)), nil
}

// SHA256HashFile calculates the SHA256 hash of a file
func SHA256HashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return CRC32Hash(file)
}
