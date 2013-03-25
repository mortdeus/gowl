package gowl

/*

   A group of keyboards, pointer (mice, for example) and touch
   devices . This object is published as a global during start up,
   or when such a device is hot plugged.  A seat typically has a
   pointer and maintains a keyboard_focus and a pointer_focus.

*/

type Seat struct{}

/*

        This is a bitmask of capabilities this seat has; if a member is
	set, then it is present on the seat.

*/
var Capability int

const (
	Pointer = 1

	Keyboard = 2

	Touch = 4
)

func (*Seat) GetPointer(id new_id) {
}

func (*Seat) GetKeyboard(id new_id) {
}

func (*Seat) GetTouch(id new_id) {
}

func (*Seat) Capabilities(capabilities uint) {
}
