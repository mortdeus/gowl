package gowl

	/*
	
      A surface.  This is an image that is displayed on the screen.
      It has a location, size and pixel contents.
    
	*/
	type Surface struct{}

		

		
			/*
			
	Deletes the surface and invalidates its object id.
      */
			func (*Surface) Destroy(){
		}
		
			/*
			
	Set the contents of a buffer into this surface. The x and y
	arguments specify the location of the new pending buffer's upper
	left corner, relative to the current buffer's upper left corner. In
	other words, the x and y, and the width and height of the wl_buffer
	together define in which directions the surface's size changes.

	Surface contents are double-buffered state, see wl_surface.commit.

	The initial surface contents are void; there is no content.
	wl_surface.attach assigns the given wl_buffer as the pending wl_buffer.
	wl_surface.commit applies the pending wl_buffer as the new
	surface contents, and the size of the surface becomes the size of
	the wl_buffer. The wl_buffer is also kept as pending, until
	changed by wl_surface.attach or the wl_buffer is destroyed.

	Committing a pending wl_buffer allows the compositor to read the
	pixels in the wl_buffer. The compositor may access the pixels at any
	time after the wl_surface.commit request. When the compositor will
	not access the pixels anymore, it will send the wl_buffer.release
	event. Only after receiving wl_buffer.release, the client may re-use
	the wl_buffer. A wl_buffer, that has been attached and then replaced
	by another attach instead of committed, will not receive a release
	event, and is not used by the compositor.

	Destroying the wl_buffer after wl_buffer.release does not change the
	surface contents, even if the wl_buffer is still pending for the
	next commit. In such case, the next commit does not change the
	surface contents. However, if the client destroys the wl_buffer
	before receiving wl_buffer.release, the surface contents become
	undefined immediately.

	Only if wl_surface.attach is sent with a nil wl_buffer, the
	following wl_surface.commit will remove the surface content.
      */
			func (*Surface) Attach(buffer object,x int,y int,){
		}
		
			/*
			
	This request is used to describe the regions where the pending
	buffer (or if pending buffer is none, the current buffer as updated
	in-place) on the next wl_surface.commit will be different from the
	current buffer, and needs to be repainted. The pending buffer can be
	set by wl_surface.attach. The compositor ignores the parts of the
	damage that fall outside of the surface.

	Damage is double-buffered state, see wl_surface.commit.

	The initial value for pending damage is empty: no damage.
	wl_surface.damage adds pending damage: the new pending damage is the
	union of old pending damage and the given rectangle.
	wl_surface.commit assigns pending damage as the current damage, and
	clears pending damage. The server will clear the current damage as
	it repaints the surface.
      */
			func (*Surface) Damage(x int,y int,width int,height int,){
		}
		
			/*
			
	Request notification when the next frame is displayed. Useful
	for throttling redrawing operations, and driving animations.
	The frame request will take effect on the next wl_surface.commit.
	The notification will only be posted for one frame unless
	requested again.

	A server should avoid signalling the frame callbacks if the
	surface is not visible in any way, e.g. the surface is off-screen,
	or completely obscured by other opaque surfaces.

	A client can request a frame callback even without an attach,
	damage, or any other state changes. wl_surface.commit triggers a
	display update, so the callback event will arrive after the next
	output refresh where the surface is visible.
      */
			func (*Surface) Frame(callback new_id,){
		}
		
			/*
			
	This request sets the region of the surface that contains
	opaque content.  The opaque region is an optimization hint for
	the compositor that lets it optimize out redrawing of content
	behind opaque regions.  Setting an opaque region is not
	required for correct behaviour, but marking transparent
	content as opaque will result in repaint artifacts.
	The compositor ignores the parts of the opaque region that fall
	outside of the surface.

	Opaque region is double-buffered state, see wl_surface.commit.

	wl_surface.set_opaque_region changes the pending opaque region.
	wl_surface.commit copies the pending region to the current region.
	Otherwise the pending and current regions are never changed.

	The initial value for opaque region is empty. Setting the pending
	opaque region has copy semantics, and the wl_region object can be
	destroyed immediately. A nil wl_region causes the pending opaque
	region to be set to empty.
      */
			func (*Surface) SetOpaqueRegion(region object,){
		}
		
			/*
			
	This request sets the region of the surface that can receive
	pointer and touch events. Input events happening outside of
	this region will try the next surface in the server surface
	stack. The compositor ignores the parts of the input region that
	fall outside of the surface.

	Input region is double-buffered state, see wl_surface.commit.

	wl_surface.set_input_region changes the pending input region.
	wl_surface.commit copies the pending region to the current region.
	Otherwise the pending and current regions are never changed,
	except cursor and icon surfaces are special cases, see
	wl_pointer.set_cursor and wl_data_device.start_drag.

	The initial value for input region is infinite. That means the whole
	surface will accept input. Setting the pending input region has copy
	semantics, and the wl_region object can be destroyed immediately. A
	nil wl_region causes the input region to be set to infinite.
      */
			func (*Surface) SetInputRegion(region object,){
		}
		
			/*
			
	Surface state (input, opaque, and damage regions, attached buffers,
	etc.) is double-buffered. Protocol requests modify the pending
	state, as opposed to current state in use by the compositor. Commit
	request atomically applies all pending state, replacing the current
	state. After commit, the new pending state is as documented for each
	related request.

	On commit, a pending wl_buffer is applied first, all other state
	second. This means that all coordinates in double-buffered state are
	relative to the new wl_buffer coming into use, except for
	wl_surface.attach itself. If the pending wl_buffer is none, the
	coordinates are relative to the current surface contents.

	All requests that need a commit to become effective are documented
	to affect double-buffered state.

	Other interfaces may add further double-buffered surface state.
      */
			func (*Surface) Commit(){
		}
		
			/*
			
	This request sets an optional transformation on how the compositor
	interprets the contents of the buffer attached to the surface. The
	accepted values for the transform parameter are the values for
	wl_output.transform.

	Buffer transform is double-buffered state, see wl_surface.commit.

	A newly created surface has its buffer transformation set to normal.

	The purpose of this request is to allow clients to render content
	according to the output transform, thus permiting the compositor to
	use certain optimizations even if the display is rotated. Using
	hardware overlays and scanning out a client buffer for fullscreen
	surfaces are examples of such optmizations. Those optimizations are
	highly dependent on the compositor implementation, so the use of this
	request should be considered on a case-by-case basis.

	Note that if the transform value includes 90 or 270 degree rotation,
	the width of the buffer will become the surface height and the height
	of the buffer will become the surface width.
      */
			func (*Surface) SetBufferTransform(transform int,){
		}
		
	
		
			/*
			
        This is emitted whenever a surface's creation, movement, or resizing
        results in some part of it being within the scanout region of an
        output.
      */
				func (*Surface) Enter(output object,){
			}
		
			/*
			
        This is emitted whenever a surface's creation, movement, or resizing
        results in it no longer having any part of it within the scanout region
        of an output.
      */
				func (*Surface) Leave(output object,){
			}
		
	