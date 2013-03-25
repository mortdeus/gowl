package gowl

	/*
	
      Support for shared memory buffers.
    
	*/
	type Shm struct{}

		
			
			
			const(
				Error_InvalidFormat = 0
				Error_InvalidStride = 1
				Error_InvalidFd = 2
				
			)
		
			
			
			const(
				Format_Argb8888 = 0
				Format_Xrgb8888 = 1
				
			)
		

		
			/*
			
	This creates wl_shm_pool object, which can be used to create
	shared memory based wl_buffer objects.  The server will mmap
	size bytes of the passed fd, to use as backing memory for then
	pool.
      */
			func (*Shm) CreatePool(id new_id,fd fd,size int,){
		}
		
	
		
			
				func (*Shm) Format(format uint,){
			}
		
	