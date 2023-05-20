package containerruntime

import (
	"strings"
)

// ToUnixPath turns a windows path into a unix path
func ToUnixPath(path string) string {
	driveLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, element := range driveLetters {
		if strings.HasPrefix(path, element+":\\") {
			path = strings.Replace(path, element+":\\", "/"+element+"/", 1)
		}
	}

	// replace windows path separator with linux path separator
	path = strings.Replace(path, "\\", "/", -1)

	return path
}
