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

func (*ShmPool) CreateBuffer(id new_id, offset int, width int, height int, stride int, format uint) {
}

func (*ShmPool) Destroy() {
}

func (*ShmPool) Resize(size int) {
}
