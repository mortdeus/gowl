package gowl

/*

   A compositor.  This object is a singleton global.  The
   compositor is in charge of combining the contents of multiple
   surfaces into one displayable output.

*/

type Compositor struct{}

func (*Compositor) CreateSurface(id new_id) {
}

func (*Compositor) CreateRegion(id new_id) {
}
