package gowl

	/*
	
      keyboard input device
    
	*/
	type Keyboard struct{}

		
			/*
			
        This enum specifies the format of the keymap provided to the client
	with the wl_keyboard::keymap event.
      */
			
			const(
				KeymapFormat_XkbV1 = 1
				
			)
		
			/*
			
        Describes the physical state of a key which provoked the key event.
      */
			
			const(
				KeyState_Released = 0
				KeyState_Pressed = 1
				
			)
		

		
	
		
			/*
			
        This event provides a file descriptor to the client which can be
	memory-mapped to provide a keyboard mapping description.
      */
				func (*Keyboard) Keymap(format uint,fd fd,size uint,){
			}
		
			
				func (*Keyboard) Enter(serial uint,surface object,keys array,){
			}
		
			
				func (*Keyboard) Leave(serial uint,surface object,){
			}
		
			/*
			
	A key was pressed or released.
      */
				func (*Keyboard) Key(serial uint,time uint,key uint,state uint,){
			}
		
			/*
			
        Notifies clients that the modifier and/or group state has
	changed, and it should update its local state.
      */
				func (*Keyboard) Modifiers(serial uint,modsDepressed uint,modsLatched uint,modsLocked uint,group uint,){
			}
		
	