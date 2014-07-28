package gowl

// The core global object. This is a special singleton object. It
// is used for internal Wayland protocol features.
type Display interface {
	// The sync request asks the server to emit the 'done' event
	// on the returned wl_callback object. Since requests are
	// handled in-order and events are delivered in-order, this can
	// be used as a barrier to ensure all previous requests and the
	// resulting events have been handled.
	//
	// The object returned by this request will be destroyed by the
	// compositor after the callback is fired and as such the client must not
	// attempt to use it after that point.
	//
	// The callback_data passed in the callback is the event serial.
	Sync(callback NewId)

	// This request creates a registry object that allows the client
	// to list and bind the global objects available from the
	// compositor.
	GetRegistry(registry NewId)

	// The error event is sent out when a fatal (non-recoverable)
	// error has occurred. The object_id argument is the object
	// where the error occurred, most often in response to a request
	// to that object. The code identifies the error and is defined
	// by the object interface. As such, each interface defines its
	// own set of error codes. The message is an brief description
	// of the error, for (debugging) convenience.
	Error(objectId Object, code uint32, message string)

	// This event is used internally by the object ID management
	// logic. When a client deletes an object, the server will send
	// this event to acknowledge that it has seen the delete request.
	// When the client receive this event, it will know that it can
	// safely reuse the object ID.
	DeleteId(id uint32)
}

// These errors are global and can be emitted in response to any
// server request.
type DisplayErr uint

const (
	InvalidObj DisplayErr = iota
	InvaliddMethod
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

// The global registry object. The server has a number of global
// objects that are available to all clients. These objects
// typically represent an actual object in the server (for example,
// an input device)
// or they are singleton objects that provide
// extension functionality.
//
// When a client creates a registry object, the registry object
// will emit a global event for each global currently in the
// registry. Globals come and go as a result of device or
// monitor hotplugs, reconfiguration or other events, and the
// registry will send out global and global_remove events to
// keep the client up to date with the changes. To mark the end
// of the initial burst of events, the client can use the
// wl_display.sync request immediately after calling
// wl_display.get_registry.
//
// A client can bind to a global object by using the bind
// request. This creates a client-side handle that lets the object
// emit events to the client and lets the client invoke requests on
// the object.
type Registry interface {
	// Binds a new, client-created object to the server using the
	// specified name as the identifier.
	Bind(name uint32, id NewId)

	// Notify the client of global objects.
	//
	// The event notifies the client that a global object with
	// the given name is now available, and it implements the
	// given version of the given interface.
	Global(name uint32, interface_ string, version uint32)

	// Notify the client of removed global objects.
	//
	// This event notifies the client that the global identified
	// by name is no longer available. If the client bound to
	// the global using the bind request, the client should now
	// destroy that object.
	//
	// The object remains valid and requests to the object will be
	// ignored until the client destroys it, to avoid races between
	// the global going away and a client sending a request to it.
	GlobalRemove(name uint32)
}

// Clients can handle the 'done' event to get notified when
// the related request is done.
type Callback interface {
	// Notify the client when the related request is done.
	Done(callbackData uint32)
}

// A compositor. This object is a singleton global. The
// compositor is in charge of combining the contents of multiple
// surfaces into one displayable output.
type Compositor interface {
	// Ask the compositor to create a new surface.
	CreateSurface(id NewId)

	// Ask the compositor to create a new region.
	CreateRegion(id NewId)
}

// The wl_shm_pool object encapsulates a piece of memory shared
// between the compositor and client. Through the wl_shm_pool
// object, the client can allocate shared memory wl_buffer objects.
// All objects created through the same pool share the same
// underlying mapped memory. Reusing the mapped memory avoids the
// setup/teardown overhead and is useful when interactively resizing
// a surface or for many small buffers.
type ShmPool interface {
	// Create a wl_buffer object from the pool.
	//
	// The buffer is created offset bytes into the pool and has
	// width and height as specified. The stride arguments specifies
	// the number of bytes from beginning of one row to the beginning
	// of the next. The format is the pixel format of the buffer and
	// must be one of those advertised through the wl_shm.format event.
	//
	// A buffer will keep a reference to the pool it was created from
	// so it is valid to destroy the pool immediately after creating
	// a buffer from it.
	CreateBuffer(id NewId, offset int32, width int32, height int32, stride int32, format uint32)

	// Destroy the shared memory poo
	//
	// The mmapped memory will be released when all
	// buffers that have been created from this pool
	// are gone.
	Destroy()

	// This request will cause the server to remap the backing memory
	// for the pool from the file descriptor passed when the pool was
	// created, but using the new size. This request can only be
	// used to make the pool bigger.
	Resize(size int32)
}

// A global singleton object that provides support for shared
// memory.
// Clients can create wl_shm_pool objects using the create_pool
// request.
// At connection setup time, the wl_shm object emits one or more
// format events to inform clients about the valid pixel formats
// that can be used for buffers.
type Shm interface {

	// Create a new wl_shm_pool object.
	// The pool can be used to create shared memory based buffer
	// objects. The server will mmap size bytes of the passed file
	// descriptor, to use as backing memory for the pool.
	CreatePool(id NewId, fd Fd, size int32)

	// Informs the client about a valid pixel format that
	// can be used for buffers. Known formats include
	// argb8888 and xrgb8888.
	Format(format uint32)
}

// These errors can be emitted in response to wl_shm requests.
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

// This describes the memory layout of an individual pixel.
// All renderers should support argb8888 and xrgb8888 but any other
// formats are optional and may not be supported by the particular
// renderer in use.
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

// A buffer provides the content for a wl_surface. Buffers are
// created through factory interfaces such as wl_drm, wl_shm or
// similar. It has a width and a height and can be attached to a
// wl_surface, but the mechanism by which a client provides and
// updates the contents is defined by the buffer factory interface.
type Buffer interface {
	// Destroy a buffer. If and how you need to release the backing
	// storage is defined by the buffer factory interface.
	// For possible side-effects to a surface, see wl_surface.attach.
	Destroy()

	// Sent when this wl_buffer is no longer used by the compositor.
	// The client is now free to re-use or destroy this buffer and its
	// backing storage.
	//
	// If a client receives a release event before the frame callback
	// requested in the same wl_surface.commit that attaches this
	// wl_buffer to a surface, then the client is immediately free to
	// re-use the buffer and its backing storage, and does not need a
	// second buffer for the next surface content update. Typically
	// this is possible, when the compositor maintains a copy of the
	// wl_surface contents, e.g. as a GL texture. This is an important
	// optimization for GL(ES) compositors with wl_shm clients.
	Release()
}

// A wl_data_offer represents a piece of data offered for transfer
// by another client (the source client). It is used by the
// copy-and-paste and drag-and-drop mechanisms. The offer
// describes the different mime types that the data can be
// converted to and provides the mechanism for transferring the
// data directly from the source client.
type DataOffer interface {

	// Indicate that the client can accept the given mime type, or
	// NULL for not accepted.
	//
	// Used for feedback during drag-and-drop.
	Accept(serial uint32, mimeType string)

	// To transfer the offered data, the client issues this request
	// and indicates the mime type it wants to receive. The transfer
	// happens through the passed file descriptor (typically created
	// with the pipe system call). The source client writes the data
	// in the mime type representation requested and then closes the
	// file descriptor.
	//
	// The receiving client reads from the read end of the pipe until
	// EOF and the closes its end, at which point the transfer is
	// complete.
	Receive(mimeType string, fd Fd)

	// Destroy the data offer.
	Destroy()

	// Sent immediately after creating the wl_data_offer object. One
	// event per offered mime type.
	Offer(mimeType string)
}

// The wl_data_source object is the source side of a wl_data_offer.
// It is created by the source client in a data transfer and
// provides a way to describe the offered data and a way to respond
// to requests to transfer the data.
type DataSource interface {

	// This request adds a mime type to the set of mime types
	// advertised to targets. Can be called several times to offer
	// multiple types.
	Offer(mimeType string)

	// Destroy the data source.
	Destroy()

	// Sent when a target accepts pointer_focus or motion events. If
	// a target does not accept any of the offered types, type is NULL.
	//
	// Used for feedback during drag-and-drop.
	Target(mimeType string)

	// Request for data from the client. Send the data as the
	// specified mime type over the passed file descriptor, then
	// close it.
	Send(mimeType string, fd Fd)

	// This data source has been replaced by another data source.
	// The client should clean up and destroy this data source.
	Cancelled()
}

// There is one wl_data_device per seat which can be obtained
// from the global wl_data_device_manager singleton.
//
// A wl_data_device provides access to inter-client data transfer
// mechanisms such as copy-and-paste and drag-and-drop.
type DataDevice interface {

	// This request asks the compositor to start a drag-and-drop
	// operation on behalf of the client.
	//
	// The source argument is the data source that provides the data
	// for the eventual data transfer. If source is NULL, enter, leave
	// and motion events are sent only to the client that initiated the
	// drag and the client is expected to handle the data passing
	// internally.
	//
	// The origin surface is the surface where the drag originates and
	// the client must have an active implicit grab that matches the
	// serial.
	//
	// The icon surface is an optional (can be NULL) surface that
	// provides an icon to be moved around with the cursor. Initially,
	// the top-left corner of the icon surface is placed at the cursor
	// hotspot, but subsequent wl_surface.attach request can move the
	// relative position. Attach requests must be confirmed with
	// wl_surface.commit as usual.
	//
	// The current and pending input regions of the icon wl_surface are
	// cleared, and wl_surface.set_input_region is ignored until the
	// wl_surface is no longer used as the icon surface. When the use
	// as an icon ends, the current and pending input regions become
	// undefined, and the wl_surface is unmapped.
	StartDrag(source Object, origin Object, icon Object, serial uint32)

	// This request asks the compositor to set the selection
	// to the data from the source on behalf of the client.
	//
	// To unset the selection, set the source to NULL.
	SetSelection(source Object, serial uint32)

	// The data_offer event introduces a new wl_data_offer object,
	// which will subsequently be used in either the
	// data_device.enter event (for drag-and-drop)or the
	// data_device.selection event (for selections). Immediately
	// following the data_device_data_offer event, the new data_offer
	// object will send out data_offer.offer events to describe the
	// mime types it offers.
	DataOffer(id NewId)

	// This event is sent when an active drag-and-drop pointer enters
	// a surface owned by the client. The position of the pointer at
	// enter time is provided by the x and y arguments, in surface
	// local coordinates.
	Enter(serial uint32, surface Object, x Fixed, y Fixed, id Object)

	// This event is sent when the drag-and-drop pointer leaves the
	// surface and the session ends. The client must destroy the
	// wl_data_offer introduced at enter time at this point.
	Leave()

	// This event is sent when the drag-and-drop pointer moves within
	// the currently focused surface. The new position of the pointer
	// is provided by the x and y arguments, in surface local
	// coordinates.
	Motion(time uint32, x Fixed, y Fixed)

	// The event is sent when a drag-and-drop operation is ended
	// because the implicit grab is removed.
	Drop()

	// The selection event is sent out to notify the client of a new
	// wl_data_offer for the selection for this device. The
	// data_device.data_offer and the data_offer.offer events are
	// sent out immediately before this event to introduce the data
	// offer object. The selection event is sent to a client
	// immediately before receiving keyboard focus and when a new
	// selection is set while the client has keyboard focus. The
	// data_offer is valid until a new data_offer or NULL is received
	// or until the client loses keyboard focus.
	Selection(id Object)
}

// The wl_data_device_manager is a singleton global object that
// provides access to inter-client data transfer mechanisms such as
// copy-and-paste and drag-and-drop. These mechanisms are tied to
// a wl_seat and this interface lets a client get a wl_data_device
// corresponding to a wl_seat.
type DataDeviceManager interface {

	// Create a new data source.
	CreateDataSource(id NewId)
	// Create a new data device for a given seat.
	GetDataDevice(id NewId, seat Object)
}

// This interface is implemented by servers that provide
// desktop-style user interfaces.
//
// It allows clients to associate a wl_shell_surface with
// a basic surface.
type Shell interface {

	// Create a shell surface for an existing surface.
	//
	// Only one shell surface can be associated with a given surface.
	GetShellSurface(id NewId, surface Object)
}

// An interface that may be implemented by a wl_surface, for
// implementations that provide a desktop-style user interface.
//
// It provides requests to treat surfaces like toplevel, fullscreen
// or popup windows, move, resize or maximize them, associate
// metadata like title and class, etc.
//
// On the server side the object is automatically destroyed when
// the related wl_surface is destroyed. On client side,
// wl_shell_surface_destroy() must be called before destroying
// the wl_surface object.
type ShellSurface interface {

	// A client must respond to a ping event with a pong request or
	// the client may be deemed unresponsive.
	Pong(serial uint32)

	// Start a pointer-driven move of the surface.
	//
	// This request must be used in response to a button press event.
	// The server may ignore move requests depending on the state of
	// the surface (e.g. fullscreen or maximized).
	Move(seat Object, serial uint32)

	// Start a pointer-driven resizing of the surface.
	//
	// This request must be used in response to a button press event.
	// The server may ignore resize requests depending on the state of
	// the surface (e.g. fullscreen or maximized).
	Resize(seat Object, serial uint32, edges uint32)

	// Map the surface as a toplevel surface.
	//
	// A toplevel surface is not fullscreen, maximized or transient.
	SetToplevel()

	// Map the surface relative to an existing surface.
	//
	// The x and y arguments specify the locations of the upper left
	// corner of the surface relative to the upper left corner of the
	// parent surface, in surface local coordinates.
	//
	// The flags argument controls details of the transient behaviour.
	SetTransient(parent Object, x int32, y int32, flags uint32)

	// Map the surface as a fullscreen surface.
	//
	// If an output parameter is given then the surface will be made
	// fullscreen on that output. If the client does not specify the
	// output then the compositor will apply its policy - usually
	// choosing the output on which the surface has the biggest surface
	// area.
	//
	// The client may specify a method to resolve a size conflict
	// between the output size and the surface size - this is provided
	// through the method parameter.
	//
	// The framerate parameter is used only when the method is set
	// to "driver", to indicate the preferred framerate. A value of 0
	// indicates that the app does not care about framerate. The
	// framerate is specified in mHz, that is framerate of 60000 is 60Hz.
	//
	// A method of "scale" or "driver" implies a scaling operation of
	// the surface, either via a direct scaling operation or a change of
	// the output mode. This will override any kind of output scaling, so
	// that mapping a surface with a buffer size equal to the mode can
	// fill the screen independent of buffer_scale.
	//
	// A method of "fill" means we don't scale up the buffer, however
	// any output scale is applied. This means that you may run into
	// an edge case where the application maps a buffer with the same
	// size of the output mode but buffer_scale 1 (thus making a
	// surface larger than the output). In this case it is allowed to
	// downscale the results to fit the screen.
	//
	// The compositor must reply to this request with a configure event
	// with the dimensions for the output on which the surface will
	// be made fullscreen.
	SetFullscreen(method uint32, framerate uint32, output Object)

	// Map the surface as a popup.
	//
	// A popup surface is a transient surface with an added pointer grab.
	//
	// An existing implicit grab will be changed to owner-events mode,
	// and the popup grab will continue after the implicit grab ends
	// (i.e. releasing the mouse button does not cause the popup to
	// be unmapped).
	//
	// The popup grab continues until the window is destroyed or a
	// mouse button is pressed in any other clients window. A click
	// in any of the clients surfaces is reported as normal, however,
	// clicks in other clients surfaces will be discarded and trigger
	// the callback.
	//
	// The x and y arguments specify the locations of the upper left
	// corner of the surface relative to the upper left corner of the
	// parent surface, in surface local coordinates.
	SetPopup(seat Object, serial uint32, parent Object, x int32, y int32, flags uint32)

	// Map the surface as a maximized surface.
	//
	// If an output parameter is given then the surface will be
	// maximized on that output. If the client does not specify the
	// output then the compositor will apply its policy - usually
	// choosing the output on which the surface has the biggest surface
	// area.
	//
	// The compositor will reply with a configure event telling
	// the expected new surface size. The operation is completed
	// on the next buffer attach to this surface.
	//
	// A maximized surface typically fills the entire output it is
	// bound to, except for desktop element such as panels. This is
	// the main difference between a maximized shell surface and a
	// fullscreen shell surface.
	//
	// The details depend on the compositor implementation.
	SetMaximized(output Object)

	// Set a short title for the surface.
	//
	// This string may be used to identify the surface in a task bar,
	// window list, or other user interface elements provided by the
	// compositor.
	//
	// The string must be encoded in UTF-8.
	SetTitle(title string)

	// Set a class for the surface.
	//
	// The surface class identifies the general class of applications
	// to which the surface belongs. A common convention is to use the
	// file name (or the full path if it is a non-standard location) of
	// the application's .desktop file as the class.
	SetClass(class string)

	// Ping a client to check if it is receiving events and sending
	// requests. A client is expected to reply with a pong request.
	Ping(serial uint32)

	// The configure event asks the client to resize its surface.
	//
	// The size is a hint, in the sense that the client is free to
	// ignore it if it doesn't resize, pick a smaller size (to
	// satisfy aspect ratio or resize in steps of NxM pixels).
	//
	// The edges parameter provides a hint about how the surface
	// was resized. The client may use this information to decide
	// how to adjust its content to the new size (e.g. a scrolling
	// area might adjust its content position to leave the viewable
	// content unmoved).
	//
	// The client is free to dismiss all but the last configure
	// event it received.
	//
	// The width and height arguments specify the size of the window
	// in surface local coordinates.
	Configure(edges uint32, width int32, height int32)

	// The popup_done event is sent out when a popup grab is broken,
	// that is, when the user clicks a surface that doesn't belong
	// to the client owning the popup surface.
	PopupDone()
}

// These values are used to indicate which edge of a surface
// is being dragged in a resize operation. The server may
// use this information to adapt its behavior, e.g. choose
// an appropriate cursor image.
type Resize uint

const (
	None        Resize = 0
	Top         Resize = 1
	Bottom      Resize = 2
	Left        Resize = 4
	TopLeft     Resize = 5
	BottomLeft  Resize = 6
	Right       Resize = 8
	TopRight    Resize = 9
	BottomRight Resize = 10
)

// These flags specify details of the expected behaviour
// of transient surfaces. Used in the settransient request.
type Transient uint

const (
	Inactive Transient = 0x1
)

// Hints to indicate to the compositor how to deal with a conflict
// between the dimensions of the surface and the dimensions of the
// output. The compositor is free to ignore this parameter.
type FullScreenMethod uint

const (
	Default FullScreenMethod = iota
	Scale
	Driver
	Fill
)

// A surface is a rectangular area that is displayed on the screen.
// It has a location, size and pixel contents.
//
// The size of a surface (and relative positions on it) is described
// in surface local coordinates, which may differ from the buffer
// local coordinates of the pixel content, in case a buffer_transform
// or a buffer_scale is used.
//
// Surfaces are also used for some special purposes, e.g. as
// cursor images for pointers, drag icons, etc.
type Surface interface {

	// Deletes the surface and invalidates its object ID.
	Destroy()

	// Set a buffer as the content of this surface.
	//
	// The new size of the surface is calculated based on the buffer
	// size transformed by the inverse buffer_transform and the
	// inverse buffer_scale. This means that the supplied buffer
	// must be an integer multiple of the buffer_scale.
	//
	// The x and y arguments specify the location of the new pending
	// buffer's upper left corner, relative to the current buffer's upper
	// left corner, in surface local coordinates. In other words, the
	// x and y, combined with the new surface size define in which
	// directions the surface's size changes.
	//
	// Surface contents are double-buffered state, see wl_surface.commit.
	//
	// The initial surface contents are void; there is no content.
	// wl_surface.attach assigns the given wl_buffer as the pending
	// wl_buffer. wl_surface.commit makes the pending wl_buffer the new
	// surface contents, and the size of the surface becomes the size
	// calculated from the wl_buffer, as described above. After commit,
	// there is no pending buffer until the next attach.
	//
	// Committing a pending wl_buffer allows the compositor to read the
	// pixels in the wl_buffer. The compositor may access the pixels at
	// any time after the wl_surface.commit request. When the compositor
	// will not access the pixels anymore, it will send the
	// wl_buffer.release event. Only after receiving wl_buffer.release,
	// the client may re-use the wl_buffer. A wl_buffer that has been
	// attached and then replaced by another attach instead of committed
	// will not receive a release event, and is not used by the
	// compositor.
	//
	// Destroying the wl_buffer after wl_buffer.release does not change
	// the surface contents. However, if the client destroys the
	// wl_buffer before receiving the wl_buffer.release event, the surface
	// contents become undefined immediately.
	//
	// If wl_surface.attach is sent with a NULL wl_buffer, the
	// following wl_surface.commit will remove the surface content.
	Attach(buffer Object, x int32, y int32)

	// This request is used to describe the regions where the pending
	// buffer is different from the current surface contents, and where
	// the surface therefore needs to be repainted. The pending buffer
	// must be set by wl_surface.attach before sending damage. The
	// compositor ignores the parts of the damage that fall outside of
	// the surface.
	//
	// Damage is double-buffered state, see wl_surface.commit.
	//
	// The damage rectangle is specified in surface local coordinates.
	//
	// The initial value for pending damage is empty: no damage.
	// wl_surface.damage adds pending damage: the new pending damage
	// is the union of old pending damage and the given rectangle.
	//
	// wl_surface.commit assigns pending damage as the current damage,
	// and clears pending damage. The server will clear the current
	// damage as it repaints the surface.
	Damage(x int32, y int32, width int32, height int32)

	// Request a notification when it is a good time start drawing a new
	// frame, by creating a frame callback. This is useful for throttling
	// redrawing operations, and driving animations.
	//
	// When a client is animating on a wl_surface, it can use the 'frame'
	// request to get notified when it is a good time to draw and commit the
	// next frame of animation. If the client commits an update earlier than
	// that, it is likely that some updates will not make it to the display,
	// and the client is wasting resources by drawing too often.
	//
	// The frame request will take effect on the next wl_surface.commit.
	// The notification will only be posted for one frame unless
	// requested again. For a wl_surface, the notifications are posted in
	// the order the frame requests were committed.
	//
	// The server must send the notifications so that a client
	// will not send excessive updates, while still allowing
	// the highest possible update rate for clients that wait for the reply
	// before drawing again. The server should give some time for the client
	// to draw and commit after sending the frame callback events to let them
	// hit the next output refresh.
	//
	// A server should avoid signalling the frame callbacks if the
	// surface is not visible in any way, e.g. the surface is off-screen,
	// or completely obscured by other opaque surfaces.
	//
	// The object returned by this request will be destroyed by the
	// compositor after the callback is fired and as such the client must not
	// attempt to use it after that point.
	//
	// The callback_data passed in the callback is the current time, in
	// milliseconds.
	Frame(callback NewId)

	// This request sets the region of the surface that contains
	// opaque content.
	//
	// The opaque region is an optimization hint for the compositor
	// that lets it optimize out redrawing of content behind opaque
	// regions. Setting an opaque region is not required for correct
	// behaviour, but marking transparent content as opaque will result
	// in repaint artifacts.
	//
	// The opaque region is specified in surface local coordinates.
	//
	// The compositor ignores the parts of the opaque region that fall
	// outside of the surface.
	//
	// Opaque region is double-buffered state, see wl_surface.commit.
	//
	// wl_surface.set_opaque_region changes the pending opaque region.
	// wl_surface.commit copies the pending region to the current region.
	// Otherwise, the pending and current regions are never changed.
	//
	// The initial value for opaque region is empty. Setting the pending
	// opaque region has copy semantics, and the wl_region object can be
	// destroyed immediately. A NULL wl_region causes the pending opaque
	// region to be set to empty.
	SetOpaqueRegion(region Object)

	// This request sets the region of the surface that can receive
	// pointer and touch events.
	//
	// Input events happening outside of this region will try the next
	// surface in the server surface stack. The compositor ignores the
	// parts of the input region that fall outside of the surface.
	//
	// The input region is specified in surface local coordinates.
	//
	// Input region is double-buffered state, see wl_surface.commit.
	//
	// wl_surface.set_input_region changes the pending input region.
	// wl_surface.commit copies the pending region to the current region.
	// Otherwise the pending and current regions are never changed,
	// except cursor and icon surfaces are special cases, see
	// wl_pointer.set_cursor and wl_data_device.start_drag.
	//
	// The initial value for input region is infinite. That means the
	// whole surface will accept input. Setting the pending input region
	// has copy semantics, and the wl_region object can be destroyed
	// immediately. A NULL wl_region causes the input region to be set
	// to infinite.
	SetInputRegion(region Object)

	// Surface state (input, opaque, and damage regions, attached buffers,
	// etc.) is double-buffered. Protocol requests modify the pending
	// state, as opposed to current state in use by the compositor. Commit
	// request atomically applies all pending state, replacing the current
	// state. After commit, the new pending state is as documented for each
	// related request.
	//
	// On commit, a pending wl_buffer is applied first, all other state
	// second. This means that all coordinates in double-buffered state are
	// relative to the new wl_buffer coming into use, except for
	// wl_surface.attach itself. If there is no pending wl_buffer, the
	// coordinates are relative to the current surface contents.
	//
	// All requests that need a commit to become effective are documented
	// to affect double-buffered state.
	//
	// Other interfaces may add further double-buffered surface state.
	Commit()

	// This request sets an optional transformation on how the compositor
	// interprets the contents of the buffer attached to the surface. The
	// accepted values for the transform parameter are the values for
	// wl_output.transform.
	//
	// Buffer transform is double-buffered state, see wl_surface.commit.
	//
	// A newly created surface has its buffer transformation set to normal.
	//
	// wl_surface.set_buffer_transform changes the pending buffer
	// transformation. wl_surface.commit copies the pending buffer
	// transformation to the current one. Otherwise, the pending and current
	// values are never changed.
	//
	// The purpose of this request is to allow clients to render content
	// according to the output transform, thus permiting the compositor to
	// use certain optimizations even if the display is rotated. Using
	// hardware overlays and scanning out a client buffer for fullscreen
	// surfaces are examples of such optimizations. Those optimizations are
	// highly dependent on the compositor implementation, so the use of this
	// request should be considered on a case-by-case basis.
	//
	// Note that if the transform value includes 90 or 270 degree rotation,
	// the width of the buffer will become the surface height and the height
	// of the buffer will become the surface width.
	//
	// If transform is not one of the values from the
	// wl_output.transform enum the invalid_transform protocol error
	// is raised.
	SetBufferTransform(transform int32)

	// This request sets an optional scaling factor on how the compositor
	// interprets the contents of the buffer attached to the window.
	//
	// Buffer scale is double-buffered state, see wl_surface.commit.
	//
	// A newly created surface has its buffer scale set to 1.
	//
	// wl_surface.set_buffer_scale changes the pending buffer scale.
	// wl_surface.commit copies the pending buffer scale to the current one.
	// Otherwise, the pending and current values are never changed.
	//
	// The purpose of this request is to allow clients to supply higher
	// resolution buffer data for use on high resolution outputs. Its
	// intended that you pick the same	buffer scale as the scale of the
	// output that the surface is displayed on.This means the compositor
	// can avoid scaling when rendering the surface on that output.
	//
	// Note that if the scale is larger than 1, then you have to attach
	// a buffer that is larger (by a factor of scale in each dimension)
	// than the desired surface size.
	//
	// If scale is not positive the invalid_scale protocol error is
	// raised.
	SetBufferScale(scale int32)

	// This is emitted whenever a surface's creation, movement, or resizing
	// results in some part of it being within the scanout region of an
	// output.
	//
	// Note that a surface may be overlapping with zero or more outputs.
	Enter(output Object)

	// This is emitted whenever a surface's creation, movement, or resizing
	// results in it no longer having any part of it within the scanout region
	// of an output.
	Leave(output Object)
}

// These errors can be emitted in response to wl_surface requests.
type SurfaceErr uint

const (
	InvalidScale SurfaceErr = iota
	InvalidTransform
)

func (e DisplayErr) Error() string {
	switch e {
	case InvalidScale:
		return "buffer scale value is invalid"
	case InvalidTransform:
		return "buffer transform value is invalid"
	default:
		return "unrecognized error value"
	}

}

// A seat is a group of keyboards, pointer and touch devices. This
// object is published as a global during start up, or when such a
// device is hot plugged. A seat typically has a pointer and
// maintains a keyboard focus and a pointer focus.
type Seat interface {

	// The ID provided will be initialized to the wl_pointer interface
	// for this seat.
	//
	// This request only takes effect if the seat has the pointer
	// capability.
	GetPointer(id NewId)

	// The ID provided will be initialized to the wl_keyboard interface
	// for this seat.
	//
	// This request only takes effect if the seat has the keyboard
	// capability.
	GetKeyboard(id NewId)

	// The ID provided will be initialized to the wl_touch interface
	// for this seat.
	//
	// This request only takes effect if the seat has the touch
	// capability.
	GetTouch(id NewId)

	// This is emitted whenever a seat gains or loses the pointer,
	// keyboard or touch capabilities. The argument is a capability
	// enum containing the complete set of capabilities this seat has.
	Capabilities(capabilities uint32)

	// In a multiseat configuration this can be used by the client to help
	// identify which physical devices the seat represents. Based on
	// the seat configuration used by the compositor.
	Name(name string)
}

// This is a bitmask of capabilities this seat has; if a member is
// set, then it is present on the seat.
type CapabilityMask uint

const (
	PointerCap CapabilityMask = 1 << iota
	KeyboardCap
	TouchCap
)

// The wl_pointer interface represents one or more input devices,
// such as mice, which control the pointer location and pointer_focus
// of a seat.
//
// The wl_pointer interface generates motion, enter and leave
// events for the surfaces that the pointer is located over,
// and button and axis events for button presses, button releases
// and scrolling.
type Pointer interface {

	// Set the pointer surface, i.e., the surface that contains the
	// pointer image (cursor	. This request only takes effect if the pointer
	// focus for this device is one of the requesting client's surfaces
	// or the surface parameter is the current pointer surface. If
	// there was a previous surface set with this request it is
	// replaced. If surface is NULL, the pointer image is hidden.
	//
	// The parameters hotspot_x and hotspot_y define the position of
	// the pointer surface relative to the pointer location. Its
	// top-left corner is always at (x, y)- (hotspot_x, hotspot_y),
	// where (x, y) are the coordinates of the pointer location, in surface
	// local coordinates.
	//
	// On surface.attach requests to the pointer surface, hotspot_x
	// and hotspot_y are decremented by the x and y parameters
	// passed to the request. Attach must be confirmed by
	// wl_surface.commit as usual.
	//
	// The hotspot can also be updated by passing the currently set
	// pointer surface to this request with new values for hotspot_x
	// and hotspot_y.
	//
	// The current and pending input regions of the wl_surface are
	// cleared, and wl_surface.set_input_region is ignored until the
	// wl_surface is no longer used as the cursor. When the use as a
	// cursor ends, the current and pending input regions become
	// undefined, and the wl_surface is unmapped.
	SetCursor(serial uint32, surface Object, hotspotX int32, hotspotY int32)

	Release()

	// Notification that this seat's pointer is focused on a certain
	// surface.
	//
	// When an seat's focus enters a surface, the pointer image
	// is undefined and a client should respond to this event by setting
	// an appropriate pointer image with the set_cursor request.
	Enter(serial uint32, surface Object, surfaceX Fixed, surfaceY Fixed)

	// Notification that this seat's pointer is no longer focused on
	// a certain surface.
	//
	// The leave notification is sent before the enter notification
	// for the new focus.
	Leave(serial uint32, surface Object)

	// Notification of pointer location change. The arguments
	// surface_x and surface_y are the location relative to the
	// focused surface.
	Motion(time uint32, surfaceX Fixed, surfaceY Fixed)

	// Mouse button click and release notifications.

	// The location of the click is given by the last motion or
	// enter event.
	// The time argument is a timestamp with millisecond
	// granularity, with an undefined base.
	Button(serial uint32, time uint32, button uint32, state uint32)

	// Scroll and other axis notifications.
	//
	// For scroll events (vertical and horizontal scroll axes), the
	// value parameter is the length of a vector along the specified
	// axis in a coordinate space identical to those of motion events,
	// representing a relative movement along the specified axis.
	//
	// For devices that support movements non-parallel to axes multiple
	// axis events will be emitted.
	//
	// When applicable, for example for touch pads, the server can
	// choose to emit scroll events where the motion vector is
	// equivalent to a motion event vector.
	//
	// When applicable, clients can transform its view relative to the
	// scroll distance.
	Axis(time uint32, axis uint32, value Fixed)
}

// Describes the physical state of a button which provoked the button event.
type ButtonState uint

const (
	Released uint = iota
	Pressed
)

// Describes the axis types of scroll events.
type Axis uint

const (
	VerticalScroll Axis = iota
	HorizontalScroll
)

// The wl_keyboard interface represents one or more keyboards
// associated with a seat.
type Keyboard interface {
	Release()

	// This event provides a file descriptor to the client which can be
	// memory-mapped to provide a keyboard mapping description.
	Keymap(format uint32, fd Fd, size uint32)

	// Notification that this seat's keyboard focus is on a certain
	// surface.
	Enter(serial uint32, surface Object, keys Array)

	// Notification that this seat's keyboard focus is no longer on
	// a certain surface.
	//
	// The leave notification is sent before the enter notification
	// for the new focus.
	Leave(serial uint32, surface Object)

	// A key was pressed or released.
	// The time argument is a timestamp with millisecond
	// granularity, with an undefined base.
	Key(serial uint32, time uint32, key uint32, state uint32)

	// Notifies clients that the modifier and/or group state has
	// changed, and it should update its local state.
	Modifiers(serial uint32, modsDepressed uint32, modsLatched uint32, modsLocked uint32, group uint32)

	// Informs the client about the keyboard's repeat rate and delay.
	//
	// This event is sent as soon as the wl_keyboard object has been created,
	// and is guaranteed to be received by the client before any key press
	// event.
	//
	// Negative values for either rate or delay are illegal. A rate of zero
	// will disable any repeating (regardless of the value of delay).
	//
	// This event can be sent later on as well with a new value if necessary,
	// so clients should continue listening for the event past the creation
	// of wl_keyboard.
	RepeatInfo(rate int32, delay int32)
}

// This specifies the format of the keymap provided to the
// client with the wl_keyboard.keymap event.
type KeymapFormat uint

const (
	NoKeymap KeymapFormat = iota
	XkbV1
)

// Describes the physical state of a key which provoked the key event.
type KeyState uint

const (
	Released KeyState = iota
	Pressed
)

// The wl_touch interface represents a touchscreen
// associated with a seat.
//
// Touch interactions can consist of one or more contacts.
// For each contact, a series of events is generated, starting
// with a down event, followed by zero or more motion events,
// and ending with an up event. Events relating to the same
// contact point can be identified by the ID of the sequence.

type Touch interface {
	Release()

	// A new touch point has appeared on the surface. This touch point is
	// assigned a unique @id. Future events from this touchpoint reference
	// this ID. The ID ceases to be valid after a touch up event and may be
	// re-used in the future.
	Down(serial uint32, time uint32, surface Object, id int32, x Fixed, y Fixed)

	// The touch point has disappeared. No further events will be sent for
	// this touchpoint and the touch point's ID is released and may be
	// re-used in a future touch down event.
	Up(serial uint32, time uint32, id int32)

	// A touchpoint has changed coordinates.
	Motion(time uint32, id int32, x Fixed, y Fixed)

	// Indicates the end of a contact point list.
	Frame()

	// Sent if the compositor decides the touch stream is a global
	// gesture. No further events are sent to the clients from that
	// particular gesture. Touch cancellation applies to all touch points
	// currently active on this client's surface. The client is
	// responsible for finalizing the touch points, future touch points on
	// this surface may re-use the touch point ID.
	Cancel()
}

// An output describes part of the compositor geometry. The
// compositor works in the 'compositor coordinate system' and an
// output corresponds to rectangular area in that space that is
// actually visible. This typically corresponds to a monitor that
// displays part of the compositor space. This object is published
// as global during start up, or when a monitor is hotplugged.
type Output interface {

	// The geometry event describes geometric properties of the output.
	// The event is sent when binding to the output object and whenever
	// any of the properties change.
	Geometry(x int32, y int32, physicalWidth int32, physicalHeight int32, subpixel int32, make string, model string, transform int32)

	// The mode event describes an available mode for the output.
	//
	// The event is sent when binding to the output object and there
	// will always be one mode, the current mode. The event is sent
	// again if an output changes mode, for the mode that is now
	// current. In other words, the current mode is always the last
	// mode that was received with the current flag set.
	//
	// The size of a mode is given in physical hardware units of
	// the output device. This is not necessarily the same as
	// the output size in the global compositor space. For instance,
	// the output may be scaled, as described in wl_output.scale,
	// or transformed , as described in wl_output.transform.
	Mode(flags uint32, width int32, height int32, refresh int32)

	// This event is sent after all other properties has been
	// sent after binding to the output object and after any
	// other property changes done after that. This allows
	// changes to the output properties to be seen as
	// atomic, even if they happen via multiple events.
	Done()

	// This event contains scaling geometry information
	// that is not in the geometry event. It may be sent after
	// binding the output object or if the output scale changes
	// later. If it is not sent, the client should assume a
	// scale of 1.
	//
	// A scale larger than 1 means that the compositor will
	// automatically scale surface buffers by this amount
	// when rendering. This is used for very high resolution
	// displays where applications rendering at the native
	// resolution would be too small to be legible.
	//
	// It is intended that scaling aware clients track the
	// current output of a surface, and if it is on a scaled
	// output it should use wl_surface.set_buffer_scale with
	// the scale of the output. That way the compositor can
	// avoid scaling the surface, and the client can supply
	// a higher detail image.
	Scale(factor int32)
}

// This enumeration describes how the physical
// pixels on an output are layed out.
type Subpixel uint

const (
	Unknown Subpixel = iota
	None
	HorizontalRgb
	HorizontalBgr
	VerticalRgb
	VerticalBgr
)

// This describes the transform that a compositor will apply to a
// surface to compensate for the rotation or mirroring of an
// output device.
//
// The flipped values correspond to an initial flip around a
// vertical axis followed by rotation.
//
// The purpose is mainly to allow clients render accordingly and
// tell the compositor, so that for fullscreen surfaces, the
// compositor will still be able to scan out directly from client
// surfaces.

type Transform uint

const (
	Normal Transform = iota
	R90
	R180
	R270
	Flipped
	Flipped90
	Flipped180
	Flipped270
)

// These flags describe properties of an output mode.
// They are used in the flags bitfield of the mode event.
type Mode uint

const (
	Current   Mode = 0x1
	Preferred Mode = 0x2
)

// A region object describes an area.
//
// Region objects are used to describe the opaque and input
// regions of a surface.
type Region interface {

	// Destroy the region. This will invalidate the object ID.
	Destroy()

	// Add the specified rectangle to the region.
	Add(x int32, y int32, width int32, height int32)

	// Subtract the specified rectangle from the region.
	Subtract(x int32, y int32, width int32, height int32)
}

// The global interface exposing sub-surface compositing capabilities.
// A wl_surface, that has sub-surfaces associated, is called the
// parent surface. Sub-surfaces can be arbitrarily nested and create
// a tree of sub-surfaces.
//
// The root surface in a tree of sub-surfaces is the main
// surface. The main surface cannot be a sub-surface, because
// sub-surfaces must always have a parent.
//
// A main surface with its sub-surfaces forms a (compound) window.
// For window management purposes, this set of wl_surface objects is
// to be considered as a single window, and it should also behave as
// such.
//
// The aim of sub-surfaces is to offload some of the compositing work
// within a window from clients to the compositor. A prime example is
// a video player with decorations and video in separate wl_surface
// objects. This should allow the compositor to pass YUV video buffer
// processing to dedicated overlay hardware when possible.

type Subcompositor interface {

	// Informs the server that the client will not be using this
	// protocol object anymore. This does not affect any other
	// objects, wl_subsurface objects included.
	Destroy()

	// Create a sub-surface interface for the given surface, and
	// associate it with the given parent surface. This turns a
	// plain wl_surface into a sub-surface.
	//
	// The to-be sub-surface must not already have a dedicated
	// purpose, like any shell surface type, cursor image, drag icon,
	// or sub-surface. Otherwise a protocol error is raised.
	GetSubsurface(id NewId, surface Object, parent Object)
}
type SubcompositorErr uint

const (
	BadSurface SubcompositorErr = iota
)

func (err SubcompositorErr) Error() string {
	switch err {
	case BadSurface:
		return "the to-be sub-surface is invalid"
	}
}

// An additional interface to a wl_surface object, which has been
// made a sub-surface. A sub-surface has one parent surface. A
// sub-surface's size and position are not limited to that of the parent.
// Particularly, a sub-surface is not automatically clipped to its
// parent's area.
//
// A sub-surface becomes mapped, when a non-NULL wl_buffer is applied
// and the parent surface is mapped. The order of which one happens
// first is irrelevant. A sub-surface is hidden if the parent becomes
// hidden, or if a NULL wl_buffer is applied. These rules apply
// recursively through the tree of surfaces.
//
// The behaviour of wl_surface.commit request on a sub-surface
// depends on the sub-surface's mode. The possible modes are
// synchronized and desynchronized, see methods
// wl_subsurface.set_sync and wl_subsurface.set_desync. Synchronized
// mode caches the wl_surface state to be applied when the parent's
// state gets applied, and desynchronized mode applies the pending
// wl_surface state directly. A sub-surface is initially in the
// synchronized mode.
//
// Sub-surfaces have also other kind of state, which is managed by
// wl_subsurface requests, as opposed to wl_surface requests. This
// state includes the sub-surface position relative to the parent
// surface (wl_subsurface.set_position), and the stacking order of
// the parent and its sub-surfaces (wl_subsurface.place_above and
// .place_below) This state is applied when the parent surface's
// wl_surface state is applied, regardless of the sub-surface's mode.
// As the exception, set_sync and set_desync are effective immediately.
//
// The main surface can be thought to be always in desynchronized mode,
// since it does not have a parent in the sub-surfaces sense.
//
// Even if a sub-surface is in desynchronized mode, it will behave as
// in synchronized mode, if its parent surface behaves as in
// synchronized mode. This rule is applied recursively throughout the
// tree of surfaces. This means, that one can set a sub-surface into
// synchronized mode, and then assume that all its child and grand-child
// sub-surfaces are synchronized, too, without explicitly setting them.
//
// If the wl_surface associated with the wl_subsurface is destroyed, the
// wl_subsurface object becomes inert. Note, that destroying either object
// takes effect immediately. If you need to synchronize the removal
// of a sub-surface to the parent surface update, unmap the sub-surface
// first by attaching a NULL wl_buffer, update parent, and then destroy
// the sub-surface.
//
// If the parent wl_surface object is destroyed, the sub-surface is
// unmapped.
type Subsurface interface {

	// The sub-surface interface is removed from the wl_surface object
	// that was turned into a sub-surface with
	// wl_subcompositor.get_subsurface request. The wl_surface's association
	// to the parent is deleted, and the wl_surface loses its role as
	// a sub-surface. The wl_surface is unmapped.
	Destroy()

	// This schedules a sub-surface position change.
	// The sub-surface will be moved so, that its origin (top-left
	// corner pixel) will be at the location x, y of the parent surface
	// coordinate system. The coordinates are not restricted to the parent
	// surface area. Negative values are allowed.
	//
	// The next wl_surface.commit on the parent surface will reset
	// the sub-surface's position to the scheduled coordinates.
	//
	// If more than one set_position request is invoked by the client before
	// the commit of the parent surface, the position of a new request always
	// replaces the scheduled position from any previous request.
	//
	// The initial position is 0, 0.
	SetPosition(x int32, y int32)

	// This sub-surface is taken from the stack, and put back just
	// above the reference surface, changing the z-order of the sub-surfaces.
	// The reference surface must be one of the sibling surfaces, or the
	// parent surface. Using any other surface, including this sub-surface,
	// will cause a protocol error.
	//
	// The z-order is double-buffered. Requests are handled in order and
	// applied immediately to a pending state, then committed to the active
	// state on the next commit of the parent surface.
	// See wl_surface.commit and wl_subcompositor.get_subsurface.
	//
	// A new sub-surface is initially added as the top-most in the stack
	// of its siblings and parent.
	PlaceAbove(sibling Object)

	// The sub-surface is placed just below of the reference surface.
	// See wl_subsurface.place_above.
	PlaceBelow(sibling Object)

	// Change the commit behaviour of the sub-surface to synchronized
	// mode, also described as the parent dependant mode.
	//
	// In synchronized mode, wl_surface.commit on a sub-surface will
	// accumulate the committed state in a cache, but the state will
	// not be applied and hence will not change the compositor output.
	// The cached state is applied to the sub-surface immediately after
	// the parent surface's state is applied. This ensures atomic
	// updates of the parent and all its synchronized sub-surfaces.
	// Applying the cached state will invalidate the cache, so further
	// parent surface commits do not (re-) apply old state.
	//
	// See wl_subsurface for the recursive effect of this mode.
	SetSync()

	// Change the commit behaviour of the sub-surface to desynchronized
	// mode, also described as independent or freely running mode.
	//
	// In desynchronized mode, wl_surface.commit on a sub-surface will
	// apply the pending state directly, without caching, as happens
	// normally with a wl_surface. Calling wl_surface.commit on the
	// parent surface has no effect on the sub-surface's wl_surface
	// state. This mode allows a sub-surface to be updated on its own.
	//
	// If cached state exists when wl_surface.commit is called in
	// desynchronized mode, the pending state is added to the cached
	// state, and applied as whole. This invalidates the cache.
	//
	// Note: even if a sub-surface is set to desynchronized, a parent
	// sub-surface may override it to behave as synchronized. For details,
	// see wl_subsurface.
	//
	// If a surface's parent surface behaves as desynchronized, then
	// the cached state is applied on set_desync.
	SetDesync()
}

type SubsurfaceErr uint

const (
	BadSurface SubsurfaceErr = iota
)

func (err SubcompositorErr) Error() string {
	switch err {
	case BadSurface:
		return "wl_surface is not a sibling or the parent"
	}
}
