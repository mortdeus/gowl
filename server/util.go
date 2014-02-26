package server

import (
	"container/list"
	"unsafe"
)

type Message struct {
	name      string
	signature string
	types     []Interface
}

type Interface struct {
	name    string
	version int
	methods []Message
	events  []Message
}

type Object struct {
	iface          *Interface
	implementation interface{}
	id             uint
}
type List struct {
	list.List
}

