package ci

import (
	"strings"
)

// ToUnixPath turns a windows path into a unix path
func ToUnixPath(path string) string {
	driveLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, element := range driveLetters {
		if strings.HasPrefix(path, element+":\\") {
			path = strings.Replace(path, element+":\\", "/"+element+"/", -1)
			path = strings.Replace(path, "\\", "/", -1)
		} else if strings.HasPrefix(path, element+":/") {
			path = strings.Replace(path, element+":/", "/"+element+"/", -1)
		}
	}

	// replace windows path separator with linux path separator
	path = strings.Replace(path, "\\", "/", -1)

	return path
}

// ToUnixPathArgs checks each argument and turns it into a unix path if needed
func ToUnixPathArgs(data string) string {
	argList := strings.Split(data, " ")

	for _, a := range argList {
		driveLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		for _, element := range driveLetters {
			if strings.HasPrefix(a, element+":\\") {
				data = strings.Replace(data, element+":\\", "/"+element+"/", -1)
				data = strings.Replace(data, "\\", "/", -1)
			} else if strings.HasPrefix(a, element+":/") {
				data = strings.Replace(data, element+":/", "/"+element+"/", -1)
			}
		}
	}

	return data
}
