package containerruntime

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	ciexec "github.com/cidverse/cidverseutils/exec"
)

// IsPodman checks if podman is available in the PATH
func IsPodman() bool {
	return ciexec.InPath("podman")
}

// IsDockerNative checks if docker is available in the PATH
func IsDockerNative() bool {
	return ciexec.InPath("docker")
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
