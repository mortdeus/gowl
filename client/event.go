package client

type EventQueue <-chan int

func (eq EventQueue) Create() {
	eq = make(EventQueue, 0)

	//TODO(mortdeus): implement connect logic
}
func (eq EventQueue) Destroy() {
	//TODO(mortdeus)Y implement callback logic

}
