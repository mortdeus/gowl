package gowl

import "testing"
import "unsafe"

func TestFixed(t *testing.T) {
	t.Logf("FixedFromF64:\t %b Size:%v", FixedFromF64(33283.005), unsafe.Sizeof(FixedFromF64(33283.255)))
	//t.Logf("FixedFromF64:\t %b", FixedFromF64(33283.3515625))
	t.Logf("FixedToF64:\t %v Size:%v", FixedToF64(FixedFromF64(33283.005)), unsafe.Sizeof(FixedToF64(FixedFromF64(33283.255))))
	t.Logf("FixedFromInt:\t %b Size: %v", FixedFromInt(3568), unsafe.Sizeof(FixedFromInt(3568)))
	t.Logf("FixedToInt:\t %b Size: %v", Fixed(FixedToInt(FixedFromInt(3568))), unsafe.Sizeof(FixedToInt(FixedFromInt(3568))))
}
