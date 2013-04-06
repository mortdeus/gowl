package gowl

/*

   A buffer provides the content for a wl_surface. Buffers are
   created through factory interfaces such as wl_drm, wl_shm or
   similar. It has a width and a height and can be attached to a
   wl_surface, but the mechanism by which a client provides and
   updates the contents is defined by the buffer factory interface.

*/
type Buffer struct{}

/*

	Destroy a buffer. If and how you need to release the backing
	storage is defined by the buffer factory interface.

	For possible side-effects to a surface, see wl_surface.attach.
*/
func (*Buffer) Destroy() {
}

/*

	Sent when this wl_buffer is no longer used by the compositor.
	The client is now free to re-use or destroy this buffer and its
	backing storage.

	If a client receives a release event before the frame callback
	requested in the same wl_surface.commit that attaches this
	wl_buffer to a surface, then the client is immediately free to
	re-use the buffer and its backing storage, and does not need a
	second buffer for the next surface content update. Typically
	this is possible, when the compositor maintains a copy of the
	wl_surface contents, e.g. as a GL texture. This is an important
	optimization for GL(ES) compositors with wl_shm clients.
*/
func (*Buffer) Release() {
}
