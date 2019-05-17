// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-17 11:21:57 B13156                          zr/[currency_test.go]
// -----------------------------------------------------------------------------

package zr

// # Currency Factory:
//   Test_crcy_CurrencyOf_
//
// # String Output:
//   Test_crcy_Currency_Fmt_
//   Test_crcy_Currency_InWordsEN_
//   Test_crcy_Currency_String_
//
// # Division:
//   Test_crcy_Currency_Div_
//   Test_crcy_Currency_DivFloat_
//   Test_crcy_Currency_DivInt_
//
// # Multiplication:
//   Test_crcy_Currency_Mul_
//   Test_crcy_Currency_MulFloat_
//   Test_crcy_Currency_MulInt_
//
// # Addition:
//   Test_crcy_Currency_Add_
//   Test_crcy_Currency_AddFloat_
//   Test_crcy_Currency_AddInt_
//
// # Subtraction:
//   Test_crcy_Currency_Sub_
//   Test_crcy_Currency_SubFloat_
//   Test_crcy_Currency_SubInt_
//
// # Information:
//   Test_crcy_Currency_Float64_
//   Test_crcy_Currency_Int_
//   Test_crcy_Currency_Int64_
//   Test_crcy_Currency_IsEqual_
//   Test_crcy_Currency_IsNegative_
//   Test_crcy_Currency_IsZero_
//   Test_crcy_Currency_Overflow_
//
// # JSON:
//   Test_crcy_Currency_MarshalJSON_
//   Test_crcy_Currency_UnmarshalJSON_
//
// # Helper Functions
//   Test_crcy_currencyOverflow_
//
// # Test Helper Functions
//   arC(ar ...interface{}) (ret []Currency)
//   arF(ar ...interface{}) (ret []float64)
//   arI(ar ...interface{}) (ret []int)
//   cur = CurrencyOf
//   curFloatOpTest(
//       t *testing.T,
//       opName string,
//       opFunc func(values ...float64) Currency,
//       ob Currency,
//       values []float64,
//       expect Currency,
//       expectErrors int,
//   )
//   curIntOpTest(
//       t *testing.T,
//       opName string,
//       opFunc func(values ...int) Currency,
//       ob Currency,
//       values []int,
//       expect Currency,
//       expectErrors int,
//   )
//   curOpTest(
//       t *testing.T,
//       opName string,
//       opFunc func(values ...Currency) Currency,
//       ob Currency,
//       values []Currency,
//       expect Currency,
//       expectErrors int,
//   )

//  to test all items in currency.go use:
//  	go test --run Test_crcy_
//
//  to generate a test coverage report for the whole module use:
//  	go test -coverprofile cover.out
//  	go tool cover -html=cover.out

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"testing"
)

// -----------------------------------------------------------------------------
// # Currency Factory:

// go test --run Test_crcy_CurrencyOf_
func Test_crcy_CurrencyOf_(t *testing.T) {
	TBegin(t)
	//
	// CurrencyOf(value interface{}) Currency
	//
	const cur4d = 10000 // reserved 4 decimal places
	test := func(input interface{}, expect Currency) {
		got := CurrencyOf(input)
		if got.i64 != expect.i64 {
			TFailf(t, `CurrencyOf(%v) returned %v instead of %v`,
				input, got, expect)
		}
	}
	// highest/lowest possible currency value
	test(int64(922337203685476), Currency{922337203685476 * cur4d})
	test(int64(-922337203685476), Currency{-922337203685476 * cur4d})
	//
	// strings
	test("", Currency{0})
	test(" ", Currency{0})
	test("1.01", Currency{1.01 * cur4d})
	test("922337203685476", Currency{922337203685476 * cur4d})
	test("-922337203685476", Currency{-922337203685476 * cur4d})
	//
	// Stringer interface
	{
		var ts TStringer
		ts.Set("922337203685476")
		stringer := fmt.Stringer(ts)
		test(stringer, Currency{922337203685476 * cur4d})
	}
	{
		var ts TStringer
		ts.Set("-922337203685476")
		test(ts, Currency{-922337203685476 * cur4d})
	}
	// integers
	test(int(-123456), Currency{-(123456 * cur4d)})
	test(int8(math.MaxInt8), Currency{127 * cur4d})
	test(int8(math.MinInt8), Currency{-128 * cur4d})
	test(int16(math.MaxInt16), Currency{32767 * cur4d})
	test(int16(math.MinInt16), Currency{-32768 * cur4d})
	test(int32(math.MaxInt32), Currency{2147483647 * cur4d})
	test(int32(math.MinInt32), Currency{-2147483648 * cur4d})
	// TODO: test with acceptable limit for int64
	test(int64(math.MaxInt32), Currency{2147483647 * cur4d})
	test(int64(math.MinInt32), Currency{-2147483648 * cur4d})
	//
	// unsigned integers
	test(uint(123456), Currency{123456 * cur4d})
	test(uint8(0), Currency{0})
	test(uint8(math.MaxUint8), Currency{255 * cur4d})
	test(uint16(0), Currency{0})
	test(uint16(math.MaxUint16), Currency{65535 * cur4d})
	test(uint32(0), Currency{0})
	test(uint32(math.MaxUint32), Currency{4294967295 * cur4d})
	test(uint64(0), Currency{0})
	test(uint64(922337203685476), Currency{922337203685476 * cur4d})
	//
	// floating-point numbers:
	//
	test(float32(0), Currency{0})
	test(float64(0), Currency{0})
	//
	test(float32(0.1), Currency{1000})
	//
	test(float32(-0.1), Currency{-1000})
	//
	test(float64(0.00001), Currency{0})
	test(float64(0.0001), Currency{1})
	test(float64(0.001), Currency{10})
	test(float64(0.01), Currency{100})
	test(float64(0.1), Currency{1000})
	//
	test(float64(-0.00001), Currency{0})
	test(float64(-0.0001), Currency{-1})
	test(float64(-0.001), Currency{-10})
	test(float64(-0.01), Currency{-100})
	test(float64(-0.1), Currency{-1000})
	//
	test(float32(1000), Currency{1000 * cur4d})
	test(float32(1234.5), Currency{1234.5 * cur4d})
	test(float64(1000), Currency{1000 * cur4d})
	test(float64(12345.6789), Currency{12345.6789 * cur4d})
	//
	// integer pointers
	{
		n := int(-123456)
		test(&n, Currency{-(123456 * cur4d)})
	}
	{
		n := int8(math.MaxInt8)
		test(&n, Currency{127 * cur4d})
	}
	{
		n := int8(math.MinInt8)
		test(&n, Currency{-128 * cur4d})
	}
	{
		n := int16(math.MaxInt16)
		test(&n, Currency{32767 * cur4d})
	}
	{
		n := int16(math.MinInt16)
		test(&n, Currency{-32768 * cur4d})
	}
	{
		n := int32(math.MaxInt32)
		test(&n, Currency{2147483647 * cur4d})
	}
	{
		n := int32(math.MinInt32)
		test(&n, Currency{-2147483648 * cur4d})
	}
	{
		n := int64(922337203685476)
		test(&n, Currency{922337203685476 * cur4d})
	}
	{
		n := int64(-922337203685476)
		test(&n, Currency{-922337203685476 * cur4d})
	}
	// unsigned integer pointers
	{
		n := uint(123456)
		test(&n, Currency{123456 * cur4d})
	}
	{
		n := uint8(math.MaxUint8)
		test(&n, Currency{255 * cur4d})
	}
	{
		n := uint8(0)
		test(&n, Currency{0})
	}
	{
		n := uint16(math.MaxUint16)
		test(&n, Currency{65535 * cur4d})
	}
	{
		n := uint16(0)
		test(&n, Currency{0})
	}
	{
		n := uint32(math.MaxUint32)
		test(&n, Currency{4294967295 * cur4d})
	}
	{
		n := uint32(0)
		test(&n, Currency{0})
	}
	// pointers to floating-point numbers
	{
		n := float32(1000)
		test(&n, Currency{1000 * cur4d})
	}
	{
		n := float32(1234.5)
		test(&n, Currency{1234.5 * cur4d})
	}
	{
		n := float64(1000)
		test(&n, Currency{1000 * cur4d})
	}
	{
		n := float64(12345.6789)
		test(&n, Currency{12345.6789 * cur4d})
	}
	// pointer to string
	{
		s := "922337203685476"
		test(&s, Currency{922337203685476 * cur4d})
	}
	{
		s := "-922337203685476"
		test(&s, Currency{-922337203685476 * cur4d})
	}
	// test conditions that generate errors:
	DisableErrors()
	// non-numeric string
	{
		ec1 := GetErrorCount()
		test("abc", Currency{})
		ec2 := GetErrorCount()
		if ec2 != ec1+1 {
			TFail(t, `Expected 1 error, but got `, ec2-ec1)
		}
	}
	// wrong type
	{
		ec1 := GetErrorCount()
		test(true, Currency{})
		ec2 := GetErrorCount()
		if ec2 != ec1+1 {
			TFail(t, `Expected 1 error, but got `, ec2-ec1)
		}
	}
	// overflow
	{
		ec1 := GetErrorCount()
		//
		// large floating-point numbers
		test(float32(-10e20), Currency{math.MinInt64})
		test(float32(-10e20), Currency{math.MinInt64})
		test(float64(10e20), Currency{math.MaxInt64})
		test(float64(10e20), Currency{math.MaxInt64})
		//
		// max int64
		test(int64(math.MaxInt64), Currency{math.MaxInt64})
		{
			n := int64(math.MaxInt64)
			test(&n, Currency{math.MaxInt64})
		}
		// min int64
		test(int64(math.MinInt64), Currency{math.MinInt64})
		{
			n := int64(math.MinInt64)
			test(&n, Currency{math.MinInt64})
		}
		// max uint64
		test(uint64(math.MaxUint64), Currency{math.MaxInt64})
		{
			n := uint64(math.MaxUint64)
			test(&n, Currency{math.MaxInt64})
		}
		ec2 := GetErrorCount()
		if ec2 != ec1+10 {
			TFail(t, `Expected 10 errors, but got `, ec2-ec1)
		}
	}
	EnableErrors()
} //                                                       Test_crcy_CurrencyOf_

// -----------------------------------------------------------------------------
// # String Output:

// go test --run Test_crcy_Currency_Fmt_
func Test_crcy_Currency_Fmt_(t *testing.T) {
	TBegin(t)
	// (ob Currency) Fmt(decimalPlaces int) string
	//
	test := func(input interface{}, decimalPlaces int, expect string) {
		n := cur(input)
		got := n.Fmt(decimalPlaces)
		if got != expect {
			TFailf(t, `Currency(%s).Fmt(%v) returned %q instead of %q`,
				n, decimalPlaces, got, expect)
		}
	}
	// variable decimal places
	test("1234.00000", -1, "1,234")
	test("1234.50000", -1, "1,234.5")
	test("1234.56000", -1, "1,234.56")
	test("1234.56700", -1, "1,234.567")
	test("1234.56780", -1, "1,234.5678")
	test("1234.56789", -1, "1,234.5678")
	//
	// more decimals than Currency's precision
	test("1234.56789", 4, "1,234.5678")
	test("1234.56780", 4, "1,234.5678")
	test("1234.56", 2, "1,234.56")
	test("1234.00", 2, "1,234.00")
	test("1234.0000", 2, "1,234.00")
	test("1234.5678", 0, "1,234")
	test("1234.5678", 1, "1,234.5")
	test("1234.5678", 2, "1,234.56")
	test("1234.5678", 3, "1,234.567")
	test("1234.5678", 4, "1,234.5678")
	test("1234.56789", 5, "1,234.56780")
	//
	// 0
	test("", 0, "0")
	//
	// 1
	test(-1, 0, "-1")
	test(-9, 0, "-9")
	//
	// 2
	test(-10, 0, "-10")
	test(-12, 0, "-12")
	test(-99, 0, "-99")
	//
	// 3
	test(-100, 0, "-100")
	test(-201, 0, "-201")
	test(-999, 0, "-999")
	test(-123, 0, "-123")
	//
	// 4
	test(-1000, 0, "-1,000")
	test(-1001, 0, "-1,001")
	test(-9999, 0, "-9,999")
	test(-1234, 0, "-1,234")
	//
	// 1
	test(1, 0, "1")
	test(9, 0, "9")
	//
	// 2
	test(10, 0, "10")
	test(12, 0, "12")
	test(99, 0, "99")
	//
	// 3
	test(100, 0, "100")
	test(201, 0, "201")
	test(999, 0, "999")
	test(123, 0, "123")
	//
	// 4
	test(1000, 0, "1,000")
	test(1001, 0, "1,001")
	test(9999, 0, "9,999")
	test(1234, 0, "1,234")
	//
	// large numbers
	test("22337203685477", 0, "22,337,203,685,477")
	test("100000000000000", 0, "100,000,000,000,000")
	test("900000000000000", 0, "900,000,000,000,000")
	test("99999999999999", 0, "99,999,999,999,999")
	test("922337203685477", 0, "922,337,203,685,477")
	test("-922337203685477", 0, "-922,337,203,685,477")
} //                                                     Test_crcy_Currency_Fmt_

// go test --run Test_crcy_Currency_InWordsEN_
func Test_crcy_Currency_InWordsEN_(t *testing.T) {
	// (ob Currency) InWordsEN(fmt string) string
	test := func(ob Currency, fmt, expect string) {
		got := ob.InWordsEN(fmt)
		if got != expect {
			TFail(t,
				`Currency(`+ob.String()+`).Fmt("`+fmt+`")`+
					` returned "`+got+`". must be "`+expect+`"`,
			)
		}
	}
	test(cur(int64(9000000000000)), "Unit", "Nine Trillion Units")
	// TODO: test(int64(math.MaxInt64), "Unit", "")
	// TODO: FAIL test(0, "Unit", "Zero")
	// TODO: FAIL test(1, "", "One")
	// TODO: ADD TEST FOR UNIT NAMES
	// TODO: ADD TEST FOR GROUP NAMES
	//
	test(cur(1), "Unit", "One Unit")
	test(cur(2), "Unit", "Two Units")
	//
	test(cur(11.02), ";;Cent;Only", "Two Cents Only")
	test(cur(11.02), "Dollar;;Cent", "Eleven Dollars and Two Cents")
	test(cur(11.02), "Euro", "Eleven Euros")
	test(cur(11.02), "Pound;;;Pence", "Eleven Pounds and Two Pence")
	//
	test(cur(-987), "Unit", "Nine Hundred and Eighty Seven Units")
	//
	test(
		cur(123.456),
		";Units;;Units",
		"One Hundred and Twenty Three Units and Forty Five Units",
	)
	test(
		cur(123.456),
		"Large-Single;Large-Plural;Small-Single;Small-Plural",
		"One Hundred and Twenty Three Large-Plural"+
			" and Forty Five Small-Plural",
	)
	test(
		cur(1.10),
		"Large-Single;Large-Plural;Small-Single;Small-Plural",
		"One Large-Single and Ten Small-Plural",
	)
	test(
		cur(1.01),
		"Large-Single;Large-Plural;Small-Single;Small-Plural",
		"One Large-Single and One Small-Single",
	)
	test(
		cur(1234567890),
		"Unit",
		"One Billion"+
			" Two Hundred and Thirty Four Million"+
			" Five Hundred and Sixty Seven Thousand"+
			" Eight Hundred and Ninety Units",
	)
	test(cur(int64(9000000000)), "Unit", "Nine Billion Units")
	//FAIL test(cur(9000000000000), "Unit", "Nine Trillion Units")
} //                                               Test_crcy_Currency_InWordsEN_

// go test --run Test_crcy_Currency_String_
func Test_crcy_Currency_String_(t *testing.T) {
	TBegin(t)
	// (ob Currency) String() string
	//
	for i, test := range []struct {
		input  Currency
		expect string
	}{
		// zero
		{Currency{0}, "0"},
		//
		// numbers with zero int part and some decimals
		{Currency{9999}, "0.9999"},
		{Currency{1000}, "0.1"},
		{Currency{100}, "0.01"},
		{Currency{10}, "0.001"},
		{Currency{1}, "0.0001"},
		{Currency{-1}, "-0.0001"},
		{Currency{-10}, "-0.001"},
		{Currency{-100}, "-0.01"},
		{Currency{-1000}, "-0.1"},
		{Currency{-9999}, "-0.9999"},
		//
		// numbers with both int part and decimals
		{Currency{-10100}, "-1.01"},
		{Currency{10100}, "1.01"},
		{Currency{-123456789}, "-12345.6789"},
		{Currency{123456789}, "12345.6789"},
		{Currency{-2000000000000000001}, "-200000000000000.0001"},
		{Currency{2000000000000000001}, "200000000000000.0001"},
		//
		// smallest and largest currency values:
		{Currency{-9223372036854769999}, "-922337203685476.9999"},
		{Currency{9223372036854769999}, "922337203685476.9999"},
		//
		// these are overflown values, but String() must print then anyway
		{Currency{math.MinInt64}, "-922337203685477.5808"},
		{Currency{math.MaxInt64}, "922337203685477.5807"},
	} {
		var (
			input = test.input
			init  = input
			got   = input.String()
		)
		if got != test.expect {
			TFail(t,
				`#`, i, ` (`, input, `).String()`,
				` returned "`, got, `". must be "`, test.expect, `"`,
			)
		}
		if input.i64 != init.i64 {
			TFail(t, ` mutated from `, init, ` to `, input)
		}
	}
} //                                                  Test_crcy_Currency_String_

// -----------------------------------------------------------------------------
// # Division:

// go test --run Test_crcy_Currency_Div_
func Test_crcy_Currency_Div_(t *testing.T) {
	TBegin(t)
	//
	// (ob Currency) Div(nums... Currency) Currency
	//
	test := func(ob Currency, nums []Currency, expect Currency) {
		init := ob
		for _, n := range nums {
			old := ob
			got := ob.Div(n)
			// object of invoked method must not change
			if ob.i64 != old.i64 {
				TFail(t, `(`, init, `) mutated from `, old, ` to `, ob)
			}
			ob = got
		}
		if ob.i64 != expect.i64 {
			TFail(t,
				`(`, init, `).Div(`, nums, `)`,
				` returned `, ob, `. must be `, expect,
			)
		}
	}
	//           12345.6789   /              1.0    =       12345.6789
	test(Currency{123456789}, arC(Currency{10000}), Currency{123456789})
} //                                                     Test_crcy_Currency_Div_

// go test --run Test_crcy_Currency_DivFloat_
func Test_crcy_Currency_DivFloat_(t *testing.T) {
	TBegin(t)
	// (ob Currency) DivFloat(nums...float64) Currency
	//
	test := func(ob Currency, nums []float64, expect Currency) {
		curFloatOpTest(t, "DivFloat", ob.DivFloat, ob, nums, expect, 0)
	}
	//           12345.6789 / 1.0  =            12345.6789
	test(Currency{123456789}, arF(1.0), Currency{123456789})
	//
	//           12345.6789 / 2.0  =            6172.8394
	test(Currency{123456789}, arF(2.0), Currency{61728394})
} //                                                Test_crcy_Currency_DivFloat_

// go test --run Test_crcy_Currency_DivInt_
func Test_crcy_Currency_DivInt_(t *testing.T) {
	// (ob Currency) DivInt(nums...int) Currency
	test := func(ob Currency, nums []int, expect Currency) {
		curIntOpTest(t, "DivInt", ob.DivInt, ob, nums, expect, 0)
	}
	TBegin(t)
	//
	// TODO: declaration comment
	//
	//       12345.6789 / 1   =               12345.6789
	test(Currency{123456789}, arI(1), Currency{123456789})
	//
	//       12345.6789 / 2   =               6172.8394
	test(Currency{123456789}, arI(2), Currency{61728394})
} //                                                  Test_crcy_Currency_DivInt_

// -----------------------------------------------------------------------------
// # Multiplication:

// go test --run Test_crcy_Currency_Mul_
func Test_crcy_Currency_Mul_(t *testing.T) {
	TBegin(t)
	// (ob Currency) Mul(nums... Currency) Currency
	//
	test := func(ob Currency, nums []Currency, expect Currency) {
		curOpTest(t, "Mul", ob.Mul, ob, nums, expect, 0)
	}
	test(cur(0), arC(0), cur(0))
	test(cur(1), arC(1), cur(1))
	test(cur(1), arC(1), cur(1))
	test(cur(2), arC(2), cur(4))
	test(cur(-2), arC(2), cur(-4))
	test(cur(2), arC(-2), cur(-4))
	test(cur(1000), arC(1000), cur(1000000))
	//
	// these would overflow if big.Int is not used
	test(Currency{4611686018427384999}, arC(2), Currency{9223372036854769998})
	test(Currency{922337203685476999}, arC(10), Currency{9223372036854769990})
	//
	// tests that cause overflow
	DisableErrors()
	test = func(ob Currency, nums []Currency, expect Currency) {
		curOpTest(t, "Mul", ob.Mul, ob, nums, expect, 1)
	}
	//
	// overflow: the result fits in int64, but not in acceptable Currency range
	test(Currency{MaxCurrencyI64 + 1}, arC(1), Currency{math.MaxInt64})
	//
	test(cur(int64(-627199999000000)), arC(int64(627199999000000)),
		Currency{math.MinInt64},
	)
	test(cur(int64(-922337203685476)), arC(int64(922337203685476)),
		Currency{math.MinInt64},
	)
	test(cur(int64(627199999000000)), arC(int64(627199999000000)),
		Currency{math.MaxInt64},
	)
	test(cur(int64(922337203685476)), arC(int64(922337203685476)),
		Currency{math.MaxInt64},
	)
	EnableErrors()
} //                                                     Test_crcy_Currency_Mul_

// go test --run Test_crcy_Currency_MulFloat_
func Test_crcy_Currency_MulFloat_(t *testing.T) {
	TBegin(t)
	// (ob Currency) MulFloat(nums...float64) Currency
	//
	test := func(ob Currency, nums []float64, expect Currency) {
		curFloatOpTest(t, "MulFloat", ob.MulFloat, ob, nums, expect, 0)
	}
	//            0   *   1.0   =        0
	test(Currency{0}, arF(1.0), Currency{0})
	//
	//            2   *           2.0   =        4
	test(Currency{2 * cur4d}, arF(2.0), Currency{4 * cur4d})
	//
	//           12345.6789   *   1   =       12345.6789
	test(Currency{123456789}, arF(1), Currency{123456789})
	//
	//           12345.6789   /   2   =        24691.3578
	test(Currency{123456789}, arF(2), Currency{246913578})
	//
	test(cur(-1), arF(1e6), cur(-1000000))             // -1 million
	test(cur(-1), arF(1e7), cur(-10000000))            // -10 million
	test(cur(-1), arF(1e8), cur(-100000000))           // -100 million
	test(cur(-1), arF(1e9), cur(-1000000000))          // -1 billion
	test(cur(-1), arF(1e10), cur(int64(-10000000000))) // -10 billion
	//
	test(cur(1), arF(1e6), cur(1000000))             // 1 million
	test(cur(1), arF(1e7), cur(10000000))            // 10 million
	test(cur(1), arF(1e8), cur(100000000))           // 100 million
	test(cur(1), arF(1e9), cur(1000000000))          // 1 billion
	test(cur(1), arF(1e10), cur(int64(10000000000))) // 10 billion
	//
	// Here, if Mul() didn't use big.Int, multiplication would overflow
	test(cur(1), arF(1e14), Currency{int64(1e18)})
	test(cur(-1), arF(1e14), Currency{int64(-1e18)})
	test(cur(123), arF(1e9), Currency{int64(123 * 1e13)}) // 123 billion
	//
	// overflow
	DisableErrors()
	test = func(ob Currency, nums []float64, expect Currency) {
		curFloatOpTest(t, "MulFloat", ob.MulFloat, ob, nums, expect, 1)
	}
	test(cur(1), arF(1e20), Currency{math.MaxInt64})
	test(cur(123), arF(float64(MaxCurrencyI64)), Currency{math.MaxInt64})
	test(cur(123), arF(MinCurrencyI64-0.001), Currency{math.MinInt64})
	EnableErrors()
} //                                                Test_crcy_Currency_MulFloat_

// go test --run Test_crcy_Currency_MulInt_
func Test_crcy_Currency_MulInt_(t *testing.T) {
	// (ob Currency) MulInt(nums...int) Currency
	test := func(ob Currency, nums []int, expect Currency) {
		curIntOpTest(t, "MulInt", ob.MulInt, ob, nums, expect, 0)
	}
	//            0   *   1   =        0
	test(Currency{0}, arI(1), Currency{0})
	//
	//            2           *   2   =        4
	test(Currency{2 * cur4d}, arI(2), Currency{4 * cur4d})
	//
	//           12345.6789   *   1   =       12345.6789
	test(Currency{123456789}, arI(1), Currency{123456789})
	//
	//           12345.6789   /   2   =       24691.3578
	test(Currency{123456789}, arI(2), Currency{246913578})
	//
	// TODO: Try to cause failure in unit test:
	//       pass MaxInt64 * just above currency limit.
} //                                                  Test_crcy_Currency_MulInt_

// -----------------------------------------------------------------------------
// # Addition:

// go test --run Test_crcy_Currency_Add_
func Test_crcy_Currency_Add_(t *testing.T) {
	TBegin(t)
	// (ob Currency) Add(nums... Currency) Currency
	//
	test := func(ob Currency, nums []Currency, expect Currency) {
		curOpTest(t, "Add", ob.Add, ob, nums, expect, 0)
	}
	// test additions within range of Currency:
	test(cur(0), arC(0), cur(0))
	test(cur(1), arC(1), cur(2))
	test(cur(1000), arC(1000), cur(2000))
	test(cur(-1), arC(-1), cur(-2))
	//
	// multiple additions:
	test(cur(1), arC(1.1, 2.02, 3.003, 4.0004), cur(11.1234))
	//
	// test addition close to the limits of Currency:
	//
	// largest + smallest number: must 'annihilate' to zero
	test(
		Currency{MinCurrencyI64},      // -9223372036854769999
		arC(Currency{MaxCurrencyI64}), //  9223372036854769999
		Currency{0},
	)
	test(
		Currency{MaxCurrencyI64},      //  9223372036854769999
		arC(Currency{MinCurrencyI64}), // -9223372036854769999
		Currency{0},
	)
	// -922337203685476 + -0.9999 = -922337203685476.9999
	test(
		Currency{-9223372036854760000},
		arC(Currency{-9999}),
		Currency{MinCurrencyI64},
	)
	test(
		Currency{-9999},
		arC(Currency{-9223372036854760000}),
		Currency{MinCurrencyI64},
	)
	// -922337203685476.9998 + -0.0001 = -922337203685476.9999
	test(
		Currency{-9223372036854769998},
		arC(Currency{-1}),
		Currency{MinCurrencyI64},
	)
	test(
		Currency{-1},
		arC(Currency{-9223372036854769998}),
		Currency{MinCurrencyI64},
	)
	// 922337203685476 + 0.9999 = 922337203685476.9999
	test(
		Currency{9223372036854760000},
		arC(Currency{9999}),
		Currency{MaxCurrencyI64},
	)
	test(
		Currency{9999},
		arC(Currency{9223372036854760000}),
		Currency{MaxCurrencyI64},
	)
	// 922337203685476.9998 + 0.0001 = 922337203685476.9999
	test(
		Currency{9223372036854769998},
		arC(Currency{1}),
		Currency{MaxCurrencyI64},
	)
	test(
		Currency{1},
		arC(Currency{9223372036854769998}),
		Currency{MaxCurrencyI64},
	)
	// test additions that will overflow Currency:
	DisableErrors()
	test = func(ob Currency, nums []Currency, expect Currency) {
		curOpTest(t, "Add", ob.Add, ob, nums, expect, 1)
	}
	// adding -0.0001 to the smallest number must overflow
	test(
		Currency{MinCurrencyI64}, // -922337203685476.9999
		arC(Currency{-1}),        //               -0.0001
		Currency{math.MinInt64},
	)
	// adding 0.0001 to the largest number must overflow
	test(
		Currency{MaxCurrencyI64}, //  922337203685476.9999
		arC(Currency{1}),         //                0.0001
		Currency{math.MaxInt64},
	)
	// adding two smallest numbers must overflow
	test(
		Currency{MinCurrencyI64},      // -922337203685476.9999
		arC(Currency{MinCurrencyI64}), // -922337203685476.9999
		Currency{math.MinInt64},
	)
	// adding two largest numbers must overflow
	test(
		Currency{MaxCurrencyI64},      //  922337203685476.9999
		arC(Currency{MaxCurrencyI64}), //  922337203685476.9999
		Currency{math.MaxInt64},
	)
	EnableErrors()
} //                                                     Test_crcy_Currency_Add_

// go test --run Test_crcy_Currency_AddFloat_
func Test_crcy_Currency_AddFloat_(t *testing.T) {
	TBegin(t)
	// (ob Currency) AddFloat(nums...float64) Currency
	//
	test := func(ob Currency, nums []float64, expect Currency) {
		curFloatOpTest(t, "AddFloat", ob.AddFloat, ob, nums, expect, 0)
	}
	test(cur(0), arF(0), cur(0))
	test(cur(1), arF(1), cur(2))
	//
	// test multiple arguments
	test(cur(1), arF(2, 3, 4), cur(10))
	//
	// must overflow
	DisableErrors()
	test = func(ob Currency, nums []float64, expect Currency) {
		curFloatOpTest(t, "AddFloat", ob.AddFloat, ob, nums, expect, 1)
	}
	test(cur(1), arF(1e30), Currency{math.MaxInt64})
	EnableErrors()
} //                                                Test_crcy_Currency_AddFloat_

// go test --run Test_crcy_Currency_AddInt_
func Test_crcy_Currency_AddInt_(t *testing.T) {
	// (ob Currency) AddInt(nums...int) Currency
	test := func(ob Currency, nums []int, expect Currency) {
		curIntOpTest(t, "AddInt", ob.AddInt, ob, nums, expect, 0)
	}
	test(cur(0), arI(0), cur(0))
} //                                                  Test_crcy_Currency_AddInt_

// -----------------------------------------------------------------------------
// # Subtraction:

// go test --run Test_crcy_Currency_Sub_
func Test_crcy_Currency_Sub_(t *testing.T) {
	TBegin(t)
	// (ob Currency) Sub(nums... Currency) Currency
	//
	test := func(ob Currency, nums []Currency, expect Currency) {
		curOpTest(t, "Sub", ob.Sub, ob, nums, expect, 0)
	}
	//
	// test subtractions within range of Currency:
	//
	test(cur(0), arC(0), cur(0))
	test(cur(0.0001), arC(0.0001), cur(0))
	test(cur(0), arC(0.0001), cur(-0.0001))
	test(cur(1), arC(1), cur(0))
	test(cur(10), arC(1, 2, 3), cur(4))
	test(cur(int64(1234567891234)), arC(67891234), cur(int64(1234500000000)))
	test(cur(2), arC(1), cur(1))
	test(cur(2000), arC(1000), cur(1000))
	test(cur(-2), arC(-1), cur(-1))
	//
	// multiple subtractions:
	test(cur(11.1234), arC(4.0004, 3.003, 2.02, 1.1), cur(1))
	//
	// test subtraction close to the limits of Currency:
	//
	// largest - largest number: must be to zero
	test(
		Currency{MaxCurrencyI64},      // 9223372036854769999
		arC(Currency{MaxCurrencyI64}), // 9223372036854769999
		Currency{0},
	)
	// -922337203685476.9999 - -0.9999 = -922337203685476
	test(
		Currency{-9223372036854769999},
		arC(Currency{-9999}),
		Currency{-9223372036854760000},
	)
	test(
		Currency{MinCurrencyI64},
		arC(Currency{-9223372036854760000}),
		Currency{-9999},
	)
	// -922337203685476.9999 - -922337203685476.9998 = -0.0001
	test(
		Currency{MinCurrencyI64},
		arC(Currency{-9223372036854769998}),
		Currency{-1},
	)
	// -922337203685476.9999 - 0.0001 = -922337203685476.9998
	test(
		Currency{MinCurrencyI64},
		arC(Currency{-1}),
		Currency{-9223372036854769998},
	)
	// 922337203685476.9999 - 922337203685476 = 0.9999
	test(
		Currency{MaxCurrencyI64},
		arC(Currency{9223372036854760000}),
		Currency{9999},
	)
	// 922337203685476.9999 - -0.9999 = 922337203685476
	test(
		Currency{9223372036854760000},
		arC(Currency{-9999}),
		Currency{MaxCurrencyI64},
	)
	// 922337203685476.9999 - 922337203685476.9998 = 0.0001
	test(
		Currency{MaxCurrencyI64},
		arC(Currency{9223372036854769998}),
		Currency{1},
	)
	// 922337203685476.9999 - 0.0001 = 922337203685476.9998
	test(
		Currency{MaxCurrencyI64},
		arC(Currency{1}),
		Currency{9223372036854769998},
	)
	// test subtraction that will overflow Currency:
	DisableErrors()
	test = func(ob Currency, nums []Currency, expect Currency) {
		curOpTest(t, "Sub", ob.Sub, ob, nums, expect, 1)
	}
	// subtracting 0.0001 from the smallest number must overflow
	test(
		Currency{MinCurrencyI64}, // -922337203685476.9999
		arC(Currency{1}),         //               -0.0001
		Currency{math.MinInt64},
	)
	// subtracting -0.0001 from the largest number must overflow
	test(
		Currency{MaxCurrencyI64}, // 922337203685476.9999
		arC(Currency{-1}),        //               0.0001
		Currency{math.MaxInt64},
	)
	// subtracting largest from smallest number must overflow
	test(
		Currency{MinCurrencyI64},      // -922337203685476.9999
		arC(Currency{MaxCurrencyI64}), //  922337203685476.9999
		Currency{math.MinInt64},
	)
	// subtracting smallest from largest number must overflow
	test(
		Currency{MaxCurrencyI64},      //  922337203685476.9999
		arC(Currency{MinCurrencyI64}), // -922337203685476.9999
		Currency{math.MaxInt64},
	)
	EnableErrors()
} //                                                     Test_crcy_Currency_Sub_

// go test --run Test_crcy_Currency_SubFloat_
func Test_crcy_Currency_SubFloat_(t *testing.T) {
	TBegin(t)
	// (ob Currency) SubFloat(nums...float64) Currency
	//
	test := func(ob Currency, nums []float64, expect Currency) {
		curFloatOpTest(t, "SubFloat", ob.SubFloat, ob, nums, expect, 0)
	}
	test(cur(0), arF(0), cur(0))
} //                                                Test_crcy_Currency_SubFloat_

// go test --run Test_crcy_Currency_SubInt_
func Test_crcy_Currency_SubInt_(t *testing.T) {
	TBegin(t)
	// (ob Currency) SubInt(nums...int) Currency
	//
	test := func(ob Currency, nums []int, expect Currency) {
		curIntOpTest(t, "SubInt", ob.SubInt, ob, nums, expect, 0)
	}
	test(cur(0), arI(0), cur(0))
	// TODO: more unit test cases
} //                                                  Test_crcy_Currency_SubInt_

// -----------------------------------------------------------------------------
// # Information:

// go test --run Test_crcy_Currency_Float64_
func Test_crcy_Currency_Float64_(t *testing.T) {
	TBegin(t)
	// (ob Currency) Float64() float64
	//
	test := func(ob Currency, expect float64) {
		got := ob.Float64()
		if got != expect {
			TFail(t,
				`(`, ob, `).Float64() returned `, got, `. must be `, expect,
			)
		}
	}
	test(cur(0), 0.0)
	test(Currency{MinCurrencyI64}, -9.22337203685477e+14)
	test(Currency{MaxCurrencyI64}, 9.22337203685477e+14)
	test(cur(1234567890), 1234567890)
	test(cur(987654321), 987654321)
	test(Currency{1}, 0.0001)
	test(Currency{-1}, -0.0001)
} //                                                 Test_crcy_Currency_Float64_

// go test --run Test_crcy_Currency_Int_
func Test_crcy_Currency_Int_(t *testing.T) {
	TBegin(t)
	// (ob Currency) Int() int64
	//
	test := func(ob Currency, expect int64) {
		if ob.Int() != expect {
			TFail(
				t, `(`, ob.String(), `).Int()`,
				` returned`, ob.Int(), `, must be `, expect,
			)
		}
	}
	test(cur(-1), -1)
	test(cur(0), 0)
	test(cur(1), 1)
	// TODO: more unit test cases
} //                                                     Test_crcy_Currency_Int_

// go test --run Test_crcy_Currency_Int64_
func Test_crcy_Currency_Int64_(t *testing.T) {
	TBegin(t)
	// (ob Currency) Int64() int64
	//
	test := func(ob Currency, expect int64) {
		if ob.Int64() != expect {
			TFail(
				t, `(`, ob.String(), `).Int64()`,
				` returned`, ob.Int64(), `, must be `, expect,
			)
		}
	}
	test(cur(-1), -1)
	test(cur(0), 0)
	test(cur(1), 1)
	// TODO: more unit test cases
} //                                                   Test_crcy_Currency_Int64_

// go test --run Test_crcy_Currency_IsEqual_
func Test_crcy_Currency_IsEqual_(t *testing.T) {
	TBegin(t)
	// (ob Currency) IsEqual() bool
	//
	test := func(ob, n Currency, expect bool) {
		got := ob.IsEqual(n)
		if got != expect {
			TFail(
				t, `(`, ob, `).IsEqual(`, n, `)`,
				` returned `, got, `. must be `, expect,
			)
		}
	}
	test(cur(0), cur(0), true)
	test(cur(1), cur(1), true)
	test(cur(0), cur(1), false)
	test(cur(1), cur(0), false)
	// TODO: more unit test cases
} //                                                 Test_crcy_Currency_IsEqual_

// go test --run Test_crcy_Currency_IsNegative_
func Test_crcy_Currency_IsNegative_(t *testing.T) {
	TBegin(t)
	// (ob Currency) IsNegative() bool
	//
	test := func(ob Currency, expect bool) {
		got := ob.IsNegative()
		if got != expect {
			TFail(
				t, `(`, ob, `).IsNegative()`,
				` returned`, got, `. must be `, expect,
			)
		}
	}
	test(Currency{MinCurrencyI64}, true)
	test(cur(-1000), true)
	test(cur(-1), true)
	//
	test(cur(0), false)
	//
	test(cur(1), false)
	test(cur(1000), false)
	test(Currency{MaxCurrencyI64}, false)
	// TODO: more unit test cases
} //                                              Test_crcy_Currency_IsNegative_

// go test --run Test_crcy_Currency_IsZero_
func Test_crcy_Currency_IsZero_(t *testing.T) {
	TBegin(t)
	// (ob Currency) IsZero() bool
	//
	test := func(ob Currency, expect bool) {
		got := ob.IsZero()
		if got != expect {
			TFail(
				t, `(`, ob, `).IsZero() returned`, got, `. must be `, expect,
			)
		}
	}
	test(Currency{MinCurrencyI64}, false)
	test(cur(-1000), false)
	test(cur(-1), false)
	test(cur(1), false)
	test(cur(1000), false)
	test(Currency{MaxCurrencyI64}, false)
	//
	test(cur(0), true)
	test(cur(""), true)
	test(cur("0"), true)
	test(cur("00"), true)
	test(Currency{0}, true)
	test(cur(float32(0)), true)
	test(cur(float64(0)), true)
	test(cur("-0.00001"), true) // because this is > 4 decimal places
	//
	test(cur("-0.0001"), false)
	test(cur("0.0001"), false)
	test(Currency{1}, false)
	test(Currency{-1}, false)
	test(cur("1"), false)
	test(cur("-922337203685476.9999"), false) // lowest currency value
	test(cur("922337203685476.9999"), false)  // highest currency value
	//
	// TODO: more unit test cases
	//
	// logged with errors
	// TODO: catch error here
	//       {"XYZ", true}, // non-numeric strings are initialized to zero
} //                                                  Test_crcy_Currency_IsZero_

// go test --run Test_crcy_Currency_Overflow_
func Test_crcy_Currency_Overflow_(t *testing.T) {
	TBegin(t)
	// (ob Currency) Overflow() int
	//
	test := func(ob Currency, expect int) {
		got := ob.Overflow()
		if got != expect {
			TFail(t, `cur(`, ob, `) returned `, got, `. must be `, expect)
		}
	}
	test(Currency{math.MinInt64}, -1)
	test(Currency{MinCurrencyI64 - 1}, -1)
	test(Currency{MinCurrencyI64}, 0)
	test(Currency{0}, 0)
	test(Currency{MaxCurrencyI64}, 0)
	test(Currency{MaxCurrencyI64 + 1}, 1)
	test(Currency{math.MaxInt64}, 1)
} //                                                Test_crcy_Currency_Overflow_

// -----------------------------------------------------------------------------
// # JSON:

// go test --run Test_crcy_Currency_MarshalJSON_
func Test_crcy_Currency_MarshalJSON_(t *testing.T) {
	TBegin(t)
	// (ob Currency) MarshalJSON() ([]byte, error)
	//
	test := func(input interface{}, expect string) {
		type T struct {
			Val Currency
		}
		var ty T
		ty.Val = cur(input)
		jsn, _ := json.MarshalIndent(ty, "", " ")
		//                 ^  calls the object's MarshalIndent() method
		got := string(jsn)
		got = strings.Replace(got, "\n", "", -1)
		got = strings.Replace(got, "{ ", "{", -1)
		TEqual(t, got, (expect))
	}
	test(-1, `{"Val": -1}`)
	test(0, `{"Val": 0}`)
	test(0.1, `{"Val": 0.1}`)
	test(0.01, `{"Val": 0.01}`)
	test(0.001, `{"Val": 0.001}`)
	test(0.0001, `{"Val": 0.0001}`)
	test(1, `{"Val": 1}`)
	//
	test(-100000, `{"Val": -100000}`)   // -100 thousand
	test(100000, `{"Val": 100000}`)     // 100 thousand
	test(-1000000, `{"Val": -1000000}`) // -1 million
	test(1000000, `{"Val": 1000000}`)   // 1 million
	//
	test(math.MaxInt32, `{"Val": 2147483647}`)
	test(math.MinInt32, `{"Val": -2147483648}`)
	//
	test(123.456, `{"Val": 123.456}`)
	test("2.0001", `{"Val": 2.0001}`)
	test("1,000,000,000,000.0001", `{"Val": 1000000000000.0001}`)
	//
	//          Tn  Bn  Millions
	test("1,234,567,890,123.4321", `{"Val": 1234567890123.4321}`)
	test("9,999,999,999,999.9999", `{"Val": 9999999999999.9999}`)
	//
	// overflow
	DisableErrors()
	test(int64(922337203685477), `{"Val": 922337203685477.5807}`)
	test(int64(-922337203685476), `{"Val": -922337203685476}`)
	EnableErrors()
} //                                             Test_crcy_Currency_MarshalJSON_

// go test --run Test_crcy_Currency_UnmarshalJSON_
func Test_crcy_Currency_UnmarshalJSON_(t *testing.T) {
	TBegin(t)
	// (ob *Currency) UnmarshalJSON(data []byte) error
	//
	{
		var n Currency
		err := n.UnmarshalJSON([]byte("123"))
		if err != nil {
			TFail(t, err)
		}
		if n != cur(123) {
			TFail(t, err)
		}
	}
	// error with nil receiver
	{
		DisableErrors()
		var ob *Currency
		err := ob.UnmarshalJSON([]byte("123"))
		if err == nil {
			TFail(t, `Did not return an error when called on nil receiver.`)
		}
		EnableErrors()
	}
	// test for error
	{
		DisableErrors()
		//
		// mock jsonUnmarshal() so it returns an error (UnmarshalJSON calls it)
		const ERRM = "our error message"
		mod.json.Unmarshal = func([]byte, interface{}) error {
			return fmt.Errorf(ERRM)
		}
		defer mod.Reset() // undo the mock!
		//
		// do the test
		var n Currency
		err := n.UnmarshalJSON([]byte("123"))
		//
		if err == nil {
			TFail(t, err)
		}
		if err.Error() != ERRM {
			TFail(t, `returned error "`, err.Error(), `". must be "`, ERRM, `"`)
		}
		if n != cur(0) {
			TFail(t, err)
		}
		EnableErrors()
	}
} //                                           Test_crcy_Currency_UnmarshalJSON_

// -----------------------------------------------------------------------------
// # Helper Functions

// go test --run Test_crcy_currencyOverflow_
func Test_crcy_currencyOverflow_(t *testing.T) {
	TBegin(t)
	// currencyOverflow(isNegative bool, a ...interface{}) int64
	//
	// mock Error() function called by tested function
	var errorCalled bool
	var errm string
	mod.Error = func(a ...interface{}) error {
		errm = fmt.Sprint(a...)
		errorCalled = true
		return fmt.Errorf(errm)
	}
	defer mod.Reset() // restore standard functions!
	//
	errorCalled = false
	errm = ""
	TEqual(t, currencyOverflow(true, "negative overflow ", -1),
		Currency{math.MinInt64},
	)
	TEqual(t, errm, "negative overflow -1")
	TTrue(t, errorCalled)
	//
	errorCalled = false
	errm = ""
	TEqual(t, currencyOverflow(false, "positive overflow ", 1),
		Currency{math.MaxInt64},
	)
	TEqual(t, errm, "positive overflow 1")
	TTrue(t, errorCalled)
} //                                                 Test_crcy_currencyOverflow_

// -----------------------------------------------------------------------------
// # Test Helper Functions

// arC is a convenience function to create an array of Currency numbers.
// That is, instead of having to specify
// '[]Currency{Currency{10000}, Currency{20000}, Currency{30000}}'
// you can just use 'arC(1, 2, 3)'.
func arC(ar ...interface{}) (ret []Currency) {
	for _, v := range ar {
		ret = append(ret, cur(v))
	}
	return ret
} //                                                                         arC

// arF is a convenience function to create an array of float64 numbers.
// That is, instead of having to specify '[]float64{1.0, 2.0, 3.0}'
// you can just use 'arF(1.0, 2.0, 3.0)'.
func arF(ar ...interface{}) (ret []float64) {
	for _, v := range ar {
		ret = append(ret, Float64(v))
	}
	return ret
} //                                                                         arF

// arI is a convenience function to create an array of int numbers.
// That is, instead of having to specify '[]int{1, 2, 3}'
// you can just use 'arI(1, 2, 3)'.
func arI(ar ...interface{}) (ret []int) {
	for _, v := range ar {
		ret = append(ret, Int(v))
	}
	return ret
} //                                                                         arI

var cur = CurrencyOf // a short CurrencyOf() alias used in many unit tests here

// curFloatOpTest tests basic arithmetic
// operations with Currency and float64.
func curFloatOpTest(
	t *testing.T,
	opName string,
	opFunc func(values ...float64) Currency,
	ob Currency,
	values []float64,
	expect Currency,
	expectErrors int,
) {
	var (
		old = ob
		ec1 = GetErrorCount()
		got = opFunc(values...)
		ec2 = GetErrorCount()
	)
	if expectErrors == 0 {
		// must have no errors
		if ec2 != ec1 {
			TFail(t, `Got `, ec2-ec1, ` error(s)`)
		}
	} else {
		// must have one error
		if ec2 != ec1+1 {
			TFail(t, `Expected 1 error, but got `, ec2-ec1)
		}
	}
	// object of invoked method must not change
	if ob.i64 != old.i64 {
		TFail(t, `(`, old, `) mutated to `, ob)
	}
	// check if returned value matches expected value
	if got.i64 != expect.i64 {
		TFail(t, `(`, old, `).`, opName, `(`, values, `)`,
			` returned `, got, `. must be `, expect,
		)
	}
} //                                                              curFloatOpTest

// curIntOpTest tests basic arithmetic operations with Currency and int.
func curIntOpTest(
	t *testing.T,
	opName string,
	opFunc func(values ...int) Currency,
	ob Currency,
	values []int,
	expect Currency,
	expectErrors int,
) {
	var (
		old = ob
		ec1 = GetErrorCount()
		got = opFunc(values...)
		ec2 = GetErrorCount()
	)
	if expectErrors == 0 {
		// must have no errors
		if ec2 != ec1 {
			TFail(t, `Got `, ec2-ec1, ` error(s)`)
		}
	} else {
		// must have one error
		if ec2 != ec1+1 {
			TFail(t, `Expected 1 error, but got `, ec2-ec1)
		}
	}
	// object of invoked method must not change
	if ob.i64 != old.i64 {
		TFail(t, `(`, old, `) mutated to `, ob)
	}
	// check if returned value matches expected value
	if got.i64 != expect.i64 {
		TFail(t, `(`, old, `).`, opName, `(`, values, `)`,
			` returned `, got, `. must be `, expect,
		)
	}
} //                                                                curIntOpTest

// curOpTest tests basic arithmetic operations with Currency.
func curOpTest(
	t *testing.T,
	opName string,
	opFunc func(values ...Currency) Currency,
	ob Currency,
	values []Currency,
	expect Currency,
	expectErrors int,
) {
	var (
		old = ob
		ec1 = GetErrorCount()
		got = opFunc(values...)
		ec2 = GetErrorCount()
	)
	if expectErrors == 0 {
		// must have no errors
		if ec2 != ec1 {
			TFail(t, `Got `, ec2-ec1, ` error(s)`)
		}
	} else {
		// must have one error
		if ec2 != ec1+1 {
			TFail(t, `Expected 1 error, but got `, ec2-ec1)
		}
	}
	// object of invoked method must not change
	if ob.i64 != old.i64 {
		TFail(t, `(`, old, `) mutated to `, ob)
	}
	// check if returned value matches expected value
	if got.i64 != expect.i64 {
		TFail(t, `(`, old, `).`, opName, `(`, values, `)`,
			` returned `, got, `. must be `, expect,
		)
	}
} //                                                                   curOpTest

//end
