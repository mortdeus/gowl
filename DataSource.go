package gowl

/*

   The wl_data_source object is the source side of a wl_data_offer.
   It is created by the source client in a data transfer and
   provides a way to describe the offered data and a way to respond
   to requests to transfer the data.

*/

type DataSource struct{}

func (*DataSource) Offer(Type string) {
}

func (*DataSource) Destroy() {
}

func (*DataSource) Target(MimeType string) {
}

func (*DataSource) Send(MimeType string, Fd fd) {
}

func (*DataSource) Cancelled() {
}
