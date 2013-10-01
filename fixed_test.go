package gowl

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Second()))
}

func TestFixedFromF64(t *testing.T) {
	assert := func(d float64, val int) {
		f := FixedFromF64(d)
		if f != Fixed(val) {
			t.Fail()
			t.Logf("d:%f, val:%v [0x%x] failed.", d, val, val)
			t.Logf("FixedFromF64(%f) != %v [0x%x]", d, f, f)
		}
		if FixedToF64(f) != d {
			t.Fail()
			t.Logf("d:%f, val:%v [0x%x] failed.", d, val, val)
			t.Logf("FixedToF64(%v [0x%x] ) != %f", f, f, d)
		}
		if !t.Failed() {
			t.Logf("FixedFromF64 pass: (d:%f, f:%v [0x%x] val:%v [0x%x])", d, f, f, val, val)
		}
	}

	assert(62.125, 0x3e20)
	assert(-1200.625, -0x4b0a0)
}
func TestFixedToF64(t *testing.T) {
	assert := func(f Fixed, val float64) {
		d := FixedToF64(f)
		if d != val {
			t.Fail()
			t.Logf("f:%v [0x%x], val:%f failed.", f, f, val)
			t.Logf("FixedToF64(%f) != %f", d, val)

		}
		if FixedFromF64(d) != f {
			t.Fail()
			t.Logf("f:%v [0x%x], val:%f failed.", f, f, val)
			t.Logf("FixedFromF64(%f) != %v [0x%x]", d, f, f)
		}
		if !t.Failed() {
			t.Logf("FixedToF64 pass: (f:%v [0x%x], d:%f, val:%f)", f, f, d, val)
		}
	}
	f := Fixed(rand.Int31())
	assert(f, float64(f)/256.0)
	assert(0x012030, 288.1875)
	f = Fixed(0x70000000)
	assert(f, float64(f)/256.0)
	assert(-0x012030, -288.1875)
	f = Fixed(-0x80000000)
	assert(f, float64(f)/256.0)

}

func TestFixedFromInt(t *testing.T) {
	assert := func(i int32, val int32) {

		f := FixedFromInt(i)
		if f != Fixed(val) {
			t.Fail()
			t.Logf("i:%v [0x%x], val:%v [0x%x] failed.", i, i, val, val)
			t.Logf("FixedFromInt(%v [0x%x]) != %v [0x%x]", i, i, val, val)
		}

		if !t.Failed() {
			t.Logf("FixedFromInt pass: [ i:%v [0x%x], f:%v [0x%x],  val:%v [0x%x]]", i, i, f, f, val, val)
		}
	}
	assert(62, 62*256)
	assert(-2080, -2080*256)
}

func TestFixedToInt(t *testing.T) {
	assert := func(f Fixed, val int32) {
		i := FixedToInt(f)
		if i != val {
			t.Fail()
			t.Logf("f:%v [0x%x], val:%v [0x%x] failed.", f, f, val, val)
			t.Logf("FixedToInt(%v [0x%x]) != %v [0x%x]", i, i, val, val)
		}
		if !t.Failed() {
			t.Logf("FixedToInt pass: [ f:%v [0x%x], i:%v [0x%x], val:%v [0x%x]]", f, f, i, i, val, val)
		}
	}
	assert(0x277013, 0x2770)
	assert(-0x5044, -0x50)
}
