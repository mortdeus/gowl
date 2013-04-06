package gowl

/*

   Region.

*/
type Region struct{}

/*

	Destroy the region.  This will invalidate the object id.
*/
func (*Region) Destroy() {
}

/*

	Add the specified rectangle to the region
*/
func (*Region) Add(x int32, y int32, width int32, height int32) {
}

/*

	Subtract the specified rectangle from the region
*/
func (*Region) Subtract(x int32, y int32, width int32, height int32) {
}
