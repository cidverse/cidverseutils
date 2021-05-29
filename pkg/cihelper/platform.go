package cihelper

import (
	"os"
	"os/exec"
	"reflect"
)

// IsCIEnvironment checks if this process in running as part of a CI process
func IsCIEnvironment() (val bool) {
	value, found := os.LookupEnv("CI")
	if found && value == "true" {
		return true
	}

	return false
}

// IsMinGW checks if the current execution context is a Minimalist GNU for Windows environment (cygwin / git bash)
func IsMinGW() bool {
	value, _ := os.LookupEnv("MSYSTEM")
	if value == "MINGW64" {
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

/**
 * Checks if a object is part of a array
 */
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
