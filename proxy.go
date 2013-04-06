package gowl

type proxy struct {
	object  Object
	display *Display
	//queue *EventQueue
	flags    uint32
	refcount int
	userData interface{}
}
