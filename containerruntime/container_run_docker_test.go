package containerruntime

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cidverse/cidverseutils/exec"
	"github.com/stretchr/testify/assert"
)

func TestDockerSetParamsInteractive(t *testing.T) {
	if !exec.IsInteractiveTerminal() {
		return
	}
	container := Container{}
	_ = os.Unsetenv("CI")

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, "-ti", "params should include -ti")
}

func TestDockerSetParamsCI(t *testing.T) {
	container := Container{}
	_ = os.Setenv("CI", "true")

	containerCmd, _ := container.GetRunCommand("docker")
	assert.NotContains(t, containerCmd, "-ti", "params should not include -ti")
}

func TestDockerSetName(t *testing.T) {
	container := Container{
		Name: "testCase",
	}

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, "--name \"testCase\"", "name not set correctly")
}

func TestDockerSetEntrypoint(t *testing.T) {
	entrypoint := "/bin/test"
	container := Container{
		Entrypoint: &entrypoint,
	}

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, "--entrypoint \"/bin/test\"", "entrypoint not set correctly")
}

func TestDockerSetEnvironment(t *testing.T) {
	container := Container{}
	container.AddEnvironmentVariable("DEBUG", "true")

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, fmt.Sprintf("-e %s=%s", "DEBUG", strconv.Quote("true")), "env not set correctly")
}

func TestDockerPublishPort(t *testing.T) {
	container := Container{
		ContainerPorts: []ContainerPort{
			{
				Source: 80,
				Target: 80,
			},
		},
	}

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, fmt.Sprintf("-p %d:%d", 80, 80), "publish port not set correctly")
}

func TestDockerSetWorkingDirectory(t *testing.T) {
	container := Container{
		WorkingDirectory: "/home/app",
	}

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, fmt.Sprintf("-w %s", strconv.Quote("/home/app")), "workdir not set correctly")
}

func TestDockerAddVolume(t *testing.T) {
	container := Container{}
	container.AddVolume(ContainerMount{MountType: "directory", Source: "/root", Target: "/root"})

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, "-v \"/root:/root\"", "mount not set correctly")
}

func TestDockerAddVolumeReadOnly(t *testing.T) {
	container := Container{}
	container.AddVolume(ContainerMount{MountType: "directory", Source: "/root", Target: "/root", Mode: ReadMode})

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "-v \"/root:/root:ro\"", "mount not set correctly")
}

func TestDockerSetUserArgs(t *testing.T) {
	container := Container{
		UserArgs: "--link hello:world",
	}

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, "--link hello:world", "user args nto set correctly")
}

func TestDockerSetImage(t *testing.T) {
	container := Container{
		Image: "alpine:latest",
	}

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, "alpine:latest", "image not set correctly")
}

func TestDockerSetCommand(t *testing.T) {
	container := Container{
		CommandShell: "sh",
		Command:      "printenv",
	}

	containerCmd, _ := container.GetRunCommand("docker")
	assert.Contains(t, containerCmd, "\"/usr/bin/env\" \"sh\" \"-c\" \"printenv\"", "container command invalid")
}
