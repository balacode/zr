// -----------------------------------------------------------------------------
// ZR Library                                                 zr/[bytes_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # Special Constructors
//   Test_byts_BytesAlloc_
//   Test_byts_BytesAllocUTF8_
//   Test_byts_BytesWrap_
//
// # Read-Only Properties / Methods
//   Test_byts_Bytes_Cap_
//   Test_byts_Bytes_FindByte_
//   Test_byts_Bytes_GetByte_
//   Test_byts_Bytes_GetChar_
//   Test_byts_Bytes_Size_
//
// # fmt.Stringer Interface
//   Test_byts_Bytes_String_
//
// # Methods
//   Test_byts_Bytes_Append_
//   Test_byts_Bytes_AppendChar_
//   Test_byts_Bytes_GetBytes_
//   Test_byts_Bytes_Insert_
//   Test_byts_Bytes_Remove_
//   Test_byts_Bytes_Reset_
//   Test_byts_Bytes_Resize_
//   Test_byts_Bytes_SetByte_
//   Test_byts_Bytes_Slice_

//  to test all items in bytes.go use:
//      go test --run Test_byts_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"
	// "bytes"
	// "compress/zlib"
	// "io"
	// "unicode/utf8"
)

// -----------------------------------------------------------------------------
// # Special Constructors

// go test --run Test_byts_BytesAlloc_
func Test_byts_BytesAlloc_(t *testing.T) {
	TBegin(t)
	//
	// BytesAlloc(size int) Bytes
	//
	o := BytesAlloc(4096) // 4KiB
	TEqual(t, cap(o.ar), 4096)
} //                                                       Test_byts_BytesAlloc_

// go test --run Test_byts_BytesAllocUTF8_
func Test_byts_BytesAllocUTF8_(t *testing.T) {
	TBegin(t)
	//
	// BytesAllocUTF8(s string) Bytes
	//
	o := BytesAllocUTF8("Привет !")
	TEqual(t, len(o.ar), 14)
} //                                                   Test_byts_BytesAllocUTF8_

// go test --run Test_byts_BytesWrap_
func Test_byts_BytesWrap_(t *testing.T) {
	TBegin(t)
	//
	// BytesWrap(ar []byte) Bytes
	//
	ar := []byte{'H', 'e', 'l', 'l', 'o', '!'}
	o := BytesWrap(ar)
	TEqual(t, len(o.ar), 6)
	//
	// TODO: Test how 'ar' is affected when object is changed
} //                                                        Test_byts_BytesWrap_

// -----------------------------------------------------------------------------
// # Read-Only Properties / Methods

// go test --run Test_byts_Bytes_Cap_
func Test_byts_Bytes_Cap_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) Cap() int
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.Cap()
		TCheckError(t, ENilReceiver)
		TEqual(t, result, 0)
	}
	const CAP = 1234
	o := BytesAlloc(CAP)
	TEqual(t, cap(o.ar), CAP)
	TEqual(t, o.Cap(), CAP)
} //                                                        Test_byts_Bytes_Cap_

// go test --run Test_byts_Bytes_FindByte_
func Test_byts_Bytes_FindByte_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) FindByte(b byte) int
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.FindByte('A')
		TCheckError(t, ENilReceiver)
		TEqual(t, result, 0)
	}
	// TODO: TEST return bytes.IndexByte(ob.ar, b)
} //                                                   Test_byts_Bytes_FindByte_

// go test --run Test_byts_Bytes_GetByte_
func Test_byts_Bytes_GetByte_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) GetByte(index int) byte
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.GetByte(0)
		TCheckError(t, ENilReceiver)
		TEqual(t, result, 0)
	}
	// TODO: TEST return ob.ar[index]
} //                                                    Test_byts_Bytes_GetByte_

// go test --run Test_byts_Bytes_GetChar_
func Test_byts_Bytes_GetChar_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) GetChar(index int) CharSize
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.GetChar(0)
		TCheckError(t, ENilReceiver)
		TEqual(t, result, CharSize{})
	}
	// TODO: TEST r, size := utf8.DecodeRune(ob.ar[index:])
	// TODO: TEST return CharSize{Val: r, Size: size}
} //                                                    Test_byts_Bytes_GetChar_

// go test --run Test_byts_Bytes_Size_
func Test_byts_Bytes_Size_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) Size() int
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.Size()
		TCheckError(t, ENilReceiver)
		TEqual(t, result, 0)
	}
	// TODO: TEST return len(ob.ar)
} //                                                       Test_byts_Bytes_Size_

// -----------------------------------------------------------------------------
// # fmt.Stringer Interface

// go test --run Test_byts_Bytes_String_
func Test_byts_Bytes_String_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) String() string
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.String()
		TCheckError(t, ENilReceiver)
		TEqual(t, result, "")
	}
	// TODO: TEST return string(ob.ar)
} //                                                     Test_byts_Bytes_String_

// -----------------------------------------------------------------------------
// # Methods

// go test --run Test_byts_Bytes_Append_
func Test_byts_Bytes_Append_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) Append(b Bytes)
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		np.Append(Bytes{})
		TCheckError(t, ENilReceiver)
	}
	// TODO: TEST ob.ar = append(ob.ar, b.ar...)
} //                                                     Test_byts_Bytes_Append_

// go test --run Test_byts_Bytes_AppendChar_
func Test_byts_Bytes_AppendChar_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) AppendChar(ch rune) int
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		np.AppendChar('A')
		TCheckError(t, ENilReceiver)
	}
	//  TODO: TEST Bytes.AppendChar():
	// 	size := utf8.RuneLen(ch)
	// 	if size == -1 {
	// 		Error("utf8.RuneLen(ch) == -1")
	// 		return -1
	// 	}
	// 	var buf [utf8.UTFMax]byte
	// 	ret := utf8.EncodeRune(buf[:], ch)
	// 	ob.ar = append(ob.ar, buf[:ret]...)
	// 	return ret
} //                                                 Test_byts_Bytes_AppendChar_

// go test --run Test_byts_Bytes_GetBytes_
func Test_byts_Bytes_GetBytes_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) GetBytes() []byte
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.GetBytes()
		TCheckError(t, ENilReceiver)
		TEqual(t, result, []byte{})
	}
	// TODO: TEST return ob.ar
} //                                                   Test_byts_Bytes_GetBytes_

// go test --run Test_byts_Bytes_Insert_
func Test_byts_Bytes_Insert_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) Insert(index int, data Bytes, count int) int
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.Insert(0, Bytes{[]byte{'A'}}, 1)
		TCheckError(t, ENilReceiver)
		TEqual(t, result, 0)
	}
	// TODO: TEST Bytes.Insert():
	// 	if index < 0 || index > len(ob.ar) {
	// 		Error("Index", index, "out of range; array:", len(ob.ar))
	// 		return 0
	// 	}
	// 	src := data.ar
	// 	if count != -1 {
	// 		src = data.ar[:count]
	// 	}
	// 	srcLen := len(src)
	// 	if srcLen == 0 {
	// 		return 0
	// 	}
	// 	ob.ar = append(ob.ar, src...)
	// 	copy(ob.ar[index+srcLen:], ob.ar[index:])
	// 	copy(ob.ar[index:], src)
	// 	return srcLen
} //                                                     Test_byts_Bytes_Insert_

// go test --run Test_byts_Bytes_Remove_
func Test_byts_Bytes_Remove_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) Remove(index, count int)
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		np.Remove(0, 0)
		TCheckError(t, ENilReceiver)
	}
	// TODO: TEST Bytes.Remove():
	// 	if index < 0 || index >= len(ob.ar) {
	// 		Error("Index", index, "out of range; array:", len(ob.ar))
	// 		return
	// 	}
	// 	if count == 0 {
	// 		return
	// 	}
	// 	ob.ar = ob.ar[:index+copy(ob.ar[index:], ob.ar[index+count:])]
	//
} //                                                     Test_byts_Bytes_Remove_

// go test --run Test_byts_Bytes_Reset_
func Test_byts_Bytes_Reset_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) Reset()
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		np.Reset()
		TCheckError(t, ENilReceiver)
	}
	// TODO: TEST ob.ar = ob.ar[:0]
} //                                                      Test_byts_Bytes_Reset_

// TODO: Bytes.Resize() can return old size and error
// go test --run Test_byts_Bytes_Resize_
func Test_byts_Bytes_Resize_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) Resize(size int)
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		np.Resize(1024)
		TCheckError(t, ENilReceiver)
	}
	// TODO: TEST Bytes.Resize():
	// 	current := len(ob.ar)
	// 	if size == current {
	// 		return
	// 	}
	// 	if size < current {
	// 		ob.ar = ob.ar[:size]
	// 	} else if size > current {
	// 		extra := size - current
	// 		zeros := bytes.Repeat([]byte{0}, extra)
	// 		ob.ar = append(ob.ar, zeros...)
	// 	}
} //                                                     Test_byts_Bytes_Resize_

// go test --run Test_byts_Bytes_SetByte_
func Test_byts_Bytes_SetByte_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) SetByte(index int, val byte)
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		np.SetByte(0, 128)
		TCheckError(t, ENilReceiver)
	}
	// TODO: Implement Bytes.SetByte()
	// TODO: Bytes.SetByte() can return previous byte
} //                                                    Test_byts_Bytes_SetByte_

// go test --run Test_byts_Bytes_Slice_
func Test_byts_Bytes_Slice_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Bytes) Slice(beginIndex, endIndex int) Bytes
	//
	{
		// method call on a nil receiver must log an error
		TBeginError()
		var np *Bytes
		result := np.Slice(0, 0)
		TCheckError(t, ENilReceiver)
		TEqual(t, result, Bytes{[]byte{}})
	}
	// TODO: TEST Bytes.Slice():
	// 	if beginIndex < 0 {
	// 		Error("Treating beginIndex", beginIndex, "as 0")
	// 		beginIndex = 0
	// 	}
	// 	if endIndex < 0 || endIndex > len(ob.ar) {
	// 		// -1 is acceptable for endIndex
	// 		if endIndex != -1 {
	// 			Error("Treating endIndex", endIndex, "as length", len(ob.ar))
	// 		}
	// 		endIndex = len(ob.ar)
	// 	}
	// 	return Bytes(Bytes{ar: ob.ar[beginIndex:endIndex]})
} //                                                      Test_byts_Bytes_Slice_

//end
