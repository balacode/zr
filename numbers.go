// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 17:47:59 0EBFD8                                zr/[numbers.go]
// -----------------------------------------------------------------------------

package zr

// # Constants
//   MaxInt
//   MaxUint
//   MinInt
//
// # Numeric Functions
//   Float64(val interface{}) float64
//   Int(val interface{}) int
//   IsNumber(val interface{}) bool
//   MaxIntOf(values []int) (max int, found bool)
//   MinMaxGap(values []int) (min, max int)
//
// # Formatting Functions
//   BlankZero(s string) string
//   CommaDelimit(number string, decimalPlaces int) string
//   IntInWordsEN(number int64) string

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// -----------------------------------------------------------------------------
// # Constants

// MaxInt __
const MaxInt = int(MaxUint >> 1)

// MaxUint __
const MaxUint = ^uint(0)

// MinInt __
const MinInt = -MaxInt - 1

// DigitNamesEN are English names of decimal digits 0 to 9.
// These constants are mainly used by IntInWordsEN().
var DigitNamesEN = []string{
	"Zero", "One", "Two", "Three", "Four",
	"Five", "Six", "Seven", "Eight", "Nine",
}

// TeensEN are English names of numbers from 11 to 19.
// These constants are mainly used by IntInWordsEN().
var TeensEN = []string{
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen",
	"Sixteen", "Seventeen", "Eighteen", "Nineteen",
}

// TensEN are English names of tens (10, 20,.. 90)
// These constants are mainly used by IntInWordsEN().
var TensEN = []string{
	"Ten", "Twenty", "Thirty", "Forty", "Fifty",
	"Sixty", "Seventy", "Eighty", "Ninety",
}

// -----------------------------------------------------------------------------
// # Numeric Functions

// Float64 converts primitive types, fmt.Stringer and
// fmt.GoStringer to a floating point number:
//
// - Dereferences pointers (but not pointers to pointers).
// - Converts nil to 0.
// - Converts boolean true to 1, false to 0.
// - Converts numeric strings to float64.
// - Converts string types using strconv.ParseFloat().
//   If a string can't be converted, returns 0.
//
// This function can be used in cases where a simple cast won't work,
// and to easily convert interface{} to a float64.
//
// If the conversion fails, does not return any error, but logs an
// issue in the log file if the type of 'val' can not be handled.
func Float64(val interface{}) float64 {
	switch val := val.(type) {
	//
	// nil and boolean:
	case nil:
		return 0.0
	case bool:
		if val {
			return 1.0
		}
		return 0.0
	case *bool:
		if val != nil {
			return Float64(*val)
		}
	// strings
	case string:
		ret, err := strconv.ParseFloat(val, 64)
		if err != nil {
			ret = 0.0
		}
		return ret
	case *string:
		if val != nil {
			return Float64(*val)
		}
	case fmt.Stringer:
		return Float64(val.String())
	case fmt.GoStringer:
		return Float64(val.GoString())
	//
	// signed integers
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
		//
		// pointers to signed integers
	case *int:
		if val != nil {
			return float64(*val)
		}
	case *int8:
		if val != nil {
			return float64(*val)
		}
	case *int16:
		if val != nil {
			return float64(*val)
		}
	case *int32:
		if val != nil {
			return float64(*val)
		}
	case *int64:
		if val != nil {
			return float64(*val)
		}
		// unsigned integers
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
		//
		// pointers to unsigned integers
	case *uint:
		if val != nil {
			return float64(*val)
		}
	case *uint8:
		if val != nil {
			return float64(*val)
		}
	case *uint16:
		if val != nil {
			return float64(*val)
		}
	case *uint32:
		if val != nil {
			return float64(*val)
		}
	case *uint64:
		if val != nil {
			return float64(*val)
		}
		// floating-point numbers
	case float32:
		return float64(val)
	case float64:
		return val
		//
		// pointers to floating-point numbers
	case *float32:
		if val != nil {
			return float64(*val)
		}
	case *float64:
		if val != nil {
			return *val
		}
	}
	mod.Error("Can not convert",
		reflect.TypeOf(val), "to float64:", val)
	return 0.0
} //                                                                     Float64

// Int converts primitive types, fmt.Stringer and
// fmt.GoStringer to an integer number (int type):
//
// - Dereferences pointers (but not pointers to pointers).
// - Converts nil to 0.
// - Converts boolean true to 1, false to 0.
// - Converts numeric strings to int.
//   Parsing continues until the first non-numeric character.
//   Therefore a string like '123AA456' converts to '123'.
//
// This function can be used in cases where a simple cast won't work,
// and to easily convert interface{} to an int.
//
// If the type of 'val' can not be handled, does not
// return an error but logs an issue in the log.
func Int(val interface{}) int {
	switch val := val.(type) {
	//
	// nil and bool types
	case nil:
		return 0
	case bool:
		if val {
			return 1
		}
		return 0
	case *bool:
		if val != nil {
			return Int(*val)
		}
	// strings
	case string:
		ret := 0
		var hasDigit, hasMinus, hasPlus bool
	loop:
		for _, ch := range val {
			//
			// ignore leading spaces
			if !(hasDigit || hasMinus || hasPlus) {
				for _, sp := range SPACES {
					if ch == sp {
						continue loop
					}
				}
			}
			// handle '-' and '+' signs
			if ch == '-' || ch == '+' {
				if hasMinus || hasPlus || hasDigit {
					return ret
				}
				if ch == '-' {
					hasMinus = true
				} else {
					hasPlus = true
				}
				continue loop
			}
			// add digits to result
			if ch >= '0' && ch <= '9' {
				hasDigit = true
				ret = ret*10 + int(ch-'0')
				continue loop
			}
			break
		}
		if hasMinus {
			ret = -ret
		}
		return ret
	case *string:
		if val != nil {
			return Int(*val)
		}
	case fmt.Stringer:
		return Int(val.String())
	case fmt.GoStringer:
		return Int(val.GoString())
	//
	// signed integers
	case int:
		return int(val)
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	//
	// pointers to signed integers
	case *int:
		if val != nil {
			return int(*val)
		}
	case *int8:
		if val != nil {
			return int(*val)
		}
	case *int16:
		if val != nil {
			return int(*val)
		}
	case *int32:
		if val != nil {
			return int(*val)
		}
	case *int64:
		if val != nil {
			return int(*val)
		}
	// unsigned integers
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	//
	// pointers to unsigned integers
	case *uint:
		if val != nil {
			return int(*val)
		}
	case *uint8:
		if val != nil {
			return int(*val)
		}
	case *uint16:
		if val != nil {
			return int(*val)
		}
	case *uint32:
		if val != nil {
			//TODO: check for overflow (in other locations too)
			return int(*val)
		}
	case *uint64:
		if val != nil {
			return int(*val)
		}
	// floating-point numbers
	case float32:
		//TODO: find how to find out the limit of int
		//if val < -float32(CurrencyIntLimit) ||val > float32(m) {
		//	return currencyOverflow(val < 0, "uint64 overflow:", val)
		//}
		return int(val)
	case float64:
		if val < -float64(math.MinInt64) || val > float64(math.MaxInt64) {
			mod.Error("overflow")
			if val < 0 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		return int(val)
	//
	// pointers to floating-point numbers
	case *float32:
		if val != nil {
			return int(*val)
		}
	case *float64:
		if val != nil {
			return int(*val)
		}
	}
	mod.Error("Can not convert", reflect.TypeOf(val), "to int:", val)
	return 0
} //                                                                         Int

// IsNumber returns true if 'val' is a number or numeric string,
// or false otherwise. It also accepts pointers to numeric types
// and strings and Stringer. Always returns false if val is nil
// or bool, even though Int() can convert bool to 1 or 0.
func IsNumber(val interface{}) bool {
	const groupSeparatorChar = ','
	const decimalPointChar = '.'
	//
	switch val := val.(type) {
	case int, int8, int16, int32, int64,
		*int, *int8, *int16, *int32, *int64,
		uint, uint8, uint16, uint32, uint64,
		*uint, *uint8, *uint16, *uint32, *uint64,
		float32, float64,
		*float32, *float64:
		return true
	case string:
		s := strings.Trim(val, SPACES)
		if len(s) < 1 {
			return false
		}
		var hasDecPoint, hasDigit, hasSign, prevSep bool
		for _, r := range s {
			switch {
			case r >= '0' && r <= '9':
				hasDigit = true
			case r == groupSeparatorChar:
				// two consecutive group separators make string non-numeric
				if prevSep || !hasDigit {
					return false
				}
				prevSep = true
				continue
			case r == '-' || r == '+':
				if hasSign || hasDigit {
					return false
				}
				hasSign = true
			case r == decimalPointChar:
				if hasDecPoint {
					return false
				}
				hasDecPoint = true
			default:
				return false
			}
			prevSep = false
		}
		return hasDigit
	case *string:
		if val != nil {
			return IsNumber(*val)
		}
	case fmt.Stringer:
		return IsNumber(val.String())
	}
	return false
} //                                                                    IsNumber

// MinMaxGap returns the lowest and highest unique integer that can
// fit in a gap in a series of integers. E.g. given 1, 4, and 7
// this would be 2 and 6. Returns the resulting integers if there
// is a gap, or MaxInt and MinInt if there is no gap in the series.
func MinMaxGap(values []int) (min, max int) {
	//
	// return immediately when the slice has less than 2 values:
	if len(values) < 2 {
		return MaxInt, MinInt
	}
	// copy and sort the input slice, so original is unchanged
	ar := make([]int, len(values))
	copy(ar, values)
	sort.Ints(ar)
	//
	// find the lowest unused integer in range
	min = MaxInt
	for i := 0; i < len(ar)-1; i++ {
		if ar[i+1] != ar[i] && ar[i+1] != ar[i]+1 {
			min = ar[i] + 1
			break
		}
	}
	// find the highest unused integer in range
	max = MinInt
	for i := len(ar) - 1; i > 0; i-- {
		if ar[i-1] != ar[i] && ar[i-1] != ar[i]-1 {
			max = ar[i] - 1
			break
		}
	}
	return min, max
} //                                                                   MinMaxGap

// MaxIntOf returns the maximum value in a slice of integers (and true
// in the second returned value) or 0 and false if the slice is empty.
func MaxIntOf(values []int) (max int, found bool) {
	if len(values) == 0 {
		return 0, false
	}
	for _, n := range values {
		if n > max || !found {
			max = n
			found = true
		}
	}
	return max, found
} //                                                                    MaxIntOf

// -----------------------------------------------------------------------------
// # Formatting Functions

// BlankZero returns a blank string when given a string
// containing only zeros, decimal points and white-spaces.
// Any string that doesn't contain '0' is returned unchanged.
func BlankZero(s string) string {
	if !strings.Contains(s, "0") {
		return s
	}
	for _, ch := range s {
		if ch != ' ' && ch != '.' &&
			ch != '\a' && ch != '\b' &&
			ch != '\f' && ch != '\n' &&
			ch != '\r' && ch != '\t' &&
			ch != '\v' && ch != '0' {
			return s
		}
	}
	return ""
} //                                                                   BlankZero

//TODO: CommaDelimit should accept interface{} in number
// CommaDelimit delimits a numeric string with commas (grouped every
// three digits) and also sets the required number of decimal places.
// Numbers are not rounded, just cut at the required number of decimals.
func CommaDelimit(number string, decimalPlaces int) string {
	var (
		retBuf = bytes.NewBuffer(make([]byte, 0, 32))
		ws     = retBuf.WriteString
		intLen = 0
		decAt  = strings.Index(number, ".")
	)
	// calculate length of number's integer part
	if decAt == -1 {
		intLen = len(number)
	} else if decAt != 0 {
		intLen = decAt
	}
	{ // write delimited integer part
		groups := (intLen / 3) + 1
		digits := intLen % 3
		at := 0
		for groups > 0 {
			ws(number[at : at+digits])
			if groups > 1 && digits != 0 {
				ws(",")
			}
			at += digits
			digits = 3
			groups--
		}
	}
	if intLen == 0 {
		ws("0")
	}
	if decimalPlaces > 0 { // write fractional part
		ws(".")
		decLen := 0
		if decAt != -1 {
			decLen = len(number[decAt+1:])
		}
		if decLen > decimalPlaces {
			decLen = decimalPlaces
		}
		if decLen > 0 {
			ws(number[decAt+1 : decAt+1+decLen])
		}
		for decLen < decimalPlaces {
			ws("0")
			decLen++
		}
	}
	return retBuf.String()
} //                                                                CommaDelimit

// IntInWordsEN returns the given number as a description in words.
// Uses English language names, hence the 'EN' suffix.
// This function is useful for showing amounts in invoices, etc.
// 'number' must be a positive integer in the range of 0 to 1 trillion.
// E.g. IntInWordsEN(256) returns "Two Hundred and Fifty Six"
func IntInWordsEN(number int64) string {
	if number == 0 {
		return DigitNamesEN[0]
	}
	// divide number into billions, millions, thousands, units, etc.
	groups := []struct {
		n    int64
		unit string
		base int64
	}{
		{0, "", 1},               // units
		{0, "Thousand", 1e3},     // 10^3 (10^3 has 3 zeros, etc)
		{0, "Million", 1e6},      // 10^6
		{0, "Billion", 1e9},      // 10^9
		{0, "Trillion", 1e12},    // 10^12
		{0, "Quadrillion", 1e15}, // 10^15
		{0, "Quintillion", 1e18}, // 10^18
		// math.MaxInt64 = 9223372036854775807
	}
	{
		n := number
		for i := len(groups) - 1; i >= 0; i-- {
			gr := &groups[i]
			if n < gr.base {
				continue
			}
			gr.n = n / gr.base
			n -= gr.n * gr.base
		}
	}
	var retBuf bytes.Buffer
	ws := retBuf.WriteString
	for i := len(groups) - 1; i >= 0; i-- {
		n := groups[i].n
		if n == 0 {
			continue
		}
		unit := groups[i].unit
		//
		// count hundreds, tens and units
		n100 := n / 100
		n -= n100 * 100
		n10 := n / 10
		n -= n10 * 10
		n1 := n
		//
		// append names of hundreds to result
		if n100 != 0 {
			ws(" ")
			ws(DigitNamesEN[n100])
			ws(" Hundred")
		}
		if (n1 != 0 || n10 != 0) && n100 != 0 {
			ws(" and")
		}
		// append tens and units (or teen numbers) to result
		switch {
		case n10 == 1 && n1 != 0:
			if n1 >= 0 && n1 <= 9 {
				ws(" ")
				ws(TeensEN[n1-1])
			}
		case n10 == 0 && n1 != 0:
			ws(" ")
			ws(DigitNamesEN[n1])
		default:
			if n10 != 0 {
				ws(" ")
				ws(TensEN[n10-1])
			}
			if n1 != 0 {
				ws(" ")
				ws(DigitNamesEN[n1])
			}
		}
		ws(" ")
		ws(unit)
	}
	return strings.Trim(retBuf.String(), SPACES)
} //                                                                IntInWordsEN

//end
