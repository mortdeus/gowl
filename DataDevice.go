package gowl

type DataDevice struct{}

func (*DataDevice) StartDrag(Source object, Origin object, Icon object, Serial uint) {
}

func (*DataDevice) SetSelection(Source object, Serial uint) {
}

func (*DataDevice) DataOffer(Id new_id) {
}

func (*DataDevice) Enter(Serial uint, Surface object, X fixed, Y fixed, Id object) {
}

func (*DataDevice) Leave() {
}

func (*DataDevice) Motion(Time uint, X fixed, Y fixed) {
}

func (*DataDevice) Drop() {
}

func (*DataDevice) Selection(Id object) {
}
