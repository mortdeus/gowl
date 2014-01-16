package client

type Proxy interface {
	Create()
	Destroy()
	AddListener(listener *func(), data interface{}) error
	GetListener() Listener
	AddDispatcher(dispatcher func(), implementation, data interface{}) error
	Marshal(opcode ...uint)
	MarshalArray(opcode uint, args ...interface{})
	SetUserData(data interface{})
	GetUserData() interface{}
	Getid() uint
	GetType() string
	SetQueue(q *EventQueue)
}
