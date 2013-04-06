package gowl

/*

   An interface implemented by a wl_surface.  On server side the
   object is automatically destroyed when the related wl_surface is
   destroyed.  On client side, wl_shell_surface_destroy() must be
   called before destroying the wl_surface object.

*/
type ShellSurface struct{}

const (
	Resize_None        = 0
	Resize_Top         = 1
	Resize_Bottom      = 2
	Resize_Left        = 4
	Resize_TopLeft     = 5
	Resize_BottomLeft  = 6
	Resize_Right       = 8
	Resize_TopRight    = 9
	Resize_BottomRight = 10
)

const (
	Transient_Inactive = 0x1
)

/*

   Hints to indicate compositor how to deal with a conflict between the
   dimensions for the surface and the dimensions of the output. As a hint
   the compositor is free to ignore this parameter.

   "default" The client has no preference on fullscreen behavior,
   policies are determined by compositor.

   "scale" The client prefers scaling by the compositor. Scaling would
   always preserve surface's aspect ratio with surface centered on the
   output

   "driver" The client wants to switch video mode to the smallest mode
   that can fit the client buffer. If the sizes do not match the
   compositor must add black borders.

   "fill" The surface is centered on the output on the screen with no
   scaling. If the surface is of insufficient size the compositor must
   add black borders.
*/

const (
	FullscreenMethod_Default = 0
	FullscreenMethod_Scale   = 1
	FullscreenMethod_Driver  = 2
	FullscreenMethod_Fill    = 3
)

/*

	A client must respond to a ping event with a pong request or
	the client may be deemed unresponsive.
*/
func (*ShellSurface) Pong(serial uint32) {
}

func (*ShellSurface) Move(seat Object, serial uint32) {
}

func (*ShellSurface) Resize(seat Object, serial uint32, edges uint32) {
}

/*

	Make the surface a toplevel window.
*/
func (*ShellSurface) SetToplevel() {
}

/*

	Map the surface relative to an existing surface. The x and y
	arguments specify the locations of the upper left corner of
	the surface relative to the upper left corner of the parent
	surface.  The flags argument controls overflow/clipping
	behaviour when the surface would intersect a screen edge,
	panel or such.  And possibly whether the offset only
	determines the initial position or if the surface is locked to
	that relative position during moves.
*/
func (*ShellSurface) SetTransient(parent Object, x int32, y int32, flags uint32) {
}

/*

   Map the surface as a fullscreen surface. If an output parameter is
   given then the surface will be made fullscreen on that output. If the
   client does not specify the output then the compositor will apply its
   policy - usually choosing the output on which the surface has the
   biggest surface area.

   The client may specify a method to resolve a size conflict between the
   output size and the surface size - this is provided through the
   fullscreen_method parameter.

   The framerate parameter is used only when the fullscreen_method is set
   to "driver", to indicate the preferred framerate. framerate=0 indicates
   that the app does not care about framerate.  The framerate is
   specified in mHz, that is framerate of 60000 is 60Hz.

   The compositor must reply to this request with a configure event with
   the dimensions for the output on which the surface will be made fullscreen.
*/
func (*ShellSurface) SetFullscreen(method uint32, framerate uint32, output Object) {
}

/*

	Popup surfaces.  Will switch an implicit grab into
	owner-events mode, and grab will continue after the implicit
	grab ends (button released).  Once the implicit grab is over,
	the popup grab continues until the window is destroyed or a
	mouse button is pressed in any other clients window.  A click
	in any of the clients surfaces is reported as normal, however,
	clicks in other clients surfaces will be discarded and trigger
	the callback.

	TODO: Grab keyboard too, maybe just terminate on any click
	inside or outside the surface?
*/
func (*ShellSurface) SetPopup(seat Object, serial uint32, parent Object, x int32, y int32, flags uint32) {
}

/*

	A request from the client to notify the compositor the maximized
	operation. The compositor will reply with a configure event telling
        the expected new surface size. The operation is completed on the
        next buffer attach to this surface.
        A maximized client will fill the fullscreen of the output it is bound
        to, except the panel area. This is the main difference between
        a maximized shell surface and a fullscreen shell surface.
*/
func (*ShellSurface) SetMaximized(output Object) {
}

/*

*/
func (*ShellSurface) SetTitle(title string) {
}

/*

	The surface class identifies the general class of applications
	to which the surface belongs.  The class is the file name of
	the applications .desktop file (absolute path if non-standard
	location).
*/
func (*ShellSurface) SetClass(class string) {
}

/*

	Ping a client to check if it is receiving events and sending
	requests. A client is expected to reply with a pong request.
*/
func (*ShellSurface) Ping(serial uint32) {
}

/*

	The configure event asks the client to resize its surface.
	The size is a hint, in the sense that the client is free to
	ignore it if it doesn't resize, pick a smaller size (to
	satisfy aspect ratio or resize in steps of NxM pixels).  The
	client is free to dismiss all but the last configure event it
	received.
*/
func (*ShellSurface) Configure(edges uint32, width int32, height int32) {
}

/*

	The popup_done event is sent out when a popup grab is broken,
	that is, when the users clicks a surface that doesn't belong
	to the client owning the popup surface.
*/
func (*ShellSurface) PopupDone() {
}
