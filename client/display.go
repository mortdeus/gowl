package client

import (
	"os"
	"sync"
)

type Display struct {
	sync.Mutex
}

func (d *Display) Marshal()                      {}
func (d *Display) Connect(name string)           {}
func (d *Display) ConnectToFile(f os.File) error { return nil }

func (d *Display) Disconnect() {}

func (d *Display) Getfd() {}

func (d *Display) Roundtrip() int { return -1 }

func (d *Display) ReadEvents() error                  { return nil }
func (d *Display) PrepareRead() error                 { return nil }
func (d *Display) PrepareReadQueue(q *EventQueue) int { return -1 }
func (d *Display) CancelRead() error                  { return nil }

func (d *Display) CreateQueue() *EventQueue  { return nil }
func (d *Display) DispatchQueue() int        { return -1 }
func (d *Display) DispatchQueuePending() int { return -1 }

func (d *Display) Dispatch() int        { return -1 }
func (d *Display) DispatchPending() int { return -1 }

func (d *Display) Flush() int { return -1 }
