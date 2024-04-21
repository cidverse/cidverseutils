package exec

import (
	"os/exec"
)

// InPath checks if an executable is in the PATH
func InPath(executable string) bool {
	_, err := exec.LookPath(executable)
	if err != nil {
		return false
	}

	return true
}
