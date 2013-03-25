package gowl

type DataDevice struct{}

func (*DataDevice) StartDrag(source object, origin object, icon object, serial uint) {
}

func (*DataDevice) SetSelection(source object, serial uint) {
}

func (*DataDevice) DataOffer(id new_id) {
}

func (*DataDevice) Enter(serial uint, surface object, x fixed, y fixed, id object) {
}

func (*DataDevice) Leave() {
}

func (*DataDevice) Motion(time uint, x fixed, y fixed) {
}

func (*DataDevice) Drop() {
}

func (*DataDevice) Selection(id object) {
}
