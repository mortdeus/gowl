package gowl

	/*
	
      The wl_data_source object is the source side of a wl_data_offer.
      It is created by the source client in a data transfer and
      provides a way to describe the offered data and a way to respond
      to requests to transfer the data.
    
	*/
	type DataSource struct{}

		

		
			/*
			
	This request adds a mime-type to the set of mime-types
	advertised to targets.  Can be called several times to offer
	multiple types.
      */
			func (*DataSource) Offer(type_ string,){
		}
		
			/*
			
	Destroy the data source.
      */
			func (*DataSource) Destroy(){
		}
		
	
		
			/*
			
	Sent when a target accepts pointer_focus or motion events.  If
	a target does not accept any of the offered types, type is NULL.
      */
				func (*DataSource) Target(mimeType string,){
			}
		
			/*
			
	Request for data from another client.  Send the data as the
	specified mime-type over the passed fd, then close the fd.
      */
				func (*DataSource) Send(mimeType string,fd fd,){
			}
		
			/*
			
	This data source has been replaced by another data source.
	The client should clean up and destroy this data source.
      */
				func (*DataSource) Cancelled(){
			}
		
	