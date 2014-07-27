package client

import (
	//"encoding/binary"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
)

type Display struct {
	sync.Mutex
	Conn   net.Conn
	Events []EventQueue
}

func (d *Display) Sync() {
	fmt.Fprint(d.Conn, uint32(1), uint32(3)<<16|(uint32(0)))
}

func (d *Display) Connect(name string) (err error) {
	if name == "" {
		name = os.Getenv("GOWL_DISPLAY")
	}
	if name == "" {
		name = runtime.GOROOT() + "/gowl.sock"
	}
	d.Conn, err = net.Dial("unix", name)
	if err != nil {
		return err
	}

	d.Conn.(*net.UnixConn).SetReadBuffer(4096)
	d.Conn.(*net.UnixConn).SetWriteBuffer(4096)
	return nil
}
func (d *Display) ConnectTo(f *os.File) (err error) {
	d.Conn, err = net.FileConn(f)
	return
}
func (d *Display) Disconnect() error {
	return d.Conn.Close()
}

func (d *Display) File() (*os.File, error) { return d.Conn.(*net.UnixConn).File() }

func (d *Display) Roundtrip() int { return -1 }

func (d *Display) Flush() int {
	return 0
}
func (d *Display) Read(p []byte) (n int, err error) {
	return d.Conn.Read(p)
}
func (d *Display) Write(p []byte) (n int, err error) {
	return d.Conn.Write(p)
}
func (d *Display) prepareRead() error                 { return nil }
func (d *Display) prepareReadQueue(q *EventQueue) int { return -1 }
func (d *Display) CancelRead() error                  { return nil }

func (d *Display) Dispatch() int {
	return 0
}
