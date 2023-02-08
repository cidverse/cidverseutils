package cihelper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCiEnvironmentTrue(t *testing.T) {
	_ = os.Setenv("CI", "true")
	assert.Equal(t, true, IsCIEnvironment())
}

func TestCiEnvironmentFalse(t *testing.T) {
	_ = os.Unsetenv("CI")
	assert.Equal(t, false, IsCIEnvironment())
}

func TestToUnixPath(t *testing.T) {
	args := "H:/example3/test"
	unixPath := ToUnixPath(args)

	assert.Equal(t, "/H/example3/test", unixPath)
}

func TestToUnixPathArgs(t *testing.T) {
	args := "H:\\example1 H:\\example2 H:/example3/test"
	unixPath := ToUnixPathArgs(args)

	assert.Equal(t, "/H/example1 /H/example2 /H/example3/test", unixPath)
}
