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

func (*ShmPool) CreateBuffer(Id new_id, Offset int, Width int, Height int, Stride int, Format uint) {
}

func (*ShmPool) Destroy() {
}

func (*ShmPool) Resize(Size int) {
}
