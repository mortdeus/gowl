package gowl

/*


*/

type Touch struct{}

func (*Touch) Down(Serial uint, Time uint, Surface object, Id int, X fixed, Y fixed) {
}

func (*Touch) Up(Serial uint, Time uint, Id int) {
}

func (*Touch) Motion(Time uint, Id int, X fixed, Y fixed) {
}

func (*Touch) Frame() {
}

func (*Touch) Cancel() {
}
