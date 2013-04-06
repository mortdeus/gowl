package gowl

import "unsafe"
type message struct {
	name      string
	signature string
	types     []interface_
}

type interface_ struct {
	name    string
	version int
	methods []message
	events  []message
}

type Object struct {
	iface          *interface_
	implementation interface{}
	id             uint32
}

type list struct {
	prev *list
	next *list
}

func (l *list) init() {

}
func (l *list) insert(elm *list) {

}
func (l *list) remove(elm *list) {

}
func (l *list) length() {

}
func (l *list) empty() {

}
func (l *list) insertList(other *list) {

}

//Consider changing this to []byte
type Array struct {
	size  uintptr
	alloc uintptr
	data  []byte
}

func (a *Array) init() {

}
func (a *Array) release() {

}
func (a *Array) add(size uintptr) {

}
func (a *Array) copy(source *Array) {

}

type Fixed int32

func FixedToF64(f Fixed) float64 {
	d := int64(1023+44)<<52 +
		(1 << 51) + int64(f)
	return *(*float64)(unsafe.Pointer(&d)) - float64(3<<43)
}

func FixedFromF64(d float64) Fixed {
	f := d + float64(uint64(3<<(51-8)))
	return *(*Fixed)(unsafe.Pointer(&f))
}

func FixedToInt(f Fixed) int32 {
	return int32(f / 256)
}
func FixedFromInt(i int) Fixed {
	return Fixed(i * 256)
}
