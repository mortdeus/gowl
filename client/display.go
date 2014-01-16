package client

import (
	"os"
)

type Display struct{}

func (d *Display) Connect(name string)           {}
func (d *Display) ConnectToFile(f os.File) error {}

func (d *Display) Disconnect() {}

func (d *Display) Getfd() {}

func (d *Display) Roundtrip() int {}

func (d *Display) ReadEvents() error  {}
func (d *Display) PrepareRead() error {}
func (d *Display) PrepareReadQueue(q *EventQueue) int
func (d *Display) CancelRead() error {}

func (d *Display) CreateQueue() *EventQueue  {}
func (d *Display) DispatchQueue() int        {}
func (d *Display) DispatchQueuePending() int {}

func (d *Display) Dispatch() int        {}
func (d *Display) DispatchPending() int {}

func (d *Display) Flush() int {}
