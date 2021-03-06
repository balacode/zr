// -----------------------------------------------------------------------------
// ZR Library                                                    zr/[go_lang.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # Interface
//   GoStringerEx interface
//
// # Functions
//   GoName(s string) string
//   GoString(value interface{}, optIndentAt ...int) string
//   WriteGoString(value interface{}, indentAt int, buf *bytes.Buffer)
//
// # Helper Function
//   indentPos(optIndentAt []int) int

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// -----------------------------------------------------------------------------
// # Interface

// GoStringerEx interface is implemented by objects that can output
// their definitions in Go language syntax. It extends the standard
// fmt.GoStringer interface with an optional indent parameter.
type GoStringerEx interface {
	GoStringEx(indentAt int) string
} //                                                                GoStringerEx

// -----------------------------------------------------------------------------
// # Functions

// GoName converts a name to a Go language convention name.
// It removes underscores from names and changes names to 'TitleCase'.
func GoName(s string) string {
	ret := strings.TrimSpace(s)
	if len(ret) == 0 {
		return ""
	}
	// replace undserscores with spaces to isolate words
	if strings.Contains(ret, "_") {
		ret = strings.ReplaceAll(ret, "_", " ")
	}
	// capitalize the fist letter of every word
	ret = TitleCase(ret)
	//
	// the word 'ID' should be all-capital
	if ContainsI(ret, "id") {
		ret = ReplaceWord(ret, "id", "ID", IgnoreCase)
	}
	// remove spaces to get a camel-case string
	if strings.Contains(ret, " ") {
		ret = strings.ReplaceAll(ret, " ", "")
	}
	return ret
} //                                                                      GoName

// GoString converts fundamental types to strings in Go Language syntax.
// You can copy its output and paste in source code if needed.
//
// If the type of value implements fmt.GoStringer or zr.GoStringerEx
// interfaces, uses the method provided by the interface.
//
// optIndentAt: omit this optional argument to place all output
//              on one line, or specify 0 or more tab positions
//              to indent the output on multiple lines.
//
func GoString(value interface{}, optIndentAt ...int) string {
	var (
		useGoStringer = true
		indentAt      = indentPos(optIndentAt)
		buf           bytes.Buffer
	)
	WriteGoString(value, useGoStringer, indentAt, &buf)
	return buf.String()
} //                                                                    GoString

// WriteGoString writes a fundamental type
// in Go language syntax to a buffer.
//
// It is called by zr.GoString() function and various
// types' GoString() methods to generate their results.
//
// value: the value being read
//
// useGoStringer: when true, calls GoString() or GoStringEx() if
//                value implements any of these methods.
//
// indentAt: specifies if output should be on a single
//           line (-1) or indented to a number of tab stops.
//
// buf: pointer to output buffer
func WriteGoString(
	value interface{},
	useGoStringer bool,
	indentAt int,
	buf *bytes.Buffer,
) {
	// write multiple strings to buffer
	ws := func(a ...string) {
		for _, s := range a {
			buf.WriteString(s)
		}
	}
	writeGoString := func(value interface{}) {
		WriteGoString(value, useGoStringer, indentAt, buf)
	}
	if value == nil {
		ws("nil")
		return
	}
	if useGoStringer {
		switch v := value.(type) {
		case GoStringerEx:
			{
				ws(v.GoStringEx(indentAt))
				return
			}
		case fmt.GoStringer:
			ws(v.GoString())
			return
		}
	}
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	switch v.Kind() {
	case reflect.Bool:
		{
			if v.Bool() {
				ws("true")
			} else {
				ws("false")
			}
			return
		}
	case reflect.Int, reflect.Int64, reflect.Int32,
		reflect.Int16, reflect.Int8:
		{
			ws(String(v.Int()))
			return
		}
	case reflect.Uint, reflect.Uint64, reflect.Uint32,
		reflect.Uint16, reflect.Uint8:
		{
			ws(String(v.Uint()))
			return
		}
	case reflect.Uintptr:
		{
			// TODO: handle Uintptr
			break
		}
	case reflect.Float64, reflect.Float32:
		{
			ws(String(v.Float()))
			return
		}
	case reflect.Complex64, reflect.Complex128, reflect.Array,
		reflect.Chan, reflect.Func, reflect.Interface:
		{
			// TODO: handle multiple types
			break
		}
	case reflect.Map:
		{
			ws("map[", t.Key().String(), "]", t.Elem().String(), "{")
			//
			// since MapKeys are returned in no specific order,
			// append each key-value in map to a string array
			// then sort it to ensure the result is consistent
			lines := make([]string, 0, v.Len())
			for _, key := range v.MapKeys() {
				lines = append(lines,
					TabSpace+GoString(key.Interface())+": "+
						GoString(v.MapIndex(key).Interface())+",")
			}
			sort.Strings(lines)
			//
			// write out the array
			for _, s := range lines {
				ws("\n", s)
			}
			ws("\n}")
			return
		}
	case reflect.Ptr:
		{
			writeGoString(v.Elem().Interface())
			return
		}
	case reflect.Slice:
		{
			ws(t.String(), "{")
			manyLines := v.Len() > 0 && v.Index(0).Kind() == reflect.Slice
			for i, n := 0, v.Len(); i < n; i++ {
				if i > 0 {
					ws(",")
				}
				if manyLines {
					ws("\n", TabSpace)
				} else if i > 0 {
					ws(" ")
				}
				writeGoString(v.Index(i).Interface())
			}
			if manyLines {
				ws(",\n")
			}
			ws("}")
			return
		}
	case reflect.String:
		{
			ws(`"`, strings.ReplaceAll(value.(string), `"`, `\"`), `"`)
			return
		}
	case reflect.Struct:
		{
			ws(t.String(), "{")
			for i, n := 0, t.NumField(); i < n; i++ {
				if !v.Field(i).CanInterface() {
					continue
				}
				if i > 0 {
					ws(", ")
				}
				ws(t.Field(i).Name, ": ")
				writeGoString(v.Field(i).Interface())
			}
			ws("}")
			return
		}
	case reflect.UnsafePointer:
		break // TODO: reflect.UnsafePointer
	}
	// finally, try using fmt.Stringer (treat value as a string)
	if v, ok := value.(fmt.Stringer); ok {
		ws(GoString(v.String()))
		return
	}
	// if value is still not processed, log an error, try to use fmt.Sprint()
	mod.Error("Type", t, "(kind:", v.Kind(), ") not handled:", value)
	ws("(", fmt.Sprint(value), ")")
} //                                                               WriteGoString

// -----------------------------------------------------------------------------
// # Helper Function

// indentPos helps to read the indent position from the optinal argument
func indentPos(optIndentAt []int) int {
	ret := -1
	n := len(optIndentAt)
	switch {
	case n == 1:
		{
			ret = optIndentAt[0]
		}
	case n > 1:
		mod.Error(EInvalidArg, "optIndentAt", ":", optIndentAt)
	}
	return ret
} //                                                                   indentPos

//eof
