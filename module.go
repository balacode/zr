// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 62ED08                                 [zr/module.go]
// -----------------------------------------------------------------------------

package zr

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"time"
)

// # Library Version
//   VersionTime() string
//
// # String Constants
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
	var VersionTime = "2018-02-27 17:03"
	return VersionTime
} //                                                                 VersionTime

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

// ModReset restored all mocked functions to the original standard functions.
func (ob *thisMod) Reset() {
	ob.Error = Error
	ob.json.Unmarshal = json.Unmarshal
	ob.rand.Read = rand.Read
} //                                                                       Reset

//end
