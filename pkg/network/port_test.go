package network

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFreePort(t *testing.T) {
	// bind a port to make it unavailable
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	port := ln.Addr().(*net.TCPAddr).Port

	// Check if the bound port is not free
	if IsFreePort(port) {
		t.Errorf("Expected port %d to be not free, but it is free", port)
	}

	// Check if a different port is free
	if !IsFreePort(port + 1) {
		t.Errorf("Expected port %d to be free, but it is not free", port+1)
	}
}

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	assert.Nil(t, err)
	assert.True(t, port > 0)
	assert.True(t, IsFreePort(port))
}
