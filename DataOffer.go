package gowl

/*

   A wl_data_offer represents a piece of data offered for transfer
   by another client (the source client).  It is used by the
   copy-and-paste and drag-and-drop mechanisms.  The offer
   describes the different mime types that the data can be
   converted to and provides the mechanism for transferring the
   data directly from the source client.

*/
type DataOffer struct{}

/*

	Indicate that the client can accept the given mime-type, or
	NULL for not accepted.  Use for feedback during drag and drop.
*/
func (*DataOffer) Accept(serial uint32, type_ string) {
}

/*

	To transfer the offered data, the client issues this request
	and indicates the mime-type it wants to receive.  The transfer
	happens through the passed fd (typically a pipe(7) file
	descriptor).  The source client writes the data in the
	mime-type representation requested and then closes the fd.
	The receiving client reads from the read end of the pipe until
	EOF and the closes its end, at which point the transfer is
	complete.
*/
func (*DataOffer) Receive(mimeType string, fd Fd) {
}

func (*DataOffer) Destroy() {
}

/*

	Sent immediately after creating the wl_data_offer object.  One
	event per offered mime type.
*/
func (*DataOffer) Offer(type_ string) {
}
