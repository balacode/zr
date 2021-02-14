// -----------------------------------------------------------------------------
// ZR Library                                                  zr/[bool_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

//   Test_bool_Bool_
//   Test_bool_TrueCount_

import (
	"math"
	"reflect"
	"testing"
)

//  to test all items in bool.go use:
//      go test --run Test_bool_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

// go test --run Test_bool_Bool_
func Test_bool_Bool_(t *testing.T) {
	TBegin(t)
	//
	// Bool(value interface{}) bool
	//
	test := func(value interface{}, expect bool) {
		got := Bool(value)
		if got != expect {
			TFailf(t, "Bool(%v) [Type:%v] returned %v instead of %v",
				value, reflect.TypeOf(value), got, expect)
		}
	}
	{ // nil:
		test(nil, false)
	}
	{ // boolean:
		test(false, false)
		test(true, true)
		//
		t := true
		test(&t, true)
		//
		f := false
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
		t := "true"
		test(&t, true)
		//
		f := "false"
		test(&f, false)
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
		fi := int(0)
		test(&fi, false)
		//
		fi8 := int8(0)
		test(&fi8, false)
		//
		fi16 := int16(0)
		test(&fi16, false)
		//
		fi32 := int32(0)
		test(&fi32, false)
		//
		fi64 := int64(0)
		test(&fi64, false)
		//
		ti := int(1)
		test(&ti, true)
		//
		ti8 := int8(1)
		test(&ti8, true)
		//
		ti16 := int16(1)
		test(&ti16, true)
		//
		ti32 := int32(1)
		test(&ti32, true)
		//
		ti64 := int64(1)
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
		fu := uint(0)
		test(&fu, false)
		//
		fu8 := uint8(0)
		test(&fu8, false)
		//
		fu16 := uint16(0)
		test(&fu16, false)
		//
		fu32 := uint32(0)
		test(&fu32, false)
		//
		fu64 := uint64(0)
		test(&fu64, false)
		//
		tu := uint(1)
		test(&tu, true)
		//
		tu8 := uint8(1)
		test(&tu8, true)
		//
		tu16 := uint16(1)
		test(&tu16, true)
		//
		tu32 := uint32(1)
		test(&tu32, true)
		//
		tu64 := uint64(1)
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
		false32 := float32(0)
		test(&false32, false)
		//
		false64 := float64(0)
		test(&false64, false)
		//
		true32 := float32(1)
		test(&true32, true)
		//
		true64 := float64(1)
		test(&true64, true)
	}
	{ // error conditions
		type T struct{ v bool }
		var bad T
		bad.v = true
		TBeginError()
		test(bad, false)
		TCheckError(t, "Can not convert")
	}
} //                                                             Test_bool_Bool_

// go test --run Test_bool_IsBool_
func Test_bool_IsBool_(t *testing.T) {
	TBegin(t)
	//
	// IsBool(value interface{}) bool
	//
	test := func(value interface{}, expect bool) {
		got := IsBool(value)
		if got != expect {
			TFailf(t, "Bool(%v) [Type:%v] returned %v instead of %v",
				value, reflect.TypeOf(value), got, expect)
		}
	}
	{ // nil:
		test(nil, false)
	}
	{ // boolean:
		test(true, true)
		test(false, true)
		//
		t := true
		test(&t, true)
		//
		f := false
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
		f := "false"
		t := "true"
		test(&f, true)
		test(&t, true)
		//
		test("", false)
		test(" ", false)
		test("XYZ", false)
		test("0", true)
		test("1", true)
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
		fi := int(0)
		test(&fi, true)
		//
		fi8 := int8(0)
		test(&fi8, true)
		//
		fi16 := int16(0)
		test(&fi16, true)
		//
		fi32 := int32(0)
		test(&fi32, true)
		//
		fi64 := int64(0)
		test(&fi64, true)
		//
		ti := int(1)
		test(&ti, true)
		//
		ti8 := int8(1)
		test(&ti8, true)
		//
		ti16 := int16(1)
		test(&ti16, true)
		//
		ti32 := int32(1)
		test(&ti32, true)
		//
		ti64 := int64(1)
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
		fu := uint(0)
		test(&fu, true)
		//
		fu8 := uint8(0)
		test(&fu8, true)
		//
		fu16 := uint16(0)
		test(&fu16, true)
		//
		fu32 := uint32(0)
		test(&fu32, true)
		//
		fu64 := uint64(0)
		test(&fu64, true)
		//
		tu := uint(1)
		test(&tu, true)
		//
		tu8 := uint8(1)
		test(&tu8, true)
		//
		tu16 := uint16(1)
		test(&tu16, true)
		//
		tu32 := uint32(1)
		test(&tu32, true)
		//
		tu64 := uint64(1)
		test(&tu64, true)
	}
	{ // floating-point numbers:
		float32ptr := func(n float32) *float32 { return &n }
		float64ptr := func(n float64) *float64 { return &n }
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
		type T struct{ v bool }
		var bad T
		bad.v = true
		test(bad, false)
		test(&bad, false)
	}
} //                                                           Test_bool_IsBool_

// go test --run Test_bool_TrueCount_
func Test_bool_TrueCount_(t *testing.T) {
	TBegin(t)
	//
	// TrueCount(values ...bool) int
	//
	test := func(expect int, values ...bool) {
		got := TrueCount(values...)
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

// end
