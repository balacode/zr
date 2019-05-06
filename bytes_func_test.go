// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-06 06:25:43 D5DD22                        zr/[bytes_func_test.go]
// -----------------------------------------------------------------------------

package zr

//   Test_bytf_AppendRuneBytes_
//   Test_bytf_ByteSizeString_
//   Test_bytf_ClearBytes_
//   Test_bytf_CompressBytes_
//   Test_bytf_FoldXorBytes_
//   Test_bytf_HexStringOfBytes_
//   Test_bytf_InsertBytes_
//   Test_bytf_RemoveBytes_
//   Test_bytf_RuneOffset_
//   Test_bytf_UncompressBytes_
//   Test_bytf_XorBytes_

/*
to test all items in bytes_func.go use:
    go test --run Test_bytf_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"bytes"
	"math"
	"testing"
)

// -----------------------------------------------------------------------------
// # Tests

// go test --run Test_bytf_AppendRuneBytes_
func Test_bytf_AppendRuneBytes_(t *testing.T) {
	TBegin(t)
	//   AppendRuneBytes(dest *[]byte, r rune) int
	//
	r := 'Б'
	buf := make([]byte, 0, 10)
	AppendRuneBytes(&buf, r)
	expect := []byte{208, 145}
	if bytes.Compare(buf, expect) != 0 {
		t.Errorf("FAILED: buf became %v instead of %v", buf, expect)
	}
} //                                                  Test_bytf_AppendRuneBytes_

// go test --run Test_bytf_ByteSizeString_
func Test_bytf_ByteSizeString_(t *testing.T) {
	TBegin(t)
	// ByteSizeString(sizeInBytes int64, useSI bool) string
	//
	test := func(sizeInBytes int64, useSI bool, expect string) {
		got := ByteSizeString(sizeInBytes, useSI)
		if got != expect {
			TFail(t,
				`ByteSizeString(`, sizeInBytes, `, `, useSI, `)`,
				`returned "`, got, `". must be "`, expect, `".`,
			)
		}
	}
	const KiB = 1024
	const MiB = KiB * 1024
	const GiB = MiB * 1024
	const TiB = GiB * 1024
	const PiB = TiB * 1024
	const EiB = PiB * 1024
	//
	const KB = 1000
	const MB = KB * 1000
	const GB = MB * 1000
	const TB = GB * 1000
	const PB = TB * 1000
	const EB = PB * 1000
	//
	// binary prefixes (use powers of 2)
	{
		//
		// zero
		test(0, false, "0 B")
		//
		// limits
		test(1, false, "1 B")
		test(math.MaxInt64, false, "7.9 EiB")
		test(math.MaxInt64-1, false, "7.9 EiB")
		//
		test(math.MinInt64+1, false, "-7.9 EiB")
		test(math.MinInt64, false, "-7.9 EiB")
		//
		// one exbibyte less one byte
		// (this demonstrates that 1 decimal cut-off works)
		//
		test(EiB-1, false, "1023.9 PiB")
		test(-EiB+1, false, "-1023.9 PiB")
		//
		// one pebibyte less one byte
		test(PiB-1, false, "1023.9 TiB")
		test(-PiB+1, false, "-1023.9 TiB")
		//
		test(-1, false, "-1 B")
		test(math.MinInt64, false, "-7.9 EiB")
		//
		// smaller than 1 Kib
		test(1023, false, "1023 B")
		test(-1023, false, "-1023 B")
		//
		// positive
		test(KiB, false, "1 KiB")
		test(MiB, false, "1 MiB")
		test(GiB, false, "1 GiB")
		test(TiB, false, "1 TiB")
		test(PiB, false, "1 PiB")
		test(EiB, false, "1 EiB")
		//
		test(-KiB, false, "-1 KiB")
		test(-MiB, false, "-1 MiB")
		test(-GiB, false, "-1 GiB")
		test(-TiB, false, "-1 TiB")
		test(-PiB, false, "-1 PiB")
		test(-EiB, false, "-1 EiB")
		//
		// positive fraction
		test(int64(float64(KiB)*1.5), false, "1.5 KiB")
		test(int64(float64(MiB)*1.5), false, "1.5 MiB")
		test(int64(float64(GiB)*1.5), false, "1.5 GiB")
		test(int64(float64(TiB)*1.5), false, "1.5 TiB")
		test(int64(float64(PiB)*1.5), false, "1.5 PiB")
		test(int64(float64(EiB)*1.5), false, "1.5 EiB")
		//
		test(-int64(float64(KiB)*1.5), false, "-1.5 KiB")
		test(-int64(float64(MiB)*1.5), false, "-1.5 MiB")
		test(-int64(float64(GiB)*1.5), false, "-1.5 GiB")
		test(-int64(float64(TiB)*1.5), false, "-1.5 TiB")
		test(-int64(float64(PiB)*1.5), false, "-1.5 PiB")
		test(-int64(float64(EiB)*1.5), false, "-1.5 EiB")
		//
		// edge cases
		test(MiB-1, false, "1023.9 KiB")
	}
	// TODO: Test_bytf_ByteSizeString_(): test with SI units
} //                                                   Test_bytf_ByteSizeString_

// go test --run Test_bytf_ClearBytes_
func Test_bytf_ClearBytes_(t *testing.T) {
	TBegin(t)
	//   ClearBytes(slice *[]byte)
	//
	slice := make([]byte, 0, 0)
	slice = append(slice, []byte{1, 2, 3}...)
	initialCap := cap(slice)
	ClearBytes(&slice)
	if len(slice) != 0 {
		t.Errorf("len(slice) != 0 (%v != 0)", len(slice))
	}
	if cap(slice) != initialCap {
		t.Errorf("cap(slice) != initialCap. (%v != %v)",
			cap(slice), initialCap)
	}
} //                                                       Test_bytf_ClearBytes_

// go test --run Test_bytf_CompressBytes_
func Test_bytf_CompressBytes_(t *testing.T) {
	TBegin(t)
	// CompressBytes(data []byte) []byte
	//
	// TODO: TEST:
	// 	if len(data) == 0 {
	// 		return []byte{}
	// 	}
	// 	var retBuf bytes.Buffer
	// 	w := zlib.NewWriter(&retBuf)
	// 	defer w.Close()
	// 	var_, err = w.Write([]byte(data))
	// 	if err != nil {
	// 		Error("Compressing:", err)
	// 		return []byte{}
	// 	}
	// 	return retBuf.Bytes()
} //                                                    Test_bytf_CompressBytes_

// go test --run Test_bytf_FoldXorBytes_
func Test_bytf_FoldXorBytes_(t *testing.T) {
	TBegin(t)
	//   FoldXorBytes(ar []byte, returnLen int) []byte
	//
	for i, test := range []struct {
		ar        []byte
		returnLen int
		expect    []byte
	}{
		{[]byte{0xAA, 0xBB, 0xCC}, 1, []byte{0xDD}}, // UNTESTED
	} {
		got := FoldXorBytes(test.ar, test.returnLen)
		if !TArrayEqual(t, got, test.expect) {
			TFail(t, "#", i,
				` ar `, GoString(test.ar),
				` returnLen `, GoString(test.returnLen),
				` expected `, GoString(test.expect),
				` returned `, GoString(got),
			)
		}
	}
} //                                                     Test_bytf_FoldXorBytes_

// go test --run Test_bytf_HexStringOfBytes_
func Test_bytf_HexStringOfBytes_(t *testing.T) {
	TBegin(t)
	//   HexStringOfBytes(ar []byte) string
	//
	test := func(ar []byte, expect string) {
		got := HexStringOfBytes(ar)
		if got != expect {
			TFail(t,
				`HexStringOfBytes(`, GoString(ar), `) `,
				`returned "`, got, `". must be "`, expect, `"`,
			)
		}
	}
	test([]byte{}, "")
	test([]byte{0}, "00")
	test([]byte{1}, "01")
	test([]byte{1, 2}, "0102")
	test([]byte{1, 2, 3}, "010203")
	test([]byte{0xAA}, "AA")
	test([]byte{0xBB, 0xCC}, "BBCC")
	test([]byte{0xDD, 0xEE, 0xFF}, "DDEEFF")
} //                                                 Test_bytf_HexStringOfBytes_

// go test --run Test_bytf_InsertBytes_
func Test_bytf_InsertBytes_(t *testing.T) {
	TBegin(t)
	//   InsertBytes(dest *[]byte, pos int, src ...[]byte)
	//
	{
		dest := []byte("ABCD")
		pos := 2
		src1 := []byte("12")
		src2 := []byte("3")
		got := []byte(string(dest))
		expect := []byte("AB123CD")
		InsertBytes(&got, pos, src1, src2)
		if len(got) != 7 {
			t.Errorf("len(got) is %d instead of 7", len(got))
		}
		if !bytes.Equal(got, expect) {
			t.Errorf("InsertBytes(%q, %d, %q, %q)"+
				" returned %q instead of %q"+LB,
				got, pos, string(src1), string(src2),
				string(got), string(expect))
		}
	}
	{
		slc := make([]byte, 0, 0)
		slc = append(slc, []byte{1, 2, 3}...)
		InsertBytes(&slc, 1, []byte{10, 20, 30})
		if !bytes.Equal(slc, []byte{1, 10, 20, 30, 2, 3}) {
			t.Error("InsertBytes() test failed")
		}
		slc = []byte{1, 2, 3, 4, 5}
		InsertBytes(&slc, 0, []byte{6})
		TBytesEqual(t, slc, []byte{6, 1, 2, 3, 4, 5})
	}
	des1 := []byte("ABCDEFGH")
	src1 := []byte("123")
	des2 := []byte("ABC")
	src2 := []byte("12345")
	// TODO: use test() function instead of loop
	for _, test := range []struct {
		des    []byte
		pos    int
		src    []byte
		expect []byte
	}{
		{des1, 0, src1, []byte("123ABCDEFGH")},
		{des1, 1, src1, []byte("A123BCDEFGH")},
		{des1, 2, src1, []byte("AB123CDEFGH")},
		{des1, 3, src1, []byte("ABC123DEFGH")},
		{des1, 4, src1, []byte("ABCD123EFGH")},
		{des1, 5, src1, []byte("ABCDE123FGH")},
		{des1, 6, src1, []byte("ABCDEF123GH")},
		{des1, 7, src1, []byte("ABCDEFG123H")},
		{des1, 8, src1, []byte("ABCDEFGH123")},
		{des2, 0, src2, []byte("12345ABC")},
		{des2, 1, src2, []byte("A12345BC")},
		{des2, 2, src2, []byte("AB12345C")},
		{des2, 3, src2, []byte("ABC12345")},
	} {
		got := []byte(string(test.des))
		InsertBytes(&got, test.pos, test.src)
		if string(got) != string(test.expect) {
			t.Errorf("InsertBytes(%q, %d, %q)"+
				" returned %q instead of %q"+LB,
				string(test.des), test.pos, string(test.src),
				string(got), string(test.expect))
		}
	}
} //                                                      Test_bytf_InsertBytes_

// go test --run Test_bytf_RemoveBytes_
func Test_bytf_RemoveBytes_(t *testing.T) {
	TBegin(t)
	//   RemoveBytes(dest *[]byte, pos, count int)
	//
	if true {
		slice := []byte{1, 2, 3}
		RemoveBytes(&slice, 0, 3)
		TBytesEqual(t, slice, []byte{})
	}
	if true {
		slice := []byte{1, 2, 3}
		RemoveBytes(&slice, 0, 1)
		TBytesEqual(t, slice, []byte{2, 3})
	}
	if true {
		slice := []byte{1, 2, 3}
		RemoveBytes(&slice, 1, 1)
		TBytesEqual(t, slice, []byte{1, 3})
	}
	if true {
		slice := []byte{1, 2, 3}
		RemoveBytes(&slice, 2, 1)
		TBytesEqual(t, slice, []byte{1, 2})
	}
} //                                                      Test_bytf_RemoveBytes_

// go test --run Test_bytf_RuneOffset_
func Test_bytf_RuneOffset_(t *testing.T) {
	TBegin(t)
	//   RuneOffset(slice []byte, runeIndex int) (byteIndex int)
	//
	slice := []byte("Hello, 世界")
	TTrue(t, RuneOffset(slice, 0) == 0)
	TTrue(t, RuneOffset(slice, 1) == 1)
	TTrue(t, RuneOffset(slice, 2) == 2)
	TTrue(t, RuneOffset(slice, 3) == 3)
	TTrue(t, RuneOffset(slice, 4) == 4)
	TTrue(t, RuneOffset(slice, 5) == 5)
	TTrue(t, RuneOffset(slice, 6) == 6)
	TTrue(t, RuneOffset(slice, 7) == 7)
	TTrue(t, RuneOffset(slice, 8) == 10)
	TTrue(t, RuneOffset(slice, 9) == 13)
	//
	// try some out-of-range values:
	TTrue(t, RuneOffset(slice, -2) == -1)
	TTrue(t, RuneOffset(slice, -1) == -1)
	TTrue(t, RuneOffset(slice, 10) == -1)
} //                                                       Test_bytf_RuneOffset_

// go test --run Test_bytf_UncompressBytes_
func Test_bytf_UncompressBytes_(t *testing.T) {
	TBegin(t)
	// UncompressBytes(data []byte) []byte
	//
	// TODO: TEST UncompressBytes():
	// 	retBuf := bytes.NewReader(data)
	// 	rd, err := zlib.NewReader(retBuf)
	// 	if err != nil {
	// 		Error("Uncompressing:", err)
	// 	}
	// 	ret := bytes.NewBuffer(make([]byte, 0, 8192))
	// 	io.Copy(ret, rd)
	// 	rd.Close()
	// 	return ret.Bytes()
} //                                                  Test_bytf_UncompressBytes_

// go test --run Test_bytf_XorBytes_
func Test_bytf_XorBytes_(t *testing.T) {
	TBegin(t)
	// XorBytes(data, cipher []byte) []byte
	//
	// TODO: TEST XorBytes():
	// 	ret := bytes.NewBuffer(make([]byte, 0, len(data)))
	// 	i, l := 0, len(cipher)
	// 	for_, b := range data {
	// 		ret.WriteByte(b ^ cipher[i])
	// 		i++
	// 		if i == l {
	// 			i = 0
	// 		}
	// 	}
	// 	return ret.Bytes()
} //                                                         Test_bytf_XorBytes_

//end
