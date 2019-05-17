// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-17 11:06:49 4D703E                                zr/[numbers.go]
// -----------------------------------------------------------------------------

package zr

// # Constants
//   MaxInt
//   MaxUint
//   MinInt
//
// # Numeric Functions
//   Float64(value interface{}) float64
//   Int(value interface{}) int
//   IsNumber(value interface{}) bool
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

const (
	// MaxInt __
	MaxInt = int(MaxUint >> 1)

	// MaxUint __
	MaxUint = ^uint(0)

	// MinInt __
	MinInt = -MaxInt - 1
)

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
// issue in the log file if the type of value can not be handled.
func Float64(value interface{}) float64 {
	switch v := value.(type) {
	case nil:
		{
			return 0.0
		}
	case bool:
		{
			if v {
				return 1.0
			}
			return 0.0
		}
	case *bool:
		if v != nil {
			return Float64(*v)
		}
	// strings
	case string:
		{
			ret, err := strconv.ParseFloat(v, 64)
			if err != nil {
				ret = 0.0
			}
			return ret
		}
	case *string:
		if v != nil {
			return Float64(*v)
		}
	case fmt.Stringer:
		{
			return Float64(v.String())
		}
	case fmt.GoStringer:
		{
			return Float64(v.GoString())
		}
	// signed integers
	case int:
		{
			return float64(v)
		}
	case int64:
		{
			return float64(v)
		}
	case int32:
		{
			return float64(v)
		}
	case int16:
		{
			return float64(v)
		}
	case int8:
		{
			return float64(v)
		}
		// pointers to signed integers
	case *int:
		if v != nil {
			return float64(*v)
		}
	case *int64:
		if v != nil {
			return float64(*v)
		}
	case *int32:
		if v != nil {
			return float64(*v)
		}
	case *int16:
		if v != nil {
			return float64(*v)
		}
	case *int8:
		if v != nil {
			return float64(*v)
		}
		// unsigned integers
	case uint:
		{
			return float64(v)
		}
	case uint64:
		{
			return float64(v)
		}
	case uint32:
		{
			return float64(v)
		}
	case uint16:
		{
			return float64(v)
		}
	case uint8:
		{
			return float64(v)
		}
		// pointers to unsigned integers
	case *uint:
		if v != nil {
			return float64(*v)
		}
	case *uint64:
		if v != nil {
			return float64(*v)
		}
	case *uint32:
		if v != nil {
			return float64(*v)
		}
	case *uint16:
		if v != nil {
			return float64(*v)
		}
	case *uint8:
		if v != nil {
			return float64(*v)
		}
		// floating-point numbers
	case float64:
		{
			return v
		}
	case float32:
		{
			return float64(v)
		}
		// pointers to floating-point numbers
	case *float64:
		if v != nil {
			return *v
		}
	case *float32:
		if v != nil {
			return float64(*v)
		}
	}
	mod.Error("Can not convert", reflect.TypeOf(value), "to float64:", value)
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
// If the type of value can not be handled, does not
// return an error but logs an issue in the log.
func Int(value interface{}) int {
	switch v := value.(type) {
	case nil:
		{
			return 0
		}
	case bool:
		{
			if v {
				return 1
			}
			return 0
		}
	case *bool:
		if v != nil {
			return Int(*v)
		}
	// strings
	case string:
		{
			ret := 0
			var hasDigit, hasMinus, hasPlus bool
		loop:
			for _, ch := range v {
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
		}
	case *string:
		if v != nil {
			return Int(*v)
		}
	case fmt.Stringer:
		{
			return Int(v.String())
		}
	case fmt.GoStringer:
		{
			return Int(v.GoString())
		}
	// signed integers
	case int:
		{
			return int(v)
		}
	case int64:
		{
			return int(v)
		}
	case int32:
		{
			return int(v)
		}
	case int16:
		{
			return int(v)
		}
	case int8:
		{
			return int(v)
		}
	// pointers to signed integers
	case *int:
		if v != nil {
			return int(*v)
		}
	case *int64:
		if v != nil {
			return int(*v)
		}
	case *int32:
		if v != nil {
			return int(*v)
		}
	case *int16:
		if v != nil {
			return int(*v)
		}
	case *int8:
		if v != nil {
			return int(*v)
		}
	// unsigned integers
	case uint:
		{
			return int(v)
		}
	case uint64:
		{
			return int(v)
		}
	case uint32:
		{
			return int(v)
		}
	case uint16:
		{
			return int(v)
		}
	case uint8:
		{
			return int(v)
		}
	// pointers to unsigned integers
	case *uint:
		if v != nil {
			return int(*v)
		}
	case *uint64:
		if v != nil {
			return int(*v)
		}
	case *uint32:
		if v != nil {
			// TODO: check for overflow (in other locations too)
			return int(*v)
		}
	case *uint16:
		if v != nil {
			return int(*v)
		}
	case *uint8:
		if v != nil {
			return int(*v)
		}
	// floating-point numbers
	case float64:
		{
			if v < -float64(math.MinInt64) || v > float64(math.MaxInt64) {
				mod.Error("overflow")
				if v < 0 {
					return math.MinInt32
				}
				return math.MaxInt32
			}
			return int(v)
		}
	case float32:
		{
			// TODO: find how to find out the limit of int
			//if v < -float32(CurrencyIntLimit) ||v > float32(m) {
			//	return currencyOverflow(v < 0, "uint64 overflow:", v)
			//}
			return int(v)
		}
	// pointers to floating-point numbers
	case *float64:
		if v != nil {
			return int(*v)
		}
	case *float32:
		if v != nil {
			return int(*v)
		}
	}
	mod.Error("Can not convert", reflect.TypeOf(value), "to int:", value)
	return 0
} //                                                                         Int

// IsNumber returns true if value is a number or numeric string,
// or false otherwise. It also accepts pointers to numeric types
// and strings and Stringer. Always returns false if value is nil
// or bool, even though Int() can convert bool to 1 or 0.
func IsNumber(value interface{}) bool {
	const (
		groupSeparatorChar = ','
		decimalPointChar   = '.'
	)
	switch v := value.(type) {
	case int, int64, int32, int16, int8, float64, float32,
		uint, uint64, uint32, uint16, uint8,
		*int, *int64, *int32, *int16, *int8, *float64, *float32,
		*uint, *uint64, *uint32, *uint16, *uint8:
		{
			return true
		}
	case string:
		{
			s := strings.TrimSpace(v)
			if len(s) < 1 {
				return false
			}
			var hasDecPoint, hasDigit, hasSign, prevSep bool
			for _, r := range s {
				switch {
				case r >= '0' && r <= '9':
					{
						hasDigit = true
					}
				case r == groupSeparatorChar:
					{
						// two consecutive group separators
						// make string non-numeric
						if prevSep || !hasDigit {
							return false
						}
						prevSep = true
						continue
					}
				case r == '-' || r == '+':
					{
						if hasSign || hasDigit {
							return false
						}
						hasSign = true
					}
				case r == decimalPointChar:
					{
						if hasDecPoint {
							return false
						}
						hasDecPoint = true
					}
				default:
					return false
				}
				prevSep = false
			}
			return hasDigit
		}
	case *string:
		if v != nil {
			return IsNumber(*v)
		}
	case fmt.Stringer:
		return IsNumber(v.String())
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

// AmountInWordsEN returns the currency value as an English description in words
// This function is useful for showing amounts in invoices, etc.
//
// Uses English language names, hence the 'EN' suffix.
//
// fmt: a string specifying the currency format:
//      format_ = "<Large-Single>;<Large-Plural>;<S-Single>;<S-Plural>;<Only>"
//      Large-Single - Large Denomination, Singular. E.g. "Dollar"
//      Large-Plural - Large Denomination, Plural.   E.g. "Dollars"
//      S-Single     - Small Denomination, Singular. E.g. "Cent"
//      S-Plural     - Small Denomination, Plural.   E.g. "Cents"
//
//      All the format specifiers are optional but must follow in
//      the order specified. If a plural specifier is omitted, then
//      an "s" is added to the singular specifier for values greater
//      than 1. If the singular specifier is omitted, Then the
//      plural is used for all values (Used for some currencies).
//      If both specifiers are omitted for either denomination, the
//      denomination will not be returned. See the examples below.
//
// Returns: The amount in words as a string, including the currency
//          and the word "Only". Uses proper capitalisation.
//
// Example: PARAMETER               RESULT
//          (11.02,";;Cent;Only")   "Two Cents Only"
//          (11.02,"Dollar;;Cent")  "Eleven Dollars and Two Cents"
//          (11.02,"Euro")          "Eleven Euros"
//          (11.02,"Pound;;;Pence") "Eleven Pounds and Two Pence"
func AmountInWordsEN(n Currency, fmt string) string {
	i := n.i64
	if i < 0 {
		i = -n.i64
	}
	var (
		bigUnits = i / 1E4
		smlUnits = (i - bigUnits*1E4) / 100
		hasOnly  = strings.HasSuffix(strings.ToLower(fmt), "only")
	)
	if hasOnly {
		fmt = fmt[:len(fmt)-4]
	}
	getPart := func(partNo int) string {
		parts := strings.Split(fmt, ";")
		if partNo < 0 || partNo >= len(parts) {
			return ""
		}
		return parts[partNo]
	}
	var (
		big1 = getPart(0)
		bigN = getPart(1)
		sml1 = getPart(2)
		smlN = getPart(3)
		ret  = ""
	)
	if bigUnits > 0 && (big1+bigN) != "" {
		ret += IntInWordsEN(bigUnits) + " "
		if big1 == "" && bigN != "" {
			ret += bigN
		} else if big1 != "" && bigN == "" {
			ret += big1
			if bigUnits > 1 {
				ret += "s"
			}
		} else if big1 != "" && bigN != "" {
			if bigUnits == 1 {
				ret += big1
			}
			if bigUnits > 1 {
				ret += bigN
			}
		}
	}
	if ((sml1 + smlN) != "") && smlUnits > 0 {
		if (big1+bigN != "") && bigUnits > 0 {
			ret += " and "
		}
		ret += IntInWordsEN(smlUnits) + " "
		if sml1 == "" && smlN != "" {
			ret += smlN
		} else if sml1 != "" && smlN == "" {
			ret += sml1
			if smlUnits > 1 {
				ret += "s"
			}
		} else if sml1 != "" && smlN != "" {
			if smlUnits == 1 {
				ret += sml1
			}
			if smlUnits > 1 {
				ret += smlN
			}
		}
	}
	if hasOnly && len(strings.TrimSpace(ret)) > 0 {
		ret += " Only"
	}
	return ret
} //                                                             AmountInWordsEN

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

// TODO: CommaDelimit should accept interface{} in number
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
		var (
			groups = (intLen / 3) + 1
			digits = intLen % 3
			at     = 0
		)
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
			{
				ws(" ")
				ws(DigitNamesEN[n1])
			}
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
	return strings.TrimSpace(retBuf.String())
} //                                                                IntInWordsEN

//end
