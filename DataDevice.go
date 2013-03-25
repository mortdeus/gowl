package gowl

	
	type DataDevice struct{}

		

		
			/*
			
	This request asks the compositor to start a drag and drop
	operation on behalf of the client.

	The source argument is the data source that provides the data
	for the eventual data transfer. If source is NULL, enter, leave
	and motion events are sent only to the client that initiated the
	drag and the client is expected to handle the data passing
	internally.

	The origin surface is the surface where the drag originates and
	the client must have an active implicit grab that matches the
	serial.

	The icon surface is an optional (can be nil) surface that
	provides an icon to be moved around with the cursor.  Initially,
	the top-left corner of the icon surface is placed at the cursor
	hotspot, but subsequent wl_surface.attach request can move the
	relative position. Attach requests must be confirmed with
	wl_surface.commit as usual.

	The current and pending input regions of the icon wl_surface are
	cleared, and wl_surface.set_input_region is ignored until the
	wl_surface is no longer used as the icon surface. When the use
	as an icon ends, the the current and pending input regions
	become undefined, and the wl_surface is unmapped.
      */
			func (*DataDevice) StartDrag(source object,origin object,icon object,serial uint,){
		}
		
			
			func (*DataDevice) SetSelection(source object,serial uint,){
		}
		
	
		
			/*
			
	The data_offer event introduces a new wl_data_offer object,
	which will subsequently be used in either the
	data_device.enter event (for drag and drop) or the
	data_device.selection event (for selections).  Immediately
	following the data_device_data_offer event, the new data_offer
	object will send out data_offer.offer events to describe the
	mime-types it offers.
      */
				func (*DataDevice) DataOffer(id new_id,){
			}
		
			/*
			
	This event is sent when an active drag-and-drop pointer enters
	a surface owned by the client.  The position of the pointer at
	enter time is provided by the @x an @y arguments, in surface
	local coordinates.
      */
				func (*DataDevice) Enter(serial uint,surface object,x fixed,y fixed,id object,){
			}
		
			/*
			
	This event is sent when the drag-and-drop pointer leaves the
	surface and the session ends.  The client must destroy the
	wl_data_offer introduced at enter time at this point.
      */
				func (*DataDevice) Leave(){
			}
		
			/*
			
	This event is sent when the drag-and-drop pointer moves within
	the currently focused surface. The new position of the pointer
	is provided by the @x an @y arguments, in surface local
	coordinates.
      */
				func (*DataDevice) Motion(time uint,x fixed,y fixed,){
			}
		
			
				func (*DataDevice) Drop(){
			}
		
			/*
			
	The selection event is sent out to notify the client of a new
	wl_data_offer for the selection for this device.  The
	data_device.data_offer and the data_offer.offer events are
	sent out immediately before this event to introduce the data
	offer object.  The selection event is sent to a client
	immediately before receiving keyboard focus and when a new
	selection is set while the client has keyboard focus.  The
	data_offer is valid until a new data_offer or NULL is received
	or until the client loses keyboard focus.
      */
				func (*DataDevice) Selection(id object,){
			}
		
	