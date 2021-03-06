// -----------------------------------------------------------------------------
// ZR Library                                                     zr/[module.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
)

// # Library Version
//   VersionTime() string
//
// # String Constants
// # Global Settings
//   DebugMode() bool
//   SetDebugMode(val bool)
//
// # Module Variables

// -----------------------------------------------------------------------------
// # Library Version

// VersionTime returns the library version as a date/time string.
func VersionTime() string {
	// name capitalized to make it easy to locate
	versionTime := "2021-02-11 15:21"
	return versionTime
} //                                                                 VersionTime

// -----------------------------------------------------------------------------
// # Global Settings

// DebugMode _ _
func DebugMode() bool {
	return debugMode
} //                                                                   DebugMode

// SetDebugMode _ _
func SetDebugMode(val bool) {
	debugMode = val
} //                                                                SetDebugMode

// -----------------------------------------------------------------------------
// # Module Variables

var (
	// PL is fmt.Println() but is used only for debugging.
	PL = fmt.Println

	// VL is VerboseLog() but is used only for debugging.
	VL = VerboseLog

	// debugMode makes the library print verbose output to assist debugging.
	debugMode = false

	// TabSpace specifies the string to use as a single tab
	// (used by GoStringEx() to indent code)
	TabSpace = "    "
)

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

// end
