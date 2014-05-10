package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var Display = new(display)

type display struct {
	sync.Mutex
	conn   net.Conn
	events []EventQueue
}

func (d *display) Connect(name string) (err error) {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		return fmt.Errorf("XDG_RUNTIME_DIR is not defined.")
	}
	if name == "" {
		name = os.Getenv("WAYLAND_DISPLAY")
	}
	if name == "" {
		name = "wayland-0"
	}
	addr, err := net.ResolveUnixAddr("unix", runtimeDir+"/"+name)
	if err != nil {
		return err
	}
	d.conn, err = net.DialUnix("unix", nil, addr)
	if err != nil {
		return err
	}

	d.conn.(*net.UnixConn).SetReadBuffer(4096)
	d.conn.(*net.UnixConn).SetWriteBuffer(4096)
	return nil
}
func (d *display) ConnectTo(f *os.File) (err error) {
	d.conn, err = net.FileConn(f)
	return
}
func (d *display) Disconnect() {}

func (d *display) File() (*os.File, error) { return d.conn.(*net.UnixConn).File() }

func (d *display) Roundtrip() int { return -1 }

func (d *display) Flush() int {
	return 0
}

func (d *display) Read() error {

	fmt.Println("writing to server")
	_, err := fmt.Fprintln(d.conn, "foobar")
	if err != nil {
		return err
	}
	fmt.Println("wrote to server")

	d.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	msgs := make([]byte, 4096)

	br := bufio.NewReader(d.conn)
	if br.Buffered() == 0 {
		return nil
	}
	if _, err := br.Read(msgs); err != nil {
		return err
	}
	for i := 0; i < len(msgs)%4; i++ {
		if len(msgs[i:]) > 4 {
			fmt.Println(msgs[i : i+4])
		} else {
			return fmt.Errorf("alignment is off! %v", msgs)
		}
	}
	return nil
}
func (d *display) prepareRead() error                 { return nil }
func (d *display) prepareReadQueue(q *EventQueue) int { return -1 }
func (d *display) CancelRead() error                  { return nil }

func (d *display) Dispatch() int {
	return 0
}
