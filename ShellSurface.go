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

func (*ShellSurface) Pong(Serial uint) {
}

func (*ShellSurface) Move(Seat object, Serial uint) {
}

func (*ShellSurface) Resize(Seat object, Serial uint, Edges uint) {
}

func (*ShellSurface) SetToplevel() {
}

func (*ShellSurface) SetTransient(Parent object, X int, Y int, Flags uint) {
}

func (*ShellSurface) SetFullscreen(Method uint, Framerate uint, Output object) {
}

func (*ShellSurface) SetPopup(Seat object, Serial uint, Parent object, X int, Y int, Flags uint) {
}

func (*ShellSurface) SetMaximized(Output object) {
}

func (*ShellSurface) SetTitle(Title string) {
}

func (*ShellSurface) SetClass(Class string) {
}

func (*ShellSurface) Ping(Serial uint) {
}

func (*ShellSurface) Configure(Edges uint, Width int, Height int) {
}

func (*ShellSurface) PopupDone() {
}
