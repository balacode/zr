// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-09 18:09:23 55D388                           zr/[go_lang_test.go]
// -----------------------------------------------------------------------------

package zr

// # Tests
//   Test_gola_GoName_
//   Test_gola_GoString_
//
// # Test Support
//   golaTestStruct struct
//   (ob golaTestStruct) GoString() string

import (
	// "fmt"
	"math"
	"testing"
)

/*
to test all items in go_lang.go use:
	go test --run Test_gola_

to generate a test coverage report use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

// -----------------------------------------------------------------------------
// # Tests

// go test --run Test_gola_GoName_
func Test_gola_GoName_(t *testing.T) {
	TBegin(t)
	// GoName_(s string) string
	//
	//X 	ret := Trim(s, SPACES)
	//X 	if len(ret) == 0 {
	//X 		return ""
	//X 	}
	//X 	if Contains(ret, "_") {
	//X 		ret = Replace(ret, "_", " ", -1)
	//X 	}
	//X 	ret = TitleCase(ret)
	//X 	if ContainsI(ret, "id") {
	//X 		ret = ReplaceWord(ret, "id", "ID", IgnoreCase)
	//X 	}
	//X 	if Contains(ret, " ") {
	//X 		ret = Replace(ret, " ", "", -1)
	//X 	}
	//X 	return ret
} //                                                           Test_gola_GoName_

// TODO: Test GoStringerEx interface.

// go test --run Test_gola_GoString_
func Test_gola_GoString_(t *testing.T) {
	TBegin(t)
	// GoString(val interface{}, optIndentAt ...int) string
	//
	// nil
	TEqual(t, GoString(nil), "nil")
	//
	// bool
	falseBool := false
	trueBool := true
	TEqual(t, GoString(&falseBool), "false")
	TEqual(t, GoString(&trueBool), "true")
	TEqual(t, GoString(falseBool), "false")
	TEqual(t, GoString(trueBool), "true")
	//
	// int
	TEqual(t, GoString(int(0)), "0")
	maxI := int(math.MaxInt32)
	minI := int(math.MinInt32)
	TEqual(t, GoString(&maxI), "2147483647")
	TEqual(t, GoString(&minI), "-2147483648")
	TEqual(t, GoString(maxI), "2147483647")
	TEqual(t, GoString(minI), "-2147483648")
	//
	// int8
	maxI8 := int8(math.MaxInt8)
	minI8 := int8(math.MinInt8)
	TEqual(t, GoString(&maxI8), "127")
	TEqual(t, GoString(&minI8), "-128")
	TEqual(t, GoString(maxI8), "127")
	TEqual(t, GoString(minI8), "-128")
	//
	// int16
	maxI16 := int16(math.MaxInt16)
	minI16 := int16(math.MinInt16)
	TEqual(t, GoString(&maxI16), "32767")
	TEqual(t, GoString(&minI16), "-32768")
	TEqual(t, GoString(maxI16), "32767")
	TEqual(t, GoString(minI16), "-32768")
	//
	// int32
	maxI32 := int32(math.MaxInt32)
	minI32 := int32(math.MinInt32)
	TEqual(t, GoString(&maxI32), "2147483647")
	TEqual(t, GoString(&minI32), "-2147483648")
	TEqual(t, GoString(maxI32), "2147483647")
	TEqual(t, GoString(minI32), "-2147483648")
	//
	// int64
	maxI64 := int64(math.MaxInt64)
	minI64 := int64(math.MinInt64)
	TEqual(t, GoString(&maxI64), "9223372036854775807")
	TEqual(t, GoString(&minI64), "-9223372036854775808")
	TEqual(t, GoString(maxI64), "9223372036854775807")
	TEqual(t, GoString(minI64), "-9223372036854775808")
	//
	// uint
	maxU := uint(math.MaxUint32)
	minU := uint(0)
	TEqual(t, GoString(&maxU), "4294967295")
	TEqual(t, GoString(&minU), "0")
	TEqual(t, GoString(maxU), "4294967295")
	TEqual(t, GoString(minU), "0")
	//
	// uint8
	maxU8 := uint8(math.MaxUint8)
	minU8 := uint8(0)
	TEqual(t, GoString(&maxU8), "255")
	TEqual(t, GoString(&minU8), "0")
	TEqual(t, GoString(maxU8), "255")
	TEqual(t, GoString(minU8), "0")
	//
	// uint16
	maxU16 := uint16(math.MaxUint16)
	minU16 := uint16(0)
	TEqual(t, GoString(&maxU16), "65535")
	TEqual(t, GoString(&minU16), "0")
	TEqual(t, GoString(maxU16), "65535")
	TEqual(t, GoString(minU16), "0")
	//
	// uint32
	maxU32 := uint32(math.MaxUint32)
	minU32 := uint32(0)
	TEqual(t, GoString(&maxU32), "4294967295")
	TEqual(t, GoString(&minU32), "0")
	TEqual(t, GoString(maxU32), "4294967295")
	TEqual(t, GoString(minU32), "0")
	//
	// uint64
	maxU64 := uint64(math.MaxUint64)
	minU64 := uint64(0)
	TEqual(t, GoString(&maxU64), "18446744073709551615")
	TEqual(t, GoString(&minU64), "0")
	TEqual(t, GoString(maxU64), "18446744073709551615")
	TEqual(t, GoString(minU64), "0")
	//
	// TODO: uintptr
	//
	// float32
	float32a := float32(1.234567)
	float32b := float32(1.2345678)
	TEqual(t, GoString(&float32a), `1.234567`)
	TEqual(t, GoString(&float32b), `1.234568`)
	TEqual(t, GoString(float32a), `1.234567`)
	TEqual(t, GoString(float32b), `1.234568`)
	//
	// float64
	float64v := CurrencyOf(12345.6789)
	TEqual(t, GoString(&float64v), `zr.CurrencyOf(12345.6789)`)
	TEqual(t, GoString(float64v), `zr.CurrencyOf(12345.6789)`)
	//
	// TODO: handle Complex64
	// TODO: handle Complex128
	// TODO: handle Array
	// TODO: handle Chan
	// TODO: handle Func
	// TODO: handle Interface
	//
	// map
	{
		m := map[string]int{
			"a":   1,
			"bb":  22,
			"ccc": 333,
		}
		s := "" +
			`map[string]int{` + "\n" +
			`    "a": 1,` + "\n" +
			`    "bb": 22,` + "\n" +
			`    "ccc": 333,` + "\n" +
			`}`
		TEqual(t, GoString(m), s)
	}
	// TODO: handle Ptr
	//
	// slice
	{
		TEqual(t, GoString([]string{"XX", "YYY", "ZZZZ"}),
			`[]string{"XX", "YYY", "ZZZZ"}`,
		)
		TEqual(t, GoString([]string{"XX", "YYY", "ZZZZ"}, 4),
			`[]string{"XX", "YYY", "ZZZZ"}`,
		)
		TEqual(t, GoString([]uint8{11, 22, 33}), `[]uint8{11, 22, 33}`)
		//
		TEqual(t,
			GoString([]Currency{
				CurrencyOf(1), CurrencyOf(22), CurrencyOf(333),
			}), `[]zr.Currency{`+
				`zr.CurrencyOf(1), zr.CurrencyOf(22), zr.CurrencyOf(333)`+
				`}`)
	}
	// string
	{
		s := "XYZ"
		TEqual(t, GoString(&s), `"XYZ"`)
		TEqual(t, GoString(s), `"XYZ"`)
	}
	// struct
	{
		type IntType int
		type StructType struct {
			StrVal string
			IntVal IntType
		}
		itm := StructType{StrVal: "XYZ", IntVal: 123}
		TEqual(t, GoString(itm), `zr.StructType{StrVal: "XYZ", IntVal: 123}`)
	}
	{
		// once a struct implements fmt.GoStringer, GoString() should
		// use that method instead of reading the object itself
		itm := golaTestStruct{strField: "ABC", intField: 123}
		TEqual(t, GoString(itm), `golaTestStruct{<custom>}`)
	}
	// TODO: handle UnsafePointer
	//
	// TODO: test types that can't be converted
	// TBeginError()
	// TEqual(t, GoString(o), "")
	// TCheckError(t, "not handled")
	//
	// TODO: Test_gola_GoString_: test with different indent values
	//
	// TODO: Test_gola_GoString_: test with GoStringerEx
	// TODO: Test_gola_GoString_: test with fmt.GoStringer
	// TODO: Test_gola_GoString_: test with fmt.Stringer
	//
	{
		// TODO: syntactically correct, but doesn't split into multiple lines
		ar := [][]string{
			{"rst", "uvw", "xyz"},
			{"123", "456", "789"},
		}
		const expect = "" +
			`[][]string{` + "\n" +
			`    []string{"rst", "uvw", "xyz"},` + "\n" +
			`    []string{"123", "456", "789"},` + "\n" +
			`}`
		TEqual(t, GoString(ar), expect)
	}
	{
		curAr := [][]Currency{
			{CurrencyOf(321), CurrencyOf(654)},
			{CurrencyOf(123), CurrencyOf(456)},
		}
		const expect = "" +
			`[][]zr.Currency{` + "\n" +
			`    []zr.Currency{zr.CurrencyOf(321), zr.CurrencyOf(654)},` + "\n" +
			`    []zr.Currency{zr.CurrencyOf(123), zr.CurrencyOf(456)},` + "\n" +
			`}`
		TEqual(t, GoString(curAr), expect)
	}
} //                                                         Test_gola_GoString_

// -----------------------------------------------------------------------------
// # Test Support

// golaTestStruct is used to test GoString() with a private
// struct which implements the fmt.GoStringer interface
type golaTestStruct struct {
	strField string
	intField int
} //                                                              golaTestStruct

// GoString implements fmt.GoStringer interface
// note that the receiver should not be a pointer
func (ob golaTestStruct) GoString() string {
	return "golaTestStruct{<custom>}"
} //                                                                    GoString

//eof
