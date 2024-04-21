package exec

import (
	"os"

	"github.com/mattn/go-isatty"
)

// IsWindowsTerminal checks if the current execution context is the new windows terminal
func IsWindowsTerminal() bool {
	_, isPresent := os.LookupEnv("WT_SESSION")
	return isPresent
}

// IsInteractiveTerminal checks if the current session is/supports interactive
func IsInteractiveTerminal() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) || IsWindowsTerminal() {
		return true
	}

	return false
}
