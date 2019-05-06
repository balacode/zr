// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-06 06:25:43 C28530                           zr/[numbers_test.go]
// -----------------------------------------------------------------------------

package zr

// # Numeric Function Tests
//   Test_nums_BlankZero_
//   Test_nums_CommaDelimit_
//   Test_nums_Float64_ // TODO: create unit test
//   Test_nums_Int_
//   Test_nums_IntInWordsEN_
//   Test_nums_IsNumber_

/*
to test all items in numbers.go use:
	go test --run Test_nums_

to generate a test coverage report use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

import (
	"fmt"
	"math"
	"testing"
)

// go test --run Test_nums_consts_
func Test_nums_consts_(t *testing.T) {
	//
	// min/max int and uint values (on 64-bit platforms only!)
	TEqual(t, MaxInt, int64(9223372036854775807))
	TEqual(t, MaxUint, uint64(18446744073709551615))
	TEqual(t, MinInt, int64(-9223372036854775808))
} //                                                           Test_nums_consts_

// go test --run Test_nums_BlankZero_
func Test_nums_BlankZero_(t *testing.T) {
	TBegin(t)
	// BlankZero(s string) string
	//
	test := func(input, expect string) {
		got := BlankZero(input)
		if !TEqual(t, got, (expect)) {
			fmt.Printf("INPUT: %s"+LF+LF, input)
		}
	}
	// zero values should become blank
	test(" .0 ", "")
	test(" .0", "")
	test(" 0 ", "")
	test(" 0", "")
	test(".0 ", "")
	test(".0", "")
	test("0", "")
	test("0.0", "")
	test("0.00", "")
	test("00", "")
	test("00.0", "")
	//
	// non-zero strings should remain unchanged
	test(" ", " ")
	test(" .01 ", " .01 ")
	test(" .01", " .01")
	test(" 01 ", " 01 ")
	test(" 01", " 01")
	test(".01 ", ".01 ")
	test(".01", ".01")
	test("0.001", "0.001")
	test("0.01", "0.01")
	test("00.01", "00.01")
	test("001", "001")
	test("01", "01")
	test("1.00", "1.00")
	test("abc", "abc")
} //                                                        Test_nums_BlankZero_

// go test --run Test_nums_CommaDelimit_
func Test_nums_CommaDelimit_(t *testing.T) {
	TBegin(t)
	// CommaDelimit(number string, decimalPlaces int) string
	//
	test := func(s string, decimalPlaces int, expect string) {
		got := CommaDelimit(s, decimalPlaces)
		if got != expect {
			t.Errorf("CommaDelimit(%q, %d) returned %q instead of %q",
				s, decimalPlaces, got, expect)
		}
	}
	test(".1", 0, "0")
	test(".1", 1, "0.1")
	test(".1", 2, "0.10")
	test(".1", 3, "0.100")
	test("1.00", 0, "1")
	test("123", 0, "123")
	test("123.45", 0, "123")
	test("1234", 0, "1,234")
	test("1234.56", 0, "1,234")
	test("1234.567", 0, "1,234")
	test("1234.567", 1, "1,234.5")
	test("1234.567", 2, "1,234.56")
	test("1234.567", 3, "1,234.567")
	test("1234.567", 4, "1,234.5670")
	test("12345", 0, "12,345")
	test("12345.67", 0, "12,345")
	test("123456", 0, "123,456")
	test("1234567", 0, "1,234,567")
	test("12345678", 0, "12,345,678")
} //                                                     Test_nums_CommaDelimit_

// go test --run Test_nums_IntInWordsEN_
func Test_nums_IntInWordsEN_(t *testing.T) {
	TBegin(t)
	// IntInWordsEN(number int64) string
	//
	test := func(input int64, expect string) {
		got := IntInWordsEN(input)
		if got != expect {
			t.Errorf("IntInWordsEN(%d) returned %q instead of %q",
				input, got, expect)
		}
	}
	test(256, "Two Hundred and Fifty Six")
	test(0, "Zero")
	test(1, "One")
	test(2, "Two")
	test(3, "Three")
	test(4, "Four")
	test(5, "Five")
	test(6, "Six")
	test(7, "Seven")
	test(8, "Eight")
	test(9, "Nine")
	test(10, "Ten")
	test(11, "Eleven")
	test(12, "Twelve")
	test(13, "Thirteen")
	test(14, "Fourteen")
	test(15, "Fifteen")
	test(16, "Sixteen")
	test(17, "Seventeen")
	test(18, "Eighteen")
	test(19, "Nineteen")
	test(20, "Twenty")
	test(21, "Twenty One")
	test(22, "Twenty Two")
	test(23, "Twenty Three")
	test(24, "Twenty Four")
	test(25, "Twenty Five")
	test(26, "Twenty Six")
	test(27, "Twenty Seven")
	test(28, "Twenty Eight")
	test(29, "Twenty Nine")
	test(30, "Thirty")
	test(31, "Thirty One")
	test(32, "Thirty Two")
	test(33, "Thirty Three")
	test(34, "Thirty Four")
	test(35, "Thirty Five")
	test(36, "Thirty Six")
	test(37, "Thirty Seven")
	test(38, "Thirty Eight")
	test(39, "Thirty Nine")
	test(40, "Forty")
	test(41, "Forty One")
	test(42, "Forty Two")
	test(43, "Forty Three")
	test(44, "Forty Four")
	test(45, "Forty Five")
	test(46, "Forty Six")
	test(47, "Forty Seven")
	test(48, "Forty Eight")
	test(49, "Forty Nine")
	test(50, "Fifty")
	test(51, "Fifty One")
	test(52, "Fifty Two")
	test(53, "Fifty Three")
	test(54, "Fifty Four")
	test(55, "Fifty Five")
	test(56, "Fifty Six")
	test(57, "Fifty Seven")
	test(58, "Fifty Eight")
	test(59, "Fifty Nine")
	test(60, "Sixty")
	test(61, "Sixty One")
	test(62, "Sixty Two")
	test(63, "Sixty Three")
	test(64, "Sixty Four")
	test(65, "Sixty Five")
	test(66, "Sixty Six")
	test(67, "Sixty Seven")
	test(68, "Sixty Eight")
	test(69, "Sixty Nine")
	test(70, "Seventy")
	test(71, "Seventy One")
	test(72, "Seventy Two")
	test(73, "Seventy Three")
	test(74, "Seventy Four")
	test(75, "Seventy Five")
	test(76, "Seventy Six")
	test(77, "Seventy Seven")
	test(78, "Seventy Eight")
	test(79, "Seventy Nine")
	test(80, "Eighty")
	test(81, "Eighty One")
	test(82, "Eighty Two")
	test(83, "Eighty Three")
	test(84, "Eighty Four")
	test(85, "Eighty Five")
	test(86, "Eighty Six")
	test(87, "Eighty Seven")
	test(88, "Eighty Eight")
	test(89, "Eighty Nine")
	test(90, "Ninety")
	test(91, "Ninety One")
	test(92, "Ninety Two")
	test(93, "Ninety Three")
	test(94, "Ninety Four")
	test(95, "Ninety Five")
	test(96, "Ninety Six")
	test(97, "Ninety Seven")
	test(98, "Ninety Eight")
	test(99, "Ninety Nine")
	test(100, "One Hundred")
	test(101, "One Hundred and One")
	//
	//FAIL test(math.MinInt64, "-")
	test(
		math.MaxInt32, // = 2,147,483,647
		"Two Billion"+
			" One Hundred and Forty Seven Million"+
			" Four Hundred and Eighty Three Thousand"+
			" Six Hundred and Forty Seven",
	)
	test(
		math.MaxInt64, // = 9,223,372,036,854,775,807
		"Nine Quintillion"+
			" Two Hundred and Twenty Three Quadrillion"+
			" Three Hundred and Seventy Two Trillion"+
			" Thirty Six Billion"+
			" Eight Hundred and Fifty Four Million"+
			" Seven Hundred and Seventy Five Thousand"+
			" Eight Hundred and Seven",
	)
} //                                                     Test_nums_IntInWordsEN_

// go test --run Test_nums_Int_
func Test_nums_Int_(t *testing.T) {
	TBegin(t)
	// Int(val interface{}) int
	//
	test := func(val interface{}, expect int) {
		got := Int(val)
		if got != expect {
			t.Errorf("Int(%q) returned %d instead of %d"+LB,
				val, got, expect)
		}
	}
	// blank strings return 0
	test("", 0)
	test(" ", 0)
	test("  ", 0)
	//
	// non-numeric strings return 0
	test(".", 0)
	test("+", 0)
	test("-", 0)
	test("XYZ", 0)
	//
	// int32 limits
	test("-2147483648", math.MinInt32)
	test("2147483647", math.MaxInt32)
	//
	// once there is a space or any non-digit after a digit, parsing ends
	test(" 1-", 1)
	test(" 123 456", 123)
	test(" 123.456", 123)
	test(" 123/xyz", 123)
	test(" 1A", 1)
	test(" 23BC", 23)
	test(" 456DEF", 456)
	test(" A1 ", 0)
	test(" A1", 0)
	//
	test("1-", 1)
	test("123 456", 123)
	test("123.456", 123)
	test("123/xyz", 123)
	test("1A", 1)
	test("23BC", 23)
	test("456DEF", 456)
	test("A1 ", 0)
	test("A1", 0)
	//
	// zero integer (any decimals are ignored)
	test(" 0 ", 0)
	test(" 0", 0)
	test(" 00 ", 0)
	test(" 000000000 ", 0)
	test(" 000000000", 0)
	test("+0", 0)
	test("+0.99", 0)
	test("-0", 0)
	test("-0.99", 0)
	test("0 ", 0)
	test("0", 0)
	test("0.", 0)
	test("0.0", 0)
	test("0.9", 0)
	test("0.99", 0)
	test("00", 0)
	test("000", 0)
	test("000000000 ", 0)
	test("000000000", 0)
	//
	// single digits
	test("0", 0)
	test("1", 1)
	test("2", 2)
	test("3", 3)
	test("4", 4)
	test("5", 5)
	test("6", 6)
	test("7", 7)
	test("8", 8)
	test("9", 9)
	//
	// leading/trailing whitespace is OK
	test(" 123 ", 123)
	test(" 456", 456)
	test("789 ", 789)
	//
	// negative integers
	test(" -1", -1)
	test("-123", -123)
	test("-456", -456)
	test("-789", -789)
	//
	// decimals are ignored
	test("123456.789", 123456)
	test("-123456.789", -123456)
	//
	test(" 0000000001", 1)
	test(" 1-", 1)
	test(" 100000000 ", 100000000)
	test(" 123456789 ", 123456789)
	test(" 123456789", 123456789)
	test("-1", -1)
	test("0000000001", 1)
	test("100000000 ", 1e8)
	test("123456789 ", 123456789)
	test("123456789", 123456789)
	//
	// leading white-space
	test(" 0", 0)
	test("  1", 1)
	test("   2", 2)
	test("    3", 3)
	test("     4", 4)
	test("      5", 5)
	test("       6", 6)
	test("        7", 7)
	test("         8", 8)
	test("          9", 9)
	//
	// trailing white-space
	test("0 ", 0)
	test("1  ", 1)
	test("2   ", 2)
	test("3    ", 3)
	test("4     ", 4)
	test("5      ", 5)
	test("6       ", 6)
	test("7        ", 7)
	test("8         ", 8)
	test("9          ", 9)
	// overflow
	// TODO: test(10e20, MAX)
} //                                                              Test_nums_Int_

// go test --run Test_nums_IsNumber_
func Test_nums_IsNumber_(t *testing.T) {
	TBegin(t)
	// IsNumber(val interface{}) bool
	//
	test := func(input string, expect bool) {
		got := IsNumber(input)
		if got != expect {
			t.Errorf("IsNumber(%q) returned %v instead of %v",
				input, got, expect)
		}
	}
	test(" - 1 ", false)
	test(" -1 ", true)
	test(" -123,456.78 ", true)
	test(" .0. ", false)
	test(" 0 ", true)
	test(" 0. ", true)
	test(" 123,456.78 ", true)
	test("-.", false)
	test("0", true)
	test("123,456.78 ", true)
	test("123,456.78", true)
} //                                                         Test_nums_IsNumber_

// go test --run Test_nums_MinMaxGap_
func Test_nums_MinMaxGap_(t *testing.T) {
	TBegin(t)
	// MinMaxGap(values []int) (min, max int)
	//
	test := func(values []int, expectMin, expectMax int) {
		// save old array
		prevValues := make([]int, len(values))
		copy(prevValues, values)
		//
		gotMin, gotMax := MinMaxGap(values)
		if gotMin != expectMin || gotMax != expectMax {
			TFail(t,
				`MinMaxGap(`, GoString(values), `)`,
				` returned (`, gotMin, gotMax, `).`,
				` must be (`, expectMin, expectMax, `)`,
			)
		}
		// check that input slice hasn't been altered
		TArrayEqual(t, values, prevValues)
	}
	//
	// empty slice, should return MaxInt, MinInt
	test([]int{}, MaxInt, MinInt)
	//
	// a slice with one element should always return MaxInt, MinInt
	test([]int{42}, MaxInt, MinInt)
	//
	// two unique values, with one gap: 5
	test([]int{4, 6, 4, 6}, 5, 5)
	test([]int{4, 4, 6, 6}, 5, 5)
	//
	// no gap: all numbers same: return MaxInt, MinInt
	test([]int{1}, MaxInt, MinInt)
	test([]int{1, 1}, MaxInt, MinInt)
	test([]int{1, 1, 1}, MaxInt, MinInt)
	//
	// only 4 is available, return 4, true
	test([]int{7, 2, 3, 5, 6, 1}, 4, 4)
	test([]int{7, 2, 5, 1, 3, 6}, 4, 4)
	//
	// several slots available, return 2 and 6
	test([]int{1, 7, 5}, 2, 6)
	//
	// no free spaces, should return MaxInt, MinInt
	test([]int{1, 2, 3, 4, 5, 6, 7}, MaxInt, MinInt)
	test([]int{1, 1, 2, 2, 3, 3, 4}, MaxInt, MinInt)
} //                                                        Test_nums_MinMaxGap_

// go test --run Test_nums_MaxIntOf_
func Test_nums_MaxIntOf_(t *testing.T) {
	TBegin(t)
	// MaxIntOf(values []int) (max int, found bool)
	//
	test := func(values []int, expectMax int, expectFound bool) {
		// save old array
		oldValues := make([]int, len(values))
		copy(oldValues, values)
		//
		gotMax, gotFound := MaxIntOf(values)
		//
		if gotMax != expectMax || gotFound != expectFound {
			TFail(t,
				`MaxIntOf(`, GoString(values), `)`,
				` returned (`, gotMax, gotFound, `).`,
				` must be (`, expectMax, expectFound, `)`,
			)
		}
		// check that values slice hasn't been altered
		TArrayEqual(t, oldValues, values)
	}
	//
	// empty slice, should return 0, false
	test([]int{}, 0, false)
	//
	// one item, should return same
	test([]int{0}, 0, true)
	test([]int{1}, 1, true)
	test([]int{math.MinInt32}, math.MinInt32, true)
	test([]int{math.MaxInt32}, math.MaxInt32, true)
	//
	// return the highest from a range:
	test([]int{1, 5, 7}, 7, true)
	test([]int{3, 1, 6, 4, 9, 2}, 9, true)
	test([]int{789, 456, 23, 1}, 789, true)
} //                                                         Test_nums_MaxIntOf_

//end
