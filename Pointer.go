package gowl

type Pointer struct{}

/*

        Describes the physical state of a button which provoked the button
	event.
*/

const (
	ButtonState_Released = 0
	ButtonState_Pressed  = 1
)

const (
	Axis_VerticalScroll   = 0
	Axis_HorizontalScroll = 1
)

/*

	Set the pointer surface, i.e., the surface that contains the
	pointer image (cursor). This request only takes effect if the pointer
	focus for this device is one of the requesting client's surfaces
	or the surface parameter is the current pointer surface. If
	there was a previous surface set with this request it is
	replaced. If surface is NULL, the pointer image is hidden.

	The parameters hotspot_x and hotspot_y define the position of
	the pointer surface relative to the pointer location. Its
	top-left corner is always at (x, y) - (hotspot_x, hotspot_y),
	where (x, y) are the coordinates of the pointer location.

	On surface.attach requests to the pointer surface, hotspot_x
	and hotspot_y are decremented by the x and y parameters
	passed to the request. Attach must be confirmed by
	wl_surface.commit as usual.

	The hotspot can also be updated by passing the currently set
	pointer surface to this request with new values for hotspot_x
	and hotspot_y.

	The current and pending input regions of the wl_surface are
	cleared, and wl_surface.set_input_region is ignored until the
	wl_surface is no longer used as the cursor. When the use as a
	cursor ends, the current and pending input regions become
	undefined, and the wl_surface is unmapped.
*/
func (*Pointer) SetCursor(serial uint32, surface Object, hotspotX int32, hotspotY int32) {
}

/*

	Notification that this seat's pointer is focused on a certain
	surface. When an seat's focus enters a surface, the pointer image
	is undefined and a client should respond to this event by setting
	an appropriate pointer image.
*/
func (*Pointer) Enter(serial uint32, surface Object, surfaceX Fixed, surfaceY Fixed) {
}

/*

*/
func (*Pointer) Leave(serial uint32, surface Object) {
}

/*

	Notification of pointer location change. The arguments surface_[xy]
	are the location relative to the focused surface.
*/
func (*Pointer) Motion(time uint32, surfaceX Fixed, surfaceY Fixed) {
}

/*

	Mouse button click and release notifications.  The location
	of the click is given by the last motion or pointer_focus event.
*/
func (*Pointer) Button(serial uint32, time uint32, button uint32, state uint32) {
}

/*

	Scroll and other axis notifications.

	For scroll events (vertical and horizontal scroll axes), the
	value parameter is the length of a vector along the specified
	axis in a coordinate space identical to those of motion events,
	representing a relative movement along the specified axis.

	For devices that support movements non-parallel to axes multiple
	axis events will be emitted.

	When applicable, for example for touch pads, the server can
	choose to emit scroll events where the motion vector is
	equivalent to a motion event vector.

	When applicable, clients can transform its view relative to the
	scroll distance.
*/
func (*Pointer) Axis(time uint32, axis uint32, value Fixed) {
}
