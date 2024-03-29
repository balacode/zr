// -----------------------------------------------------------------------------
// ZR Library                                                   zr/[currency.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # Constants:
//   CurrencyIntLimit
//   MinCurrencyI64
//   MaxCurrencyI64
//
// # Currency Type:
//   Currency struct
//
// # Currency Factories:
//   CurrencyE(value interface{}) (Currency, error)
//   CurrencyOf(value interface{}) Currency
//   CurrencyOfS(s string) Currency
//   CurrencyRaw(i64 int64) Currency
//
// # String Output:
//   (n Currency) GoString() string
//   (n Currency) Fmt(decimalPlaces int) string
//   (n Currency) InWordsEN(fmt string) string
//   (n Currency) String() string
//
// # Division:
//   (n Currency) Div(nums... Currency) Currency
//   (n Currency) DivFloat(nums...float64) Currency
//   (n Currency) DivInt(nums...int) Currency
//
// # Multiplication:
//   (n Currency) Mul(nums... Currency) Currency
//   (n Currency) MulFloat(nums...float64) Currency
//   (n Currency) MulInt(nums...int) Currency
//
// # Addition:
//   (n Currency) Add(nums... Currency) Currency
//   (n Currency) AddFloat(nums...float64) Currency
//   (n Currency) AddInt(nums...int) Currency
//
// # Subtraction:
//   (n Currency) Sub(nums... Currency) Currency
//   (n Currency) SubFloat(nums...float64) Currency
//   (n Currency) SubInt(nums...int) Currency
//
// # Information:
//   (n Currency) Float64() float64
//   (n Currency) Int() int
//   (n Currency) Int64() int64
//   (n Currency) IsEqual(value interface{}) bool
//   (n Currency) IsGreater(value interface{}) bool
//   (n Currency) IsGreaterOrEqual(value interface{}) bool
//   (n Currency) IsLesser(value interface{}) bool
//   (n Currency) IsLesserOrEqual(value interface{}) bool
//   (n Currency) IsNegative() bool
//   (n Currency) IsZero() bool
//   (n Currency) Overflow() int
//   (n Currency) Raw() int64
//
// # JSON:
//   (n Currency) MarshalJSON() ([]byte, error)
//   (n *Currency) UnmarshalJSON(data []byte) error

// # Helper Function
//   currencyOverflow(isNegative bool, a ...interface{}) Currency

import (
	// "encoding/json" // used via mod.json.* mockable
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

// -----------------------------------------------------------------------------
// # Constants:

const (
	// CurrencyIntLimit specifies the highest (and lowest
	// when negative) integer component that currency
	// can hold, about 922.33 trillion.
	// The exact number is 922(t)337(b)203(m)685,476.
	//
	// This limit exists because an int64, on which Currency is based, can hold
	// up to 9,223,372,036,854,775,807. Out of this 4 digits are used for the
	// decimal part, i.e. 922,337,203,685,477.5807. The limit is set to this
	// number minus 1, so that all decimals from .0000 to .9999. can be used.
	CurrencyIntLimit = 922337203685476

	// MinCurrencyI64 is the lowest internal value that Currency can hold.
	MinCurrencyI64 = -9223372036854769999

	// MaxCurrencyI64 is the highest internal value that Currency can hold.
	MaxCurrencyI64 = 9223372036854769999
)

// big1E4 scales the int64 to provide 4 decimal places.
var big1E4 = big.NewInt(1e4)

// -----------------------------------------------------------------------------
// # Currency Type:

// Currency represents a currency value with up to four decimal places.
// It is stored internally as a 64-bit integer. This gives it a range
// from -922,337,203,685,477.5808 to 922,337,203,685,477.5807
type Currency struct {
	i64 int64
} //                                                                    Currency

// -----------------------------------------------------------------------------
// # Currency Factories:

// CurrencyE converts any compatible value to a Currency.
// This includes simple numeric types and strings.
//
// - Dereferences pointers to evaluate the pointed-to type.
// - Converts nil to 0.
// - Converts signed and unsigned integers, and floats to Currency.
// - Converts numeric strings to Currency.
// - Converts boolean true to 1, false to 0.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
// If the value can not be converted to Currency, returns a
// zero-value Currency and an error. Does not log the error.
//
func CurrencyE(value interface{}) (Currency, error) {
	switch v := value.(type) {
	case int, int64, int32, int16, int8:
		{
			n := reflect.ValueOf(value).Int()
			if n < -CurrencyIntLimit || n > CurrencyIntLimit {
				return currencyOverflow(n < 0, n), errors.New(EOverflow)
			}
			return Currency{n * 1e4}, nil
		}
	case float64, float32:
		{
			n := reflect.ValueOf(value).Float()
			if n < -float64(CurrencyIntLimit)-0.9999 ||
				n > float64(CurrencyIntLimit)+0.9999 {
				return currencyOverflow(n < 0, n), errors.New(EOverflow)
			}
			return Currency{int64(n * 1e4)}, nil
		}
	case string:
		{
			return CurrencyOfS(v), nil
		}
	case uint, uint64, uint32, uint16, uint8:
		{
			n := reflect.ValueOf(value).Uint()
			if n > CurrencyIntLimit {
				return currencyOverflow(n < 0, n), errors.New(EOverflow)
			}
			return Currency{int64(n) * 1e4}, nil
		}
	case Currency:
		{
			return v, nil
		}
	case bool:
		{
			if v {
				return Currency{1 * 1e4}, nil
			}
			return Currency{0}, nil
		}
	case nil:
		return Currency{0}, nil
	}
	// if not converted yet, try to dereference pointer, then convert
	ret := Currency{0}
	xv := reflect.ValueOf(value)
	if xv.Kind() == reflect.Ptr {
		if xv.IsNil() {
			return Currency{0}, nil
		}
		var err error
		ret, err = CurrencyE(xv.Elem().Interface())
		if err == nil || strings.HasPrefix(err.Error(), EOverflow) {
			return ret, nil
		}
	}
	erm := fmt.Sprintf("Can not convert %s to Currency: %v",
		reflect.TypeOf(value), value)
	err := errors.New(erm)
	return ret, err
} //                                                                   CurrencyE

// CurrencyOf converts any compatible value type to a Currency.
// This includes all numeric types and strings. If a string is
// not numeric, logs an error and sets the Currency to zero.
//
// - Dereferences pointers to evaluate the pointed-to type.
// - Converts nil to 0.
// - Converts signed and unsigned integers, and floats to Currency.
// - Converts numeric strings to Currency.
// - Converts boolean true to 1, false to 0.
//
// If the value can not be converted to Currency, returns a zero-value
// Currency and logs an error (when logging is active).
//
func CurrencyOf(value interface{}) Currency {
	n, err := CurrencyE(value)
	if err != nil && !strings.HasPrefix(err.Error(), EOverflow) {
		mod.Error(err)
	}
	return n
} //                                                                  CurrencyOf

// CurrencyOfS converts a numeric string to a Currency.
// If the string is not numeric, logs an error and sets the Currency to zero.
func CurrencyOfS(s string) Currency {
	var (
		minus bool
		fract bool
		dp    int
		ret   Currency
	)
	for _, r := range s {
		if r == '-' {
			minus = true
		} else if r == '.' {
			fract = true
		} else if r >= '0' && r <= '9' {
			if fract {
				dp++
				if dp > 4 {
					break
				}
			}
			ret.i64 *= 10
			ret.i64 += int64(r - '0')
		} else if r != ',' &&
			r != ' ' && r != '\a' && r != '\b' && r != '\f' &&
			r != '\n' && r != '\r' && r != '\t' && r != '\v' {
			mod.Error("Non-numeric string:^", s)
			break
		}
	}
	for dp < 4 {
		dp++
		ret.i64 *= 10
	}
	if minus {
		ret.i64 = -ret.i64
	}
	return ret
} //                                                                 CurrencyOfS

// CurrencyRaw initializes a currency value from a scaled 64-bit integer.
// The decimal point is moved left 4 decimal places.
// For example, a i64 value of 15500 results in a currency value of 1.55
func CurrencyRaw(i64 int64) Currency {
	return Currency{i64: i64}
} //                                                                 CurrencyRaw

// -----------------------------------------------------------------------------
// # String Output:

// GoString outputs the value as a Go language string,
// It implements the fmt.GoStringer interface.
func (n Currency) GoString() string {
	return "zr.CurrencyOf(" + n.String() + ")"
} //                                                                    GoString

// Fmt returns the currency value as a a string
// delimited with commas (grouped every three digits)
// and having the specified number of decimal places.
// When decimalPlaces is negative, the resulting
// number's decimals will vary.
func (n Currency) Fmt(decimalPlaces int) string {
	var (
		retBuf   = bytes.NewBuffer(make([]byte, 0, 25))
		ws       = retBuf.WriteString
		wr       = retBuf.WriteRune
		intLen   = 0
		intPart  = n.i64 / 1e4         // integer part of the number
		decPart  = n.i64 - intPart*1e4 // decimal part (as an int)
		negative = n.i64 < 0           // is it negative? use absolute value
	)
	if negative {
		intPart = -intPart
		ws("-")
	}
	// calculate length of number's integer part
	for limit := int64(0); intLen < 15; intLen++ {
		if intPart <= limit {
			break
		}
		limit *= 10
		limit += 9
	}
	// write delimited integer part
	if intLen == 0 {
		ws("0")
	} else {
		var (
			write  = false
			power  = int64(100000000000000) // 10^14
			digits = intLen % 3
		)
		if digits == 0 {
			digits = 3
		}
		for power > 0 {
			x := intPart / power
			if x > 0 {
				write = true
			}
			intPart -= x * power
			power /= 10
			if !write {
				continue
			}
			wr(rune(x + 48))
			digits--
			if power > 0 && digits <= 0 {
				ws(",")
				digits = 3
			}
		}
	}
	// write fractional part
	if decimalPlaces != 0 {
		power := int64(1000) // 10^3
		unfixed := decPart > 0 && decimalPlaces < 0
		if unfixed {
			decimalPlaces = 4
		}
		if decimalPlaces > 0 {
			ws(".")
		}
		for decimalPlaces > 0 {
			decimalPlaces--
			x := int64(0)
			if power > 0 {
				x = decPart / power
				decPart -= x * power
				power /= 10
			}
			wr(rune(x + 48))
			if unfixed && decPart == 0 {
				break
			}
		}
	}
	return retBuf.String()
} //                                                                         Fmt

// InWordsEN returns the currency value as an English description in words
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
func (n Currency) InWordsEN(fmt string) string {
	i := n.i64
	if i < 0 {
		i = -n.i64
	}
	var (
		bigUnits = i / 1e4
		smlUnits = (i - bigUnits*1e4) / 100
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
} //                                                                   InWordsEN

// String returns a string representing the currency value
// and implements the fmt.Stringer interface.
func (n Currency) String() string {
	var (
		i    = n.i64 / 1e4              // integer value
		d    = n.i64 - i*1e4            // decimal value
		sint = strconv.FormatInt(i, 10) // integer part
		sdec string                     // decimal part
	)
	if d != 0 { // format decimals
		if d < 0 { // adjust for negative value
			d = -d
			if i == 0 {
				sint = "-" + sint
			}
		}
		sdec = "." + strings.TrimRight(
			strconv.FormatInt(d+1e4, 10)[1:],
			"0",
		)
	}
	return sint + sdec
} //                                                                      String

// -----------------------------------------------------------------------------
// # Division:

// Div divides a currency object by one or more currency values
// and returns the result. The object's value isn't changed.
func (n Currency) Div(nums ...Currency) Currency {
	for _, num := range nums {
		n.i64 *= 1e4
		n.i64 /= num.i64
	}
	return n
} //                                                                         Div

// DivFloat divides a currency object by one or more floating-point
// numbers and returns the result. The object's value isn't changed.
func (n Currency) DivFloat(nums ...float64) Currency {
	for _, num := range nums {
		n.i64 *= 1e4
		n.i64 /= int64(num * 1e4)
	}
	return n
} //                                                                    DivFloat

// DivInt divides a currency object by one or more integer values
// and returns the result. The object's value isn't changed.
func (n Currency) DivInt(nums ...int) Currency {
	for _, num := range nums {
		n.i64 *= 1e4
		n.i64 /= (int64(num) * 1e4)
	}
	return n
} //                                                                      DivInt

// -----------------------------------------------------------------------------
// # Multiplication:

// Mul multiplies a currency object by one or more currency values
// and returns the result. The object's value isn't changed.
func (n Currency) Mul(nums ...Currency) Currency {
	for _, num := range nums {
		var (
			a = n.i64
			b = num.i64
		)
		// return zero if either number is zero
		if a == 0 || b == 0 {
			return Currency{0}
		}
		// if direct multiplication will overflow int64, use big.Int
		lim := MaxCurrencyI64 / a
		if lim < 0 {
			lim = -lim
		}
		if b >= lim || b <= -lim {
			x := big.NewInt(a)
			x.Mul(x, big.NewInt(b))
			x.Div(x, big1E4)
			//
			// if result can't be stored in Currency, return overflow
			//
			// TODO: IsInt64() is not available in older Go versions ``
			overflow := !x.IsInt64()
			var ret int64
			if !overflow {
				ret = x.Int64()
				if ret < MinCurrencyI64 || ret > MaxCurrencyI64 {
					overflow = true
				}
			}
			if overflow {
				return currencyOverflow(
					(a < 0 || b < 0) && (a > 0 || b > 0),
					a, " * ", b, " = ", x)
			}
			return Currency{ret}
		}
		n.i64 = a * b / 1e4
	}
	return n
} //                                                                         Mul

// MulFloat multiplies a currency object by one or more floating-point
// numbers and returns the result. The object's value isn't changed.
func (n Currency) MulFloat(nums ...float64) Currency {
	a := float64(n.i64)
	for _, b := range nums {
		// check for negative or positive overflow
		lim := MaxCurrencyI64 / a
		if lim < 0 {
			lim = -lim
		}
		if b < -lim || b > lim {
			return currencyOverflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				a, " * ", b, " = ", a*b)
		}
		// multiply using int64, if there is no overflow
		n.i64 = int64(a * b)
	}
	return n
} //                                                                    MulFloat

// MulInt multiplies a currency object by one or more integer values
// and returns the result. The object's value isn't changed.
func (n Currency) MulInt(nums ...int) Currency {
	for _, num := range nums {
		n = n.Mul(Currency{int64(num * 1e4)})
	}
	return n
} //                                                                      MulInt

// -----------------------------------------------------------------------------
// # Addition:

// Add adds one or more currency values and returns a new Currency object.
// The value in the object to which this method is applied isn't changed.
// If there is an overflow, sets the Currency's internal value to
// math.MinInt64 or math.MaxInt64 depending on if the result is negative.
func (n Currency) Add(nums ...Currency) Currency {
	for _, num := range nums {
		var (
			a = n.i64
			b = num.i64
			c = a + b
		)
		// check for overflow
		if c < MinCurrencyI64 || (a < 0 && b < 0 && b < (MinCurrencyI64-a)) {
			return currencyOverflow(true, a, " + ", b)
		}
		if c > MaxCurrencyI64 || (a > 0 && b > 0 && b > (MaxCurrencyI64-a)) {
			return currencyOverflow(false, a, " + ", b)
		}
		n.i64 = c
	}
	return n
} //                                                                         Add

// AddFloat adds one or more floating-point numbers to a currency
// object and returns the result. The object's value isn't changed.
func (n Currency) AddFloat(nums ...float64) Currency {
	const lim float64 = math.MaxInt64 / 1e4
	a := n.i64
	for _, b := range nums {
		//
		// check for overflow
		if b < -lim || b > lim {
			return currencyOverflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				a, " * ", b, " = ", float64(a)*b)
		}
		// use Add() because it has other overflow checks
		n = n.Add(Currency{int64(b * 1e4)})
	}
	return n
} //                                                                    AddFloat

// AddInt adds one or more integer values to a currency object
// and returns the result. The object's value isn't changed.
func (n Currency) AddInt(nums ...int) Currency {
	for _, num := range nums {
		n = n.Add(Currency{int64(num) * 1e4})
	}
	return n
} //                                                                      AddInt

// -----------------------------------------------------------------------------
// # Subtraction:

// Sub subtracts one or more currency values from a currency object
// and returns the result. The object's value isn't changed.
func (n Currency) Sub(nums ...Currency) Currency {
	for _, num := range nums {
		var (
			a = n.i64
			b = num.i64
			c = a - b
		)
		// check for overflow
		if c < MinCurrencyI64 || (a < 0 && b > 0 && b > (-MinCurrencyI64+a)) {
			return currencyOverflow(true, a, " - ", b)
		}
		if c > MaxCurrencyI64 || (a > 0 && b < 0 && b < (-MaxCurrencyI64+a)) {
			return currencyOverflow(false, a, " - ", b)
		}
		n.i64 = c
	}
	return n
} //                                                                         Sub

// SubFloat subtracts one or more floating-point numbers from a currency
// object and returns the result. The object's value isn't changed.
func (n Currency) SubFloat(nums ...float64) Currency {
	for _, num := range nums {
		n.i64 -= int64(num * 1e4)
	}
	return n
} //                                                                    SubFloat

// SubInt subtracts one or more integer values from a currency object
// and returns the result. The object's value isn't changed.
func (n Currency) SubInt(nums ...int) Currency {
	for _, num := range nums {
		n.i64 -= int64(num) * 1e4
	}
	return n
} //                                                                      SubInt

// -----------------------------------------------------------------------------
// # Information:

// Float64 returns the currency value as a float64 value.
func (n Currency) Float64() float64 {
	return float64(n.i64) / 1e4
} //                                                                     Float64

// Int returns the currency value as an int value.
func (n Currency) Int() int64 {
	return n.i64 / 1e4
} //                                                                         Int

// Int64 returns the currency value as an int64 value.
func (n Currency) Int64() int64 {
	return n.i64 / 1e4
} //                                                                       Int64

// IsEqual returns true if the value of the currency object is negative.
func (n Currency) IsEqual(value interface{}) bool {
	return n.i64 == CurrencyOf(value).i64
} //                                                                     IsEqual

// IsGreater returns true if the object is greater than value.
func (n Currency) IsGreater(value interface{}) bool {
	return n.i64 > CurrencyOf(value).i64
} //                                                                   IsGreater

// IsGreaterOrEqual returns true if the object is greater or equal to value.
func (n Currency) IsGreaterOrEqual(value interface{}) bool {
	return n.i64 >= CurrencyOf(value).i64
} //                                                            IsGreaterOrEqual

// IsLesser returns true if the object is lesser than value.
func (n Currency) IsLesser(value interface{}) bool {
	return n.i64 < CurrencyOf(value).i64
} //                                                                    IsLesser

// IsLesserOrEqual returns true if the object is lesser or equal to value.
func (n Currency) IsLesserOrEqual(value interface{}) bool {
	return n.i64 <= CurrencyOf(value).i64
} //                                                             IsLesserOrEqual

// IsNegative returns true if the value of the currency object is negative.
func (n Currency) IsNegative() bool {
	return n.i64 < 0
} //                                                                  IsNegative

// IsZero returns true if the value of the currency object is zero.
func (n Currency) IsZero() bool {
	return n.i64 == 0
} //                                                                      IsZero

// Overflow returns 1 if the currency contains a positive overflow,
// -1 if it contains a negative overflow,
// or 0 if there is no overflow
//
// Overflows occur when an arithmeric operation's
// result exceeds the storage capacity of Currency.
//
func (n Currency) Overflow() int {
	if n.i64 > MaxCurrencyI64 {
		return 1
	}
	if n.i64 < MinCurrencyI64 {
		return -1
	}
	return 0
} //                                                                    Overflow

// Raw returns the internal int64 used to store the currency value.
func (n Currency) Raw() int64 {
	return n.i64
} //                                                                         Raw

// -----------------------------------------------------------------------------
// # JSON:

// MarshalJSON returns the JSON encoding of zr.Currency.
func (n Currency) MarshalJSON() ([]byte, error) {
	// TODO: using fmt.Sprintf() may slow down performance.
	//       There are faster ways to build a number with 4 decimals.
	//       Create a benchmark to find the fastest method.
	var (
		i   = n.i64 / 1e4   // integer part
		d   = n.i64 - i*1e4 // decimal part
		ret = fmt.Sprintf("%d", i)
	)
	if d != 0 {
		ret += strings.TrimRight(
			fmt.Sprintf("%0.4f", float32(d)/1e4)[1:],
			"0",
		)
	}
	return []byte(ret), nil
} //                                                                 MarshalJSON

// UnmarshalJSON unmarshals a JSON description of zr.Currency.
// This method alters the object's value.
func (n *Currency) UnmarshalJSON(data []byte) error {
	//   ^  don't remove pointer receiver, it is necessary
	var num float64
	err := mod.json.Unmarshal(data, &num)
	if err != nil {
		return err
	}
	n.i64 = int64(num * 1e4)
	return nil
} //                                                               UnmarshalJSON

// -----------------------------------------------------------------------------
// # Helper Function

// currencyOverflow returns a negative (math.MinInt64)
// or positive (math.MaxInt64) overflow value.
//
// It also calls Error() to write an error in the log. The error is logged,
// but the function returns an int64 overflow value instead of an error.
//
// isNegative: should specify if the number is negative or positive.
// a: an array of values used to build the error message.
func currencyOverflow(isNegative bool, a ...interface{}) Currency {
	ar := []interface{}{EOverflow + ": "}
	ar = append(ar, a...)
	mod.Error(ar...)
	if isNegative {
		return Currency{math.MinInt64}
	}
	return Currency{math.MaxInt64}
} //                                                            currencyOverflow

// end
