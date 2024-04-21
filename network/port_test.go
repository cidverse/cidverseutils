package network

import (
	"net"
	"testing"
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
	port, err := FreePort()
	if err != nil {
		t.Fatal(err)
	}
	if port <= 0 {
		t.Errorf("Expected port to be greater than 0, but got %d", port)
	}
	if !IsFreePort(port) {
		t.Errorf("Expected port %d to be free, but it is not free", port)
	}
}
