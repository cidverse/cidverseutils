package ci

import (
	"os"
)

// IsCI checks if the current environment is a CI environment
func IsCI() (val bool) {
	value, found := os.LookupEnv("CI")
	if found && value == "true" {
		return true
	}

	return false
}
