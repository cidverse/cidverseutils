package containerruntime

import (
	"bytes"
	"fmt"
	"strconv"
)

// GetDockerCommand renders the command needed the run the container using docker
func (c *Container) GetDockerCommand() string {
	var shellCommand bytes.Buffer

	// build command
	shellCommand.WriteString("docker run --rm ")
	// - interactive, tty
	if c.Interactive {
		shellCommand.WriteString("-i ")
	}
	if c.TTY {
		shellCommand.WriteString("-t ")
	}
	// - name
	if c.Name != "" {
		shellCommand.WriteString(fmt.Sprintf("--name %s ", strconv.Quote(c.Name)))
	}
	// - user
	if c.User != "" {
		shellCommand.WriteString(fmt.Sprintf("-u %s ", strconv.Quote(c.User)))
	}
	// - entrypoint
	if c.Entrypoint != nil {
		if len(*c.Entrypoint) > 0 {
			shellCommand.WriteString(fmt.Sprintf("--entrypoint %s ", strconv.Quote(*c.Entrypoint)))
		} else {
			shellCommand.WriteString("--entrypoint= ")
		}
	}
	// - environment variables
	for _, envVariable := range c.Environment {
		shellCommand.WriteString(fmt.Sprintf("-e %s=%s ", envVariable.Name, strconv.Quote(envVariable.Value)))
	}
	// - publish ports
	for _, publishVariable := range c.ContainerPorts {
		shellCommand.WriteString(fmt.Sprintf("-p %d:%d ", publishVariable.Source, publishVariable.Target))
	}
	// - capabilities / privileged
	if c.Privileged == true {
		shellCommand.WriteString(fmt.Sprintf("--privileged "))
	} else {
		for _, capability := range c.Capabilities {
			shellCommand.WriteString(fmt.Sprintf("--cap-add %s ", strconv.Quote(capability)))
		}
	}
	// - set working directory
	if len(c.WorkingDirectory) > 0 {
		shellCommand.WriteString(fmt.Sprintf("-w %s ", strconv.Quote(c.WorkingDirectory)))
	}
	// - volume mounts
	for _, containerMount := range c.Volumes {
		if containerMount.MountType == "directory" || containerMount.MountType == "volume" {
			var mountSource = containerMount.Source
			var mountTarget = containerMount.Target

			suffix := ""
			if containerMount.Mode == ReadMode {
				suffix = ":ro"
			}

			shellCommand.WriteString(fmt.Sprintf("-v %q ", mountSource+":"+mountTarget+suffix))
		}
	}
	// - userArgs
	if c.UserArgs != "" {
		shellCommand.WriteString(c.UserArgs + " ")
	}
	// - image
	shellCommand.WriteString(fmt.Sprintf("%s ", c.Image))
	// - command to run inside the container
	shellCommand.WriteString(sanitizeCommand(c.CommandShell, c.Command))

	return shellCommand.String()
}
