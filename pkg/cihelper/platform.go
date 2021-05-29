package cihelper

import (
	"github.com/mattn/go-isatty"
	"os"
	"os/exec"
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

// IsInteractiveTerminal checks if the current session is/supports interactive
func IsInteractiveTerminal() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) || IsWindowsTerminal() {
		return true
	}

	return false
}
