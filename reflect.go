// -----------------------------------------------------------------------------
// ZR Library                                                    zr/[reflect.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # Slices
//   IndexSliceStructValue(
//       input interface{},
//       fieldName string,
//       checker func(value interface{}) bool,
//   ) (int, error)
//
// # Structs
//   DescribeStruct(structPtr interface{}) string
//   GetStructInt(structPtr interface{}, field string) (int, bool)
//   GetStructString(structPtr interface{}, field string) (string, bool)
//   GetStructType(structPtr interface{}) (reflect.Type, bool)
//   SetStructInt(structPtr interface{}, field string, val int) bool
//   SetStructString(structPtr interface{}, field, val string) bool

import (
	"fmt"
	"reflect"
	"strings"
)

// -----------------------------------------------------------------------------
// # Slices

// IndexSliceStructValue _ _
func IndexSliceStructValue(
	input interface{},
	fieldName string,
	checker func(value interface{}) bool,
) (int, error) {
	if input == nil {
		return -1, Error(EInvalidArg, "^input", "is nil")
	}
	if fieldName == "" {
		return -1, Error(EInvalidArg, "^fieldName", "is nil")
	}
	if checker == nil {
		return -1, Error(EInvalidArg, "^checker", "is nil")
	}
	slice := reflect.ValueOf(input)
	if slice.Kind() != reflect.Slice {
		return -1, Error("Input is not a slice: It is", slice.Kind())
	}
	for i := 0; i < slice.Len(); i++ {
		it := slice.Index(i)
		if it.Kind() != reflect.Struct {
			return -1, Error("Item", i, "is not a struct: it is:", it.Kind())
		}
		val := reflect.Indirect(it)
		for j := 0; j < val.NumField(); j++ {
			if val.Type().Field(j).Name != fieldName {
				continue
			}
			if checker(val.Field(j).Interface()) {
				return i, nil
			}
		}
	}
	return -1, nil
} //                                                       IndexSliceStructValue

// -----------------------------------------------------------------------------
// # Structs

// DescribeStruct _ _
func DescribeStruct(structPtr interface{}) string {
	// the function is recursive, so it needs to be declared before use
	var unwrap func(interface{}, int) string
	unwrap = func(structPtr interface{}, indentAt int) string {
		ty, ok := GetStructType(structPtr)
		if !ok {
			return ""
		}
		max := 0 // maximum width of a column
		for i, count := 0, ty.NumField(); i < count; i++ {
			l := len(ty.Field(i).Name)
			if l > max {
				max = l
			}
		}
		s := fmt.Sprintf("%s {\r\n", ty.Name())
		printf := func(format string, a ...interface{}) {
			s += strings.Repeat(TabSpace, indentAt) + fmt.Sprintf(format, a...)
		}
		indentAt += 1
		for i, count := 0, ty.NumField(); i < count; i++ {
			var (
				name = ty.Field(i).Name
				pad  = strings.Repeat(" ", max-len(name))
				val  = reflect.ValueOf(structPtr).Elem().Field(i)
			)
			if val.Kind() == reflect.Slice {
				{
					stype := reflect.TypeOf(val.Interface()).String()
					if strings.Contains(stype, "[]main.") {
						stype = strings.ReplaceAll(stype, "[]main.", "[]")
					}
					printf("%s: %s {\r\n", name, stype)
					indentAt += 1
				}
				for rowNo := 0; rowNo < val.Len(); rowNo++ {
					val := val.Index(rowNo).Addr().Interface()
					printf("%s", unwrap(val, indentAt)+",\r\n")
				}
				indentAt -= 1
				printf("},\r\n")
			} else if val.Kind() == reflect.String {
				printf("%s:%s %q,\r\n", name, pad, val)
			} else {
				printf("%s:%s %v,\r\n", name, pad, val)
			}
		}
		indentAt -= 1
		printf("}")
		return s
	}
	return unwrap(structPtr, 0) + "\r\n"
} //                                                              DescribeStruct

// GetStructInt returns the string value from the named field in a struct.
// If the field is not found, returns ("", false).
func GetStructInt(structPtr interface{}, field string) (int, bool) {
	ty, ok := GetStructType(structPtr)
	if !ok {
		return 0, false
	}
	for i, count := 0, ty.NumField(); i < count; i++ {
		if ty.Field(i).Name == field {
			val := reflect.ValueOf(structPtr).Elem().Field(i)
			if val, ok := val.Interface().(int); ok {
				return val, true
			}
		}
	}
	return 0, false
} //                                                                GetStructInt

// GetStructString returns the string value from
// the named field in a struct. If the field is not found,
// returns ("", false).
func GetStructString(structPtr interface{}, field string) (string, bool) {
	ty, ok := GetStructType(structPtr)
	if !ok {
		return "", false
	}
	for i, count := 0, ty.NumField(); i < count; i++ {
		if ty.Field(i).Name == field {
			val := reflect.ValueOf(structPtr).Elem().Field(i)
			if s, ok := val.Interface().(string); ok {
				return s, true
			}
		}
	}
	return "", false
} //                                                             GetStructString

// GetStructType gets the reflection type of a pointer to a struct,
// or returns (nil, false) if it does not point to a struct.
func GetStructType(structPtr interface{}) (reflect.Type, bool) {
	ty := reflect.TypeOf(structPtr)
	if ty.Kind() != reflect.Ptr {
		mod.Error(EInvalidArg, "^structPtr", "is not a pointer to a struct;",
			"it is^", ty.Kind())
		return nil, false
	}
	if ty.Elem().Kind() != reflect.Struct {
		mod.Error(EInvalidArg, "^structPtr", "is not a pointer to a struct;",
			"it is^", ty.Elem().Kind())
		return nil, false
	}
	ty = ty.Elem()
	return ty, true
} //                                                               GetStructType

// SetStructInt _ _
func SetStructInt(structPtr interface{}, field string, val int) bool {
	ty, ok := GetStructType(structPtr)
	if !ok {
		return false
	}
	for i, count := 0, ty.NumField(); i < count; i++ {
		if ty.Field(i).Name == field {
			reflect.ValueOf(structPtr).Elem().Field(i).SetInt(int64(val))
			return true
		}
	}
	return false
} //                                                                SetStructInt

// SetStructString _ _
func SetStructString(structPtr interface{}, field, val string) bool {
	ty, ok := GetStructType(structPtr)
	if !ok {
		return false
	}
	for i, count := 0, ty.NumField(); i < count; i++ {
		if ty.Field(i).Name == field {
			reflect.ValueOf(structPtr).Elem().Field(i).SetString(val)
			return true
		}
	}
	return false
} //                                                             SetStructString

// end
