package server

import (
	"net"
	"os"
	"runtime"
	//"strings"
	//"github.com/mortdeus/gowl"
)

type Display struct {
	net.Listener
	id, serial uint
}

func NewDisplay() (*Display, error) {
	d := new(Display)
	d.id = 1
	v := os.Getenv("GOWL_DISPLAY")
	if v == "" {
		v = runtime.GOROOT() + "/gowl.sock"
	}
	if l, err := net.Listen("unix", v); err != nil {
		return nil, err
	} else {
		d.Listener = l

	}
	return d, nil

}
func (d *Display) Destroy() error {
	if err := d.Listener.Close(); err != nil {
		return err
	}
	return nil
}

func (d *Display) GetEventLoop() {}

func (d *Display) Terminate() {

}
func (d *Display) Run() {

}
func (d *Display) FlushClients() {

}
func (d *Display) GetSerial() uint {
	return d.serial
}
func (d *Display) NextSerial() uint {
	d.serial++
	return d.serial

}
func (d *Display) AddDestroyListener(l *Listener) {

}
func (d *Display) GetDestroyListener(notify func()) {

}
func (d *Display) AddShmFormat(fmt uint)           {}
func (d *Display) GetAdditionalShmFormats() []uint { return nil }
