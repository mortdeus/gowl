package gowl

	/*
	
      A group of keyboards, pointer (mice, for example) and touch
      devices . This object is published as a global during start up,
      or when such a device is hot plugged.  A seat typically has a
      pointer and maintains a keyboard_focus and a pointer_focus.
    
	*/
	type Seat struct{}

		
			/*
			
        This is a bitmask of capabilities this seat has; if a member is
	set, then it is present on the seat.
      */
			
			const(
				Capability_Pointer = 1
				Capability_Keyboard = 2
				Capability_Touch = 4
				
			)
		

		
			/*
			
        The ID provided will be initialized to the wl_pointer interface
	for this seat.
      */
			func (*Seat) GetPointer(id new_id,){
		}
		
			/*
			
        The ID provided will be initialized to the wl_keyboard interface
	for this seat.
      */
			func (*Seat) GetKeyboard(id new_id,){
		}
		
			/*
			
        The ID provided will be initialized to the wl_touch interface
	for this seat.
      */
			func (*Seat) GetTouch(id new_id,){
		}
		
	
		
			/*
			
        This is emitted whenever a seat gains or loses the pointer,
	keyboard or touch capabilities.  The argument is a wl_seat_caps_mask
	enum containing the complete set of capabilities this seat has.
      */
				func (*Seat) Capabilities(capabilities uint,){
			}
		
	