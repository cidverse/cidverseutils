package filesystem

import (
	"os"
)

// FileExists checks if the file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

// DirectoryExists checks if the dir exists
func DirectoryExists(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return info.IsDir()
}
