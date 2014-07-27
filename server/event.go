package server

type EventQueue struct {
}

func (EventQueue) Create() *EventQueue {
	return new(EventQueue)
}
func (eq *EventQueue) Destroy() {}
