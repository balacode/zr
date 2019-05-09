// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-09 18:09:23 DD6957                                  zr/[debug.go]
// -----------------------------------------------------------------------------

package zr

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

// ConsumeGB consumes, then releases a large amount of RAM for testing.
func ConsumeGB(gigabytes float64) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("before allocation. mem: %s\r\n",
		ByteSizeString(int64(mem.HeapSys), false))
	{
		fmt.Printf("allocating and filling %0.1f GB\r\n", gigabytes)
		max := int64(gigabytes * 1024 * 1024 * 1024)
		big := make([]byte, max)
		b := byte(0)
		for i := int64(0); i < max; i++ {
			big[i] = b
			b++
			if b == 255 {
				b = 0
			}
		}
		runtime.ReadMemStats(&mem)
		fmt.Printf("after allocating, mem: %s\r\n",
			ByteSizeString(int64(mem.HeapSys), false))
		fmt.Printf("some data: %v\r\n", big[:100])
		fmt.Println("press ENTER to release memory")
		bufio.NewReader(os.Stdin).ReadString('\n')
		big = nil
	}
	debug.FreeOSMemory()
	time.Sleep(3 * time.Second)
	runtime.ReadMemStats(&mem)
	fmt.Printf("after releasing. mem: %s\r\n",
		ByteSizeString(int64(mem.HeapSys), false))
	fmt.Println("press ENTER to continue")
	bufio.NewReader(os.Stdin).ReadString('\n')
} //                                                                   ConsumeGB

// DebugString takes any kind of value and returns a string.
// Calls itself when describing slices, etc.
func DebugString(val interface{}, optIndentAt ...int) string {
	var indentAt int
	switch n := len(optIndentAt); {
	case n == 1:
		indentAt = optIndentAt[0]
	case n > 1:
		mod.Error(EInvalidArg, "optIndentAt", ":", optIndentAt)
	}
	var retBuf bytes.Buffer
	wr := retBuf.WriteRune
	ws := retBuf.WriteString
	switch val := val.(type) {
	// nil:
	case nil:
		return "nil"
	case bool:
		if val {
			return "true"
		}
		return "false"
	case CaseMode:
		if val == IgnoreCase {
			return "IgnoreCase"
		}
		if val == MatchCase {
			return "MatchCase"
		}
		return "INVALID" + fmt.Sprintf("%v", val)
	//
	// delegated to String():
	case *bool,
		int, *int, int8, *int8, int16, *int16, int32, *int32, int64, *int64,
		uint, *uint,
		uint8, *uint8, uint16, *uint16, uint32, *uint32, uint64, *uint64,
		float32, *float32, float64, *float64:
		return String(val)
	//
	// string types:
	case string:
		vals := []rune{'\a', '\b', '\f', '\n', '\r', 't', '\v'}
		chars := []rune{'a', 'b', 'f', 'n', 'r', 't', 'v'}
	mainLoop:
		for _, ch := range val {
			for i, cc := range vals {
				if ch == cc {
					wr('\\')
					wr(chars[i])
					continue mainLoop
				}
			}
			if ch < 32 || ch > 127 {
				wr('.')
				continue
			}
			wr(ch)
		}
	case []string:
		ws(fmt.Sprintf("[%d] ", len(val)))
		isMany := len(val) > 1
		for i, val := range val {
			if isMany {
				ws("\n")
				ws(strings.Repeat(TabSpace, indentAt+1))
				ws(fmt.Sprintf("%d:", i))
			}
			ws(DebugString(val, indentAt))
		}
	case [][]string:
		ws(fmt.Sprintf("[%d] ", len(val)))
		isMany1 := len(val) > 1
		for i1, val := range val { // range [][]string
			if isMany1 {
				ws("\n")
				ws(strings.Repeat(TabSpace, indentAt+1))
			}
			isMany2 := len(val) > 1
			ws(fmt.Sprintf("%d: [%d]", i1, len(val)))
			for i2, val := range val { // range []string
				if isMany2 {
					ws("\n")
					ws(strings.Repeat(" ", indentAt+2))
				}
				ws(fmt.Sprintf(" %d:", i2))
				ws(DebugString(val, indentAt+3)) // string
			}
		}
	case []byte:
		ws(fmt.Sprintf("[%d] ", len(val)))

	default:
		// TODO: remove this code later
		a := fmt.Sprintf("%v", val)
		b := fmt.Sprint(val)
		if a != b {
			fmt.Printf("Sprintf->%q Sprint->%q %s", a, b, Callers())
		}
		mod.Error("Type", reflect.TypeOf(val), "not handled; =", val)
		return "(" + fmt.Sprint(val) + ")"
	}
	return retBuf.String()
} //                                                                 DebugString

// DV displays human-friendly values for debugging.
//
// label:  the value's label, you should usually specify
//         the name of a variable or some tag here.
//
// values: one or more values you want to display.
func DV(label string, values ...interface{}) {
	if len(values) == 0 {
		fmt.Println(label)
		return
	}
	fmt.Print(label)
	for _, val := range values {
		typeName := "<" + reflect.TypeOf(val).String() + ">"
		changeType := func(find, repl string) {
			for strings.Contains(typeName, find) {
				typeName = strings.Replace(typeName, find, repl, -1)
			}
		}
		changeType("bool", "t")
		changeType("byte", "b")
		changeType("int", "i")
		changeType("string", "s")
		fmt.Print(" " + typeName + ": " + DebugString(val, 0))
	}
	fmt.Println()
} //                                                                          DV

//eof
