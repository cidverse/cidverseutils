package ci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
