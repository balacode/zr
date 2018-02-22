// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      [zr/func.go]
// -----------------------------------------------------------------------------

package zr

// # Utility Functions
//   AppendToTextFile(filename, text string)
//   ByteSizeString(sizeInBytes int64, useSI bool) string
//   ConsumeGB(gigabytes float64)
//   IsNumber(val interface{}) bool

import "bufio"         // standard
import "fmt"           // standard
import "math"          // standard
import "os"            // standard
import "runtime"       // standard
import "runtime/debug" // standard
import "time"          // standard
import "math/big"      // standard
import str "strings"   // standard

// -----------------------------------------------------------------------------
// # Utility Functions

// AppendToTextFile appends 'text' to file named 'filename'.
func AppendToTextFile(filename, text string) {
	var file *os.File
	var err error
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0)
	defer file.Close()
	if err != nil {
		Error("Opening file", filename, ":", err)
		return
	}
	file.WriteString(text)
} //                                                            AppendToTextFile

// ByteSizeString returns a human-friendly byte count string.
// E.g. 1024 gives '1 KiB', 1048576 gives '1 MiB'.
// If you set 'useSI' to true, uses multiples of 1000 (SI units)
// instead of 1024 (binary units) and suffixes 'KB' instead of 'KiB', etc.
func ByteSizeString(sizeInBytes int64, useSI bool) string {
	var n int64 = 1024
	if useSI {
		n = 1000
	}
	if sizeInBytes > -n && sizeInBytes < n {
		return fmt.Sprintf("%d B", sizeInBytes)
	}
	if !useSI && sizeInBytes == math.MinInt64 {
		return "-7.9 EiB"
	}
	var neg = sizeInBytes < 0
	var ret string
	if neg {
		sizeInBytes = -sizeInBytes
		ret = "-"
	}
	for _, group := range []struct {
		unit  string
		scale int64
	}{ //                                 binary:   SI:
		{"E", n * n * n * n * n * n}, //  exabyte   exbibyte
		{"P", n * n * n * n * n},     //  petabyte  pebibyte
		{"T", n * n * n * n},         //  terabyte  tebibyte
		{"G", n * n * n},             //  gigabyte  gibibyte
		{"M", n * n},                 //  megabyte  mebibyte
		{"K", n},                     //  kilobyte  kibibyte
	} {
		if sizeInBytes < group.scale {
			continue
		}
		// because Sprintf() rounds numbers up, cut at 1dp before calling it
		// (use either regular arithmetic or math.BigInt if size is very large)
		var cut float64
		if sizeInBytes < math.MaxInt64/1024 {
			cut = float64(sizeInBytes) / float64(group.scale)
			cut = float64(int64(cut*10)) / 10
		} else {
			var sz = big.NewInt(sizeInBytes)
			sz.Mul(sz, big.NewInt(10))
			sz.Div(sz, big.NewInt(group.scale))
			cut = float64(sz.Int64()) / 10
		}
		ret += fmt.Sprintf("%0.1f", cut)
		//
		// remove trailing zero decimal
		if str.HasSuffix(ret, ".0") {
			ret = ret[:len(ret)-2]
		}
		// append SI or binary units
		ret += " " + group.unit
		if useSI {
			ret += "B"
		} else {
			ret += "iB"
		}
		break
	}
	return ret
} //                                                              ByteSizeString

// ConsumeGB consumes, then releases a large amount of RAM for testing.
func ConsumeGB(gigabytes float64) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("before allocation. mem: %s"+LB,
		ByteSizeString(int64(mem.HeapSys), false))
	{
		fmt.Printf("allocating and filling %0.1f GB"+LB, gigabytes)
		var max = int64(gigabytes * 1024 * 1024 * 1024)
		var big = make([]byte, max)
		var b = byte(0)
		for i := int64(0); i < max; i++ {
			big[i] = b
			b++
			if b == 255 {
				b = 0
			}
		}
		runtime.ReadMemStats(&mem)
		fmt.Printf("after allocating, mem: %s"+LB,
			ByteSizeString(int64(mem.HeapSys), false))
		fmt.Printf("some data: %v"+LB, big[:100])
		fmt.Println("press ENTER to release memory")
		bufio.NewReader(os.Stdin).ReadString('\n')
		big = nil
	}
	debug.FreeOSMemory()
	time.Sleep(3 * time.Second)
	runtime.ReadMemStats(&mem)
	fmt.Printf("after releasing. mem: %s"+LB,
		ByteSizeString(int64(mem.HeapSys), false))
	fmt.Println("press ENTER to continue")
	bufio.NewReader(os.Stdin).ReadString('\n')
} //                                                                   ConsumeGB

// IsNumber returns true if 'val' is a number or numeric string,
// or false otherwise. It also accepts pointers to numeric types
// and strings and Stringer. Always returns false if val is nil
// or bool, even though Int() can convert bool to 1 or 0.
func IsNumber(val interface{}) bool {
	const groupSeparatorChar = ','
	const decimalPointChar = '.'
	//
	switch val := val.(type) {
	case int, int8, int16, int32, int64,
		*int, *int8, *int16, *int32, *int64,
		uint, uint8, uint16, uint32, uint64,
		*uint, *uint8, *uint16, *uint32, *uint64,
		float32, float64,
		*float32, *float64:
		return true
	case string:
		var s = str.Trim(val, SPACES)
		if len(s) < 1 {
			return false
		}
		var hasDecPoint, hasDigit, hasSign, prevSep bool
		for _, r := range s {
			switch {
			case r >= '0' && r <= '9':
				hasDigit = true
			case r == groupSeparatorChar:
				// two consecutive group separators make string non-numeric
				if prevSep || !hasDigit {
					return false
				}
				prevSep = true
				continue
			case r == '-' || r == '+':
				if hasSign || hasDigit {
					return false
				}
				hasSign = true
			case r == decimalPointChar:
				if hasDecPoint {
					return false
				}
				hasDecPoint = true
			default:
				return false
			}
			prevSep = false
		}
		return hasDigit
	case *string:
		return IsNumber(*val)
	case fmt.Stringer:
		return IsNumber(val.String())
	}
	return false
} //                                                                    IsNumber

//end
