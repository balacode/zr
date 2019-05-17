// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-17 10:58:17 1998FF                                   zr/[bool.go]
// -----------------------------------------------------------------------------

package zr

//   Bool(value interface{}) bool
//   IsBool(value interface{}) bool
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
func Bool(value interface{}) bool {
	switch v := value.(type) {
	case nil:
		{
			return false
		}
	// boolean:
	case bool:
		{
			return v
		}
	case *bool:
		if v != nil {
			return *v
		}
	// strings:
	case string:
		switch strings.ToUpper(strings.TrimSpace(v)) {
		case "FALSE", "0", "":
			{
				return false
			}
		case "TRUE", "1", "-1":
			return true
		}
	case *string:
		if v != nil {
			return Bool(*v)
		}
	case fmt.Stringer:
		{
			return Bool(v.String())
		}
	// signed integers:
	case int:
		{
			return v != 0
		}
	case int8:
		{
			return v != 0
		}
	case int16:
		{
			return v != 0
		}
	case int32:
		{
			return v != 0
		}
	case int64:
		{
			return v != 0
		}
	case *int:
		if v != nil {
			return *v != 0
		}
	case *int8:
		if v != nil {
			return *v != 0
		}
	case *int16:
		if v != nil {
			return *v != 0
		}
	case *int32:
		if v != nil {
			return *v != 0
		}
	case *int64:
		if v != nil {
			return *v != 0
		}
	// unsigned integers:
	case uint:
		{
			return v != 0
		}
	case uint64:
		{
			return v != 0
		}
	case uint32:
		{
			return v != 0
		}
	case uint16:
		{
			return v != 0
		}
	case uint8:
		{
			return v != 0
		}
	case *uint:
		if v != nil {
			return *v != 0
		}
	case *uint64:
		if v != nil {
			return *v != 0
		}
	case *uint32:
		if v != nil {
			return *v != 0
		}
	case *uint16:
		if v != nil {
			return *v != 0
		}
	case *uint8:
		if v != nil {
			return *v != 0
		}
	// floating-point numbers:
	case float64:
		{
			return v != 0
		}
	case float32:
		{
			return v != 0
		}
	case *float64:
		if v != nil {
			return *v != 0
		}
	case *float32:
		if v != nil {
			return *v != 0
		}
	}
	mod.Error("Can not convert", reflect.TypeOf(value), "to bool:", value)
	return false
} //                                                                        Bool

// IsBool returns true if value can be converted to a boolean.
// If value is any numeric type or bool, returns true.
// If value is nil, returns false.
// If value is a string, returns true if it is "TRUE", "FALSE", "0", or "1".
// If value is any other type, returns false.
func IsBool(value interface{}) bool {
	switch v := value.(type) {
	case nil:
		{
			return false
		}
	// strings:
	case string:
		switch strings.ToUpper(strings.TrimSpace(v)) {
		case "FALSE", "TRUE", "0", "1":
			{
				return true
			}
		default:
			return false
		}
	case *string:
		if v != nil {
			return IsBool(*v)
		}
	case fmt.Stringer:
		{
			return IsBool(v.String())
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
