package gowl

/*

   An output describes part of the compositor geometry.  The
   compositor work in the 'compositor coordinate system' and an
   output corresponds to rectangular area in that space that is
   actually visible.  This typically corresponds to a monitor that
   displays part of the compositor space.  This object is published
   as global during start up, or when a screen is hot plugged.

*/

type Output struct{}

var Subpixel int

const (
	Unknown = 0

	None = 1

	HorizontalRgb = 2

	HorizontalBgr = 3

	VerticalRgb = 4

	VerticalBgr = 5
)

/*

	This describes the transform that a compositor will apply to a
	surface to compensate for the rotation or mirroring of an
	output device.

	The flipped values correspond to an initial flip around a
	vertical axis followed by rotation.

	The purpose is mainly to allow clients render accordingly and
	tell the compositor, so that for fullscreen surfaces, the
	compositor will still be able to scan out directly from client
	surfaces.

*/
var Transform int

const (
	Normal = 0

	Rot90 = 1

	Rot180 = 2

	Rot270 = 3

	Flipped = 4

	Flipped90 = 5

	Flipped180 = 6

	Flipped270 = 7
)

var Mode int

const (
	Current = 0x1

	Preferred = 0x2
)

func (*Output) Geometry(x int, y int, physicalWidth int, physicalHeight int, subpixel int, make string, model string, transform int) {
}

func (*Output) Mode(flags uint, width int, height int, refresh int) {
}
