// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-16 17:40:50 A77356                                   zr/[bool.go]
// -----------------------------------------------------------------------------

package zr

//   Bool(val interface{}) bool
//   IsBool(val interface{}) bool
//   TrueCount(values ...bool) int

import (
	"fmt"
	"reflect"
	"strings"
)

// Bool converts primitive types to a boolean value:
// Converts nil to false.
// Converts all nonzero numbers to true and zeros to false.
// Converts only strings that equal 'true' (case-insensitive) to true.
func Bool(val interface{}) bool {
	switch val := val.(type) {
	case nil:
		{
			return false
		}
	// boolean:
	case bool:
		{
			return val
		}
	case *bool:
		if val != nil {
			return *val
		}
	// strings:
	case string:
		switch strings.ToUpper(strings.TrimSpace(val)) {
		case "FALSE", "0", "":
			{
				return false
			}
		case "TRUE", "1", "-1":
			return true
		}
	case *string:
		if val != nil {
			return Bool(*val)
		}
	case fmt.Stringer:
		{
			return Bool(val.String())
		}
	// signed integers:
	case int:
		{
			return val != 0
		}
	case int8:
		{
			return val != 0
		}
	case int16:
		{
			return val != 0
		}
	case int32:
		{
			return val != 0
		}
	case int64:
		{
			return val != 0
		}
	case *int:
		if val != nil {
			return *val != 0
		}
	case *int8:
		if val != nil {
			return *val != 0
		}
	case *int16:
		if val != nil {
			return *val != 0
		}
	case *int32:
		if val != nil {
			return *val != 0
		}
	case *int64:
		if val != nil {
			return *val != 0
		}
	// unsigned integers:
	case uint:
		{
			return val != 0
		}
	case uint64:
		{
			return val != 0
		}
	case uint32:
		{
			return val != 0
		}
	case uint16:
		{
			return val != 0
		}
	case uint8:
		{
			return val != 0
		}
	case *uint:
		if val != nil {
			return *val != 0
		}
	case *uint64:
		if val != nil {
			return *val != 0
		}
	case *uint32:
		if val != nil {
			return *val != 0
		}
	case *uint16:
		if val != nil {
			return *val != 0
		}
	case *uint8:
		if val != nil {
			return *val != 0
		}
	// floating-point numbers:
	case float64:
		{
			return val != 0
		}
	case float32:
		{
			return val != 0
		}
	case *float64:
		if val != nil {
			return *val != 0
		}
	case *float32:
		if val != nil {
			return *val != 0
		}
	}
	mod.Error("Can not convert", reflect.TypeOf(val), "to bool:", val)
	return false
} //                                                                        Bool

// IsBool returns true if 'val' can be converted to a boolean.
// If 'val' is any numeric type or bool, returns true.
// If 'val' is nil, returns false.
// If 'val' is a string, returns true if it is "TRUE", "FALSE", "0", or "1".
// If 'val' is any other type, returns false.
func IsBool(val interface{}) bool {
	switch val := val.(type) {
	case nil:
		{
			return false
		}
	// strings:
	case string:
		switch strings.ToUpper(strings.TrimSpace(val)) {
		case "FALSE", "TRUE", "0", "1":
			{
				return true
			}
		default:
			return false
		}
	case *string:
		if val != nil {
			return IsBool(*val)
		}
	case fmt.Stringer:
		{
			return IsBool(val.String())
		}
	case int, int64, int32, int16, int8, float64, float32,
		uint, uint64, uint32, uint16, uint8,
		bool,
		*int, *int64, *int32, *int16, *int8, *float64, *float32,
		*uint, *uint64, *uint32, *uint16, *uint8,
		*bool:
		{
			return true
		}
	}
	return false
} //                                                                      IsBool

// TrueCount __
func TrueCount(values ...bool) int {
	ret := 0
	for _, v := range values {
		if v {
			ret++
		}
	}
	return ret
} //                                                                   TrueCount

//end
