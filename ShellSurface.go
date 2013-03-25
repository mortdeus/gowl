package gowl

/*

   An interface implemented by a wl_surface.  On server side the
   object is automatically destroyed when the related wl_surface is
   destroyed.  On client side, wl_shell_surface_destroy() must be
   called before destroying the wl_surface object.

*/

type ShellSurface struct{}

var Resize int

const (
	None = 0

	Top = 1

	Bottom = 2

	Left = 4

	TopLeft = 5

	BottomLeft = 6

	Right = 8

	TopRight = 9

	BottomRight = 10
)

var Transient int

const (
	Inactive = 0x1
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
var FullscreenMethod int

const (
	Default = 0

	Scale = 1

	Driver = 2

	Fill = 3
)

func (*ShellSurface) Pong(serial uint) {
}

func (*ShellSurface) Move(seat object, serial uint) {
}

func (*ShellSurface) Resize(seat object, serial uint, edges uint) {
}

func (*ShellSurface) SetToplevel() {
}

func (*ShellSurface) SetTransient(parent object, x int, y int, flags uint) {
}

func (*ShellSurface) SetFullscreen(method uint, framerate uint, output object) {
}

func (*ShellSurface) SetPopup(seat object, serial uint, parent object, x int, y int, flags uint) {
}

func (*ShellSurface) SetMaximized(output object) {
}

func (*ShellSurface) SetTitle(title string) {
}

func (*ShellSurface) SetClass(class string) {
}

func (*ShellSurface) Ping(serial uint) {
}

func (*ShellSurface) Configure(edges uint, width int, height int) {
}

func (*ShellSurface) PopupDone() {
}
