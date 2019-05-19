// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-19 20:25:37 9DC41A                                   zr/[bool.go]
// -----------------------------------------------------------------------------

package zr

//   Bool(value interface{}) bool
//   BoolE(value interface{}) (bool, error)
//   IsBool(value interface{}) bool
//   TrueCount(values ...bool) int

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Bool converts simple types to a boolean value:
//
// Converts nil to false.
// Converts all nonzero numbers to true and zeros to false.
// Converts only strings that equal 'true' (case-insensitive) to true.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
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
	// signed integers:
	case int:
		{
			return v != 0
		}
	case int64:
		{
			return v != 0
		}
	case int32:
		{
			return v != 0
		}
	case int16:
		{
			return v != 0
		}
	case int8:
		{
			return v != 0
		}
	case *int:
		if v != nil {
			return *v != 0
		}
	case *int64:
		if v != nil {
			return *v != 0
		}
	case *int32:
		if v != nil {
			return *v != 0
		}
	case *int16:
		if v != nil {
			return *v != 0
		}
	case *int8:
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

// BoolE converts any simple numeric type or string to bool.
//
// - Dereferences pointers to evaluate the pointed-to type.
// - Converts nil to false.
// - Converts all nonzero numbers to true and zeros to false.
// - Converts only strings that equal 'true', '1' or '-1'
//   (case-insensitive) to true, all other strings to false.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
func BoolE(value interface{}) (bool, error) {
	switch v := value.(type) {
	case nil:
		{
			return false, nil
		}
	case bool:
		{
			return v, nil
		}
	case string:
		switch strings.ToUpper(strings.TrimSpace(v)) {
		case "FALSE", "0", "":
			{
				return false, nil
			}
		case "TRUE", "1", "-1":
			return true, nil
		}
	case int, int64, int32, int16, int8:
		{
			n := reflect.ValueOf(value).Int()
			return n != 0, nil
		}
	case uint, uint64, uint32, uint16, uint8:
		{
			n := reflect.ValueOf(value).Uint()
			return n != 0, nil
		}
	case float64, float32:
		{
			n := reflect.ValueOf(value).Float()
			return n != 0, nil
		}
	}
	// if not converted yet, try to dereference pointer, then convert
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return false, nil
		}
		ret, err := BoolE(v.Elem().Interface())
		if err == nil {
			return ret, nil
		}
	}
	erm := fmt.Sprintf("Can not convert %s to bool: %v",
		reflect.TypeOf(value), value)
	err := errors.New(erm)
	return false, err
} //                                                                       BoolE

// IsBool returns true if value can be converted to a boolean.
//
// If value is any numeric type or bool, returns true.
// If value is nil, returns false.
// If value is a string, returns true if it is "TRUE", "FALSE", "0", or "1".
// If value is any other type, returns false.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
func IsBool(value interface{}) bool {
	switch v := value.(type) {
	case nil:
		{
			return false
		}
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
