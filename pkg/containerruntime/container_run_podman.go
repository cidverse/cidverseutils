package containerruntime

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// GetPodmanCommand renders the command needed the run the container using podman
func (c *Container) GetPodmanCommand() string {
	var shellCommand bytes.Buffer

	// build command
	shellCommand.WriteString("podman run --rm ")
	// - terminal
	setTerminalParameters(&shellCommand)
	// - name
	if len(c.name) > 0 {
		shellCommand.WriteString(fmt.Sprintf("--name %s ", strconv.Quote(c.name)))
	}
	// - entrypoint
	if c.entrypoint != nil {
		if len(*c.entrypoint) > 0 {
			shellCommand.WriteString(fmt.Sprintf("--entrypoint %s ", strconv.Quote(*c.entrypoint)))
		} else {
			shellCommand.WriteString("--entrypoint= ")
		}
	}
	// - environment variables
	setEnvironmentVariables(&shellCommand, &c.environment)
	// - publish ports
	publishPorts(&shellCommand, &c.containerPorts)
	// - capabilities / privileged
	if c.privileged == true {
		shellCommand.WriteString(fmt.Sprintf("--privileged "))
	} else {
		for _, cap := range c.capabilities {
			shellCommand.WriteString(fmt.Sprintf("--cap-add %s ", strconv.Quote(cap)))
		}
	}
	// - set working directory
	if len(c.workingDirectory) > 0 {
		shellCommand.WriteString(fmt.Sprintf("--workdir %s ", strconv.Quote(c.workingDirectory)))
	}
	// - volume mounts
	for _, containerMount := range c.volumes {
		if containerMount.MountType == "directory" {
			var mountSource = containerMount.Source
			var mountTarget = containerMount.Target

			suffix := ""
			if containerMount.Mode == ReadMode {
				suffix = ":ro"
			}

			shellCommand.WriteString(fmt.Sprintf("-v %q ", mountSource+":"+mountTarget+suffix))
		} else if containerMount.MountType == "volume" {
			var mountSource = containerMount.Source
			var mountTarget = containerMount.Target
			mountSourceDir := filepath.Join(os.TempDir(), "podman-volume", mountSource)
			_ = os.MkdirAll(mountSourceDir, os.ModePerm)

			shellCommand.WriteString(fmt.Sprintf("-v %s ", strconv.Quote(mountSourceDir+":"+mountTarget)))
		}
	}
	// - userArgs
	if len(c.userArgs) > 0 {
		shellCommand.WriteString(c.userArgs + " ")
	}
	// - image
	shellCommand.WriteString(fmt.Sprintf("%s ", c.image))
	// - command to run inside of the container
	shellCommand.WriteString(sanitizeCommand(c.commandShell, c.command))

	return shellCommand.String()
}
