// -----------------------------------------------------------------------------
// ZR Library                                                 zr/[bytes_func.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

//   AppendRuneBytes(dest *[]byte, r rune) int
//   ByteSizeString(sizeInBytes int64, useSI bool) string
//   ClearBytes(slice *[]byte)
//   CompressBytes(data []byte) []byte
//   FoldXorBytes(ar []byte, returnLen int) []byte
//   HexStringOfBytes(ar []byte) string
//   InsertBytes(dest *[]byte, pos int, src ...[]byte)
//   RandomBytes(length int) []byte
//   RemoveBytes(dest *[]byte, pos, count int)
//   RuneOffset(slice []byte, runeIndex int) (byteIndex int)
//   UncompressBytes(data []byte) []byte
//   XorBytes(data, cipher []byte) []byte

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"math"
	"math/big"
	"strings"
	"unicode/utf8"
)

// AppendRuneBytes appends a rune to the specified buffer, encoded as UTF-8.
// Note that 'dest' is a pointer to a byte slice, to allow the slice to be
// updated. Returns the number of bytes used to encode the rune. If the rune
// is not valid, returns -1 and doesn't change 'dest'.
func AppendRuneBytes(dest *[]byte, r rune) int {
	if size := utf8.RuneLen(r); size == -1 {
		mod.Error("utf8.RuneLen(r) == -1")
		return -1
	}
	var buf [utf8.UTFMax]byte
	ret := utf8.EncodeRune(buf[:], r)
	(*dest) = append((*dest), buf[:ret]...)
	return ret
} //                                                             AppendRuneBytes

// ByteSizeString returns a human-friendly byte count string.
// E.g. 1024 gives '1 KiB', 1048576 gives '1 MiB'.
// If you set 'useSI' to true, uses multiples of 1000 (SI units)
// instead of 1024 (binary units) and suffixes 'KB' instead of 'KiB', etc.
func ByteSizeString(sizeInBytes int64, useSI bool) string {
	n := int64(1024)
	if useSI {
		n = 1000
	}
	if sizeInBytes > -n && sizeInBytes < n {
		return fmt.Sprintf("%d B", sizeInBytes)
	}
	if !useSI && sizeInBytes == math.MinInt64 {
		return "-7.9 EiB"
	}
	neg := sizeInBytes < 0
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
			sz := big.NewInt(sizeInBytes)
			sz.Mul(sz, big.NewInt(10))
			sz.Div(sz, big.NewInt(group.scale))
			cut = float64(sz.Int64()) / 10
		}
		ret += fmt.Sprintf("%0.1f", cut)
		//
		// remove trailing zero decimal
		if strings.HasSuffix(ret, ".0") {
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

// ClearBytes removes all bytes from a byte slice,
// retaining its underlying array and its allocated capacity.
func ClearBytes(slice *[]byte) {
	(*slice) = (*slice)[0:0]
} //                                                                  ClearBytes

// CompressBytes compresses an array of
// bytes returns the ZLIB-compressed bytes.
func CompressBytes(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}
	// zip data in standard manner
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write(data)
	w.Close()
	//
	// log any problem
	const ERRM = "Failed compressing data with zlib:"
	if err != nil {
		mod.Error(ERRM, err)
		return []byte{}
	}
	ret := b.Bytes()
	if len(ret) < 3 {
		mod.Error(ERRM, "length < 3")
		return []byte{}
	}
	return ret
} //                                                               CompressBytes

// FoldXorBytes creates a shorter byte array from a longer byte array
// by overwriting the shorter array's elements using a XOR operation.
// This function is mainly used to shorten long hashes to get a shorter ID
// (only where the full hash precision is not necessary).
func FoldXorBytes(ar []byte, returnLen int) []byte {
	if returnLen < 1 {
		mod.Error(EInvalidArg, "^returnLen", returnLen)
		return []byte{}
	}
	ret := make([]byte, returnLen)
	i := 0
	for _, bt := range ar {
		if i >= returnLen {
			i = 0
		}
		ret[i] ^= bt
		i++
	}
	return ret
} //                                                                FoldXorBytes

// HexStringOfBytes converts a byte array to a string of hexadecimal digits.
func HexStringOfBytes(ar []byte) string {
	return fmt.Sprintf("%0X", ar)
} //                                                            HexStringOfBytes

// InsertBytes inserts a copy of a byte slice into another byte slice.
func InsertBytes(dest *[]byte, pos int, src ...[]byte) {
	for _, part := range src {
		destLen, partLen := len(*dest), len(part)
		if pos < 0 || pos > destLen {
			mod.Error("Position", pos, "out of range; len(*dest):", destLen)
		} else if partLen != 0 {
			*dest = append(*dest, part...)             // grow destination
			copy((*dest)[pos+partLen:], (*dest)[pos:]) // shift to make space
			copy((*dest)[pos:], part)                  // copy source to dest.
		}
		pos += len(part)
	}
} //                                                                 InsertBytes

// RandomBytes generates and returns a random slice of bytes
func RandomBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	ret := make([]byte, length)
	_, err := mod.rand.Read(ret)
	if err != nil {
		return []byte{}
	}
	return ret
} //                                                                 RandomBytes

// RemoveBytes removes the specified number of bytes from a byte slice.
func RemoveBytes(dest *[]byte, pos, count int) {
	if pos < 0 || pos >= len(*dest) {
		mod.Error("Position", pos, "out of range; len(*dest):", len(*dest))
		return
	}
	if count == 0 {
		return
	}
	(*dest) = (*dest)[:pos+copy((*dest)[pos:], (*dest)[pos+count:])]
} //                                                                 RemoveBytes

// RuneOffset reads a byte slice and returns the byte position of
// the rune at runeIndex, or -1 if the index is out of range.
func RuneOffset(slice []byte, runeIndex int) (byteIndex int) {
	for runeIndex > 0 {
		runeIndex--
		_, size := utf8.DecodeRune(slice) // get size of the rune in bytes
		if size == 0 {
			break
		}
		byteIndex += size
		slice = slice[size:]
		if len(slice) == 0 {
			break
		}
	}
	if runeIndex != 0 {
		byteIndex = -1
	}
	return byteIndex
} //                                                                  RuneOffset

// UncompressBytes uncompresses a ZLIB-compressed array of bytes.
func UncompressBytes(data []byte) []byte {
	readCloser, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		mod.Error("UncompressBytes:", err)
		return []byte{}
	}
	ret := bytes.NewBuffer(make([]byte, 0, 8192))
	io.Copy(ret, readCloser)
	readCloser.Close()
	return ret.Bytes()
} //                                                             UncompressBytes

// XorBytes _ _
func XorBytes(data, cipher []byte) []byte {
	if len(data) == 0 || len(cipher) == 0 {
		return data
	}
	var (
		ret  = bytes.NewBuffer(make([]byte, 0, len(data)))
		i, l = 0, len(cipher)
	)
	for _, b := range data {
		ret.WriteByte(b ^ cipher[i])
		i++
		if i == l {
			i = 0
		}
	}
	return ret.Bytes()
} //                                                                    XorBytes

// end
