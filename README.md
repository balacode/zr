## [Zr] Zircon-Go Library  
[![Go Report Card](https://goreportcard.com/badge/github.com/balacode/zr)](https://goreportcard.com/report/github.com/balacode/zr)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)  
**Functions and classes to manipulate basic data types, unit test, etc.**  

This library module works on all build platforms. There are other related library packages, some of which may depend on a particular OS, mainly github.com/balacode/zr_win.

## To Install:
Just run:

    go get github.com/balacode/zr

## Summary:

**bool.go**: functions to work with boolean values.

**bytes.go**: a class to handle a block of bytes, with methods to insert, read, delete, etc.

**bytes_func.go**: functions to manipulate byte slices.

**calendar.go**: a class to present multiple dates and values in a calendar grid format.

**currency.go**: a fast data type for working with currency values. It is an int64 adjusted to give 4 fixed decimal places.

**dates.go**: functions to work with dates

**debug.go**: functions to help debugging

**go_lang.go**: convert any value to its representation in Go Language syntax 

**int_tuple.go**: type that provides an integer tuple (a struct made up of two integers)

**logging.go**: provides error/warning logging and related functions.

**numbers.go**: functions to convert numeric types, check if a string is numeric and format numbers.

**reflect.go**: various functions to work with reflection.

**settings.go**: a simple container and interface to read and write settings.

**strings.go**: various functions to work with strings, that are not found in the standard library, for example functions to replace words in strings and make multiple replacements simultaneously.

**string_aligner.go**: aligns strings in columns.

**timer.go**: a class to capture starting and ending times of multiple events and generate a report of total time spent at each stage.

**unittest.go**: various functions to help unit testing.

**uuid.go**: generate UUIDs with UUID() or check if a string is a valid UUID with IsUUID()
