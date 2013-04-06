package gowl

import (
	"bufio"
	"log"
	"net"
	"os"
)

type Connection struct {
	unixconn *net.UnixConn
	addr     *net.UnixAddr
	reader   *bufio.Reader
	writer   *bufio.Writer
	alive    bool
}

var conn Connection

func connectToSocket(name string) {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		log.Fatalln("XDG_RUNTIME_DIR doesnt exist.")
	}
	if name == "" {
		name = os.Getenv("WAYLAND_DISPLAY")
		if name == "" {
			name = "wayland-0"
		}
	}
	c, err := net.DialUnix("unix", conn.addr, &net.UnixAddr{runtimeDir + name, "unix"})
	if err != nil {
		log.Fatalln(err.Error())
	}
	conn.reader = bufio.NewReader(c)
	conn.writer = bufio.NewWriter(c)
	conn.unixconn = c
	conn.alive = true
}
