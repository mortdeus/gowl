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

func (*Keyboard) Keymap(Format uint, Fd fd, Size uint) {
}

func (*Keyboard) Enter(Serial uint, Surface object, Keys array) {
}

func (*Keyboard) Leave(Serial uint, Surface object) {
}

func (*Keyboard) Key(Serial uint, Time uint, Key uint, State uint) {
}

func (*Keyboard) Modifiers(Serial uint, ModsDepressed uint, ModsLatched uint, ModsLocked uint, Group uint) {
}
