package gowl

type Display interface {
	Sync(callback uint32)
	GetRegistry(registry uint32)

	//Events
	Error(objID uint32, err DisplayErr)
	DeleteID(uint32)
}

type DisplayErr uint

const (
	InvalidObj DisplayErr = iota
	InvalidMethod
	OOM
)

func (e DisplayErr) Error() string {
	switch e {
	case InvalidObj:
		return "server couldn't find object"
	case InvalidMethod:
		return "method doesn't exist on the specified interface"
	case OOM:
		return "server is out of memory"
	default:
		return "unrecognized error value"
	}

}

type Registry interface {
	// BUG(mortdeus): The xml protocol spec says Bind takes two args, but
	// the standard C implementation libwayland uses this signature.
	Bind(name uint32, iface string, ver, id uint32)

	// Events
	Global(name uint32, iface string, ver uint32)
	GlobalRemove(name uint32)
}

type Callback interface {
	Done(data uint32)
}
type Compositor interface {
	CreateSurface(id uint32)
	CreateRegion(id uint32)
}
type ShmPool interface {
	CreateBuffer(id uint32, offset, w, h, stride int32, format ShmFormat)
	Destroy()
	Resize(sz int32)
}
type Shm interface {
	CreatePool(id uint32, fd uintptr, sz int32)

	//Events
	Format(f ShmFormat)
}
type ShmErr uint

const (
	InvalidFormat ShmErr = iota
	InvalidStride
	InvalidFd
)

func (e ShmErr) Error() string {
	switch e {
	case InvalidFormat:
		return "buffer format is not known"
	case InvalidStride:
		return "invalid size or stride during pool or buffer creation"
	case InvalidFd:
		return "mmapping the file descriptor failed"
	default:
		return "unrecognized error value"
	}

}

type ShmFormat uint

const (
	ARGB8888    ShmFormat = 0
	XRGB8888    ShmFormat = 1
	C8          ShmFormat = 0x20203843
	RGB332      ShmFormat = 0x38424752
	BGR233      ShmFormat = 0x38524742
	XRGB4444    ShmFormat = 0x32315258
	XBGR4444    ShmFormat = 0x32314258
	RGBX4444    ShmFormat = 0x32315852
	BGRX4444    ShmFormat = 0x32315842
	ARGB4444    ShmFormat = 0x32315241
	ABGR4444    ShmFormat = 0x32314241
	RGBA4444    ShmFormat = 0x32314152
	BGRA4444    ShmFormat = 0x32314142
	XRGB1555    ShmFormat = 0x35315258
	XBGR1555    ShmFormat = 0x35314258
	RGBX5551    ShmFormat = 0x35315852
	BGRX5551    ShmFormat = 0x35315842
	ARGB1555    ShmFormat = 0x35315241
	ABGR1555    ShmFormat = 0x35314241
	RGBA5551    ShmFormat = 0x35314152
	BGRA5551    ShmFormat = 0x35314142
	RGB565      ShmFormat = 0x36314752
	BGR565      ShmFormat = 0x36314742
	RGB888      ShmFormat = 0x34324752
	BGR888      ShmFormat = 0x34324742
	XBGR8888    ShmFormat = 0x34324258
	RGBX8888    ShmFormat = 0x34325852
	BGRX8888    ShmFormat = 0x34325842
	ABGR8888    ShmFormat = 0x34324241
	RGBA8888    ShmFormat = 0x34324152
	BGRA8888    ShmFormat = 0x34324142
	XRGB2101010 ShmFormat = 0x30335258
	XBGR2101010 ShmFormat = 0x30334258
	RGBX1010102 ShmFormat = 0x30335852
	BGRX1010102 ShmFormat = 0x30335842
	ARGB2101010 ShmFormat = 0x30335241
	ABGR2101010 ShmFormat = 0x30334241
	RGBA1010102 ShmFormat = 0x30334152
	BGRA1010102 ShmFormat = 0x30334142
	YUYV        ShmFormat = 0x56595559
	YVYU        ShmFormat = 0x55595659
	UYVY        ShmFormat = 0x59565955
	VYUY        ShmFormat = 0x59555956
	AYUV        ShmFormat = 0x56555941
	NV12        ShmFormat = 0x3231564e
	NV21        ShmFormat = 0x3132564e
	NV16        ShmFormat = 0x3631564e
	NV61        ShmFormat = 0x3136564e
	YUV410      ShmFormat = 0x39565559
	YVU410      ShmFormat = 0x39555659
	YUV411      ShmFormat = 0x31315559
	YVU411      ShmFormat = 0x31315659
	YUV420      ShmFormat = 0x32315559
	YVU420      ShmFormat = 0x32315659
	YUV422      ShmFormat = 0x36315559
	YVU422      ShmFormat = 0x36315659
	YUV444      ShmFormat = 0x34325559
	YVU444      ShmFormat = 0x34325659
)

type Buffer interface {
	Destroy()
	Release()
}
type DataOffer interface {
	Accept(serial uint32, mimeType string)
	Receive(mimeType string, fd uintptr)
	Destroy()

	//Events
	Offer(mimeType string)
}
type DataSource interface {
	Offer(mimeType string)
	Destroy()

	//Events
	Target(mimeType string)
	Send(mimeType string, fd uintptr)
	Cancelled()
}
type DataDevice interface {
	StartDrag(source, origin, icon, serial uint32)
}
type DataDeviceManager interface{}
type Shell interface{}
type ShellSurface interface{}
type Surface interface{}
type Seat interface{}
type Pointer interface{}
type Keyboard interface{}
type Touch interface{}
type Output interface{}
type Region interface{}
type Subcompositor interface{}
type Subsurface interface{}
