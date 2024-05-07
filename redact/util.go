package redact

import (
	"encoding/base64"
)

func isBase64(s string) bool {
	if s == "" {
		return false
	}

	if len(s)%4 != 0 { // strings with a length not divisible by 4 are not valid base64
		return false
	}

	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}
