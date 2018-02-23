// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-23 17:50:57 DAF4C0                              [zr/constants.go]
// -----------------------------------------------------------------------------

// # Error Message Constants
// # String Constants

package zr

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
const ENil = "Value is nil"

// ENilReceiver indicates a method call on a nil object.
const ENilReceiver = "nil receiver"

// ENoDef __
const ENoDef = "Not defined"

// ENotFound __
const ENotFound = "Not found"

// ENotHandled __
const ENotHandled = "Not handled"

// EOverflow indicates an arithmetic overflow.
const EOverflow = "Overflow"

// -----------------------------------------------------------------------------
// # String Constants

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

//end
