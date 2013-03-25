package gowl

	/*
	
    
	*/
	type Touch struct{}

		

		
	
		
			
				func (*Touch) Down(serial uint,time uint,surface object,id int,x fixed,y fixed,){
			}
		
			
				func (*Touch) Up(serial uint,time uint,id int,){
			}
		
			
				func (*Touch) Motion(time uint,id int,x fixed,y fixed,){
			}
		
			/*
			
	Indicates the end of a contact point list.
      */
				func (*Touch) Frame(){
			}
		
			/*
			
	Sent if the compositor decides the touch stream is a global
	gesture. No further events are sent to the clients from that
	particular gesture.
      */
				func (*Touch) Cancel(){
			}
		
	