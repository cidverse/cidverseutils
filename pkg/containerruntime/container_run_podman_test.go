package containerruntime

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cidverse/cidverseutils/pkg/cihelper"
	"github.com/stretchr/testify/assert"
)

func TestPodmanSetParamsInteractive(t *testing.T) {
	if !cihelper.IsInteractiveTerminal() {
		return
	}
	container := Container{}
	_ = os.Unsetenv("CI")

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "-ti", "params should include -ti")
}

func TestPodmanSetParamsCI(t *testing.T) {
	container := Container{}
	_ = os.Setenv("CI", "true")

	containerCmd, _ := container.GetRunCommand("podman")
	assert.NotContains(t, containerCmd, "-ti", "params should not include -ti")
}

func TestPodmanSetName(t *testing.T) {
	container := Container{
		Name: "testCase",
	}

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "--name \"testCase\"", "name not set correctly")
}

func TestPodmanSetEntrypoint(t *testing.T) {
	entrypoint := "/bin/test"
	container := Container{
		Entrypoint: &entrypoint,
	}

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "--entrypoint \"/bin/test\"", "entrypoint not set correctly")
}

func TestPodmanSetEnvironment(t *testing.T) {
	container := Container{}
	container.AddEnvironmentVariable("DEBUG", "true")

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, fmt.Sprintf("-e %s=%s", "DEBUG", strconv.Quote("true")), "env not set correctly")
}

func TestPodmanPublishPort(t *testing.T) {
	container := Container{
		ContainerPorts: []ContainerPort{
			{
				Source: 80,
				Target: 80,
			},
		},
	}

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, fmt.Sprintf("-p %d:%d", 80, 80), "publish port not set correctly")
}

func TestPodmanSetWorkingDirectory(t *testing.T) {
	container := Container{
		WorkingDirectory: "/home/app",
	}

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, fmt.Sprintf("-w %s", strconv.Quote("/home/app")), "workdir not set correctly")
}

func TestPodmanAddVolume(t *testing.T) {
	container := Container{}
	container.AddVolume(ContainerMount{MountType: "directory", Source: "/root", Target: "/root"})

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "-v \"/root:/root\"", "mount not set correctly")
}

func TestPodmanAddVolumeReadOnly(t *testing.T) {
	container := Container{}
	container.AddVolume(ContainerMount{MountType: "directory", Source: "/root", Target: "/root", Mode: ReadMode})

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "-v \"/root:/root:ro\"", "mount not set correctly")
}

func TestPodmanSetUserArgs(t *testing.T) {
	container := Container{
		UserArgs: "--link hello:world",
	}

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "--link hello:world", "user args nto set correctly")
}

func TestPodmanSetImage(t *testing.T) {
	container := Container{
		Image: "alpine:latest",
	}

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "alpine:latest", "image not set correctly")
}

func TestPodmanSetCommand(t *testing.T) {
	container := Container{
		CommandShell: "sh",
		Command:      "printenv",
	}

	containerCmd, _ := container.GetRunCommand("podman")
	assert.Contains(t, containerCmd, "\"/usr/bin/env\" \"sh\" \"-c\" \"printenv\"", "container command invalid")
}
