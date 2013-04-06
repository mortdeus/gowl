package gowl

/*


*/
type Touch struct{}

func (*Touch) Down(serial uint32, time uint32, surface Object, id int32, x Fixed, y Fixed) {
}

func (*Touch) Up(serial uint32, time uint32, id int32) {
}

func (*Touch) Motion(time uint32, id int32, x Fixed, y Fixed) {
}

/*

	Indicates the end of a contact point list.
*/
func (*Touch) Frame() {
}

/*

	Sent if the compositor decides the touch stream is a global
	gesture. No further events are sent to the clients from that
	particular gesture.
*/
func (*Touch) Cancel() {
}
