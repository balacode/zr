// -----------------------------------------------------------------------------
// ZR Library                                                      zr/[bytes.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # Special Constructors
//   BytesAlloc(size int) Bytes
//   BytesAllocUTF8(s string) Bytes
//   BytesWrap(ar []byte) Bytes
//
// # Read-Only Properties / Methods (ob *Bytes)
//   ) Cap() int
//   ) FindByte(b byte) int
//   ) GetByte(index int) byte
//   ) GetChar(index int) CharSize
//   ) Size() int
//
// # fmt.Stringer Interface (ob *Bytes)
//   ) String() string
//
// # Methods (ob *Bytes)
//   ) Append(b Bytes)
//   ) AppendChar(ch rune) int
//   ) GetBytes() []byte
//   ) Insert(index int, data Bytes, count int) int
//   ) Remove(index, count int)
//   ) Reset()
//   ) Resize(size int)
//   ) SetByte(index int, val byte)
//   ) Slice(beginIndex, endIndex int) Bytes

import (
	"bytes"
	"unicode/utf8"
)

// Bytes wraps an array of bytes and provides byte-manipulation methods.
type Bytes struct {
	ar []byte
} //                                                                       Bytes

// CharSize holds a character value and its UTF-8 encoded size in bytes.
type CharSize struct {
	Val  rune
	Size int
} //                                                                    CharSize

// -----------------------------------------------------------------------------
// # Special Constructors

// BytesAlloc creates a new Bytes object with a pre-allocated size.
func BytesAlloc(size int) Bytes {
	return Bytes{ar: make([]byte, 0, size)}
} //                                                                  BytesAlloc

// BytesAllocUTF8 creates a new Bytes object
// by UTF-8 encoding the string 's'.
func BytesAllocUTF8(s string) Bytes {
	return Bytes{ar: []byte(s)}
} //                                                              BytesAllocUTF8

// BytesWrap creates a new Bytes object by wrapping a byte array 'data'.
func BytesWrap(ar []byte) Bytes {
	return Bytes{ar: ar}
} //                                                                   BytesWrap

// -----------------------------------------------------------------------------
// # Read-Only Properties / Methods (ob *Bytes)

// Cap returns the allocated capacity of this object.
func (ob *Bytes) Cap() int {
	if ob == nil {
		mod.Error(ENilReceiver)
		return 0
	}
	return cap(ob.ar)
} //                                                                         Cap

// FindByte returns the index of the first occurrence of byte 'b'.
func (ob *Bytes) FindByte(b byte) int {
	if ob == nil {
		mod.Error(ENilReceiver)
		return 0
	}
	return bytes.IndexByte(ob.ar, b)
} //                                                                    FindByte

// GetByte returns the byte at the specified byte index.
func (ob *Bytes) GetByte(index int) byte {
	if ob == nil {
		mod.Error(ENilReceiver)
		return 0
	}
	return ob.ar[index]
} //                                                                     GetByte

// GetChar returns the Unicode character at the specified byte index.
// Assumes that the data is encoded in UTF-8 format.
// Note that the index is the byte index, not the character number.
func (ob *Bytes) GetChar(index int) CharSize {
	if ob == nil {
		mod.Error(ENilReceiver)
		return CharSize{}
	}
	r, size := utf8.DecodeRune(ob.ar[index:])
	return CharSize{Val: r, Size: size}
} //                                                                     GetChar

// Size returns the written size of this object. (Always less than capacity)
func (ob *Bytes) Size() int {
	if ob == nil {
		mod.Error(ENilReceiver)
		return 0
	}
	return len(ob.ar)
} //                                                                        Size

// -----------------------------------------------------------------------------
// # fmt.Stringer Interface (ob *Bytes)

// String returns a string based on previously-written bytes
// and implements the fmt.Stringer interface.
func (ob *Bytes) String() string {
	if ob == nil {
		mod.Error(ENilReceiver)
		return ""
	}
	return string(ob.ar)
} //                                                                      String

// -----------------------------------------------------------------------------
// # Methods (ob *Bytes)

// Append appends another Bytes object to the end of this object.
func (ob *Bytes) Append(b Bytes) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	ob.ar = append(ob.ar, b.ar...)
} //                                                                      Append

// AppendChar appends a Unicode character encoded in
// UTF-8 format to the end. Returns the number of bytes
// used to encode the character. If the character is not
// valid, returns -1 and doesn't change the object.
func (ob *Bytes) AppendChar(ch rune) int {
	if ob == nil {
		mod.Error(ENilReceiver)
		return 0
	}
	size := utf8.RuneLen(ch)
	if size == -1 {
		mod.Error("utf8.RuneLen(ch) == -1")
		return -1
	}
	var buf [utf8.UTFMax]byte
	ret := utf8.EncodeRune(buf[:], ch)
	ob.ar = append(ob.ar, buf[:ret]...)
	return ret
} //                                                                  AppendChar

// GetBytes returns a raw byte array of bytes in this object.
func (ob *Bytes) GetBytes() []byte {
	if ob == nil {
		mod.Error(ENilReceiver)
		return []byte{}
	}
	return ob.ar
} //                                                                    GetBytes

// Insert inserts a sequence of 'data' bytes
// at position 'index' into object.
// Returns the number of bytes inserted.
func (ob *Bytes) Insert(index int, data Bytes, count int) int {
	if ob == nil {
		mod.Error(ENilReceiver)
		return 0
	}
	if index < 0 || index > len(ob.ar) {
		mod.Error("Index", index, "out of range; array:", len(ob.ar))
		return 0
	}
	src := data.ar
	if count != -1 {
		src = data.ar[:count]
	}
	srcLen := len(src)
	if srcLen == 0 {
		return 0
	}
	ob.ar = append(ob.ar, src...)             // grow dest. by len. of insertion
	copy(ob.ar[index+srcLen:], ob.ar[index:]) // make space for insertion
	copy(ob.ar[index:], src)                  // copy source to destination
	return srcLen
} //                                                                      Insert

// Remove removes 'count' bytes from position starting at 'index'.
func (ob *Bytes) Remove(index, count int) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	if index < 0 || index >= len(ob.ar) {
		mod.Error("Index", index, "out of range; array:", len(ob.ar))
		return
	}
	if count == 0 {
		return
	}
	ob.ar = ob.ar[:index+copy(ob.ar[index:], ob.ar[index+count:])]
} //                                                                      Remove

// Reset empties the contents of this object making Size() equal zero,
// but does not release the allocated memory, so Cap() remains the same.
func (ob *Bytes) Reset() {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	ob.ar = ob.ar[:0]
} //                                                                       Reset

// Resize resizes the Bytes object so that Size() equals 'size'.
// If the new size is smaller than the old size, the contents is truncated.
// If it is bigger, appends zeros to pad the contents.
func (ob *Bytes) Resize(size int) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	current := len(ob.ar)
	if size == current {
		return
	}
	if size < current {
		ob.ar = ob.ar[:size]
	} else if size > current {
		extra := size - current
		zeros := bytes.Repeat([]byte{0}, extra)
		ob.ar = append(ob.ar, zeros...)
	}
} //                                                                      Resize

// SetByte _ _
func (ob *Bytes) SetByte(index int, val byte) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	IMPLEMENT()
} //                                                                     SetByte

// Slice returns a Bytes object that is a slice of this
// object. The slice references the original array
// so no new memory allocation is made.
//
// beginIndex is the starting position.
// Must not be less than zero
// in which the function treats it as zero and logs an error.
//
// endIndex is the ending position, just after the last required element.
// If endIndex is -1, returns everything from beginIndex up to the end.
func (ob *Bytes) Slice(beginIndex, endIndex int) Bytes {
	if ob == nil {
		mod.Error(ENilReceiver)
		return Bytes{}
	}
	if beginIndex < 0 {
		mod.Error("Treating beginIndex", beginIndex, "as 0")
		beginIndex = 0
	}
	if endIndex < 0 || endIndex > len(ob.ar) {
		// -1 is acceptable for endIndex
		if endIndex != -1 {
			mod.Error("Treating endIndex", endIndex, "as length", len(ob.ar))
		}
		endIndex = len(ob.ar)
	}
	return Bytes(Bytes{ar: ob.ar[beginIndex:endIndex]})
} //                                                                       Slice

//end
