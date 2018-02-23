// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-23 17:50:57 43BA59                            [zr/module_test.go]
// -----------------------------------------------------------------------------

package zr

// # Library Version
//   Test_mdle_VersionTime_
//
// # Constants and Variables
//   Test_mdle_consts_
//   Test_mdle_vars_
//
// # Global Settings
//   Test_mdle_GetDebugMode_
//   Test_mdle_SetDebugMode_

/*
to test all items in strings.go use:
	go test --run Test_mdle_

to generate a test coverage report use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

import "fmt"     // standard
import "strconv" // standard
import "testing" // standard

// -----------------------------------------------------------------------------
// # Library Version

// go test --run Test_mdle_VersionTime_
func Test_mdle_VersionTime_(t *testing.T) {
	TBegin(t)
	// VersionTime_() string
	//
	var s = VersionTime()
	TTrue(t, IsDate(s))
	TTrue(t, len(VersionTime()) == 16)
	//
	// The format expected from VersionTime is:
	// "DDDD-MM-YY hh:mm"
	var yr, mt, dy, hr, mn int
	yr, _ = strconv.Atoi(s[0:4])
	mt, _ = strconv.Atoi(s[5:7])
	dy, _ = strconv.Atoi(s[8:10])
	hr, _ = strconv.Atoi(s[11:13])
	mn, _ = strconv.Atoi(s[14:16])
	TTrue(t, yr >= 2017 && yr <= 9999)
	TTrue(t, s[4:5] == "-")
	TTrue(t, mt >= 1 && mt <= 9999)
	TTrue(t, s[7:8] == "-")
	TTrue(t, dy >= 1 && dy <= 31)
	TTrue(t, s[10:11] == " ")
	TTrue(t, hr >= 0 && hr <= 59)
	TTrue(t, s[10:11] == " ")
	TTrue(t, mn >= 0 && hr <= 59)
} //                                                      Test_mdle_VersionTime_

// -----------------------------------------------------------------------------
// # Constants and Variables

// go test --run Test_mdle_consts_
func Test_mdle_consts_(t *testing.T) {
	TBegin(t)
	//
	// error message constants
	TEqual(t, EFailedParsing, ("Failed parsing"))
	TEqual(t, EFailedReading, ("Failed reading"))
	TEqual(t, EFailedWriting, ("Failed writing"))
	TEqual(t, EInvalid, ("Invalid"))
	TEqual(t, EInvalidArg, ("Invalid argument"))
	TEqual(t, EInvalidType, ("Invalid type"))
	TEqual(t, ENil, ("Value is nil"))
	TEqual(t, ENilReceiver, ("nil receiver"))
	TEqual(t, ENoDef, ("Not defined"))
	TEqual(t, ENotFound, ("Not found"))
	TEqual(t, ENotHandled, ("Not handled"))
	//
	// other constants
	TEqual(t, LB, ("\r\n"))
} //                                                           Test_mdle_consts_

// go test --run Test_mdle_vars_
func Test_mdle_vars_(t *testing.T) {
	TBegin(t)
	//
	TEqual(t, PL, (fmt.Println))
	TEqual(t, VL, (VerboseLog))
	//
	TFalse(t, lastTraceCallTime.IsZero())
	TTrue(t, lastTraceCallTime.Year() >= 2018)
} //                                                             Test_mdle_vars_

// -----------------------------------------------------------------------------
// # Global Settings

// go test --run Test_mdle_GetDebugMode_
func Test_mdle_GetDebugMode_(t *testing.T) {
	TBegin(t)
	// GetDebugMode() bool
	//
	debugMode = true
	TEqual(t, GetDebugMode(), (true))
	//
	debugMode = false
	TEqual(t, GetDebugMode(), (false))
} //                                                     Test_mdle_GetDebugMode_

// go test --run Test_mdle_SetDebugMode_
func Test_mdle_SetDebugMode_(t *testing.T) {
	TBegin(t)
	// SetDebugMode(val bool)
	//
	SetDebugMode(true)
	TEqual(t, debugMode, (true))
	//
	SetDebugMode(false)
	TEqual(t, debugMode, (false))
} //                                                     Test_mdle_SetDebugMode_

//end
