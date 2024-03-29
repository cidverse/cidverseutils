package cihelper

import (
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-isatty"
)

// IsCIEnvironment checks if this process in running as part of a CI process
func IsCIEnvironment() (val bool) {
	value, found := os.LookupEnv("CI")
	if found && value == "true" {
		return true
	}

	return false
}

// IsWindowsTerminal checks if the current execution context is the new windows terminal
func IsWindowsTerminal() bool {
	_, isPresent := os.LookupEnv("WT_SESSION")
	return isPresent
}

// IsExecutableInPath checks if a executable is available in PATH
func IsExecutableInPath(executable string) bool {
	_, err := exec.LookPath(executable)
	if err != nil {
		return false
	}

	return true
}

// IsInteractiveTerminal checks if the current session is/supports interactive
func IsInteractiveTerminal() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) || IsWindowsTerminal() {
		return true
	}

	return false
}

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
