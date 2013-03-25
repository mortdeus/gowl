package gowl

/*

   A surface.  This is an image that is displayed on the screen.
   It has a location, size and pixel contents.

*/

type Surface struct{}

func (*Surface) Destroy() {
}

func (*Surface) Attach(buffer object, x int, y int) {
}

func (*Surface) Damage(x int, y int, width int, height int) {
}

func (*Surface) Frame(callback new_id) {
}

func (*Surface) SetOpaqueRegion(region object) {
}

func (*Surface) SetInputRegion(region object) {
}

func (*Surface) Commit() {
}

func (*Surface) SetBufferTransform(transform int) {
}

func (*Surface) Enter(output object) {
}

func (*Surface) Leave(output object) {
}
