// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-03-04 20:18:51 FFA6AD                                zr/[go_lang.go]
// -----------------------------------------------------------------------------

package zr

// # Interface
//   GoStringerEx interface
//
// # Functions
//   GoName(s string) string
//   GoString(val interface{}, optIndentAt ...int) string
//   WriteGoString(val interface{}, indentAt int, buf *bytes.Buffer)
//
// # Helper Function
//   indentPos(optIndentAt []int) int

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
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
	var ret = str.Trim(s, SPACES)
	if len(ret) == 0 {
		return ""
	}
	// replace undserscores with spaces to isolate words
	if str.Contains(ret, "_") {
		ret = str.Replace(ret, "_", " ", -1)
	}
	// capitalize the fist letter of every word
	ret = TitleCase(ret)
	//
	// the word 'ID' should be all-capital
	if ContainsI(ret, "id") {
		ret = ReplaceWord(ret, "id", "ID", IgnoreCase)
	}
	// remove spaces to get a camel-case string
	if str.Contains(ret, " ") {
		ret = str.Replace(ret, " ", "", -1)
	}
	return ret
} //                                                                      GoName

// GoString converts fundamental types to strings in Go Language syntax.
// You can copy its output and paste in source code if needed.
//
// If the type of 'val' implements GoStringer or zr.GoStrngerEx
// interfaces, uses the method provided by the interface.
//
// optIndentAt: omit this optional argument to place all output
//              on one line, or specify 0 or more tab positions
//              to indent the output on multiple lines.
//
func GoString(val interface{}, optIndentAt ...int) string {
	var useGoStringer = true
	var indentAt = indentPos(optIndentAt)
	var buf bytes.Buffer
	WriteGoString(val, useGoStringer, indentAt, &buf)
	return buf.String()
} //                                                                    GoString

// WriteGoString writes a fundamental type
// in Go language syntax to a buffer.
//
// It is called by zr.GoString() function and various
// types' GoString() methods to generate their results.
//
// val: the value being read
//
// useGoStringer: when true, calls GoString() or GoStringEx() if
//                val implements any of these methods.
//
// indentAt: specifies if output should be on a single
//           line (-1) or indented to a number of tab stops.
//
// buf: pointer to output buffer
func WriteGoString(
	val interface{},
	useGoStringer bool,
	indentAt int,
	buf *bytes.Buffer,
) {
	// write multiple strings to buffer
	var ws = func(a ...string) {
		for _, s := range a {
			buf.WriteString(s)
		}
	}
	var writeGoString = func(val interface{}) {
		WriteGoString(val, useGoStringer, indentAt, buf)
	}
	if val == nil {
		ws("nil")
		return
	}
	if useGoStringer {
		switch val := val.(type) {
		case GoStringerEx:
			ws(val.GoStringEx(indentAt))
			return
		case fmt.GoStringer:
			ws(val.GoString())
			return
		}
	}
	var v = reflect.ValueOf(val)
	var t = reflect.TypeOf(val)
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			ws("true")
		} else {
			ws("false")
		}
		return
	case reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ws(String(v.Int()))
		return
	case reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ws(String(v.Uint()))
		return
	case reflect.Uintptr:
		//TODO: handle Uintptr
		break
	case reflect.Float32, reflect.Float64:
		ws(String(v.Float()))
		return
	case reflect.Complex64, reflect.Complex128, reflect.Array,
		reflect.Chan, reflect.Func, reflect.Interface:
		//TODO: handle multiple types
		break
	case reflect.Map:
		ws("map[", t.Key().String(), "]", t.Elem().String(), "{")
		//
		// since MapKeys are returned in no specific order,
		// append each key-value in map to a string array
		// then sort it to ensure the result is consistent
		var lines = make([]string, 0, v.Len())
		for _, key := range v.MapKeys() {
			lines = append(lines,
				TabSpace+GoString(key.Interface())+": "+
					GoString(v.MapIndex(key).Interface())+",")
		}
		sort.Strings(lines)
		//
		// write out the array
		for _, s := range lines {
			ws(LF, s)
		}
		ws(LF, "}")
		return
	case reflect.Ptr:
		writeGoString(v.Elem().Interface())
		return
	case reflect.Slice:
		ws(t.String(), "{")
		var manyLines = v.Len() > 0 && v.Index(0).Kind() == reflect.Slice
		for i, n := 0, v.Len(); i < n; i++ {
			if i > 0 {
				ws(",")
			}
			if manyLines {
				ws(LF, TabSpace)
			} else if i > 0 {
				ws(" ")
			}
			writeGoString(v.Index(i).Interface())
		}
		if manyLines {
			ws(",", LF)
		}
		ws("}")
		return
	case reflect.String:
		ws(`"`, str.Replace(val.(string), `"`, `\"`, -1), `"`)
		return
	case reflect.Struct:
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
	case reflect.UnsafePointer:
		break //TODO: reflect.UnsafePointer
	}
	// finally, try using fmt.Stringer (treat 'val' as a string)
	if val, ok := val.(fmt.Stringer); ok {
		ws(GoString(val.String()))
		return
	}
	// if 'val' is still not processed, log an error, try to use fmt.Sprint()
	mod.Error("Type", t, "(kind:", v.Kind(), ") not handled:", val)
	ws("(", fmt.Sprint(val), ")")
} //                                                               WriteGoString

// -----------------------------------------------------------------------------
// # Helper Function

// indentPos helps to read the indent position from the optinal argument
func indentPos(optIndentAt []int) int {
	var ret = -1
	var n = len(optIndentAt)
	switch {
	case n == 1:
		ret = optIndentAt[0]
	case n > 1:
		mod.Error(EInvalidArg, "optIndentAt", ":", optIndentAt)
	}
	return ret
} //                                                                   indentPos

//eof
