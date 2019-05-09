// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-09 18:09:23 8DFDDA                               zr/[calendar.go]
// -----------------------------------------------------------------------------

package zr

// # Types
//   Calendar struct
//   calendarDay struct
//   calendarMonth struct
//   calendarWeekdaysEN = []string
//
// # Methods (ob *Calendar)
//   ) AddMonth(year int, month time.Month) error
//   ) HasMonth(year int, month time.Month) bool
//   ) Set(date, value interface{})
//   ) String() string
//
// # Internal Methods/Functions
//  (*Calendar) firstWeekday(year int, month time.Month) time.Weekday
//  (ob *Calendar) getMonth(year int, month time.Month) *calendarMonth

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

// -----------------------------------------------------------------------------
// # Types

// Calendar provides logic for generating
// calendar grids from dates and values.
type Calendar struct {
	months []calendarMonth
} //                                                                    Calendar

// calendarDay holds the calendar entry for a single day
// 'day' is day of the month (the date)
// 'val' is the value shown on the specified date
type calendarDay struct {
	day int
	val interface{}
} //                                                                 calendarDay

// calendarMonth holds the data for a single month,
// with days arranged in a 6 x 7 grid.
type calendarMonth struct {
	year  int
	month time.Month
	cells [6][7]calendarDay
} //                                                               calendarMonth

// calendarWeekdaysEN defines English weekday names
var calendarWeekdaysEN = []string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
} //                                                          calendarWeekdaysEN

// -----------------------------------------------------------------------------
// # Methods (ob *Calendar)

// AddMonth adds a month to the calendar, without setting any values.
// The year must range from 1 to 9999.
func (ob *Calendar) AddMonth(year int, month time.Month) error {
	if year < 1 || year > 9999 {
		return Error(EInvalidArg, "^year", ":", year)
	}
	// check if month was added already
	if ob.HasMonth(year, month) {
		return Error("Month", month, year, "already added")
	}
	mth := calendarMonth{year: year, month: month}
	weekday1 := ob.firstWeekday(year, month)
	last := DaysInMonth(year, month)
	day := 1
loop:
	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			if day == 1 && col < int(weekday1)-1 {
				continue
			}
			mth.cells[row][col].day = day
			if day == last {
				break loop
			}
			day++
		}
	}
	ob.months = append(ob.months, mth)
	return nil
} //                                                                    AddMonth

// HasMonth returns true if the month specified by
// year and month has been added to the calendar
func (ob *Calendar) HasMonth(year int, month time.Month) bool {
	for _, m := range ob.months {
		if m.year == year && m.month == month {
			return true
		}
	}
	return false
} //                                                                    HasMonth

// Set assigns the specified value to the specified date.
// It automatically converts 'date' to time.Time
func (ob *Calendar) Set(date, value interface{}) {
	dt := DateOf(date)
	year, month := dt.Year(), dt.Month()
	//
	// if not existing already, add the month to the calendar
	if !ob.HasMonth(year, month) {
		ob.AddMonth(year, month)
	}
	// get a pointer to the calendar
	mth := ob.getMonth(year, month)
	if mth == nil {
		// ^ this condition should never occur
		// (no warning here: a warning is already logged by getMonth)
		return
	}
	// find the specified day and set its value
	day := dt.Day()
	found := false
	for row := 0; !found && row < 6; row++ {
		for col := 0; !found && col < 7; col++ {
			if mth.cells[row][col].day == day {
				mth.cells[row][col].val = value
				found = true
			}
		}
	}
} //                                                                         Set

// String returns the calendar as a text string
// and implements the fmt.Stringer interface.
//
// The output may contain multiple months, in which
// case the months are arranged in ascending order.
//
// See the sample output in the body of the function.
//
func (ob *Calendar) String() string {
	// Sample output:
	//
	// 2018 FEBRUARY
	// *--------------------------------------------------------------*
	// |  Mon   |  Tue   |  Wed   |  Thu   |  Fri   |  Sat   |  Sun   |
	// |--------|--------|--------|--------|--------|--------|--------|
	// |        |        |        | 1      | 2      | 3      | 4      |
	// |        |        |        |   8.44 |   7.55 |   6.66 |   5.77 |
	// |--------|--------|--------|--------|--------|--------|--------|
	// | 5      | 6      | 7      | 8      | 9      | 10     | 11     |
	// |   4.88 |   3.99 |   2.15 |   1.54 |      0 |      1 |      2 |
	// |--------|--------|--------|--------|--------|--------|--------|
	// | 12     | 13     | 14     | 15     | 16     | 17     | 18     |
	// |      3 |      4 |      5 |      6 |      7 |      8 |      9 |
	// |--------|--------|--------|--------|--------|--------|--------|
	// | 19     | 20     | 21     | 22     | 23     | 24     | 25     |
	// |     10 |        |        |        |        |        |        |
	// |--------|--------|--------|--------|--------|--------|--------|
	// | 26     | 27     | 28     |        |        |        |        |
	// |        |        |        |        |        |        |        |
	// |--------|--------|--------|--------|--------|--------|--------|
	// |        |        |        |        |        |        |        |
	// |        |        |        |        |        |        |        |
	// *--------------------------------------------------------------*
	// 95.98
	const EDGE = "*"
	const HDIV = "-"
	const VDIV = "|"
	const CELLWIDTH = 8
	//
	var retBuf bytes.Buffer
	ws := func(a ...string) {
		for _, s := range a {
			retBuf.WriteString(s)
		}
	}
	ws("\n")
	blank := strings.Repeat(" ", CELLWIDTH)
	wdayFmt := fmt.Sprintf("  %%-%ds", CELLWIDTH-2)
	dayFmt := fmt.Sprintf(" %%-%dd", CELLWIDTH-1)
	valFmt := fmt.Sprintf(" %%%dv ", CELLWIDTH-2)
	//
	// draws the outer (top or bottom) horizontal divider
	outerHLine := func() {
		ws(EDGE)
		for i := 0; i < 7; i++ {
			if i > 0 {
				ws(HDIV)
			}
			ws(strings.Repeat(HDIV, CELLWIDTH))
		}
		ws(EDGE, "\n")
	}
	// draws the inner horizontal divider
	innerHLine := func() {
		for i := 0; i < 7; i++ {
			ws(VDIV, strings.Repeat(HDIV, CELLWIDTH))
		}
		ws(VDIV, "\n")
	}
	// formats numbers
	numStr := func(val float64) string {
		ret := fmt.Sprintf("%5.2f", val)
		if strings.Contains(ret, ".") {
			ret = strings.TrimRight(ret, "0")
			ret = strings.TrimRight(ret, ".")
		}
		return ret
	}
	for _, mth := range ob.months {
		//
		// month heading
		ws(strings.ToUpper(fmt.Sprintf("%d %v", mth.year, mth.month)), "\n")
		outerHLine()
		//
		// weekday names
		for i := 0; i < 7; i++ {
			ws(VDIV)
			ws(fmt.Sprintf(wdayFmt, calendarWeekdaysEN[i][:3]))
		}
		ws(VDIV, "\n")
		var sum float64
		//
		// draw the grid
		for row := 0; row < 6; row++ {
			innerHLine()
			//
			// days on current row
			for col := 0; col < 7; col++ {
				ws(VDIV)
				day := mth.cells[row][col].day
				if day == 0 {
					ws(blank)
				} else {
					ws(fmt.Sprintf(dayFmt, day))
				}
			}
			ws(VDIV, "\n")
			//
			// values on current row
			for col := 0; col < 7; col++ {
				ws(VDIV)
				val := mth.cells[row][col].val
				if val == nil {
					ws(blank)
				} else {
					if val, ok := val.(float64); ok {
						s := numStr(val)
						sum += val
						ws(fmt.Sprintf(valFmt, s))
					} else {
						ws(fmt.Sprintf(valFmt, val))
					}
				}
			}
			ws(VDIV, "\n")
		}
		outerHLine()
		ws(numStr(sum), "\n\n")
	} // mth
	return retBuf.String()
} //                                                                      String

// -----------------------------------------------------------------------------
// # Internal Methods/Functions

// firstWeekday returns the day of week on the first of the given month
func (*Calendar) firstWeekday(year int, month time.Month) time.Weekday {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return date.Weekday()
} //                                                                firstWeekday

// getMonth returns a pointer to the month specified by 'year' and
// 'month', or nil if the month has not been added to this calendar
func (ob *Calendar) getMonth(year int, month time.Month) *calendarMonth {
	for i, m := range ob.months {
		if m.year == year && m.month == month {
			return &ob.months[i]
		}
	}
	Error(ENotFound, "month", month, year)
	return nil
} //                                                                    getMonth

//end
