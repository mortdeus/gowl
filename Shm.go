package gowl

/*

   Support for shared memory buffers.

*/

type Shm struct{}

var Error int

const (
	InvalidFormat = 0

	InvalidStride = 1

	InvalidFd = 2
)

var Format int

const (
	Argb8888 = 0

	Xrgb8888 = 1
)

func (*Shm) CreatePool(id new_id, fd fd, size int) {
}

func (*Shm) Format(format uint) {
}
