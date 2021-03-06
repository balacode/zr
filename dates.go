// -----------------------------------------------------------------------------
// ZR Library                                                      zr/[dates.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # English Month Names
//   MonthNamesEN = []string
//
// # Day-Month-Year Date Formats
//   DateFormatsDMY = []struct
//
// # Date Regular Expressions
//   ISODateTimeEx
//   ISODateEx
//   ISOTimeEx
//
// # DateRange
//   DateRange struct
//   (ob DateRange) IsNull() bool
//   (ob DateRange) String() string
//
// # Functions
//   DateOf(value interface{}) time.Time
//   DateRangeOf(s string) DateRange
//   DayMth(value interface{}) string
//   DaysInMonth(year int, month time.Month) int
//   FormatDateEN(format string, date time.Time) string
//   IsDate(value interface{}) bool
//   IsDateOnly(tm time.Time) bool
//   MonthNameEN(monthNo int, shortName ...bool) string
//   MonthNumberEN(monthName string) int
//   MthYear(value interface{}) string
//   ParseDate(s string) (year, month, day int)
//   StringDateDMY(s string) string
//   StringDateYMD(s string) string
//   StringYear(value interface{}) string
//   Timestamp(optWithMS ...bool) string
//   YMD(t time.Time) string
//
// # Private Functions
//   stringDate(value interface{}, format string) string

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Tip: Go uses the following date format reference:
//      Mon Jan 2 15:04:05 -0700 MST 2006

// -----------------------------------------------------------------------------
// # English Month Names

// MonthNamesEN is a string array of English month names.
var MonthNamesEN = []string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
} //                                                                MonthNamesEN

// -----------------------------------------------------------------------------
// # Day-Month-Year Date Formats

// DateFormatsDMY _ _
var DateFormatsDMY = []struct {
	Pat string // regular expression pattern to match
	In  string // use this date format to parse inupt
	Out string // use this date format to output
}{
	{
		Pat: `^\d{4}-\d{1,2}-\d{1,2} \d{1,2}:\d{1,2}:\d{1,2}$`,
		In:  "2006-01-02 15:04:05",
		Out: "2/Jan/2006 03:04pm",
	},
	{
		Pat: `^\d{4}-\d{1,2}-\d{1,2}$`,
		In:  "2006-01-02",
		Out: "2/Jan/2006",
	},
} //                                                              DateFormatsDMY

// -----------------------------------------------------------------------------
// # Date Regular Expressions

var (
	// ISODateEx is a regular expression that matches
	// date strings in "YYYY-MM-DD" format.
	ISODateEx = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d$`)

	// ISODateTimeEx is a regular expression that matches
	// date and time strings in "YYYY-MM-DDThh:mm:ss"
	// format or "YYYY-MM-DD hh:mm:ss" format.
	ISODateTimeEx = regexp.MustCompile(
		`^\d\d\d\d-\d\d-\d\d[ T]\d\d:\d\d:\d\d$`)

	// ISODateTimeEx is a regular expression that matches
	// time strings in "hh:mm:ss" format.
	ISOTimeEx = regexp.MustCompile(`^\d\d:\d\d:\d\d$`)
)

// -----------------------------------------------------------------------------
// # DateRange

// DateRange represents a time period.
type DateRange struct {
	From time.Time
	To   time.Time
} //                                                                   DateRange

// IsNull returns true if the date range doesn't represent any period.
func (ob DateRange) IsNull() bool {
	return ob.From.IsZero() || ob.To.IsZero()
} //                                                                      IsNull

// String returns a string representation of the DateRange structure
// and implements the fmt.Stringer interface.
func (ob DateRange) String() string {
	return ob.From.Format("2006-01-02") + " " + ob.To.Format("2006-01-02")
} //                                                                      String

// -----------------------------------------------------------------------------
// # Functions

// DateOf converts any string-like value to time.Time without
// returning an error if the conversion failed, in which case
// it logs an error and returns a zero-value time.Time.
//
// If value is a zero-length string, returns a zero-value
// time.Time but does not log a warning.
//
// It also accepts a time.Time as input.
//
// In both cases the returned Time type will contain only
// the date part without the time or time zone components.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
func DateOf(value interface{}) time.Time {
	switch v := value.(type) {
	case time.Time: // remove the time component from dates
		{
			return time.Date(v.Year(), v.Month(), v.Day(),
				0, 0, 0, 0, time.UTC)
		}
	case string:
		{
			if v == "" {
				return time.Time{}
			}
			if len(v) > 10 {
				v = v[:10]
			}
			ret, err := time.Parse(time.RFC3339, v+"T00:00:00Z")
			if err != nil || ret.IsZero() {
				if err != nil {
					mod.Error(err)
				}
				return time.Time{}
			}
			return DateOf(ret)
		}
	case *string:
		if v != nil {
			return DateOf(*v)
		}
	}
	mod.Error("Can not convert", reflect.TypeOf(value), "to int:", value)
	return time.Time{}
} //                                                                      DateOf

// DateRangeOf creates and returns a DateRange structure from a string. _ _
func DateRangeOf(s string) DateRange {
	// pre-format
	s = strings.ToUpper(strings.TrimSpace(s))
	for _, sep := range []string{" ", ".", "/", "\\", "_"} {
		for strings.Contains(s, sep) {
			s = strings.ReplaceAll(s, sep, "-")
		}
	}
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	for strings.Contains(s, "-TO-") {
		s = strings.ReplaceAll(s, "-TO-", "~")
	}
	var (
		y1, m1, d1 = 0, time.January, 1
		y2, m2, d2 = 0, time.December, 31
	)
	// date range
	if i := strings.Index(s, "~"); i != -1 {
		var (
			r1 = DateRangeOf(s[:i])
			r2 = DateRangeOf(s[i+1:])
		)
		y1, m1, d1 = r1.From.Year(), r1.From.Month(), r1.From.Day()
		y2, m2, d2 = r2.To.Year(), r2.To.Month(), r2.To.Day()
	}
	// year only
	if len(s) == 4 {
		dt, err := time.Parse("2006", s)
		if err == nil {
			y1, m1, d1 = dt.Year(), 1, 1
			y2, m2, d2 = dt.Year(), 12, 31
		}
	}
	// month only
	if y1 == 0 {
		for _, format := range []string{"Jan", "January"} {
			dt, err := time.Parse(format, s)
			if err == nil {
				now := time.Now()
				y1, m1, d1 = now.Year(), dt.Month(), 1
				y2, m2, d2 = now.Year(), dt.Month(),
					DaysInMonth(now.Year(), dt.Month())
				break
			}
		}
	}
	// complete date
	if y1 == 0 {
		for _, format := range []string{
			"2-1-2006",
			"2-Jan-2006",
			"2-January-2006",
			"2006-01-02",
			"2006-2-Jan",
			"2006-2-January",
			"2006-Jan-2",
			"2006-January-2",
			"Jan-2-2006",
			"Jan-2006-2",
		} {
			dt, err := time.Parse(format, s)
			if err == nil {
				y1, m1, d1 = dt.Year(), dt.Month(), dt.Day()
				y2, m2, d2 = dt.Year(), dt.Month(), dt.Day()
				break
			}
		}
	}
	// month and year
	if y1 == 0 {
		for _, format := range []string{
			"1-2006",
			"2006-1",
			"2006-Jan",
			"2006-January",
			"2006Jan",
			"2006January",
			"Jan-2006",
			"Jan2006",
			"January-2006",
			"January2006",
		} {
			dt, err := time.Parse(format, s)
			if err == nil {
				y1, m1, d1 = dt.Year(), dt.Month(), 1
				y2, m2, d2 = dt.Year(), dt.Month(),
					DaysInMonth(dt.Year(), dt.Month())
				break
			}
		}
	}
	// day and month
	if y1 == 0 {
		for _, format := range []string{
			"2-Jan",
			"2-January",
			"2Jan",
			"2January",
			"Jan-2",
			"Jan2",
			"January-2",
			"January2",
		} {
			dt, err := time.Parse(format, s)
			if err == nil {
				now := time.Now()
				y1, m1, d1 = now.Year(), dt.Month(), dt.Day()
				y2, m2, d2 = now.Year(), dt.Month(), dt.Day()
				break
			}
		}
	}
	if y1 == 0 {
		return DateRange{From: time.Time{}, To: time.Time{}}
	}
	return DateRange{From: time.Date(y1, m1, d1, 0, 0, 0, 0, time.UTC),
		To: time.Date(y2, m2, d2, 23, 59, 59, 999999999, time.UTC)}
} //                                                                 DateRangeOf

// DayMth returns a day-and-month string of the format
// "d mmm" when given a time.Time value or a date string.
func DayMth(value interface{}) string {
	return stringDate(value, "2 Jan")
} //                                                                      DayMth

// DaysInMonth returns the number of days in the specified year and month.
// (If year is less than 1 or greater than 9999, returns 0)
func DaysInMonth(year int, month time.Month) int {
	mth := int(month)
	if year < 0 || year > 9999 || mth < 1 || mth > 12 {
		return 0
	}
	if year < 50 {
		year += 2000
	} else if year < 100 {
		year += 1900
	}
	// a year is a leap year if it is divisible by 4 but not by 100,
	// unless it is also divisible by 400, which makes it a leap year.
	if mth == 2 && (((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)) {
		return 29
	}
	mdays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	return mdays[mth-1]
} //                                                                 DaysInMonth

// FormatDateEN formats date using the specified format.
// Uses English language names, hence the 'EN' suffix.
func FormatDateEN(format string, date time.Time) string {
	var (
		year  = date.Year()
		month = int(date.Month())
		day   = date.Day()
		ret   = format
	)
	// TODO: implement day formats in FormatDateEN()
	// ReplaceWord(ret, "dddd", "d", MatchCase)
	// ReplaceWord(ret, "Dddd", "f", MatchCase)
	// ReplaceWord(ret, "DDDD", "h", MatchCase)
	// ReplaceWord(ret, "ddd", "c", MatchCase)
	// ReplaceWord(ret, "Ddd", "e", MatchCase)
	// ReplaceWord(ret, "DDD", "g", MatchCase)
	/*TODO: create new code that can use this function
	  change := func(word string, caseMode CaseMode, to string) {
	      has := Contains
	      if caseMode == IgnoreCase {
	          has = ContainsI
	      }
	      if has(ret, word) {
	          ret = ReplaceWord(ret, word, to, caseMode)
	      }
	  }
	*/
	// day (2 digit)
	if ContainsI(ret, "dd") {
		ret = ReplaceWord(ret, "dd", fmt.Sprintf("%02d", day), IgnoreCase)
	}
	// day (1 or 2 digits)
	if ContainsI(ret, "d") {
		ret = ReplaceWord(ret, "d", String(day), IgnoreCase)
	}
	// month (full)
	if strings.Contains(ret, "MMMM") {
		ret = ReplaceWord(ret, "MMMM",
			strings.ToUpper(MonthNameEN(month)), MatchCase)
	}
	if strings.Contains(ret, "Mmmm") {
		ret = ReplaceWord(ret, "Mmmm", MonthNameEN(month), MatchCase)
	}
	if strings.Contains(ret, "mmmm") {
		ret = ReplaceWord(ret, "mmmm",
			strings.ToLower(MonthNameEN(month)), MatchCase)
	}
	// month (3 letters)
	if strings.Contains(ret, "MMM") {
		ret = ReplaceWord(ret, "MMM",
			strings.ToUpper(First(MonthNameEN(month), 3)), MatchCase)
	}
	if strings.Contains(ret, "Mmm") {
		ret = ReplaceWord(ret, "Mmm", First(MonthNameEN(month), 3), MatchCase)
	}
	if strings.Contains(ret, "mmm") {
		ret = ReplaceWord(ret, "mmm",
			strings.ToLower(First(MonthNameEN(month), 3)), MatchCase)
	}
	// month (2 digits)
	if ContainsI(ret, "mm") {
		ret = ReplaceWord(ret, "mm", fmt.Sprintf("%02d", month), IgnoreCase)
	}
	// month (1 or 2 digits)
	if ContainsI(ret, "m") {
		ret = ReplaceWord(ret, "m", String(month), IgnoreCase)
	}
	// year (4 digits)
	if ContainsI(ret, "YYYY") {
		ret = ReplaceWord(ret, "YYYY", fmt.Sprintf("%04d", year), IgnoreCase)
	}
	// year (2 digits)
	if ContainsI(ret, "YY") {
		ret = ReplaceWord(ret, "YY", Last(String(year), 2), IgnoreCase)
	}
	return ret
} //                                                                FormatDateEN

// IsDate returns true if the specified value can be converted to a date.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
func IsDate(value interface{}) bool {
	ret, reason := func(value interface{}) (bool, int) {
		switch v := value.(type) {
		case time.Time:
			{
				return true, 1
			}
		case string:
			{
				if v == "" {
					return false, 2
				}
				{ // try to use time.Parse() to quickly parse 'yyyy-mm-dd' dates
					s := v
					if len(s) > 10 {
						s = s[:10]
					}
					parsed, err := time.Parse(time.RFC3339, s+"T00:00:00Z")
					if err == nil && !parsed.IsZero() {
						return true, 3
					}
				}
				// if time.Parse() can't parse the date, try using ParseDate()
				y, m, d := ParseDate(v)
				if y == 0 || m == 0 || d == 0 {
					return false, 4
				}
				return true, 5
			}
		}
		return false, 7
	}(value)
	VL("IsDate(", value, ") returned ", ret, " ", reason)
	return ret
} //                                                                      IsDate

// IsDateOnly returns true if 'tm' does not have a time portion,
// I.e. the hour, minute, second and nanosecond are all zero
func IsDateOnly(tm time.Time) bool {
	return tm.Hour() == 0 && tm.Minute() == 0 && tm.Second() == 0 &&
		tm.Nanosecond() == 0
} //                                                                  IsDateOnly

// MonthNameEN returns the English month name given a month number.
// E.g. "January", "February", etc.
// Uses English language names, hence the 'EN' suffix.
func MonthNameEN(monthNo int, shortName ...bool) string {
	isShortName := len(shortName) > 0 && shortName[0]
	if monthNo < 1 || monthNo > 12 {
		mod.Error("Month", monthNo, "out of range")
		return ""
	}
	ret := MonthNamesEN[monthNo-1]
	if isShortName {
		ret = ret[:3]
	}
	return ret
} //                                                                 MonthNameEN

// MonthNumberEN returns a month number from 1 to 12, given an English
// month name. Accepts either a full month name like 'December',
// or a 3-character string like 'Dec'. The case is not important.
// If the string is not a month name, returns zero.
// Uses English language names, hence the 'EN' suffix.
func MonthNumberEN(monthName string) int {
	monthName = strings.ToUpper(strings.TrimSpace(monthName))
	for i, s := range MonthNamesEN {
		s = strings.ToUpper(s)
		if monthName == s || monthName == s[:3] {
			return i + 1
		}
	}
	return 0
} //                                                               MonthNumberEN

// MthYear returns a date string of the format "Mmm yyyy" when given
// a time.Time value or a date string.
func MthYear(value interface{}) string {
	return stringDate(value, "Jan 2006")
} //                                                                     MthYear

// ParseDate reads a date string and returns the year, month and day number.
func ParseDate(s string) (year, month, day int) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, 0, 0
	}
	change := func(from, to string) {
		if strings.Contains(s, from) {
			s = strings.ReplaceAll(s, from, to)
		}
	}
	change("-", "/")
	change(".", "/")
	change("\\", "/")
	parts := strings.Split(s, "/")
	if len(parts) != 3 {
		return 0, 0, 0
	}
	patterns := []struct {
		pat     string
		y, m, d int
	}{
		{`^\d{1,2}/[[:alpha:]]{3,9}/\d{1,4}$`, 2, 1, 0}, // dd/mmm/yyyy
		{`^\d{1,2}/\d{1,2}/\d{1,4}$`, 2, 1, 0},          // dd/mm/yyyy
		{`^\d{4}/\d{1,2}/\d{1,2}$`, 0, 1, 2},            // yyyy/mm/dd
	}
	for _, t := range patterns {
		match, err := regexp.MatchString(t.pat, s)
		if err != nil {
			mod.Error("Failed matching^", s, "to^", t.pat, ":", err)
		}
		if !match {
			continue
		}
		year := Int(parts[t.y])
		if year < 1 || year > 9999 {
			year = 0
		}
		if year >= 1 && year < 100 {
			if year >= 70 {
				year += 1900
			} else {
				year += 2000
			}
		}
		month := Int(parts[t.m])
		if month < 1 || month > 12 {
			month = MonthNumberEN(parts[t.m])
		}
		day := Int(parts[t.d])
		if day < 1 || day > 31 {
			day = 0
		}
		if year == 0 || month == 0 || day == 0 {
			return 0, 0, 0
		}
		return year, month, day
	}
	return 0, 0, 0
} //                                                                   ParseDate

// StringDateDMY returns a short date using the "dd/mmm/yyyy" format.
// given an ISO-8601 formatted date string.
// E.g. given "2017-04-18" it returns "18/Apr/2017".
func StringDateDMY(s string) string {
	if s == "" || s == "null" {
		return ""
	}
	// try the faster parsing method
	tm, err := time.Parse(time.RFC3339, s+"T00:00:00Z")
	if err == nil {
		return fmt.Sprintf("%d/%s/%d",
			tm.Day(), tm.Month().String()[:3], tm.Year())
	}
	// if that didn't work, try the slower method
	y, m, d := ParseDate(s)
	if m == 0 || d == 0 {
		return ""
	}
	mth := MonthNameEN(m)[:3]
	return fmt.Sprintf("%d/%s/%04d", d, mth, y)
} //                                                               StringDateDMY

// StringDateYMD returns a short date using the "yyyy-mm-dd" format.
func StringDateYMD(s string) string {
	y, m, d := ParseDate(s)
	if m == 0 || d == 0 {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", y, m, d)
} //                                                               StringDateYMD

// StringYear _ _
func StringYear(value interface{}) string {
	if IsNumber(value) {
		year := Int(value)
		if year < 1 || year > 9999 {
			mod.Error("Numeric year out of range:", year)
		}
		return String(value)
	}
	return stringDate(String(value), "2006")
} //                                                                  StringYear

// Timestamp returns a timestamp string using the current local time.
// The format is: 'YYYY-MM-DD hh:mm:ss' (18 characters).
// I.e. date, and 24-hour time with seconds, and optional milliseconds
// The time zone is not included.
func Timestamp(optWithMS ...bool) string {
	var withMS bool
	switch len(optWithMS) {
	case 0:
		{
			// do nothing: false already
		}
	case 1:
		{
			withMS = optWithMS[0]
		}
	default:
		Error(EInvalidArg + ": Too many 'optWithMS' values")
		withMS = optWithMS[0]
	}
	if withMS {
		ret := time.Now().Round(0).String()
		if len(ret) > 24 {
			ret = ret[:24]
		}
		for len(ret) < 24 { // may need to add trailing zeros in milliseconds
			ret += "0"
		}
		return ret + " "
	}
	return time.Now().String()[:19]
	//
	// longer way to get the same result:
	// ret := time.Now().Format(time.RFC3339)[:19]
	// ret = strings.ReplaceAll(ret, "T", " ")
} //                                                                   Timestamp

// YMD returns a date using the 'yyyy-mm-dd' format.
func YMD(t time.Time) string {
	y, m, d := t.Year(), int(t.Month()), t.Day()
	if m == 0 || d == 0 {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", y, m, d)
} //                                                                         YMD

// -----------------------------------------------------------------------------
// # Private Functions

// stringDate _ _
func stringDate(value interface{}, format string) string {
	const erv = ""
	var date time.Time
	switch v := value.(type) {
	case time.Time:
		{
			date = v
		}
	case string:
		{
			if v == "" {
				return erv
			}
			if len(v) > 10 {
				v = v[:10]
			}
			parsed, err := time.Parse(time.RFC3339, v+"T00:00:00Z")
			if err != nil || parsed.IsZero() {
				if err != nil {
					mod.Error(EFailedParsing, "string^", v, ":", err)
				}
				return erv
			}
			date = parsed
		}
	default:
		mod.Error("Invalid value:", v)
		return erv
	}
	return date.Format(format)
} //                                                                  stringDate

// TODO: merge StringDateYMD() and YMD()

// TODO: DateOf(): add unit test for zero-length string

// end
