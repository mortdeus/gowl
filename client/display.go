package client

import (
	"os"
	"sync"
)

type Display interface {
	sync.Locker
	Connect()
	ConnectToFd(os.File)
	Disconnect()
	Fd() uintptr
	Flush() int
	Roundtrip() int
}
type display struct {
	sync.Mutex
	p      Proxy
	events []EventQueue
}

func (d *display) Connect(name string) {}
func (d *display) ConnectToFd(fd uint) {}
func (d *display) Disconnect()         {}
func (d *display) Fd() uint            {}
func (d *display) Roundtrip() int      { return -1 }
func (d *display) Flush() int          { return -1 }

func (d *display) Read(cancel chan bool) error        { return nil }
func (d *display) prepareRead() error                 { return nil }
func (d *display) prepareReadQueue(q *EventQueue) int { return -1 }
func (d *display) cancelRead() error                  { return nil }

type bitmask uint

const Pending bitmask = 1 << 30

/*
When the Queue bit is set, bits 0-29 are interpreted as
a uint30 index into display.events slice of EventQueues; This allows developers to
execute dispatch using an alternative EventsQueue than Display's main EventsQueue.
*/
const Queue bitmask = 1 << 31

/*
display.Dispatch's mask has reserved bits 30 and 31 for
overriding Dispatch's default functionality. This allows us to simplify wayland's API
from...

wl_display_dispatch
wl_display_dispatch_queue,
wl_display_dispatch_queue_pending,
wl_display_dispatch_pending

to a single Dispatch method that satisifies the Dispatcher interface.
*/
func (d *display) Dispatch(mode bitmask) int {
	var q DataDevice

	if mode&Queue > 0 {
		q = d.events[mask&(0x3ffffff)]
	}
	if mode&Pending > 0 {
	}

}
