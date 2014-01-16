package server

type Event uint

const (
	READABLE Event = 1 << iota
	WRITEABLE
	HANGUP
	ERROR
)

type EventLoop struct {
	FdFunc     func(fd int, mask uint, data interface{})
	TimerFunc  func(data interface{})
	SignalFunc func(sigNum int, data interface{})
	IdleFunc   func(data interface{})
}

func (e *EventLoop) Create() {

}
func (e *EventLoop) Destroy() {}

func (e *EventLoop) Addfd(fd int, mask uint, data interface{}) *EventSource

func (e *EventLoop) AddTimer(data interface{}) *EventSource
func (e *EventLoop) AddSignal(sigNum int, data interface{}) *EventSource {}

func (e *EventLoop) Dispatch(timeout int) int              {}
func (e *EventLoop) DispatchIdle()                         {}
func (e *EventLoop) AddIdle(data interface{}) *EventSource {}
func (e *EventLoop) Getfd() int                            {}

type EventSource struct{}

func (e *EventSource) UpdateFd(mask uint) int {

}
func (e *EventSource) UpdateTimer(msDelay int) int {

}
func (e *EventSource) Remove() int {}

func (e *EventSource) Check() {}
