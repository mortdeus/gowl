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

func (*DataOffer) Accept(serial uint, type_ string) {
}

func (*DataOffer) Receive(mimeType string, fd fd) {
}

func (*DataOffer) Destroy() {
}

func (*DataOffer) Offer(type_ string) {
}
