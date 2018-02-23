// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-23 18:08:54 42669E                               [zr/calendar.go]
// -----------------------------------------------------------------------------

package zr

import "bytes" // standard
import "fmt"   // standard
import "time"  // standard

// -----------------------------------------------------------------------------
// # Types

// cldrWeekdaysEN __
var cldrWeekdaysEN = []string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
} //                                                              cldrWeekdaysEN

// Calendar __
// a 6 x 7 grid
type Calendar struct {
	months []cldrMonth
} //                                                                    Calendar

// cldrMonth __
type cldrMonth struct {
	year  int
	month time.Month
	cells [6][7]cldrDay
} //                                                                   cldrMonth

// cldrDay __
type cldrDay struct {
	day int
	val interface{}
} //                                                                     cldrDay

// -----------------------------------------------------------------------------
// # Factory

// NewCalendar __
func NewCalendar() Calendar {
	return Calendar{}
} //                                                                 NewCalendar

// AddMonth __
func (ob *Calendar) AddMonth(year int, month time.Month) {
	//
	// check if month was added already
	if ob.HasMonth(year, month) {
		Error("Month", month, year, "already added")
		return
	}
	var mth = cldrMonth{year: year, month: month}
	var weekday1 = cldrFirstWeekday(year, month)
	var last = DaysInMonth(year, month)
	var day = 1
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

// Set __
func (ob *Calendar) Set(date, value interface{}) {
	var dt = DateOf(date)
	var year, month = dt.Year(), dt.Month()
	//
	// if not existing already, add the month to the calendar
	if !ob.HasMonth(year, month) {
		ob.AddMonth(year, month)
	}
	// get a pointer to the calendar
	var mth = ob.getMonth(year, month)
	if mth == nil {
		return // (a warning is logged by getMonth)
	}
	// find the specified day and set its value
	var day = dt.Day()
	var found = false
	for row := 0; !found && row < 6; row++ {
		for col := 0; !found && col < 7; col++ {
			if mth.cells[row][col].day == day {
				mth.cells[row][col].val = value
				found = true
			}
		}
	}
} //                                                                         Set

// String __
func (ob *Calendar) String() string {
	const EDGE = "*"
	const HDIV = "-"
	const VDIV = "|"
	const CELLWIDTH = 8
	//
	var retBuf bytes.Buffer
	var ws = func(a ...string) {
		for _, s := range a {
			retBuf.WriteString(s)
		}
	}
	ws(LF)
	var blank = str.Repeat(" ", CELLWIDTH)
	var wdayFmt = fmt.Sprintf("  %%-%ds", CELLWIDTH-2)
	var dayFmt = fmt.Sprintf(" %%-%dd", CELLWIDTH-1)
	var valFmt = fmt.Sprintf(" %%%dv ", CELLWIDTH-2)
	//
	// draws the outer (top or bottom) horizontal divider
	var outerHLine = func() {
		ws(EDGE)
		for i := 0; i < 7; i++ {
			if i > 0 {
				ws(HDIV)
			}
			ws(str.Repeat(HDIV, CELLWIDTH))
		}
		ws(EDGE, LF)
	}
	// draws the inner horizontal divider
	var innerHLine = func() {
		for i := 0; i < 7; i++ {
			ws(VDIV, str.Repeat(HDIV, CELLWIDTH))
		}
		ws(VDIV, LF)
	}
	// formats numbers
	var numStr = func(val float64) string {
		var ret = fmt.Sprintf("%5.2f", val)
		if str.Contains(ret, ".") {
			ret = str.TrimRight(ret, "0")
			ret = str.TrimRight(ret, ".")
		}
		return ret
	}
	for _, mth := range ob.months {
		//
		// month heading
		ws(str.ToUpper(fmt.Sprintf("%d %v", mth.year, mth.month)), LF)
		outerHLine()
		//
		// weekday names
		for i := 0; i < 7; i++ {
			ws(VDIV)
			ws(fmt.Sprintf(wdayFmt, cldrWeekdaysEN[i][:3]))
		}
		ws(VDIV, LF)
		var sum float64
		//
		// draw the grid
		for row := 0; row < 6; row++ {
			innerHLine()
			//
			// days on current row
			for col := 0; col < 7; col++ {
				ws(VDIV)
				var day = mth.cells[row][col].day
				if day == 0 {
					ws(blank)
				} else {
					ws(fmt.Sprintf(dayFmt, day))
				}
			}
			ws(VDIV, LF)
			//
			// values on current row
			for col := 0; col < 7; col++ {
				ws(VDIV)
				var val = mth.cells[row][col].val
				if val == nil {
					ws(blank)
				} else {
					if val, ok := val.(float64); ok {
						var s = numStr(val)
						sum += val
						ws(fmt.Sprintf(valFmt, s))
					} else {
						ws(fmt.Sprintf(valFmt, val))
					}
				}
			}
			ws(VDIV, LF)
		}
		outerHLine()
		ws(numStr(sum), LF, LF)
	} // mth
	return retBuf.String()
} //                                                                      String

// -----------------------------------------------------------------------------
// # Helper Method

// getMonth returns a pointer to the month specified by 'year' and
// 'month', or nil if the month has not been added to this calendar
func (ob *Calendar) getMonth(year int, month time.Month) *cldrMonth {
	for i, m := range ob.months {
		if m.year == year && m.month == month {
			return &ob.months[i]
		}
	}
	Error(ENotFound, "month", month, year)
	return nil
} //                                                                    getMonth

// -----------------------------------------------------------------------------
// # Helper Function

// cldrFirstWeekday returns the day of
// week on the first of the given month,
func cldrFirstWeekday(year int, month time.Month) time.Weekday {
	var date = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return date.Weekday()
} //                                                            cldrFirstWeekday

// end
