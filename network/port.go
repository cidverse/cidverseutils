package network

import (
	"net"
	"strconv"
)

// IsFreePort determines if a specified port is available for use.
func IsFreePort(port int) bool {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return false
	}

	_ = listen.Close()
	return true
}

// FreePort identifies and returns an available port or an error if no ports are available.
func FreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	defer func(listen *net.TCPListener) {
		_ = listen.Close()
	}(listen)
	return listen.Addr().(*net.TCPAddr).Port, nil
}
