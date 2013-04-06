package gowl

/*

   A compositor.  This object is a singleton global.  The
   compositor is in charge of combining the contents of multiple
   surfaces into one displayable output.

*/
type Compositor struct{}

/*

	Ask the compositor to create a new surface.
*/
func (*Compositor) CreateSurface(id NewId) {
}

/*

	Ask the compositor to create a new region.
*/
func (*Compositor) CreateRegion(id NewId) {
}
