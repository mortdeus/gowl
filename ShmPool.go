package gowl

/*

   The wl_shm_pool object encapsulates a piece of memory shared
   between the compositor and client.  Through the wl_shm_pool
   object, the client can allocate shared memory wl_buffer objects.
   The objects will share the same underlying mapped memory.
   Reusing the mapped memory avoids the setup/teardown overhead and
   is useful when interactively resizing a surface or for many
   small buffers.

*/
type ShmPool struct{}

/*

	Create a wl_buffer from the pool.  The buffer is created a
	offset bytes into the pool and has width and height as
	specified.  The stride arguments specifies the number of bytes
	from beginning of one row to the beginning of the next.  The
	format is the pixel format of the buffer and must be one of
	those advertised through the wl_shm.format event.

	A buffer will keep a reference to the pool it was created from
	so it is valid to destroy the pool immediately after creating
	a buffer from it.
*/
func (*ShmPool) CreateBuffer(id NewId, offset int32, width int32, height int32, stride int32, format uint32) {
}

/*

	Destroy the pool.
*/
func (*ShmPool) Destroy() {
}

/*

	This request will cause the server to remap the backing memory
	for the pool from the fd passed when the pool was creating but
	using the new size.
*/
func (*ShmPool) Resize(size int32) {
}
