package gowl

	/*
	
      An output describes part of the compositor geometry.  The
      compositor work in the 'compositor coordinate system' and an
      output corresponds to rectangular area in that space that is
      actually visible.  This typically corresponds to a monitor that
      displays part of the compositor space.  This object is published
      as global during start up, or when a screen is hot plugged.
    
	*/
	type Output struct{}

		
			
			
			const(
				Subpixel_Unknown = 0
				Subpixel_None = 1
				Subpixel_HorizontalRgb = 2
				Subpixel_HorizontalBgr = 3
				Subpixel_VerticalRgb = 4
				Subpixel_VerticalBgr = 5
				
			)
		
			/*
			
	This describes the transform that a compositor will apply to a
	surface to compensate for the rotation or mirroring of an
	output device.

	The flipped values correspond to an initial flip around a
	vertical axis followed by rotation.

	The purpose is mainly to allow clients render accordingly and
	tell the compositor, so that for fullscreen surfaces, the
	compositor will still be able to scan out directly from client
	surfaces.
      */
			
			const(
				Transform_Normal = 0
				Transform_Rot90 = 1
				Transform_Rot180 = 2
				Transform_Rot270 = 3
				Transform_Flipped = 4
				Transform_Flipped90 = 5
				Transform_Flipped180 = 6
				Transform_Flipped270 = 7
				
			)
		
			
			
			const(
				Mode_Current = 0x1
				Mode_Preferred = 0x2
				
			)
		

		
	
		
			
				func (*Output) Geometry(x int,y int,physicalWidth int,physicalHeight int,subpixel int,make string,model string,transform int,){
			}
		
			/*
			
	The mode event describes an available mode for the output.
	The event is sent when binding to the output object and there
	will always be one mode, the current mode.  The event is sent
	again if an output changes mode, for the mode that is now
	current.  In other words, the current mode is always the last
	mode that was received with the current flag set.
      */
				func (*Output) Mode(flags uint,width int,height int,refresh int,){
			}
		
	