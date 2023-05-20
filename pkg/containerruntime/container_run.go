package containerruntime

import (
	"bytes"
	"errors"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/cidverse/cidverseutils/pkg/cihelper"
	"github.com/cidverse/cidverseutils/pkg/collection"
)

// Container provides all methods to interact with the container runtime
type Container struct {
	Name             string
	Image            string
	Entrypoint       *string
	CommandShell     string
	Command          string
	WorkingDirectory string
	Volumes          []ContainerMount
	Environment      []EnvironmentProperty
	ContainerPorts   []ContainerPort
	Capabilities     []string
	User             string
	UserArgs         string
	Privileged       bool
	Interactive      bool
	TTY              bool
}

func Create() Container {
	return Container{
		Name:             "",
		Image:            "",
		Entrypoint:       nil,
		CommandShell:     "",
		Command:          "",
		WorkingDirectory: "",
		Volumes:          []ContainerMount{},
		Environment:      []EnvironmentProperty{},
		ContainerPorts:   []ContainerPort{},
		Capabilities:     []string{},
		User:             "",
		UserArgs:         "",
		Privileged:       false,
		Interactive:      false,
		TTY:              false,
	}
}

// AutoTerminalParameters automatically sets the terminal parameters
func (c *Container) AutoTerminalParameters() {
	if cihelper.IsInteractiveTerminal() && !cihelper.IsCIEnvironment() {
		c.Interactive = true
		c.TTY = true
	}
}

// AddVolume mounts a directory into a container
func (c *Container) AddVolume(mount ContainerMount) {
	mount.Target = ToUnixPath(mount.Target)
	c.Volumes = append(c.Volumes, mount)
}

// AddCacheMount adds a cache mount to the container
func (c *Container) AddCacheMount(name string, sourcePath string, targetPath string) {
	c.AddVolume(ContainerMount{MountType: "directory", Source: ToUnixPath(sourcePath), Target: targetPath})
	c.AddEnvironmentVariable("cache_"+name+"_source", ToUnixPath(sourcePath))
	c.AddEnvironmentVariable("cache_"+name+"_target", targetPath)
}

// AllowContainerRuntimeAcccess allows the container to access the container runtime
func (c *Container) AllowContainerRuntimeAcccess() {
	socketPath := "/var/run/docker.sock"
	if runtime.GOOS == "windows" {
		// docker desktop
		if IsDockerNative() {
			socketPath = "//var/run/docker.sock"
		}
	}

	c.AddVolume(ContainerMount{MountType: "directory", Source: socketPath, Target: "/var/run/docker.sock"})
}

// AddContainerPorts adds multiple published ports
func (c *Container) AddContainerPorts(ports []string) {
	for _, p := range ports {
		pair := strings.SplitN(p, ":", 2)
		sourcePort, _ := strconv.Atoi(pair[0])
		targetPort, _ := strconv.Atoi(pair[1])

		c.ContainerPorts = append(c.ContainerPorts, ContainerPort{Source: sourcePort, Target: targetPort})
	}
}

// AddEnvironmentVariable adds a environment variable
func (c *Container) AddEnvironmentVariable(name string, value string) {
	c.Environment = append(c.Environment, EnvironmentProperty{Name: name, Value: value})
}

// AddEnvironmentVariables adds multiple environment variables
func (c *Container) AddEnvironmentVariables(variables []string) {
	for _, e := range variables {
		pair := strings.SplitN(e, "=", 2)
		var envName = pair[0]
		var envValue = pair[1]

		c.AddEnvironmentVariable(envName, envValue)
	}
}

// AddAllEnvironmentVariables adds all environment variables, but filters a few irrelevant ones (like PATH, HOME, etc.)
func (c *Container) AddAllEnvironmentVariables() {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		var envName = pair[0]
		var envValue = pair[1]

		// filter vars
		var systemVars = []string{
			"",
			// unix
			"_",
			"PWD",
			"OLDPWD",
			"PATH",
			"HOME",
			"HOSTNAME",
			"TERM",
			"SHLVL",
			// windows
			"PROGRAMDATA",
			"PROGRAMFILES",
			"PROGRAMFILES(x86)",
			"PROGRAMW6432",
			"COMMONPROGRAMFILES",
			"COMMONPROGRAMFILES(x86)",
			"COMMONPROGRAMW6432",
			"PATHEXT",
			// proxy
			"HTTP_PROXY",
			"HTTPS_PROXY",
		}
		isExcluded, _ := collection.InArray(strings.ToUpper(envName), systemVars)
		// recent issue of 2009 about git bash / mingw setting invalid unix variables with `var(86)=...`
		isInvalidName := strings.Contains(envName, "(") || strings.Contains(envName, ")")
		if !isExcluded && !isInvalidName {
			c.AddEnvironmentVariable(envName, envValue)
		}
	}
}

// DetectRuntime returns the first available container runtime
func (c *Container) DetectRuntime() string {
	// autodetect container runtime
	if IsPodman() {
		return "podman"
	} else if IsDockerNative() {
		return "docker"
	}

	return "unknown"
}

// GetPullCommand gets the command to pull the required image
func (c *Container) GetPullCommand(runtime string) (string, error) {
	// autodetect container runtime
	if runtime == "podman" {
		return "podman pull " + c.Image, nil
	} else if runtime == "docker" {
		return "docker pull " + c.Image, nil
	} else {
		return "", errors.New("No supported container runtime found (podman, docker)! [" + runtime + "]")
	}
}

// GetRunCommand gets the run command for the specified container runtime
func (c *Container) GetRunCommand(runtime string) (string, error) {
	var shellCommand bytes.Buffer

	// autodetect container runtime
	if runtime == "podman" {
		shellCommand.WriteString(c.GetPodmanCommand())
	} else if runtime == "docker" {
		shellCommand.WriteString(c.GetDockerCommand())
	} else {
		return "", errors.New("container runtime [" + runtime + "] is not supported!")
	}

	return shellCommand.String(), nil
}

// StartContainer starts the Container
func (c *Container) StartContainer() error {
	var shellCommand bytes.Buffer

	// - command
	runCmd, runCmdErr := c.GetRunCommand(c.DetectRuntime())
	if runCmdErr != nil {
		return runCmdErr
	}
	shellCommand.WriteString(runCmd)

	// execute command
	return systemExec(shellCommand.String())
}

// PullImage pulls the image for the container
func (c *Container) PullImage() error {
	pullCmd, pullCmdErr := c.GetPullCommand(c.DetectRuntime())
	if pullCmdErr == nil {
		return systemExec(pullCmd)
	} else {
		return errors.New("can't pull image")
	}
}
