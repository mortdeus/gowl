package gowl

type Pointer struct{}

/*

        Describes the physical state of a button which provoked the button
	event.

*/
var ButtonState int

const (
	Released = 0

	Pressed = 1
)

var Axis int

const (
	VerticalScroll = 0

	HorizontalScroll = 1
)

func (*Pointer) SetCursor(serial uint, surface object, hotspotX int, hotspotY int) {
}

func (*Pointer) Enter(serial uint, surface object, surfaceX fixed, surfaceY fixed) {
}

func (*Pointer) Leave(serial uint, surface object) {
}

func (*Pointer) Motion(time uint, surfaceX fixed, surfaceY fixed) {
}

func (*Pointer) Button(serial uint, time uint, button uint, state uint) {
}

func (*Pointer) Axis(time uint, axis uint, value fixed) {
}
