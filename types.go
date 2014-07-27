package gowl

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"unsafe"
)

type (
	Int    int32
	Uint   uint32
	Fixed  int32
	String string
	Object uint32
	NewID  uint32
	Array  []byte
	FD     uintptr
)

func (s *String) GobEncode() ([]byte, error) {
	buf := new(bytes.Buffer)

	bs, err := syscall.ByteSliceFromString(string(*s))
	if err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, uint32(len(bs))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, bs); err != nil {
		return nil, err
	}
	pad := (4 - (len(bs) % 4)) &^ 04
	if err := binary.Write(buf, binary.LittleEndian, uint32(pad)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (s1 *String) GobDecode(s []byte) error {

	var slen uint32
	if err := binary.Read(bytes.NewReader(s[:4]), binary.LittleEndian, &slen); err != nil {
		return err
	}
	var bs = make([]byte, slen)
	if err := binary.Read(bytes.NewReader(s[4:4+slen]), binary.LittleEndian, bs); err != nil {
		return err
	}
	*s1 = String(bs[:len(bs)-1])
	return nil
}

func (Fixed) ToFloat64(f Fixed) float64 {
	d := int64(1023+44)<<52 + (1 << 51) + int64(f)
	return *(*float64)(unsafe.Pointer(&d)) - float64(3<<43)
}
func (Fixed) FromFloat64(d float64) Fixed {
	f := d + float64(uint64(3<<(51-8)))
	return *(*Fixed)(unsafe.Pointer(&f))
}
func (Fixed) ToInt32(f Fixed) int32   { return int32(f) / 256 }
func (Fixed) FromInt32(i int32) Fixed { return Fixed(i) * 256 }
