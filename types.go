package gowl

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"unsafe"
)

type Int int32
type Uint uint32

type String []byte

func StringEncode(s string) (String, error) {
	b, err := syscall.ByteSliceFromString(s)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, uint32(len(b))); err != nil {
		return nil, err
	}
	pad := (4 - (len(b) % 4)) &^ 04
	
	if _, err := buf.Write(append(b, make([]byte, pad)...)); err != nil {
		return nil, err
	}

	return String(buf.Bytes()), nil
}

func StringDecode(s String) (string, error) {
	var slen uint32
	if err := binary.Read(bytes.NewBuffer(s[:4]), binary.LittleEndian, &slen); err != nil {
		return "", err
	}
	return string(s[4 : (slen+4)-1]), nil
}

type Fixed int32

func FixedToF64(f Fixed) float64 {
	d := int64(1023+44)<<52 + (1 << 51) + int64(f)
	return *(*float64)(unsafe.Pointer(&d)) - float64(3<<43)
}
func FixedFromF64(d float64) Fixed {
	f := d + float64(uint64(3<<(51-8)))
	return *(*Fixed)(unsafe.Pointer(&f))
}
func FixedToInt(f Fixed) int   { return int(f) / 256 }
func FixedFromInt(i int) Fixed { return Fixed(i) * 256 }
