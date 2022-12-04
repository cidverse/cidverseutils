package containerruntime

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/cidverse/cidverseutils/pkg/cihelper"
)

func setTerminalParameters(shellCommand *bytes.Buffer) {
	if cihelper.IsCIEnvironment() {
		// env variable CI is set, we can't use --tty or --interactive here
	} else if cihelper.IsInteractiveTerminal() {
		shellCommand.WriteString("-ti ") // --tty --interactive
	}
}

func publishPorts(shellCommand *bytes.Buffer, publish *[]ContainerPort) {
	for _, publishVariable := range *publish {
		shellCommand.WriteString(fmt.Sprintf("-p %d:%d ", publishVariable.Source, publishVariable.Target))
	}
}

func setEnvironmentVariables(shellCommand *bytes.Buffer, environment *[]EnvironmentProperty) {
	for _, envVariable := range *environment {
		shellCommand.WriteString(fmt.Sprintf("-e %s=%s ", envVariable.Name, strconv.Quote(envVariable.Value)))
	}
}

// ToUnixPath turns a windows path into a unix path
func ToUnixPath(path string) string {
	driveLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, element := range driveLetters {
		if strings.HasPrefix(path, element+":\\") {
			path = strings.Replace(path, element+":\\", "/"+element+"/", 1)
		}
	}

	// replace windows path separator with linux path separator
	path = strings.Replace(path, "\\", "/", -1)

	return path
}
