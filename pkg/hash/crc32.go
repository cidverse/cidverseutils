package hash

import (
	"encoding/hex"
	"hash/crc32"
	"io"
)

// CRC32Hash returns a CRC32 hash
func CRC32Hash(r io.Reader) (string, error) {
	hashFunc := crc32.NewIEEE()
	if _, err := io.Copy(hashFunc, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(hashFunc.Sum(nil)), nil
}
