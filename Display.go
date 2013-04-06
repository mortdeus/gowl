package gowl

/*

   The core global object.  This is a special singleton object.  It
   is used for internal Wayland protocol features.

*/
type Display struct{}

/*

	These errors are global and can be emitted in response to any
	server request.
*/

const (
	Error_InvalidObject = 0
	Error_InvalidMethod = 1
	Error_NoMemory      = 2
)

/*

	The sync request asks the server to emit the 'done' event
	on the provided wl_callback object.  Since requests are
	handled in-order, this can be used as a barrier to ensure all
	previous requests have been handled.
*/
func (*Display) Sync(callback NewId) {
}

/*

	This request creates a registry object that allows the client
	to list and bind the global objects available from the
	compositor.
*/
func (*Display) GetRegistry(callback NewId) {
}

/*

	The error event is sent out when a fatal (non-recoverable)
	error has occurred.  The @object_id argument is the object
	where the error occurred, most often in response to a request
	to that object.  The @code identifies the error and is defined
	by the object interface.  As such, each interface defines its
	own set of error codes.  The @message is an brief description
	of the error, for (debugging) convenience.
*/
func (*Display) Error(objectId Object, code uint32, message string) {
}

/*

	This event is used internally by the object ID management
	logic.  When a client deletes an object, the server will send
	this event to acknowledge that it has seen the delete request.
	When the client receive this event, it will know that it can
	safely reuse the object ID
*/
func (*Display) DeleteId(id uint32) {
}
