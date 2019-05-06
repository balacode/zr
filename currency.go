// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-06 06:25:43 CC9049                               zr/[currency.go]
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
//   CurrencyOf(value interface{}) Currency
//   CurrencyOfS(s string) Currency
//   CurrencyRaw(raw int64) Currency
//
// # String Output:
//   (ob Currency) GoString() string
//   (ob Currency) Fmt(decimalPlaces int) string
//   (ob Currency) InWordsEN(fmt string) string
//   (ob Currency) String() string
//
// # Division:
//   (ob Currency) Div(divide ...Currency) Currency
//   (ob Currency) DivFloat(divide ...float64) Currency
//   (ob Currency) DivInt(divide ...int) Currency
//
// # Multiplication:
//   (ob Currency) Mul(multiply ...Currency) Currency
//   (ob Currency) MulFloat(multiply ...float64) Currency
//   (ob Currency) MulInt(multiply ...int) Currency
//
// # Addition:
//   (ob Currency) Add(add ...Currency) Currency
//   (ob Currency) AddFloat(add ...float64) Currency
//   (ob Currency) AddInt(add ...int) Currency
//
// # Subtraction:
//   (ob Currency) Sub(subtract ...Currency) Currency
//   (ob Currency) SubFloat(subtract ...float64) Currency
//   (ob Currency) SubInt(subtract ...int) Currency
//
// # Information:
//   (ob Currency) Float64() float64
//   (ob Currency) Int() int
//   (ob Currency) Int64() int64
//   (ob Currency) IsEqual(val interface{}) bool
//   (ob Currency) IsGreater(val interface{}) bool
//   (ob Currency) IsGreaterOrEqual(val interface{}) bool
//   (ob Currency) IsLesser(val interface{}) bool
//   (ob Currency) IsLesserOrEqual(val interface{}) bool
//   (ob Currency) IsNegative() bool
//   (ob Currency) IsZero() bool
//   (ob Currency) Overflow() int
//   (ob Currency) Raw() int64
//
// # JSON:
//   (ob Currency) MarshalJSON() ([]byte, error)
//   (ob *Currency) UnmarshalJSON(data []byte) error

// # Helper Function
//   currencyOverflow(isNegative bool, a ...interface{}) Currency

import (
	// "encoding/json" // used via mod.json.* mockable
	"bytes"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

// -----------------------------------------------------------------------------
// # Constants:

// CurrencyIntLimit specifies the highest (and lowest
// when negative) integer component that currency
// can hold, about 922.33 trillion.
// The exact number is 922(t)337(b)203(m)685,476.
//
// This limit exists because an int64, on which Currency is based, can hold
// up to 9,223,372,036,854,775,807. Out of this 4 digits are used for the
// decimal part, i.e. 922,337,203,685,477.5807. The limit is set to this
// number minus 1, so that all decimals from .0000 to .9999. can be used.
const CurrencyIntLimit = 922337203685476

// MinCurrencyI64 is the lowest internal value that Currency can hold.
const MinCurrencyI64 = -9223372036854769999

// MaxCurrencyI64 is the highest internal value that Currency can hold.
const MaxCurrencyI64 = 9223372036854769999

// cur4d scales the int64 to provide 4 decimal places.
const cur4d = 10000

// bigCur4d scales the int64 to provide 4 decimal places.
var bigCur4d = big.NewInt(cur4d)

// -----------------------------------------------------------------------------
// # Currency Type:

// Currency represents a currency value with up to four decimal places.
// It is stored internally as a 64-bit integer. This gives it a range
// from -922,337,203,685,477.5808 to 922,337,203,685,477.5807
type Currency struct {
	val int64
} //                                                                    Currency

// -----------------------------------------------------------------------------
// # Currency Factories:

// CurrencyOf converts any compatible value type to a Currency.
// This includes all numeric types and strings. If a string is
// not numeric, logs an error and sets the Currency to zero.
func CurrencyOf(value interface{}) Currency {
	switch val := value.(type) {
	//
	// Currency already?
	case Currency:
		return val
	//
	// integers
	case int8:
		return Currency{int64(val) * cur4d}
	case int16:
		return Currency{int64(val) * cur4d}
	case int32: //                              int32 and int could exceed range
		return CurrencyOf(int64(val))
	case int:
		return CurrencyOf(int64(val))
	case int64:
		if val < -CurrencyIntLimit || val > CurrencyIntLimit {
			return currencyOverflow(val < 0, EOverflow, ": ", val)
		}
		return Currency{int64(val) * cur4d}
	//
	// unsigned integers
	case uint8:
		return Currency{int64(val) * cur4d}
	case uint16:
		return Currency{int64(val) * cur4d}
	case uint32: //                           uint32 and uint could exceed range
		return CurrencyOf(uint64(val))
	case uint:
		return CurrencyOf(uint64(val))
	case uint64:
		if val > CurrencyIntLimit {
			return currencyOverflow(false, EOverflow, "uint64: ", val)
		}
		return Currency{int64(val) * cur4d}
	//
	// float
	case float32:
		if val < -float32(CurrencyIntLimit)-0.9999 ||
			val > float32(CurrencyIntLimit)+0.9999 {
			return currencyOverflow(val < 0, EOverflow, "float32: ", val)
		}
		return Currency{int64(float64(val) * cur4d)}
	case float64:
		if val < -float64(CurrencyIntLimit)-0.9999 ||
			val > float64(CurrencyIntLimit)+0.9999 {
			return currencyOverflow(val < 0, EOverflow, "float64: ", val)
		}
		return Currency{int64(val * cur4d)}
	//
	// integer pointers
	case *int:
		if val != nil {
			return CurrencyOf(int64(*val))
		}
	case *int8:
		if val != nil {
			return CurrencyOf(*val)
		}
	case *int16:
		if val != nil {
			return CurrencyOf(*val)
		}
	case *int32:
		if val != nil {
			return CurrencyOf(*val)
		}
	case *int64:
		if val != nil {
			return CurrencyOf(*val)
		}
	// unsigned integer pointers
	case *uint:
		if val != nil {
			return CurrencyOf(uint64(*val))
		}
	case *uint8:
		if val != nil {
			return CurrencyOf(*val)
		}
	case *uint16:
		if val != nil {
			return CurrencyOf(*val)
		}
	case *uint32:
		if val != nil {
			return CurrencyOf(*val)
		}
	case *uint64:
		if val != nil {
			return CurrencyOf(*val)
		}
	// float pointers
	case *float32:
		if val != nil {
			return CurrencyOf(*val)
		}
	case *float64:
		if val != nil {
			return CurrencyOf(*val)
		}
	// strings
	case string:
		return CurrencyOfS(val)
	case *string:
		if val != nil {
			return CurrencyOfS(*val)
		}
	case fmt.Stringer:
		return CurrencyOfS(val.String())
	}
	mod.Error("Type", reflect.TypeOf(value), "not handled; =", value)
	return Currency{}
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
			ret.val *= 10
			ret.val += int64(r - '0')
		} else if r != ',' &&
			r != ' ' && r != '\a' && r != '\b' && r != '\f' &&
			r != '\n' && r != '\r' && r != '\t' && r != '\v' {
			mod.Error("Non-numeric string:^", s)
			break
		}
	}
	for dp < 4 {
		dp++
		ret.val *= 10
	}
	if minus {
		ret.val = -ret.val
	}
	return ret
} //                                                                 CurrencyOfS

// CurrencyRaw initializes a currency value from a scaled 64-bit integer.
// The decimal point is moved left 4 decimal places. For example, a raw value
// of 15500 results in a currency value of 1.55
func CurrencyRaw(raw int64) Currency {
	return Currency{val: raw}
} //                                                                 CurrencyRaw

// -----------------------------------------------------------------------------
// # String Output:

// GoString outputs the value as a Go language string,
// It implements the fmt.GoStringer interface.
func (ob Currency) GoString() string {
	return "zr.CurrencyOf(" + ob.String() + ")"
} //                                                                    GoString

// Fmt returns the currency value as a a string
// delimited with commas (grouped every three digits)
// and having the specified number of decimal places.
// When decimalPlaces is negative, the resulting
// number's decimals will vary.
func (ob Currency) Fmt(decimalPlaces int) string {
	var (
		retBuf  = bytes.NewBuffer(make([]byte, 0, 25))
		ws      = retBuf.WriteString
		wr      = retBuf.WriteRune
		intLen  = 0
		intPart = ob.val / cur4d         // integer part of the number
		decPart = ob.val - intPart*cur4d // decimal part (as an int)
		neg     = ob.val < 0             // is it negative? use absolute value
	)
	if neg {
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
		write := false
		power := int64(100000000000000) // 10^14
		digits := intLen % 3
		if digits == 0 {
			digits = 3
		}
		for power > 0 {
			n := intPart / power
			if n > 0 {
				write = true
			}
			intPart -= n * power
			power /= 10
			if !write {
				continue
			}
			wr(rune(n + 48))
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
			n := int64(0)
			if power > 0 {
				n = decPart / power
				decPart -= n * power
				power /= 10
			}
			wr(rune(n + 48))
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
func (ob Currency) InWordsEN(fmt string) string {
	n := ob.val
	if n < 0 {
		n = -ob.val
	}
	bigUnits := n / cur4d
	smlUnits := (n - bigUnits*cur4d) / 100
	hasOnly := strings.HasSuffix(strings.ToLower(fmt), "only")
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
	if hasOnly && len(strings.Trim(ret, SPACES)) > 0 {
		ret += " Only"
	}
	return ret
} //                                                                   InWordsEN

// String returns a string representing the currency value
// and implements the Stringer Interface.
func (ob Currency) String() string {
	var (
		i    = ob.val / cur4d           // integer value
		d    = ob.val - i*cur4d         // decimal value
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
			strconv.FormatInt(d+cur4d, 10)[1:],
			"0",
		)
	}
	return sint + sdec
} //                                                                      String

// -----------------------------------------------------------------------------
// # Division:

// Div divides a currency object by one or more currency values
// and returns the result. The object's value isn't changed.
func (ob Currency) Div(divide ...Currency) Currency {
	for _, n := range divide {
		ob.val *= cur4d
		ob.val /= n.val
	}
	return ob
} //                                                                         Div

// DivFloat divides a currency object by one or more floating-point
// numbers and returns the result. The object's value isn't changed.
func (ob Currency) DivFloat(divide ...float64) Currency {
	for _, n := range divide {
		ob.val *= cur4d
		ob.val /= int64(n * cur4d)
	}
	return ob
} //                                                                    DivFloat

// DivInt divides a currency object by one or more integer values
// and returns the result. The object's value isn't changed.
func (ob Currency) DivInt(divide ...int) Currency {
	for _, val := range divide {
		ob.val *= cur4d
		ob.val /= (int64(val) * cur4d)
	}
	return ob
} //                                                                      DivInt

// -----------------------------------------------------------------------------
// # Multiplication:

// Mul multiplies a currency object by one or more currency values
// and returns the result. The object's value isn't changed.
func (ob Currency) Mul(multiply ...Currency) Currency {
	for _, cur := range multiply {
		a := ob.val
		b := cur.val
		//
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
			n := big.NewInt(a)
			n.Mul(n, big.NewInt(b))
			n.Div(n, bigCur4d)
			//
			// if result can't be stored in Currency, return overflow
			//
			// TODO: IsInt64() is not available in older Go versions ``
			overflow := !n.IsInt64()
			var ret int64
			if !overflow {
				ret = n.Int64()
				if ret < MinCurrencyI64 || ret > MaxCurrencyI64 {
					overflow = true
				}
			}
			if overflow {
				return currencyOverflow(
					(a < 0 || b < 0) && (a > 0 || b > 0),
					EOverflow, ": ", a, " * ", b, " = ", n,
				)
			}
			return Currency{ret}
		}
		ob.val = a * b / cur4d
	}
	return ob
} //                                                                         Mul

// MulFloat multiplies a currency object by one or more floating-point
// numbers and returns the result. The object's value isn't changed.
func (ob Currency) MulFloat(multiply ...float64) Currency {
	a := float64(ob.val)
	for _, b := range multiply {
		// check for negative or positive overflow
		lim := MaxCurrencyI64 / a
		if lim < 0 {
			lim = -lim
		}
		if b < -lim || b > lim {
			return currencyOverflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				"Overflow: ", a, " * ", b, " = ", a*b,
			)
		}
		// multiply using int64, if there is no overflow
		ob.val = int64(a * b)
	}
	return ob
} //                                                                    MulFloat

// MulInt multiplies a currency object by one or more integer values
// and returns the result. The object's value isn't changed.
func (ob Currency) MulInt(multiply ...int) Currency {
	for _, n := range multiply {
		ob = ob.Mul(Currency{int64(n * cur4d)})
	}
	return ob
} //                                                                      MulInt

// -----------------------------------------------------------------------------
// # Addition:

// Add adds one or more currency values and returns a new Currency object.
// The value in the object to which this method is applied isn't changed.
// If there is an overflow, sets the Currency's internal value to
// math.MinInt64 or math.MaxInt64 depending on if the result is negative.
func (ob Currency) Add(add ...Currency) Currency {
	for _, itm := range add {
		a := ob.val
		b := itm.val
		c := a + b
		//
		// check for overflow
		if c < MinCurrencyI64 || (a < 0 && b < 0 && b < (MinCurrencyI64-a)) {
			return currencyOverflow(true, EOverflow, ": ", a, " + ", b)
		}
		if c > MaxCurrencyI64 || (a > 0 && b > 0 && b > (MaxCurrencyI64-a)) {
			return currencyOverflow(false, EOverflow, ": ", a, " + ", b)
		}
		ob.val = c
	}
	return ob
} //                                                                         Add

// AddFloat adds one or more floating-point numbers to a currency
// object and returns the result. The object's value isn't changed.
func (ob Currency) AddFloat(add ...float64) Currency {
	const lim float64 = math.MaxInt64 / cur4d
	a := ob.val
	for _, b := range add {
		//
		// check for overflow
		if b < -lim || b > lim {
			return currencyOverflow(
				(a < 0 || b < 0) && (a > 0 || b > 0),
				EOverflow, ": ", a, " * ", b, " = ", float64(a)*b,
			)
		}
		// use Add() because it has other overflow checks
		ob = ob.Add(Currency{int64(b * cur4d)})
	}
	return ob
} //                                                                    AddFloat

// AddInt adds one or more integer values to a currency object
// and returns the result. The object's value isn't changed.
func (ob Currency) AddInt(add ...int) Currency {
	for _, n := range add {
		ob = ob.Add(Currency{int64(n) * cur4d})
	}
	return ob
} //                                                                      AddInt

// -----------------------------------------------------------------------------
// # Subtraction:

// Sub subtracts one or more currency values from a currency object
// and returns the result. The object's value isn't changed.
func (ob Currency) Sub(subtract ...Currency) Currency {
	for _, n := range subtract {
		a := ob.val
		b := n.val
		c := a - b
		//
		// check for overflow
		if c < MinCurrencyI64 || (a < 0 && b > 0 && b > (-MinCurrencyI64+a)) {
			return currencyOverflow(true, EOverflow, ": ", a, " - ", b)
		}
		if c > MaxCurrencyI64 || (a > 0 && b < 0 && b < (-MaxCurrencyI64+a)) {
			return currencyOverflow(false, EOverflow, ": ", a, " - ", b)
		}
		ob.val = c
	}
	return ob
} //                                                                         Sub

// SubFloat subtracts one or more floating-point numbers from a currency
// object and returns the result. The object's value isn't changed.
func (ob Currency) SubFloat(subtract ...float64) Currency {
	for _, n := range subtract {
		ob.val -= int64(n * cur4d)
	}
	return ob
} //                                                                    SubFloat

// SubInt subtracts one or more integer values from a currency object
// and returns the result. The object's value isn't changed.
func (ob Currency) SubInt(subtract ...int) Currency {
	for _, n := range subtract {
		ob.val -= int64(n) * cur4d
	}
	return ob
} //                                                                      SubInt

// -----------------------------------------------------------------------------
// # Information:

// Float64 returns the currency value as a float64 value.
func (ob Currency) Float64() float64 {
	return float64(ob.val) / cur4d
} //                                                                     Float64

// Int returns the currency value as an int value.
func (ob Currency) Int() int64 {
	return ob.val / cur4d
} //                                                                         Int

// Int64 returns the currency value as an int64 value.
func (ob Currency) Int64() int64 {
	return ob.val / cur4d
} //                                                                       Int64

// IsEqual returns true if the value of the currency object is negative.
func (ob Currency) IsEqual(val interface{}) bool {
	return ob.val == CurrencyOf(val).val
} //                                                                     IsEqual

// IsGreater returns true if the object is greater than val.
func (ob Currency) IsGreater(val interface{}) bool {
	return ob.val > CurrencyOf(val).val
} //                                                                   IsGreater

// IsGreaterOrEqual returns true if the object is greater or equal to val.
func (ob Currency) IsGreaterOrEqual(val interface{}) bool {
	return ob.val >= CurrencyOf(val).val
} //                                                            IsGreaterOrEqual

// IsLesser returns true if the object is lesser than val.
func (ob Currency) IsLesser(val interface{}) bool {
	return ob.val < CurrencyOf(val).val
} //                                                                    IsLesser

// IsLesserOrEqual returns true if the object is lesser or equal to val.
func (ob Currency) IsLesserOrEqual(val interface{}) bool {
	return ob.val <= CurrencyOf(val).val
} //                                                             IsLesserOrEqual

// IsNegative returns true if the value of the currency object is negative.
func (ob Currency) IsNegative() bool {
	return ob.val < 0
} //                                                                  IsNegative

// IsZero returns true if the value of the currency object is zero.
func (ob Currency) IsZero() bool {
	return ob.val == 0
} //                                                                      IsZero

// Overflow returns 1 if the currency contains a positive overflow,
// -1 if it contains a negative overflow, or zero if there is no overflow or
// underflow. Overflow and underflows occur when an arithmeric operation's
// result exceeds the storage capacity of Currency.
func (ob Currency) Overflow() int {
	if ob.val > MaxCurrencyI64 {
		return 1
	}
	if ob.val < MinCurrencyI64 {
		return -1
	}
	return 0
} //                                                                    Overflow

// Raw returns the raw int64 used to store the currency value
func (ob Currency) Raw() int64 {
	return ob.val
} //                                                                         Raw

// -----------------------------------------------------------------------------
// # JSON:

// MarshalJSON returns the JSON encoding of zr.Currency.
func (ob Currency) MarshalJSON() ([]byte, error) {
	// TODO: using fmt.Sprintf() may slow down performance.
	//       There are faster ways to build a number with 4 decimals.
	//       Create a benchmark to find the fastest method.
	//
	i := ob.val / cur4d   // integer part
	d := ob.val - i*cur4d // decimal part
	ret := fmt.Sprintf("%d", i)
	if d != 0 {
		ret += strings.TrimRight(
			fmt.Sprintf("%0.4f", float32(d)/cur4d)[1:],
			"0",
		)
	}
	return []byte(ret), nil
} //                                                                 MarshalJSON

// UnmarshalJSON unmarshals a JSON description of zr.Currency.
// This method alters the object's value.
func (ob *Currency) UnmarshalJSON(data []byte) error {
	//   ^  don't remove pointer receiver, it is necessary
	if ob == nil {
		return mod.Error(ENilReceiver)
	}
	var n float64
	err := mod.json.Unmarshal(data, &n)
	if err != nil {
		return err
	}
	ob.val = int64(n * cur4d)
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
	mod.Error(a...)
	if isNegative {
		return Currency{math.MinInt64}
	}
	return Currency{math.MaxInt64}
} //                                                            currencyOverflow

//end
