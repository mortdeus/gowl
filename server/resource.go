package server

type Resource struct {
	c *Client
	i *Interface
	version,
	id uint
}

func NewResource(c *Client, i *Interface, version, id uint) *Resource {
	return &Resource{c, i, version, id}
}

func (r *Resource) PostEventArray(opcode uint, args ...interface{})  {}
func (r *Resource) PostEvent(opcode ...uint)                         {}
func (r *Resource) QueueEventArray(opcode uint, args ...interface{}) {}
func (r *Resource) QueueEvent(opcode ...uint)                        {}
func (r *Resource) PostError(err error)                              {}
func (r *Resource) PostNoMemory()                                    {}
func (r *Resource) Destroy()                                         {}

func (r *Resource) GetID() uint        { return 0 }
func (r *Resource) GetClient() *Client { return nil }
func (r *Resource) GetVersion() int    { return 0 }

func (r *Resource) GetUserData() interface{}     { return nil }
func (r *Resource) SetUserData(data interface{}) {}

func (r *Resource) SetDestructor(d func())                     {}
func (r *Resource) GetDestroyListener(notify func()) *Listener { return nil }

//func (r *Resource) InstanceOf(i *Interface, implementation interface{}) {}
func (r *Resource) GetLink() *List                          { return nil }
func (Resource) FromLink(l *List) *Resource                 { return nil }
func (Resource) FindForClient(l *List, c *Client) *Resource { return nil }

func (r *Resource) SetImplementation(implementation, data interface{}, destroy func()) {}
func (r *Resource) SetDispatcher(dispatcher func(), implementation, data interface{}, destroy func()) {
}
