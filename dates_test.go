// -----------------------------------------------------------------------------
// ZR Library                                                 zr/[dates_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # DateRange
//   Test_date_DateRangeString_
//
// # Functions
//   Test_date_DateOf_
//   Test_date_DateRangeOf_
//   Test_date_DaysInMonth_
//   Test_date_IsDateOnly_
//   Test_date_IsDate_
//   Test_date_ParseDate_
//   Test_date_StringDateDMY_
//   Test_date_StringDateYMD_
//   Test_date_Timestamp_

//  to test all items in dates.go use:
//      go test --run Test_date_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// -----------------------------------------------------------------------------
// # DateRange

// go test --run Test_date_DateRangeString_
func Test_date_DateRangeString_(t *testing.T) {
	TBegin(t)
	//
	// TODO: declaration comment
	//
	var (
		fr = time.Date(1900, time.January, 1, 1, 0, 0, 0, time.UTC)
		to = time.Date(9999, time.December, 31, 23, 59, 59, 0, time.UTC)
		rg = DateRange{From: fr, To: to}
	)
	TEqual(t, rg.String(), ("1900-01-01 9999-12-31"))
} //                                                  Test_date_DateRangeString_

// -----------------------------------------------------------------------------
// # Functions

// go test --run Test_date_DateOf_
func Test_date_DateOf_(t *testing.T) {
	TBegin(t)
	//
	// TODO: declaration comment
	//
	test := func(input interface{}, expect string) {
		got := DateOf(input)
		TEqual(t, got.String(), (expect))
	}
	DisableErrors()
	test("", "0001-01-01 00:00:00 +0000 UTC")
	EnableErrors()
	test("2001-01-01", "2001-01-01 00:00:00 +0000 UTC")
	test("2017-03-01", "2017-03-01 00:00:00 +0000 UTC")
	test("2017-03-01 20:33:00", "2017-03-01 00:00:00 +0000 UTC")
} //                                                           Test_date_DateOf_

// go test --run Test_date_DateRangeOf_
func Test_date_DateRangeOf_(t *testing.T) {
	TBegin(t)
	//
	// DateRangeOf(s string) DateRange
	//
	// useful for manual debugging: (keep condition false):
	if false {
		s := "Oct23"
		rng := DateRangeOf(s)
		fmt.Printf("\r\n\r\n\r\n"+"DateRangeOf(%q) -> %q\r\n\r\n\r\n",
			s, rng.String())
		return
	}
	if false {
		s := "January2"
		dt, err := time.Parse(s, s)
		fmt.Printf("s:%s dt:%v err:%v", s, dt, err)
		return
	}
	test := func(input, expect string) {
		TEqual(t, DateRangeOf(input).String(), (expect))
	}
	// replaces 'yyyy' with current year
	thisYear := func(s string) string {
		year := String(int(time.Now().Year()))
		return strings.Replace(s, "yyyy", year, -1)
	}
	// month only
	test("MAY", thisYear("yyyy-05-01 yyyy-05-31"))
	//
	// date ranges
	test("1987 to 1999", "1987-01-01 1999-12-31")
	test("1799_to_2005", "1799-01-01 2005-12-31")
	test("2001~2015", "2001-01-01 2015-12-31")
	test("Jun/2008~Jul/2009", "2008-06-01 2009-07-31")
	//
	// year only
	test("1900", "1900-01-01 1900-12-31")
	test("2016", "2016-01-01 2016-12-31")
	//
	// month and day
	test("14-Jan", thisYear("yyyy-01-14 yyyy-01-14"))
	test("15-February", thisYear("yyyy-02-15 yyyy-02-15"))
	test("16Mar", thisYear("yyyy-03-16 yyyy-03-16"))
	test("17April", thisYear("yyyy-04-17 yyyy-04-17"))
	test("May-18", thisYear("yyyy-05-18 yyyy-05-18"))
	test("June-19", thisYear("yyyy-06-19 yyyy-06-19"))
	test("July20", thisYear("yyyy-07-20 yyyy-07-20"))
	test("Aug-21", thisYear("yyyy-08-21 yyyy-08-21"))
	test("September22", thisYear("yyyy-09-22 yyyy-09-22"))
	test("Oct23", thisYear("yyyy-10-23 yyyy-10-23"))
	//
	// year and month name
	test("Nov_2016", "2016-11-01 2016-11-30")
	test("nov_2016", "2016-11-01 2016-11-30")
	//
	// year and month name (for a single year)
	test("2016JAN", "2016-01-01 2016-01-31")
	test("2016FEB", "2016-02-01 2016-02-29")
	test("2016MAR", "2016-03-01 2016-03-31")
	test("2016APR", "2016-04-01 2016-04-30")
	test("2016MAY", "2016-05-01 2016-05-31")
	test("2016JUN", "2016-06-01 2016-06-30")
	test("2016JUL", "2016-07-01 2016-07-31")
	test("2016AUG", "2016-08-01 2016-08-31")
	test("2016SEP", "2016-09-01 2016-09-30")
	test("2016OCT", "2016-10-01 2016-10-31")
	test("2016NOV", "2016-11-01 2016-11-30")
	test("2016DEC", "2016-12-01 2016-12-31")
	//
	// month name and year (for a single year)
	test("JAN2016", "2016-01-01 2016-01-31")
	test("FEB2016", "2016-02-01 2016-02-29")
	test("MAR2016", "2016-03-01 2016-03-31")
	test("APR2016", "2016-04-01 2016-04-30")
	test("MAY2016", "2016-05-01 2016-05-31")
	test("JUN2016", "2016-06-01 2016-06-30")
	test("JUL2016", "2016-07-01 2016-07-31")
	test("AUG2016", "2016-08-01 2016-08-31")
	test("SEP2016", "2016-09-01 2016-09-30")
	test("OCT2016", "2016-10-01 2016-10-31")
	test("NOV2016", "2016-11-01 2016-11-30")
	test("DEC2016", "2016-12-01 2016-12-31")
	//
	// non-leap year should have February 28th
	test("2015 Feb", "2015-02-01 2015-02-28")
	test("2015 FEB", "2015-02-01 2015-02-28")
	test("2015 February", "2015-02-01 2015-02-28")
	test("2015/Feb", "2015-02-01 2015-02-28")
	test("2015/FEB", "2015-02-01 2015-02-28")
	test("2015/February", "2015-02-01 2015-02-28")
	test(`2015\Feb`, "2015-02-01 2015-02-28")
	test(`2015\FEB`, "2015-02-01 2015-02-28")
	test(`2015\February`, "2015-02-01 2015-02-28")
	test("2015Feb", "2015-02-01 2015-02-28")
	test("2015-Feb", "2015-02-01 2015-02-28")
	test("2015FEB", "2015-02-01 2015-02-28")
	test("2015-FEB", "2015-02-01 2015-02-28")
	test("2015February", "2015-02-01 2015-02-28")
	test("2015-February", "2015-02-01 2015-02-28")
	test("Feb 2015", "2015-02-01 2015-02-28")
	test("FEB 2015", "2015-02-01 2015-02-28")
	test("Feb/2015", "2015-02-01 2015-02-28")
	test("FEB/2015", "2015-02-01 2015-02-28")
	test(`Feb\2015`, "2015-02-01 2015-02-28")
	test(`FEB\2015`, "2015-02-01 2015-02-28")
	test("Feb2015", "2015-02-01 2015-02-28")
	test("Feb-2015", "2015-02-01 2015-02-28")
	test("FEB2015", "2015-02-01 2015-02-28")
	test("FEB-2015", "2015-02-01 2015-02-28")
	test("February 2015", "2015-02-01 2015-02-28")
	test("February/2015", "2015-02-01 2015-02-28")
	test(`February\2015`, "2015-02-01 2015-02-28")
	test("February2015", "2015-02-01 2015-02-28")
	test("February-2015", "2015-02-01 2015-02-28")
	//
	// leap year should have February 29th
	test("2016 Feb", "2016-02-01 2016-02-29")
	test("2016 FEB", "2016-02-01 2016-02-29")
	test("2016 February", "2016-02-01 2016-02-29")
	test("2016/Feb", "2016-02-01 2016-02-29")
	test("2016/FEB", "2016-02-01 2016-02-29")
	test("2016/February", "2016-02-01 2016-02-29")
	test(`2016\Feb`, "2016-02-01 2016-02-29")
	test(`2016\FEB`, "2016-02-01 2016-02-29")
	test(`2016\February`, "2016-02-01 2016-02-29")
	test("2016Feb", "2016-02-01 2016-02-29")
	test("2016-Feb", "2016-02-01 2016-02-29")
	test("2016FEB", "2016-02-01 2016-02-29")
	test("2016-FEB", "2016-02-01 2016-02-29")
	test("2016February", "2016-02-01 2016-02-29")
	test("2016-February", "2016-02-01 2016-02-29")
	test("Feb 2016", "2016-02-01 2016-02-29")
	test("FEB 2016", "2016-02-01 2016-02-29")
	test("Feb/2016", "2016-02-01 2016-02-29")
	test("FEB/2016", "2016-02-01 2016-02-29")
	test(`Feb\2016`, "2016-02-01 2016-02-29")
	test(`FEB\2016`, "2016-02-01 2016-02-29")
	test("Feb2016", "2016-02-01 2016-02-29")
	test("Feb-2016", "2016-02-01 2016-02-29")
	test("FEB2016", "2016-02-01 2016-02-29")
	test("FEB-2016", "2016-02-01 2016-02-29")
	test("February 2016", "2016-02-01 2016-02-29")
	test("February/2016", "2016-02-01 2016-02-29")
	test(`February\2016`, "2016-02-01 2016-02-29")
	test("February2016", "2016-02-01 2016-02-29")
	test("February-2016", "2016-02-01 2016-02-29")
	//
	// numeric year and month
	test("2016-01", "2016-01-01 2016-01-31")
	test("01-2016", "2016-01-01 2016-01-31")
	//
	// single date
	test("2016-01-02", "2016-01-02 2016-01-02")
	test("2016-DEC-31", "2016-12-31 2016-12-31")
	test("Dec. 31 2016", "2016-12-31 2016-12-31")
	//
	// single date (note DMY)
	test("02-01-2016", "2016-01-02 2016-01-02")
	//
	// invalid dates
	TTrue(t, DateRangeOf("").From.IsZero())
	TTrue(t, DateRangeOf("").To.IsZero())
	TTrue(t, DateRangeOf("1").From.IsZero())
	TTrue(t, DateRangeOf("1").To.IsZero())
	TTrue(t, DateRangeOf("abc").From.IsZero())
	TTrue(t, DateRangeOf("abc").To.IsZero())
} //                                                      Test_date_DateRangeOf_

// go test --run Test_date_DaysInMonth_
func Test_date_DaysInMonth_(t *testing.T) {
	TBegin(t)
	//
	// TODO: DaysInMonth() decl
	//
	// not a leap year
	TEqual(t, DaysInMonth(2015, 2), (28))
	TEqual(t, DaysInMonth(2016, 1), (31))
	//
	// leap year
	TEqual(t, DaysInMonth(2016, 2), (29))
	TEqual(t, DaysInMonth(2016, 3), (31))
	TEqual(t, DaysInMonth(2016, 4), (30))
	TEqual(t, DaysInMonth(2016, 5), (31))
	TEqual(t, DaysInMonth(2016, 6), (30))
	TEqual(t, DaysInMonth(2016, 7), (31))
	TEqual(t, DaysInMonth(2016, 8), (31))
	TEqual(t, DaysInMonth(2016, 9), (30))
	TEqual(t, DaysInMonth(2016, 10), (31))
	TEqual(t, DaysInMonth(2016, 11), (30))
	TEqual(t, DaysInMonth(2016, 12), (31))
} //                                                      Test_date_DaysInMonth_

// go test --run Test_date_IsDateOnly_
func Test_date_IsDateOnly_(t *testing.T) {
	TBegin(t)
	//
	// TODO: IsDateOnly decl
	//
	// false if any of the time fields are not zero
	TFalse(t, IsDateOnly(time.Date(
		2009, time.November, 4, 18, 46, 55, 0, time.UTC,
	)))
	TFalse(t, IsDateOnly(time.Date(
		2009, time.November, 4, 1, 0, 0, 0, time.UTC,
	)))
	TFalse(t, IsDateOnly(time.Date(
		2009, time.November, 4, 0, 1, 0, 0, time.UTC,
	)))
	TFalse(t, IsDateOnly(time.Date(
		2009, time.November, 4, 0, 0, 1, 0, time.UTC,
	)))
	TFalse(t, IsDateOnly(time.Date(
		2009, time.November, 4, 0, 0, 0, 1, time.UTC,
	)))
	// only true when all time fields are zero
	TTrue(t, IsDateOnly(time.Date(
		2009, time.November, 4, 0, 0, 0, 0, time.UTC,
	)))
} //                                                       Test_date_IsDateOnly_

// go test --run Test_date_IsDate_
func Test_date_IsDate_(t *testing.T) {
	TBegin(t)
	//
	// IsDate(value interface{}) bool
	//
	test := func(value interface{}, expect bool) {
		got := IsDate(value)
		if got != expect {
			t.Errorf("IsDate(%s) returned %v instead of %v",
				value, got, expect)
		}
	}
	// false
	test("", false)
	test(123, false)
	test("30-Feb-2017", true)
	//
	// true
	test("01-Jun-2015", true)
	test("01/May/2017", true)
	test("1-December-2015", true)
	test("12/05/2017", true)
	test("15-05-2015", true)
	test("15-Jul-2015", true)
	test("15-Sep-2015", true)
	test("2011-02-03", true)
	test("3-August-1999", true)
	test("4-February-15", true)
	test(time.Now(), true)
} //                                                           Test_date_IsDate_

// go test --run Test_date_ParseDate_
func Test_date_ParseDate_(t *testing.T) {
	TBegin(t)
	//
	// TODO: declaration comment
	//
	test := func(input string, expectY, expectM, expectD int) {
		y, m, d := ParseDate(input)
		if y != expectY || m != expectM || d != expectD {
			t.Errorf("ParseDate(%q)"+
				" returned (%d %d %d) instead of (%d %d %d)",
				input, y, m, d, expectY, expectM, expectD)
		}
	}
	test("", 0, 0, 0)
	test("01/May/2017", 2017, 5, 1)
	test("01-Jun-2015", 2015, 6, 1)
	test("15-Jul-2015", 2015, 7, 15)
	test("15-Sep-2015", 2015, 9, 15)
	test("12/05/2017", 2017, 5, 12)
	test("15-05-2015", 2015, 5, 15)
	test("4-February-15", 2015, 2, 4)
	test("1-December-2015", 2015, 12, 1)
	test("3-August-1999", 1999, 8, 3)
	//
	// yyyy-mm-dd
	test("2000.07.05", 2000, 7, 5)
	test("2000.7.05", 2000, 7, 5)
	test("2000.7.5", 2000, 7, 5)
	test("2000/07/05", 2000, 7, 5)
	test("2000/7/05", 2000, 7, 5)
	test("2000/7/5", 2000, 7, 5)
	test("2000-07-05", 2000, 7, 5)
	test("2000-7-05", 2000, 7, 5)
	test("2000-7-5", 2000, 7, 5)
	test("2017.05.01", 2017, 5, 1)
	test("2017.07.08", 2017, 7, 8)
	test("2017.07.12", 2017, 7, 12)
	test("2017.5.01", 2017, 5, 1)
	test("2017.5.1", 2017, 5, 1)
	test("2017.7.08", 2017, 7, 8)
	test("2017.7.12", 2017, 7, 12)
	test("2017.7.8", 2017, 7, 8)
	test("2017/05/01", 2017, 5, 1)
	test("2017/07/08", 2017, 7, 8)
	test("2017/07/12", 2017, 7, 12)
	test("2017/5/01", 2017, 5, 1)
	test("2017/5/1", 2017, 5, 1)
	test("2017/7/08", 2017, 7, 8)
	test("2017/7/12", 2017, 7, 12)
	test("2017/7/8", 2017, 7, 8)
	test("2017-05-01", 2017, 5, 1)
	test("2017-07-08", 2017, 7, 8)
	test("2017-07-12", 2017, 7, 12)
	test("2017-5-01", 2017, 5, 1)
	test("2017-5-1", 2017, 5, 1)
	test("2017-7-08", 2017, 7, 8)
	test("2017-7-12", 2017, 7, 12)
	test("2017-7-8", 2017, 7, 8)
	test("9999.12.31", 9999, 12, 31)
	test("9999/12/31", 9999, 12, 31)
	test("9999-12-31", 9999, 12, 31)
	//
	// dates that don't meet format requirements
	test(" 1/1/1/ ", 0, 0, 0)
	test("100-01-02", 0, 0, 0)
	test("100-1-2", 0, 0, 0)
	test("abc", 0, 0, 0)
} //                                                        Test_date_ParseDate_

// go test --run Test_date_StringDateDMY_
func Test_date_StringDateDMY_(t *testing.T) {
	TBegin(t)
	//
	// TODO: declaration comment
	//
	test := func(input, expect string) {
		got := StringDateDMY(input)
		if got != expect {
			t.Errorf("StringDateDMY(%q) returned %q instead of %q",
				input, got, expect)
		}
	}
	test("", "")
	test("2017-01-01", "1/Jan/2017")
	test("15/Feb/2017", "15/Feb/2017")
} //                                                    Test_date_StringDateDMY_

// go test --run Test_date_StringDateYMD_
func Test_date_StringDateYMD_(t *testing.T) {
	TBegin(t)
	//
	// TODO: declaration comment
	//
	test := func(input, expect string) {
		got := StringDateYMD(input)
		if got != expect {
			t.Errorf("StringDateYMD(%q) returned %q instead of %q",
				input, got, expect)
		}
	}
	test("", "")
	test("01/May/2017", "2017-05-01")
	test("2017-05-01", "2017-05-01")
} //                                                    Test_date_StringDateYMD_

// go test --run Test_date_Timestamp_
func Test_date_Timestamp_(t *testing.T) {
	TBegin(t)
	//
	// TODO: declaration comment
	//
	s := Timestamp()
	// Expected format: YYYY-MM-DD hh:mm:ss
	//                  0123456789012345678
	TTrue(t, len(s) == 19)
	//
	TTrue(t, s[0:4] >= "0000" && s[0:4] <= "9999") // year
	TTrue(t, s[5:7] >= "01" && s[5:7] <= "12")     // month
	TTrue(t, s[8:10] >= "01" && s[8:10] <= "31")   // day
	//
	TTrue(t, s[10:11] == " ") // space between date and time
	//
	TTrue(t, s[11:13] >= "00" && s[11:13] <= "23") // hour
	TTrue(t, s[14:16] >= "00" && s[14:16] <= "59") // minute
	TTrue(t, s[17:19] >= "00" && s[17:19] <= "59") // second
} //                                                        Test_date_Timestamp_

//end
