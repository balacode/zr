// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-23 17:50:57 DA9032                               [zr/currency.go]
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
// # Currency Factory:
//   CurrencyOf(val interface{}) Currency
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
//   (ob Currency) IsEqual(val Currency) bool
//   (ob Currency) IsNegative() bool
//   (ob Currency) IsZero() bool
//   (ob Currency) Overflow() int
//
// # JSON:
//   (ob Currency) MarshalJSON() ([]byte, error)
//   (ob *Currency) UnmarshalJSON(data []byte) error

// # Helper Function
//   currencyOverflow(isNegative bool, a ...interface{}) Currency

// import "encoding/json" // used via mod.json.* mockable
import "bytes"    // standard
import "fmt"      // standard
import "math"     // standard
import "math/big" // standard
import "reflect"  // standard
import "strconv"  // standard

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
// # Currency Factory:

// CurrencyOf converts any compatible value type to a Currency.
// This includes all numeric types and strings. If a string is not numeric,
// logs an error and sets the Currency to zero.
func CurrencyOf(val interface{}) Currency {
	switch val := val.(type) {
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
		return CurrencyOf(int64(*val))
	case *int8:
		return CurrencyOf(*val)
	case *int16:
		return CurrencyOf(*val)
	case *int32:
		return CurrencyOf(*val)
	case *int64:
		return CurrencyOf(*val)
	//
	// unsigned integer pointers
	case *uint:
		return CurrencyOf(uint64(*val))
	case *uint8:
		return CurrencyOf(*val)
	case *uint16:
		return CurrencyOf(*val)
	case *uint32:
		return CurrencyOf(*val)
	case *uint64:
		return CurrencyOf(*val)
	//
	// float pointers
	case *float32:
		return CurrencyOf(*val)
	case *float64:
		return CurrencyOf(*val)
	//
	// strings
	case string:
		var minus bool
		var fract bool
		var dp int
		var ret Currency
		for _, r := range val {
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
				mod.Error("Non-numeric string:^", val)
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
	case *string:
		return CurrencyOf(*val)
	case fmt.Stringer:
		return CurrencyOf(val.String())
	}
	mod.Error("Type", reflect.TypeOf(val), "not handled; =", val)
	return Currency{}
} //                                                                  CurrencyOf

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
	var retBuf = bytes.NewBuffer(make([]byte, 0, 25))
	var ws = retBuf.WriteString
	var wr = retBuf.WriteRune
	var intLen = 0
	var intPart = ob.val / cur4d         // integer part of the number
	var decPart = ob.val - intPart*cur4d // decimal part (as an int)
	var neg = ob.val < 0                 // is it negative? use absolute value
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
		var write = false
		var power = int64(100000000000000) // 10^14
		var digits = intLen % 3
		if digits == 0 {
			digits = 3
		}
		for power > 0 {
			var n = intPart / power
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
		var power = int64(1000) // 10^3
		var unfixed = decPart > 0 && decimalPlaces < 0
		if unfixed {
			decimalPlaces = 4
		}
		if decimalPlaces > 0 {
			ws(".")
		}
		for decimalPlaces > 0 {
			decimalPlaces--
			var n = int64(0)
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
	var n = ob.val
	if n < 0 {
		n = -ob.val
	}
	var bigUnits = n / cur4d
	var smlUnits = (n - bigUnits*cur4d) / 100
	var hasOnly = str.HasSuffix(str.ToLower(fmt), "only")
	if hasOnly {
		fmt = fmt[:len(fmt)-4]
	}
	var getPart = func(partNo int) string {
		var parts = str.Split(fmt, ";")
		if partNo < 0 || partNo >= len(parts) {
			return ""
		}
		return parts[partNo]
	}
	var big1 = getPart(0)
	var bigN = getPart(1)
	var sml1 = getPart(2)
	var smlN = getPart(3)
	var ret = ""
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
	if hasOnly && len(str.Trim(ret, SPACES)) > 0 {
		ret += " Only"
	}
	return ret
} //                                                                   InWordsEN

// String returns a string representing the currency value
// and implements the Stringer Interface.
func (ob Currency) String() string {
	var i = ob.val / cur4d              // integer value
	var d = ob.val - i*cur4d            // decimal value
	var sint = strconv.FormatInt(i, 10) // integer part
	var sdec string                     // decimal part
	if d != 0 {                         // format decimals
		if d < 0 { // adjust for negative value
			d = -d
			if i == 0 {
				sint = "-" + sint
			}
		}
		sdec = "." + str.TrimRight(
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
		var a = ob.val
		var b = cur.val
		//
		// return zero if either number is zero
		if a == 0 || b == 0 {
			return Currency{0}
		}
		// if direct multiplication will overflow int64, use big.Int
		var lim = MaxCurrencyI64 / a
		if lim < 0 {
			lim = -lim
		}
		if b >= lim || b <= -lim {
			var n = big.NewInt(a)
			n.Mul(n, big.NewInt(b))
			n.Div(n, bigCur4d)
			//
			// if result can't be stored in Currency, return overflow
			var overflow = !n.IsInt64()
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
	var a = float64(ob.val)
	for _, b := range multiply {
		// check for negative or positive overflow
		var lim = MaxCurrencyI64 / a
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
		var a = ob.val
		var b = itm.val
		var c = a + b
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
	var a = ob.val
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
		var a = ob.val
		var b = n.val
		var c = a - b
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

// -----------------------------------------------------------------------------
// # JSON:

// MarshalJSON returns the JSON encoding of zr.Currency.
func (ob Currency) MarshalJSON() ([]byte, error) {
	//TODO: using fmt.Sprintf() may slow down performance.
	//      There are faster ways to build a number with 4 decimals.
	//      Create a benchmark to find the fastest method.
	//
	var i = ob.val / cur4d   // integer part
	var d = ob.val - i*cur4d // decimal part
	var ret = fmt.Sprintf("%d", i)
	if d != 0 {
		ret += str.TrimRight(
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
	var err = mod.json.Unmarshal(data, &n)
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
