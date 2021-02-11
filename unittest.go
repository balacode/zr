// -----------------------------------------------------------------------------
// ZR Library                                                   zr/[unittest.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # Global Counter
//   TErrorCount int
//
// # TStringer
//   NewTStringer(s string) TStringer
//   (ob *TStringer) Set(s string)
//   (ob TStringer) GoString() string
//   (ob TStringer) String() string
//
// # Assertion Functions
//   TArrayEqual(t *testing.T, expect, value interface{}) bool
//   TBytesEqual(t *testing.T, a, b []byte)
//   TEqual(t *testing.T, result interface{}, expect interface{}) bool
//   TFalse(t *testing.T, result bool) bool
//   TTrue(t *testing.T, result bool) bool
//
// # Other Functions
//   TBegin(t *testing.T)
//   TBeginError()
//   TCaller() string
//   TCheckError(t *testing.T, expectMessages ...string)
//   TFail(t *testing.T, a ...interface{})
//   TFailf(t *testing.T, format string, a ...interface{})
//
// # Helper Functions
//   failedFuncAndLine() (funcName string, lineNo int)
//   failedLocation() string

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

// -----------------------------------------------------------------------------
// # Global Counter

// TErrorCount records the number of logged
// errors when TBeginError() is called. This is
// then compared when TCheckError() gets called.
var TErrorCount int

// -----------------------------------------------------------------------------
// # TStringer

// TStringer is a mock type that provides the fmt.Stringer interface.
type TStringer struct {
	s string
} //                                                                   TStringer

// NewTStringer _ _
func NewTStringer(s string) TStringer {
	return TStringer{s}
} //                                                                NewTStringer

// Set sets the string that will be returned
// by the String() and GoString() methods.
func (ob *TStringer) Set(s string) {
	ob.s = s
} //                                                                         Set

// GoString _ _
func (ob TStringer) GoString() string {
	return ob.s
} //                                                                    GoString

// String _ _
func (ob TStringer) String() string {
	return ob.s
} //                                                                      String

// -----------------------------------------------------------------------------
// # Assertion Functions

// TArrayEqual checks if two arrays are equal
func TArrayEqual(t *testing.T, expect, value interface{}) bool {
	// TODO: TArrayEqual is not necessary, can be handled by Equal()
	a := fmt.Sprintf("%v", expect)
	b := fmt.Sprintf("%v", value)
	if a != b {
		fmt.Printf("%s expected: %s got: %s\r\n", TCaller(), a, b)
		t.Fail()
		return false
	}
	return true
} //                                                                 TArrayEqual

// TBytesEqual checks if two byte slices have the same length and content
func TBytesEqual(t *testing.T, a, b []byte) {
	// TODO: TBytesEqual() is not necessary, can be handled by Equal()
	if !bytes.Equal(a, b) {
		t.Errorf("FAILED AT LINE %d: %v != %v", LineNo(2), a, b)
	}
} //                                                                 TBytesEqual

// TEqual asserts that result is equal to expect.
func TEqual(t *testing.T, result interface{}, expect interface{}) bool {
	makeStr := func(value interface{}) string {
		var ret string
		switch v := value.(type) {
		case nil:
			{
				ret = "nil"
			}
		case bool:
			if v {
				ret = "true"
			} else {
				ret = "false"
			}
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr:
			{
				ret = fmt.Sprintf("%d", v)
			}
		case float64, float32:
			ret := fmt.Sprintf("%f", v)
			if strings.Contains(ret, ".") {
				for strings.HasSuffix(ret, "0") {
					ret = ret[:len(ret)-1]
				}
				for strings.HasSuffix(ret, ".") {
					ret = ret[:len(ret)-1]
				}
			}
		case string:
			{
				return v
			}
		case time.Time: // use date part without time and time zone
			ret = v.Format(time.RFC3339)[:19] // format: 2006-01-02T15:04:05Z07:00
			if strings.HasSuffix(ret, "T00:00:00") {
				ret = ret[:10]
			}
		case fmt.Stringer:
			{
				ret = v.String()
			}
		case fmt.GoStringer:
			{
				ret = v.GoString()
			}
		// TODO: add handling of various arrays of simple types [in TEqual()]
		case []string:
			{
				var buf bytes.Buffer
				buf.WriteString("[")
				for i, s := range v {
					if i != 0 {
						buf.WriteString(", ")
					}
					buf.WriteString(`"`)
					buf.WriteString(strings.Replace(s, `"`, `\"`, -1))
					buf.WriteString(`"`)
				}
				buf.WriteString("]")
				ret = buf.String()
			}
		default:
			ret = fmt.Sprintf("(type: %v value: %v)",
				reflect.TypeOf(value), value)
		}
		return ret
	}
	if makeStr(result) != makeStr(expect) {
		t.Logf("%s\r\n"+
			"EXPECTED: %s\r\n"+
			"RETURNED: %s\r\n",
			"CALLER: "+TCaller(),
			makeStr(expect),
			makeStr(result),
		)
		t.Fail()
		return false
	}
	return true
} //                                                                      TEqual

// TFalse asserts that the result value is false.
// If the assertion fails, invokes t.Error() to fail the unit test.
// Returns true if the assertion passed, or false if it failed.
func TFalse(t *testing.T, result bool) bool {
	if result == false {
		return true
	}
	return TEqual(t, result, false)
} //                                                                      TFalse

// TTrue asserts that the result value is true.
// If the assertion fails, invokes t.Error() to fail the unit test.
// Returns true if the assertion passed, or false if it failed.
func TTrue(t *testing.T, result bool) bool {
	if result == true {
		return true
	}
	return TEqual(t, result, true)
} //                                                                       TTrue

// -----------------------------------------------------------------------------
// # Other Functions

// TBegin prints a heading with the name of the tested module.
func TBegin(t *testing.T) {
	// get list of calls on the call stack, remove calls into this file
	list := CallerList()
	for len(list) > 0 &&
		(strings.TrimSpace(list[0]) == "" ||
			strings.Contains(list[0], "TBegin")) {
		list = list[1:]
	}
	// pick the first list item as the test's name
	testName := "<test-name>"
	if len(list) > 0 {
		testName = strings.TrimSpace(list[0])
	}
	// remove package name
	if strings.Contains(testName, ".") {
		testName = strings.Split(testName, ".")[1]
	}
	// remove line number
	if strings.Contains(testName, ":") {
		testName = strings.Split(testName, ":")[0]
	}
	// align the name to the right (within 80 columns)
	if len(testName) < 80 {
		testName = strings.Repeat(" ", 80-len(testName)) + testName
	}
	t.Log(testName)
} //                                                                      TBegin

// TBeginError begins a check for an error condition.
func TBeginError() {
	DisableErrors()
	TErrorCount = GetErrorCount()
} //                                                                 TBeginError

// TCaller returns the name of the unit test function.
func TCaller() string {
	for _, funcName := range CallerList() {
		if strings.HasPrefix(funcName, "TCaller") ||
			strings.HasPrefix(funcName, "TEqual") ||
			strings.HasPrefix(funcName, "TFail") ||
			strings.HasPrefix(funcName, "TFalse") ||
			strings.HasPrefix(funcName, "TTrue") ||
			strings.Contains(funcName, ".func") ||
			!strings.Contains(funcName, ".Test") {
			continue
		}
		return funcName
	}
	return "<no-caller>"
} //                                                                     TCaller

// TCheckError finishes a check for an error. Between calls to TBeginError()
// and TCheckError() there should be one error (logged with Error()).
// If not, call TFail() to trigger t.Error() because the test has failed.
//
// Place calls to TBeginError() and TCheckError() within your unit tests.
// For example:
// func TestMyFunc(t *testing.T) {
//     TBeginError()  // <- begin check for error
//     MyFunc(-1)     // <- calling this function with -1 should log an error
//     TCheckError(t) // <- finish check (if no error, call t.Error()
// }
//
func TCheckError(t *testing.T, expectMessages ...string) {
	EnableErrors()
	count := GetErrorCount()
	if count != TErrorCount+1 {
		TFail(t, "Expected 1 error, but got ", count-TErrorCount)
		return
	}
	// this should never happen!
	if count < TErrorCount {
		TFail(t,
			"Unexpected: global logged error count changed",
			" from ", TErrorCount, " to ", count,
		)
		return
	}
	// if any error message fragments were supplied,
	// check if any is found in the logged error
	found := len(expectMessages) == 0
	errm := GetLastLogMessage()
	for _, find := range expectMessages {
		find = strings.ToUpper(find)
		if strings.Contains(strings.ToUpper(errm), find) {
			found = true
			break
		}
	}
	if found {
		return
	}
	if len(errm) == 0 {
		TFail(t, "No error message")
		return
	}
	TFail(t, "Unexpected error '", errm, "'")
} //                                                                 TCheckError

// TFail _ _
func TFail(t *testing.T, a ...interface{}) {
	msg := failedLocation() + fmt.Sprint(a...)
	t.Error(msg)
	t.Fail()
} //                                                                       TFail

// TFailf _ _
func TFailf(t *testing.T, format string, a ...interface{}) {
	format = failedLocation() + format
	t.Errorf(format, a...)
	t.Fail()
} //                                                                      TFailf

// -----------------------------------------------------------------------------
// # Helper Functions

// failedFuncAndLine returns the function
// name and line number of a failed test.
func failedFuncAndLine() (funcName string, lineNo int) {
	ar := strings.Split(TCaller(), ":")
	if len(ar) > 0 {
		funcName = ar[0]
	}
	if len(ar) > 1 {
		ln, err := strconv.Atoi(ar[1])
		if err != nil {
			ln = -1
		}
		lineNo = ln
	}
	return funcName, lineNo
} //                                                           failedFuncAndLine

// failedLocation returns the location message
// used in TFail() and TFailf()
func failedLocation() string {
	funcName, lineNo := failedFuncAndLine()
	return "\n" +
		"FAILED FUNC: " + funcName + "\n" +
		"FAILED LINE: " + strconv.Itoa(lineNo) + "\n"
} //                                                              failedLocation

// TODO: BeginError() can redefine mod.Error()
//      so there is no need for ErrorCount

// TODO: write code to test the testing module (unittest.go)

//end                                 "Who is going to test the testers?" - Anon
