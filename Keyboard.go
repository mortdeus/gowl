package gowl

/*

   keyboard input device

*/

type Keyboard struct{}

/*

        This enum specifies the format of the keymap provided to the client
	with the wl_keyboard::keymap event.

*/
var KeymapFormat int

const (
	XkbV1 = 1
)

/*

   Describes the physical state of a key which provoked the key event.

*/
var KeyState int

const (
	Released = 0

	Pressed = 1
)

func (*Keyboard) Keymap(format uint, fd fd, size uint) {
}

func (*Keyboard) Enter(serial uint, surface object, keys array) {
}

func (*Keyboard) Leave(serial uint, surface object) {
}

func (*Keyboard) Key(serial uint, time uint, key uint, state uint) {
}

func (*Keyboard) Modifiers(serial uint, modsDepressed uint, modsLatched uint, modsLocked uint, group uint) {
}
