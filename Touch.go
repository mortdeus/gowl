package gowl

/*


*/

type Touch struct{}

func (*Touch) Down(serial uint, time uint, surface object, id int, x fixed, y fixed) {
}

func (*Touch) Up(serial uint, time uint, id int) {
}

func (*Touch) Motion(time uint, id int, x fixed, y fixed) {
}

func (*Touch) Frame() {
}

func (*Touch) Cancel() {
}
