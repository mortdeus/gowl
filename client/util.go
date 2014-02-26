package client

import (
	"container/list"
	"unsafe"
)

type List struct {
	list.List
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

func FixedToInt(f Fixed) int {
	return int(f) / 256
}
func FixedFromInt(i int) Fixed {
	return Fixed(i) * 256
}
