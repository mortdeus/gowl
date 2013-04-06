package gowl

/*

   keyboard input device

*/
type Keyboard struct{}

/*

        This enum specifies the format of the keymap provided to the client
	with the wl_keyboard::keymap event.
*/

const (
	KeymapFormat_XkbV1 = 1
)

/*

   Describes the physical state of a key which provoked the key event.
*/

const (
	KeyState_Released = 0
	KeyState_Pressed  = 1
)

/*

        This event provides a file descriptor to the client which can be
	memory-mapped to provide a keyboard mapping description.
*/
func (*Keyboard) Keymap(format uint32, fd Fd, size uint32) {
}

func (*Keyboard) Enter(serial uint32, surface Object, keys Array) {
}

func (*Keyboard) Leave(serial uint32, surface Object) {
}

/*

	A key was pressed or released.
*/
func (*Keyboard) Key(serial uint32, time uint32, key uint32, state uint32) {
}

/*

        Notifies clients that the modifier and/or group state has
	changed, and it should update its local state.
*/
func (*Keyboard) Modifiers(serial uint32, modsDepressed uint32, modsLatched uint32, modsLocked uint32, group uint32) {
}
