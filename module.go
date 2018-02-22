// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    [zr/module.go]
// -----------------------------------------------------------------------------

package zr

import "crypto/rand"   // standard
import "encoding/json" // standard
import "fmt"           // standard
import "time"          // standard

// # Library Version
//   VersionTime() string
//
// # Other Constants
// # Global Settings
//   GetDebugMode() bool
//   SetDebugMode(val bool)
//
// # Module Variables

// -----------------------------------------------------------------------------
// # Library Version

// VersionTime returns the library version as a date/time string.
func VersionTime() string {
	// name capitalized to make it easy to locate
	var VersionTime = "2018-02-22 00:34"
	return VersionTime
} //                                                                 VersionTime

// -----------------------------------------------------------------------------
// # Error Message Constants

// EFailedParsing __
const EFailedParsing = "Failed parsing"

// EFailedReading __
const EFailedReading = "Failed reading"

// EFailedWriting __
const EFailedWriting = "Failed writing"

// EInvalid __
const EInvalid = "Invalid"

// EInvalidArg __
const EInvalidArg = "Invalid argument"

// EInvalidType __
const EInvalidType = "Invalid type"

// ENil __
const ENil = "Value is nil:"

// ENilObject indicates a method call on a nil object.
const ENilObject = "nil receiver"

// ENoDef __
const ENoDef = "Not defined:"

// ENotFound __
const ENotFound = "Not found:"

// ENotHandled __
const ENotHandled = "Not handled:"

// EOverflow indicates an arithmetic overflow.
const EOverflow = "Overflow"

// -----------------------------------------------------------------------------
// # Other Constants

// CR is a string with a single carriage return character (decimal 13, hex 0D)
const CR = "\r"

// LF is a string with a single line feed character (decimal 10, hex 0A)
const LF = "\n"

// LB specifies a line break string.
// On Windows it is a pair of CR and LF.
// CR is decimal 13, hex 0D.
// LF is decimal 10, hex 0A.
const LB = "\r\n"

// SPACES is a string of all white-space characters,
// which includes spaces, tabs, and newline characters.
const SPACES = " \a\b\f\n\r\t\v"

// -----------------------------------------------------------------------------
// # Global Settings

// GetDebugMode __
func GetDebugMode() bool {
	return debugMode
} //                                                                GetDebugMode

// SetDebugMode __
func SetDebugMode(val bool) {
	debugMode = val
} //                                                                SetDebugMode

// -----------------------------------------------------------------------------
// # Module Variables

// PL is fmt.Println() but is used only for debugging.
var PL = fmt.Println

// VL is VerboseLog() but is used only for debugging.
var VL = VerboseLog

// debugMode makes the library print verbose output to assist debugging.
var debugMode = false

// lastTraceCallTime __
var lastTraceCallTime = time.Now()

// TabSpace specifies the string to use as a single tab
// (used by GoStringEx() to indent code)
var TabSpace = "    "

// -----------------------------------------------------------------------------
// # Function Proxy Variables (for mocking)

type jsonMod struct {
	Unmarshal func([]byte, interface{}) error
}

type randMod struct {
	Read func([]byte) (int, error)
}

type thisMod struct {
	Error func(args ...interface{}) error
	//
	// standard library modules:
	json jsonMod
	rand randMod
}

var mod = thisMod{
	Error: Error,
	json: jsonMod{
		Unmarshal: json.Unmarshal,
	},
	rand: randMod{
		Read: rand.Read,
	},
}

// ModReset restores all mocked functions to the original standard functions.
func (ob *thisMod) Reset() {
	ob.Error = Error
	ob.json.Unmarshal = json.Unmarshal
	ob.rand.Read = rand.Read
} //                                                                       Reset

//end
