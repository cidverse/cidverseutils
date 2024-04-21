package containerruntime

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/cidverse/cidverseutils/pkg/cihelper"
)

/**
 * Detect Podman
 */
func IsPodman() bool {
	return cihelper.IsExecutableInPath("podman")
}

/**
 * Detect Docker native
 */
func IsDockerNative() bool {
	return cihelper.IsExecutableInPath("docker")
}

/**
 * Fix escaping character
 */
func sanitizeCommand(commandShell string, command string) string {
	// Shell (wrap the command within the container into a shell)
	if commandShell == "powershell" {
		// would be used for windows containers, never tested though
		command = fmt.Sprintf("powershell %s", command)
	} else if commandShell == "sh" {
		if runtime.GOOS == "windows" {
			command = fmt.Sprintf("\"/usr/bin/env\" \"sh\" \"-c\" \"%s\"", strings.Replace(command, "\"", "`\"", -1))
		} else {
			command = fmt.Sprintf("\"/usr/bin/env\" \"sh\" \"-c\" \"%s\"", strings.Replace(command, "\"", "\\\"", -1))
		}
	} else if commandShell == "bash" {
		if runtime.GOOS == "windows" {
			command = fmt.Sprintf("\"/usr/bin/env\" \"bash\" \"-l\" \"-c\" \"%s\"", strings.Replace(command, "\"", "`\"", -1))
		} else {
			command = fmt.Sprintf("\"/usr/bin/env\" \"bash\" \"-l\" \"-c\" \"%s\"", strings.Replace(command, "\"", "\\\"", -1))
		}
	}

	return command
}

/**
 * CLI Command Passthru with input/output
 */
func systemExec(command string) error {
	// Run Command
	if runtime.GOOS == "linux" {
		cmd := exec.Command("/usr/bin/env", "sh", "-c", command)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell", command)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else if runtime.GOOS == "darwin" {
		cmd := exec.Command("sh", "-c", command)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
