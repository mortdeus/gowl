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
	padN := (4 - (len(bs) % 4))
	if err != nil {
		return nil, err
	}
	//TODO(Mortdeus): Do we count padN offset in len?
	if err := binary.Write(buf, binary.LittleEndian, uint32(len(bs))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, bs); err != nil {
		return nil, err
	}

	pad := make([]byte, padN)
	if err := binary.Write(buf, binary.LittleEndian, pad); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (s *String) GobDecode(s1 []byte) error {

	var slen uint32
	if err := binary.Read(bytes.NewReader(s1[:4]), binary.LittleEndian, &slen); err != nil {
		return err
	}
	var bs = make([]byte, slen)
	if err := binary.Read(bytes.NewReader(s1[4:4+slen]), binary.LittleEndian, bs); err != nil {
		return err
	}
	*s = String(bs[:len(bs)-1])
	return nil
}
func (a *Array) GobEncode() ([]byte, error) {
	buf := new(bytes.Buffer)
	padN := (4 - (len(*a) % 4))
	//TODO(Mortdeus): Do we count padN offset in len?
	if err := binary.Write(buf, binary.LittleEndian, uint32(len(*a))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, *a); err != nil {
		return nil, err
	}
	pad := make([]byte, padN)
	if err := binary.Write(buf, binary.LittleEndian, pad); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (a *Array) GobDecode(a1 []byte) error {

	var slen uint32
	if err := binary.Read(bytes.NewReader(a1[:4]), binary.LittleEndian, &slen); err != nil {
		return err
	}
	var bs = make([]byte, slen)
	if err := binary.Read(bytes.NewReader(a1[4:4+slen]), binary.LittleEndian, bs); err != nil {
		return err
	}
	*a = Array(bs[:len(bs)])
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
