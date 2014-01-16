package server

type Resource struct{}

func NewResource(c *Client, i *Interface, version, int, id uint)

func (r *Resource) PostEventArray(opcode uint, args ...interface{})  {}
func (r *Resource) PostEvent(opcode ...uint)                         {}
func (r *Resource) QueueEventArray(opcode uint, args ...interface{}) {}
func (r *Resource) QueueEvent(opcode ...uint)                        {}
func (r *Resource) PostError(err error)                              {}
func (r *Resource) PostNoMemory()                                    {}
func (r *Resource) Destroy()                                         {}

func (r *Resource) Getid() uint32      {}
func (r *Resource) GetClient() *Client {}
func (r *Resource) GetVersion() int    {}

func (r *Resource) GetUserData() interface{}     {}
func (r *Resource) SetUserData(data interface{}) {}

func (r *Resource) SetDestructor(d func())                     {}
func (r *Resource) GetDestroyListener(notify func()) *Listener {}

//func (r *Resource) InstanceOf(i *Interface, implementation interface{}) {}
func (r *Resource) GetLink() *List                          {}
func (Resource) FromLink(l *List) *Resource                 {}
func (Resource) FindForClient(l *List, c *Client) *Resource {}

func (r *Resource) SetImplementation(implementation, data interface{}, destroy func()) {}
func (r *Resource) SetDispatcher(dispatcher func(), implementation, data interface{}, destroy func()) {
}
