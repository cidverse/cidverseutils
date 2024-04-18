package hash

import (
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"
)

// CRC32Hash returns a CRC32 hash
func CRC32Hash(r io.Reader) (string, error) {
	hashFunc := crc32.NewIEEE()
	if _, err := io.Copy(hashFunc, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(hashFunc.Sum(nil)), nil
}

// CRC32HashFile calculates the CRC32 hash of a file
func CRC32HashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return CRC32Hash(file)
}
