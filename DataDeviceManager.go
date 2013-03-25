package gowl

/*

   The wl_data_device_manager is a a singleton global object that
   provides access to inter-client data transfer mechanisms such as
   copy and paste and drag and drop.  These mechanisms are tied to
   a wl_seat and this interface lets a client get a wl_data_device
   corresponding to a wl_seat.

*/

type DataDeviceManager struct{}

func (*DataDeviceManager) CreateDataSource(Id new_id) {
}

func (*DataDeviceManager) GetDataDevice(Id new_id, Seat object) {
}
