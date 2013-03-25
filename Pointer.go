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

func (*Pointer) SetCursor(Serial uint, Surface object, HotspotX int, HotspotY int) {
}

func (*Pointer) Enter(Serial uint, Surface object, SurfaceX fixed, SurfaceY fixed) {
}

func (*Pointer) Leave(Serial uint, Surface object) {
}

func (*Pointer) Motion(Time uint, SurfaceX fixed, SurfaceY fixed) {
}

func (*Pointer) Button(Serial uint, Time uint, Button uint, State uint) {
}

func (*Pointer) Axis(Time uint, Axis uint, Value fixed) {
}
