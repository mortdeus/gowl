package gowl

/*

   A buffer provides the content for a wl_surface. Buffers are
   created through factory interfaces such as wl_drm, wl_shm or
   similar. It has a width and a height and can be attached to a
   wl_surface, but the mechanism by which a client provides and
   updates the contents is defined by the buffer factory interface.

*/

type Buffer struct{}

func (*Buffer) Destroy() {
}

func (*Buffer) Release() {
}
