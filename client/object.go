package client

import "sync"

type Object interface {
	sync.Locker
	ID() uint
	Type() struct {
		name            string
		version         int
		methods, events []message
	}
	SetUserData(data interface{})
	GetUserData() interface{}
	SetQueue(q *EventQueue)

	AddListener(Listener) error
}

type object struct {
	meta *meta
	data interface{}
	id   uint
}

type Listener interface {
	Listen()
}
type Dispatcher interface {
	Dispatch(mask uint) int
}



func (obj *object) ID() uint {

}
func (obj *object) Type() struct {
	name    string
	version int
	methods []message
	events  []message
} {

}
