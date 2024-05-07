package hash

import (
	"strings"
)

// UUIDNoDash removes all dashes from a UUID string
func UUIDNoDash(uuid string) string {
	var result strings.Builder
	for _, char := range uuid {
		if char != '-' {
			result.WriteRune(char)
		}
	}
	return result.String()
}
