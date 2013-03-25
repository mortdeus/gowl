package gowl

/*

   A surface.  This is an image that is displayed on the screen.
   It has a location, size and pixel contents.

*/

type Surface struct{}

func (*Surface) Destroy() {
}

func (*Surface) Attach(Buffer object, X int, Y int) {
}

func (*Surface) Damage(X int, Y int, Width int, Height int) {
}

func (*Surface) Frame(Callback new_id) {
}

func (*Surface) SetOpaqueRegion(Region object) {
}

func (*Surface) SetInputRegion(Region object) {
}

func (*Surface) Commit() {
}

func (*Surface) SetBufferTransform(Transform int) {
}

func (*Surface) Enter(Output object) {
}

func (*Surface) Leave(Output object) {
}
