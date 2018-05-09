// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 BDEB9E                              [zr/bool_test.go]
// -----------------------------------------------------------------------------

package zr

//   Test_bool_Bool_
//   Test_bool_TrueCount_

import (
	"math"
	"reflect"
	"testing"
)

/*
to test all items in bool.go use:
    go test --run Test_bool_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

// go test --run Test_bool_Bool_
func Test_bool_Bool_(t *testing.T) {
	TBegin(t)
	// Bool(val interface{}) bool
	//
	var test = func(val interface{}, expect bool) {
		var got = Bool(val)
		if got != expect {
			TFailf(t, "Bool(%v) [Type:%v] returned %v instead of %v",
				val, reflect.TypeOf(val), got, expect)
		}
	}
	{ // nil:
		test(nil, false)
	}
	{ // boolean:
		test(false, false)
		test(true, true)
		//
		var t = true
		test(&t, true)
		//
		var f bool
		test(&f, false)
	}
	{ // strings:
		test("true", true)
		test("TRUE", true)
		test("True", true)
		//
		test(" true ", true)
		test(" TRUE ", true)
		test(" True ", true)
		//
		test("false", false)
		test("FALSE", false)
		test("False", false)
		//
		test(" false ", false)
		test(" FALSE ", false)
		test(" False ", false)
		//
		var t = "true"
		test(&t, true)
		//
		var f = "false"
		test(&f, false)
	}
	{ // Stringer
		var ts TStringer
		//
		ts.Set("true")
		test(ts, true)
		//
		ts.Set("false")
		test(ts, false)
	}
	{ // signed integers:
		test(int(0), false)
		test(int(1), true)
		test(int(-1), true)
		//
		test(int8(0), false)
		test(int8(1), true)
		test(int8(-1), true)
		//
		test(int16(0), false)
		test(int16(1), true)
		test(int16(-1), true)
		//
		test(int32(0), false)
		test(int32(1), true)
		test(int32(-1), true)
		//
		test(int64(0), false)
		test(int64(1), true)
		test(int64(-1), true)
		//
		var fi = int(0)
		test(&fi, false)
		//
		var fi8 = int8(0)
		test(&fi8, false)
		//
		var fi16 = int16(0)
		test(&fi16, false)
		//
		var fi32 = int32(0)
		test(&fi32, false)
		//
		var fi64 = int64(0)
		test(&fi64, false)
		//
		var ti = int(1)
		test(&ti, true)
		//
		var ti8 = int8(1)
		test(&ti8, true)
		//
		var ti16 = int16(1)
		test(&ti16, true)
		//
		var ti32 = int32(1)
		test(&ti32, true)
		//
		var ti64 = int64(1)
		test(&ti64, true)
	}
	{ // unsigned integers:
		test(uint(0), false)
		test(uint(1), true)
		test(uint(0xFFFFFFFF), true)
	}
	{ // unsigned integers:
		test(uint8(0), false)
		test(uint8(1), true)
		test(uint8(0xFF), true)
		//
		test(uint16(0), false)
		test(uint16(1), true)
		test(uint16(0xFFFF), true)
		//
		test(uint32(0), false)
		test(uint32(1), true)
		test(uint32(0xFFFFFFFF), true)
		//
		test(uint64(0), false)
		test(uint64(1), true)
		test(uint64(0xFFFFFFFFFFFFFFFF), true)
		//
		var fu = uint(0)
		test(&fu, false)
		//
		var fu8 = uint8(0)
		test(&fu8, false)
		//
		var fu16 = uint16(0)
		test(&fu16, false)
		//
		var fu32 = uint32(0)
		test(&fu32, false)
		//
		var fu64 = uint64(0)
		test(&fu64, false)
		//
		var tu = uint(1)
		test(&tu, true)
		//
		var tu8 = uint8(1)
		test(&tu8, true)
		//
		var tu16 = uint16(1)
		test(&tu16, true)
		//
		var tu32 = uint32(1)
		test(&tu32, true)
		//
		var tu64 = uint64(1)
		test(&tu64, true)
	}
	{ // floating-point numbers:
		test(float32(-1), true)
		test(float32(-0.0001), true)
		test(float32(0), false)
		test(float32(0.0001), true)
		test(float32(1), true)
		test(float32(math.MaxFloat32), true)
		test(float64(-1), true)
		test(float64(-0.0001), true)
		test(float64(0), false)
		test(float64(0.0001), true)
		test(float64(1), true)
		test(float64(math.MaxFloat64), true)
		//
		var false32 = float32(0)
		test(&false32, false)
		//
		var false64 = float64(0)
		test(&false64, false)
		//
		var true32 = float32(1)
		test(&true32, true)
		//
		var true64 = float64(1)
		test(&true64, true)
	}
	{ // error conditions
		type T struct{ val bool }
		var bad T
		bad.val = true
		TBeginError()
		test(bad, false)
		TCheckError(t, "Can not convert")
	}
} //                                                             Test_bool_Bool_

// go test --run Test_bool_IsBool_
func Test_bool_IsBool_(t *testing.T) {
	TBegin(t)
	// IsBool(val interface{}) bool
	//
	var test = func(val interface{}, expect bool) {
		var got = IsBool(val)
		if got != expect {
			TFailf(t, "Bool(%v) [Type:%v] returned %v instead of %v",
				val, reflect.TypeOf(val), got, expect)
		}
	}
	{ // nil:
		test(nil, false)
	}
	{ // boolean:
		test(true, true)
		test(false, true)
		//
		var t = true
		test(&t, true)
		//
		var f bool
		test(&f, true)
	}
	{ // strings:
		test(" false ", true)
		test(" False ", true)
		test(" FALSE ", true)
		test(" True ", true)
		test(" TRUE ", true)
		test("false", true)
		test("False", true)
		test("FALSE", true)
		test("true", true)
		test("True", true)
		test("TRUE", true)
		test(" true ", true)
		//
		var f = "false"
		var t = "true"
		test(&f, true)
		test(&t, true)
		//
		test("", false)
		test(" ", false)
		test("XYZ", false)
		test("0", true)
		test("1", true)
	}
	{ // Stringer
		var ts TStringer
		//
		ts.Set("true")
		test(ts, true)
		//
		ts.Set("false")
		test(ts, true)
		//
		ts.Set("xyz")
		test(ts, false)
	}
	{ // signed integers:
		test(int(0), true)
		test(int(1), true)
		test(int(-1), true)
		//
		test(int8(0), true)
		test(int8(1), true)
		test(int8(-1), true)
		//
		test(int16(0), true)
		test(int16(1), true)
		test(int16(-1), true)
		//
		test(int32(0), true)
		test(int32(1), true)
		test(int32(-1), true)
		//
		test(int64(0), true)
		test(int64(1), true)
		test(int64(-1), true)
		//
		var fi = int(0)
		test(&fi, true)
		//
		var fi8 = int8(0)
		test(&fi8, true)
		//
		var fi16 = int16(0)
		test(&fi16, true)
		//
		var fi32 = int32(0)
		test(&fi32, true)
		//
		var fi64 = int64(0)
		test(&fi64, true)
		//
		var ti = int(1)
		test(&ti, true)
		//
		var ti8 = int8(1)
		test(&ti8, true)
		//
		var ti16 = int16(1)
		test(&ti16, true)
		//
		var ti32 = int32(1)
		test(&ti32, true)
		//
		var ti64 = int64(1)
		test(&ti64, true)
	}
	{ // unsigned integers:
		test(uint(0), true)
		test(uint(1), true)
		test(uint(0xFFFFFFFF), true)
		//
		test(uint8(0), true)
		test(uint8(1), true)
		test(uint8(0xFF), true)
		//
		test(uint16(0), true)
		test(uint16(1), true)
		test(uint16(0xFFFF), true)
		//
		test(uint32(0), true)
		test(uint32(1), true)
		test(uint32(0xFFFFFFFF), true)
		//
		test(uint64(0), true)
		test(uint64(1), true)
		test(uint64(0xFFFFFFFFFFFFFFFF), true)
		//
		var fu = uint(0)
		test(&fu, true)
		//
		var fu8 = uint8(0)
		test(&fu8, true)
		//
		var fu16 = uint16(0)
		test(&fu16, true)
		//
		var fu32 = uint32(0)
		test(&fu32, true)
		//
		var fu64 = uint64(0)
		test(&fu64, true)
		//
		var tu = uint(1)
		test(&tu, true)
		//
		var tu8 = uint8(1)
		test(&tu8, true)
		//
		var tu16 = uint16(1)
		test(&tu16, true)
		//
		var tu32 = uint32(1)
		test(&tu32, true)
		//
		var tu64 = uint64(1)
		test(&tu64, true)
	}
	{ // floating-point numbers:
		var float32ptr = func(n float32) *float32 { return &n }
		var float64ptr = func(n float64) *float64 { return &n }
		//
		test(float32(0), true)
		test(float32(math.MaxFloat32), true)
		test(float32(-math.MaxFloat32), true)
		//
		test(float64(0), true)
		test(float64(math.MaxFloat64), true)
		test(float64(-math.MaxFloat64), true)
		//
		test(float32ptr(0), true)
		test(float32ptr(-1e32), true)
		test(float32ptr(1e32), true)
		//
		test(float64ptr(0), true)
		test(float64ptr(math.MaxFloat64), true)
		test(float64ptr(-math.MaxFloat64), true)
	}
	{
		type T struct{ val bool }
		var bad T
		bad.val = true
		test(bad, false)
		test(&bad, false)
	}
} //                                                           Test_bool_IsBool_

// go test --run Test_bool_TrueCount_
func Test_bool_TrueCount_(t *testing.T) {
	TBegin(t)
	// TrueCount(values ...bool) int
	//
	var test = func(expect int, values ...bool) {
		var got = TrueCount(values...)
		if got != expect {
			t.Errorf("TrueCount(%v) returned %v", values, got)
		}
	}
	test(0)
	test(0, false)
	test(0, false, false)
	test(0, false, false, false)
	test(1, true)
	test(1, true, false)
	test(1, false, true)
	test(1, true, false, false)
	test(1, false, true, false)
	test(1, false, false, true)
} //                                                        Test_bool_TrueCount_

//end
