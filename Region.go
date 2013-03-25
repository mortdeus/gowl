package gowl

	/*
	
      Region.
    
	*/
	type Region struct{}

		

		
			/*
			
	Destroy the region.  This will invalidate the object id.
      */
			func (*Region) Destroy(){
		}
		
			/*
			
	Add the specified rectangle to the region
      */
			func (*Region) Add(x int,y int,width int,height int,){
		}
		
			/*
			
	Subtract the specified rectangle from the region
      */
			func (*Region) Subtract(x int,y int,width int,height int,){
		}
		
	
		
	