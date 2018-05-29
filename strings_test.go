// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-29 09:25:37 7D230E                           [zr/strings_test.go]
// -----------------------------------------------------------------------------

package zr

// # String Function Tests
//   Test_strs_After_
//   Test_strs_CamelCase_
//   Test_strs_CharsOf_
//   Test_strs_CompactSpaces_
//   Test_strs_ContainsI_
//   Test_strs_ContainsWord_
//   Test_strs_CountCRLF_
//   Test_strs_EqualStringSlices_
//   Test_strs_FindChar_
//   Test_strs_FindInSlice_
//   Test_strs_First_
//   Test_strs_GetPart_
//   Test_strs_IfString_
//   Test_strs_IsIdentifier_
//   Test_strs_IsWhiteSpace_
//   Test_strs_JSUnescapeStruct_
//   Test_strs_JSUnescape_
//   Test_strs_Last_
//   Test_strs_LineBeginIndexB_
//   Test_strs_LineBeginIndex_
//   Test_strs_LineEndIndexB_
//   Test_strs_LineEndIndex_
//   Test_strs_LineOfIndex_
//   Test_strs_LineOffsetUTF8_
//   Test_strs_Padf_
//   Test_strs_ReplaceEx1_
//   Test_strs_ReplaceI_
//   Test_strs_ReplaceMany_
//   Test_strs_ReplaceWord_
//   Test_strs_SetPart_
//   Test_strs_SetSlice_
//   Test_strs_ShowSpaces_
//   Test_strs_SkipChars_
//   Test_strs_SkipName_
//   Test_strs_SkipSpaces_
//   Test_strs_Slice_
//   Test_strs_SplitQuoted_
//   Test_strs_StrOneOf_
//   Test_strs_String_
//   Test_strs_Substr_
//   Test_strs_TitleCase_
//   Test_strs_TokenGetEx_
//   Test_strs_TokenGet_
//   Test_strs_WordIndex_

/*
to test all items in strings.go use:
	go test --run Test_strs_

to generate a test coverage report use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

import (
	"fmt"
	"testing"
)

// go test --run Test_strs_consts_
func Test_strs_consts_(t *testing.T) {
	TBegin(t)
	//
	TEqual(t, SPACES, " \a\b\f\n\r\t\v")
	//
	TEqual(t, IgnoreCase, CaseMode(1))
	TEqual(t, MatchCase, CaseMode(2))
	//
	TEqual(t, IgnoreWord, WordMode(1))
	TEqual(t, MatchWord, WordMode(2))
} //                                                           Test_strs_consts_

// go test --run Test_strs_After_
func Test_strs_After_(t *testing.T) {
	TBegin(t)
	// After(s string, find ...string) string
	//
	var test = func(s string, find []string, expect string) {
		var got = After(s, find...)
		TEqual(t, got, (expect))
	}
	{
		test("ABC", []string{""}, "ABC")
		test("ABC", []string{"A"}, "BC")
		test("ABC", []string{"A", "B"}, "C")
		test("ABC", []string{"A", "B", "C"}, "")
		test("ABC", []string{"D"}, "")
		test("ABC", []string{"A", "B", "C", "D"}, "")
	}
	{
		test("", []string{""}, "")
		test("", []string{"1"}, "")
		test("", []string{"12"}, "")
	}
	{
		test("122333", []string{""}, "122333")
		test("122333", []string{"1"}, "22333")
		test("122333", []string{"1", "22"}, "333")
		test("122333", []string{"1", "22", "333"}, "")
		test("122333", []string{"D"}, "")
		test("122333", []string{"1", "22", "333", "D"}, "")
	}
	{
		test("333221", []string{""}, "333221")
		test("333221", []string{"333"}, "221")
		test("333221", []string{"333", "22"}, "1")
		test("333221", []string{"333", "22", "1"}, "")
		test("333221", []string{"D"}, "")
		test("333221", []string{"333", "22", "1", "D"}, "")
	}
} //                                                            Test_strs_After_

// go test --run Test_strs_CamelCase_
func Test_strs_CamelCase_(t *testing.T) {
	TBegin(t)
	// CamelCase(s string) string
	//
	var test = func(s, expect string) {
		var got = CamelCase(s)
		TEqual(t, got, (expect))
	}
	test("", "")
	test("CAPS", "CAPS")
	test("some_name", "someName")
} //                                                        Test_strs_CamelCase_

// go test --run Test_strs_CharsOf_
func Test_strs_CharsOf_(t *testing.T) {
	TBegin(t)
	// CharsOf(s string) []rune
	//
	TArrayEqual(t, []rune{}, CharsOf(""))
	TArrayEqual(t, []rune{'A', 'B', 'C'}, CharsOf("ABC"))
} //                                                          Test_strs_CharsOf_

// go test --run Test_strs_CompactSpaces_
func Test_strs_CompactSpaces_(t *testing.T) {
	TBegin(t)
	// CompactSpaces(s string) string
	//
	TEqual(t, CompactSpaces("a  b  c"), ("a b c"))
	TEqual(t, CompactSpaces("a\r\nb\r\nc"), ("a b c"))
} //                                                    Test_strs_CompactSpaces_

// go test --run Test_strs_ContainsI_
func Test_strs_ContainsI_(t *testing.T) {
	TBegin(t)
	// ContainsI(s, substr string) bool
	//
	TTrue(t, ContainsI("abc", ""))
	TTrue(t, ContainsI("abc", "A"))
	TTrue(t, ContainsI("abc", "AB"))
	TTrue(t, ContainsI("abc", "ABC"))
	TTrue(t, ContainsI("abc", "B"))
	TTrue(t, ContainsI("abc", "BC"))
	TTrue(t, ContainsI("abc", "C"))
	TFalse(t, ContainsI("abc", " abc"))
	TFalse(t, ContainsI("abc", " b"))
	TFalse(t, ContainsI("abc", "abc "))
	TFalse(t, ContainsI("abc", "ac"))
	TFalse(t, ContainsI("abc", "b "))
	TFalse(t, ContainsI("abc", "ca"))
	TFalse(t, ContainsI("abc", "cb"))
	TFalse(t, ContainsI("abc", "d"))
} //                                                        Test_strs_ContainsI_

// go test --run Test_strs_CountCRLF_
func Test_strs_CountCRLF_(t *testing.T) {
	TBegin(t)
	// CountCRLF(s string) (count, countCR, countLF int)
	//
	var n, cr, lf int
	//
	n, cr, lf = CountCRLF("")
	TEqual(t, n, 0)
	TEqual(t, cr, 0)
	TEqual(t, lf, 0)
	//
	n, cr, lf = CountCRLF(" ")
	TEqual(t, n, 0)
	TEqual(t, cr, 0)
	TEqual(t, lf, 0)
	//
	n, cr, lf = CountCRLF("\r\n\r\n")
	TEqual(t, n, 4)
	TEqual(t, cr, 2)
	TEqual(t, lf, 2)
	//TODO: CountCRLF() may need to be changed to count actual number
	//      of CR+LF pairs, not just add count of CRs and LFs.
} //                                                        Test_strs_CountCRLF_

// go test --run Test_strs_ContainsWord_
func Test_strs_ContainsWord_(t *testing.T) {
	TBegin(t)
	// ContainsWord(s, word string, caseMode CaseMode) bool
	//
	var s = "abc defg hijk"
	{
		TTrue(t, ContainsWord(s, "abc", MatchCase))
		TTrue(t, ContainsWord(s, "defg", MatchCase))
		TTrue(t, ContainsWord(s, "hijk", MatchCase))
	}
	{
		TTrue(t, ContainsWord(s, "ABC", IgnoreCase))
		TTrue(t, ContainsWord(s, "DEFG", IgnoreCase))
		TTrue(t, ContainsWord(s, "HIJK", IgnoreCase))
	}
	{
		TFalse(t, ContainsWord(s, "ABC", MatchCase))
		TFalse(t, ContainsWord(s, "DEFG", MatchCase))
		TFalse(t, ContainsWord(s, "HIJK", MatchCase))
	}
	{
		TFalse(t, ContainsWord(s, "a", MatchCase))
		TFalse(t, ContainsWord(s, "ab", MatchCase))
		TFalse(t, ContainsWord(s, "bc", MatchCase))
	}
	{
		TFalse(t, ContainsWord(s, "def", MatchCase))
		TFalse(t, ContainsWord(s, "ef", MatchCase))
		TFalse(t, ContainsWord(s, "fg", MatchCase))
	}
	{
		TFalse(t, ContainsWord(s, "h", MatchCase))
		TFalse(t, ContainsWord(s, "hi", MatchCase))
		TFalse(t, ContainsWord(s, "hij", MatchCase))
		TFalse(t, ContainsWord(s, "ijk", MatchCase))
	}
} //                                                     Test_strs_ContainsWord_

// go test --run Test_strs_EqualStringSlices_
func Test_strs_EqualStringSlices_(t *testing.T) {
	TBegin(t)
	// EqualStringSlices(a, b []string) bool
	//
	var a = []string{"first", "second", "third"}
	var b = []string{"first", "second", "third"}
	//
	// return true when comparing an object to itself
	TTrue(t, EqualStringSlices(a, a))
	TTrue(t, EqualStringSlices(b, b))
	//
	// true when arrays are different but contain the same values
	TTrue(t, EqualStringSlices(a, b))
	//
	// return 'false' when number of items not the same
	TFalse(t, EqualStringSlices(a,
		[]string{"first", "second", "third", ""}))
	//
	// return 'false' when ordering of items not the same
	TFalse(t, EqualStringSlices(a,
		[]string{"first", "third", "second"}))
} //                                                Test_strs_EqualStringSlices_

// go test --run Test_strs_FindChar_
func Test_strs_FindChar_(t *testing.T) {
	TBegin(t)
	// FindChar(s string, ch byte, beginIndex int) int
	//
	TEqual(t, FindChar("ABC ABC", 'C', 0), (2))
	TEqual(t, FindChar("ABC ABC", 'C', 2), (2))
	TEqual(t, FindChar("ABC ABC", 'C', 3), (6))
	TEqual(t, FindChar("ABC ABC", 'D', 0), (-1))
} //                                                         Test_strs_FindChar_

// go test --run Test_strs_FindInSlice_
func Test_strs_FindInSlice_(t *testing.T) {
	TBegin(t)
	// FindInSlice(s string, start, end int, substr string) int
	//
	var test = func(s string, start, end int, substr string,
		expect int) {
		var got = FindInSlice(s, start, end, substr)
		TEqual(t, got, (expect))
	}
	test("AA BA CA", 3, 5, "A", 4)
} //                                                      Test_strs_FindInSlice_

// go test --run Test_strs_First_
func Test_strs_First_(t *testing.T) {
	TBegin(t)
	// First(s string, count int) string
	//
	{
		DisableErrors()
		TEqual(t, First("abc", -1), (""))
		EnableErrors()
	}
	TEqual(t, First("abc", 0), (""))
	TEqual(t, First("abc", 1), ("a"))
	TEqual(t, First("abc", 2), ("ab"))
	TEqual(t, First("abc", 3), ("abc"))
	{
		DisableErrors()
		TEqual(t, First("abc", 4), ("abc"))
		EnableErrors()
	}
} //                                                            Test_strs_First_

// go test --run Test_strs_GetPart_
func Test_strs_GetPart_(t *testing.T) {
	TBegin(t)
	// GetPart(s, prefix, suffix string) string
	//
	// both prefix and suffix are blank: return 's' as it is
	TEqual(t, GetPart("name:cat;", "", ""), "name:cat;")
	//
	// prefix is blank, suffix specified: return everything before the suffix
	TEqual(t, GetPart("name:cat;", "", ";"), "name:cat")
	//
	// prefix specified, suffix is blank: return everything after the prefix
	TEqual(t, GetPart("name:cat;", "name:", ""), "cat;")
	//
	// both prefix and suffix specified: return the substring between them
	TEqual(t, GetPart("name:cat;", "name:", ";"), "cat")
	//
	// non-existent prefix: return a blank string
	TEqual(t, GetPart("name:cat;", "age:", ""), "")
	//
	// non-existent suffix: return a blank string
	TEqual(t, GetPart("name:cat;", "name:", "."), "")
	//
	// additional test
	TEqual(t, GetPart("xyz class::sum; 123", "class::", ";"), "sum")
} //                                                          Test_strs_GetPart_

// go test --run Test_strs_IfString_
func Test_strs_IfString_(t *testing.T) {
	TBegin(t)
	// IfString(condition bool, trueStr, falseStr string) string
	//
	TEqual(t, IfString(true, "true", "false"), "true")
	TEqual(t, IfString(false, "true", "false"), "false")
} //                                                         Test_strs_IfString_

// go test --run Test_strs_IsIdentifier_
func Test_strs_IsIdentifier_(t *testing.T) {
	TBegin(t)
	// IsIdentifier(s string) bool
	//
	TEqual(t, IsIdentifier(""), false)
	TEqual(t, IsIdentifier(" "), false)
	TEqual(t, IsIdentifier("abc 123"), false)
	//
	TEqual(t, IsIdentifier("123"), true)
	TEqual(t, IsIdentifier("123abc"), true)
	TEqual(t, IsIdentifier("abc123"), true)
} //                                                     Test_strs_IsIdentifier_

// go test --run Test_strs_IsWhiteSpace_
func Test_strs_IsWhiteSpace_(t *testing.T) {
	TBegin(t)
	// IsWhiteSpace(s string) bool
	//
	TEqual(t, IsWhiteSpace(""), false)
	TEqual(t, IsWhiteSpace("ABC"), false)
	TEqual(t, IsWhiteSpace(" -"), false)
	//
	TEqual(t, IsWhiteSpace(" "), true)
	TEqual(t, IsWhiteSpace("\a"), true)
	TEqual(t, IsWhiteSpace("\b"), true)
	TEqual(t, IsWhiteSpace("\f"), true)
	TEqual(t, IsWhiteSpace("\n"), true)
	TEqual(t, IsWhiteSpace("\r"), true)
	TEqual(t, IsWhiteSpace("\t"), true)
	TEqual(t, IsWhiteSpace("\v"), true)
} //                                                     Test_strs_IsWhiteSpace_

// go test --run Test_strs_JSUnescape_
func Test_strs_JSUnescape_(t *testing.T) {
	TBegin(t)
	// JSUnescape(s string) string
	//
	TEqual(t, JSUnescape("%25"), "%")
	TEqual(t, JSUnescape("%20"), " ")
	//
	//TODO: more unit test cases for JSUnescape()
} //                                                       Test_strs_JSUnescape_

// go test --run Test_strs_JSUnescapeStruct_
func Test_strs_JSUnescapeStruct_(t *testing.T) {
	TBegin(t)
	// JSUnescapeStruct(structPtr interface{})
	//
	type Item struct {
		Sub1 string
		Sub2 int
	}
	type Test struct {
		Main1 int
		Main2 string
		Main3 string
		Items []Item
	}
	var v = Test{
		Main1: 123,
		Main2: "ABC%20DEF",
		Main3: "GHI%20JKL",
		Items: []Item{
			{Sub1: "mno%20pqr", Sub2: 456},
			{Sub1: "stu%20vwx", Sub2: 789},
		},
	}
	JSUnescapeStruct(&v)
	var got = fmt.Sprintf("%v", v)
	if got != "{123 ABC DEF GHI JKL [{mno pqr 456} {stu vwx 789}]}" {
		t.Errorf("test failed")
	}
} //                                                 Test_strs_JSUnescapeStruct_

// go test --run Test_strs_Last_
func Test_strs_Last_(t *testing.T) {
	TBegin(t)
	// Last(s string, count int) string
	//
	{
		DisableErrors()
		TEqual(t, Last("abc", -1), (""))
		EnableErrors()
	}
	TEqual(t, Last("abc", 0), (""))
	TEqual(t, Last("abc", 1), ("c"))
	TEqual(t, Last("abc", 2), ("bc"))
	TEqual(t, Last("abc", 3), ("abc"))
	{
		DisableErrors()
		TEqual(t, Last("abc", 4), ("abc"))
		EnableErrors()
	}
} //                                                             Test_strs_Last_

// go test --run Test_strs_LineBeginIndexB_
func Test_strs_LineBeginIndexB_(t *testing.T) {
	TBegin(t)
	// LineBeginIndexB(s []byte, index int) int
	//
	var ar = []byte{
		'a', '\r', 'b', 'c', '\r', '\n', 'd', 'e', 'f'}
	//   0     1    2    3     4     5    6    7    8
	TEqual(t, LineBeginIndexB(ar, 0), (0))
	TEqual(t, LineBeginIndexB(ar, 1), (0))
	TEqual(t, LineBeginIndexB(ar, 2), (2))
	TEqual(t, LineBeginIndexB(ar, 3), (2))
	TEqual(t, LineBeginIndexB(ar, 4), (2))
	//
	//TODO: fix this wrong value (should be 2):
	TEqual(t, LineBeginIndexB(ar, 5), (5))
	//
	TEqual(t, LineBeginIndexB(ar, 6), (6))
	TEqual(t, LineBeginIndexB(ar, 7), (6))
	TEqual(t, LineBeginIndexB(ar, 8), (6))
	TEqual(t, LineBeginIndexB(ar, 9), (6))
	//
	// -1 is acceptable, it signifies end of byte array
	TEqual(t, LineBeginIndexB(ar, -1), (6))
	//
	// test exceptions:
	DisableErrors()
	TEqual(t, LineBeginIndexB(ar, 10), (6))
	EnableErrors()
} //                                                  Test_strs_LineBeginIndexB_

// go test --run Test_strs_LineBeginIndex_
func Test_strs_LineBeginIndex_(t *testing.T) {
	TBegin(t)
	// LineBeginIndex(s string, index int) int
	//
	var s = "A\rBC\rDEF"
	//
	// test with CR:
	TEqual(t, LineBeginIndex(s, 0), (0))
	TEqual(t, LineBeginIndex(s, 1), (0))
	TEqual(t, LineBeginIndex(s, 2), (2))
	TEqual(t, LineBeginIndex(s, 3), (2))
	TEqual(t, LineBeginIndex(s, 4), (2))
	TEqual(t, LineBeginIndex(s, 2), (2))
	TEqual(t, LineBeginIndex(s, 3), (2))
	TEqual(t, LineBeginIndex(s, 4), (2))
	TEqual(t, LineBeginIndex(s, 5), (5))
	TEqual(t, LineBeginIndex(s, 6), (5))
	TEqual(t, LineBeginIndex(s, 7), (5))
	TEqual(t, LineBeginIndex(s, 8), (5))
	//
	// test with LF:
	TEqual(t, LineBeginIndex(s, 0), (0))
	TEqual(t, LineBeginIndex(s, 1), (0))
	TEqual(t, LineBeginIndex(s, 2), (2))
	TEqual(t, LineBeginIndex(s, 3), (2))
	TEqual(t, LineBeginIndex(s, 4), (2))
	TEqual(t, LineBeginIndex(s, 2), (2))
	TEqual(t, LineBeginIndex(s, 3), (2))
	TEqual(t, LineBeginIndex(s, 4), (2))
	TEqual(t, LineBeginIndex(s, 5), (5))
	TEqual(t, LineBeginIndex(s, 6), (5))
	TEqual(t, LineBeginIndex(s, 7), (5))
	TEqual(t, LineBeginIndex(s, 8), (5))
	//
	// -1 is acceptable, signifies end of line
	TEqual(t, LineBeginIndex(s, -1), (5))
	//
	// 9 is longer than length of string
	// treat it as -1, but issue a warning
	DisableErrors()
	TEqual(t, LineBeginIndex(s, 9), (5))
	EnableErrors()
} //                                                   Test_strs_LineBeginIndex_

// go test --run Test_strs_LineEndIndexB_
func Test_strs_LineEndIndexB_(t *testing.T) {
	TBegin(t)
	// LineEndIndexB(s []byte, index int) int
	//
	TEqual(t, LineEndIndexB([]byte("abc"), 0), 3)
	//
	//TODO: more unit test cases for LineEndIndexB()
} //                                                    Test_strs_LineEndIndexB_

// go test --run Test_strs_LineEndIndex_
func Test_strs_LineEndIndex_(t *testing.T) {
	TBegin(t)
	// LineEndIndex(s string, index int) int
	//
	// test with CR:
	//       0 123 4567
	var s = "A\nBC\nDEF"
	TEqual(t, LineEndIndex(s, 0), (1))
	TEqual(t, LineEndIndex(s, 1), (1))
	TEqual(t, LineEndIndex(s, 2), (4))
	TEqual(t, LineEndIndex(s, 3), (4))
	TEqual(t, LineEndIndex(s, 4), (4))
	TEqual(t, LineEndIndex(s, 5), (8))
	TEqual(t, LineEndIndex(s, 6), (8))
	TEqual(t, LineEndIndex(s, 7), (8))
	TEqual(t, LineEndIndex(s, 8), (8))
	//
	// test with LF:
	//   0 123 4567
	s = "A\rBC\rDEF"
	TEqual(t, LineEndIndex(s, 0), (1))
	TEqual(t, LineEndIndex(s, 1), (1))
	TEqual(t, LineEndIndex(s, 2), (4))
	TEqual(t, LineEndIndex(s, 3), (4))
	TEqual(t, LineEndIndex(s, 4), (4))
	TEqual(t, LineEndIndex(s, 5), (8))
	TEqual(t, LineEndIndex(s, 6), (8))
	TEqual(t, LineEndIndex(s, 7), (8))
	TEqual(t, LineEndIndex(s, 8), (8))
	//
	// test exceptions:
	TEqual(t, LineEndIndex(s, -1), (8))
	DisableErrors()
	TEqual(t, LineEndIndex(s, 9), (8))
	EnableErrors()
} //                                                     Test_strs_LineEndIndex_

// go test --run Test_strs_LineOfIndex_
func Test_strs_LineOfIndex_(t *testing.T) {
	TBegin(t)
	// LineOfIndex(s string, index int) string
	//
	var s = "01\n345\n789A"
	TEqual(t, LineOfIndex(s, 0), ("01"))
	TEqual(t, LineOfIndex(s, 1), ("01"))
	TEqual(t, LineOfIndex(s, 2), ("01"))
	TEqual(t, LineOfIndex(s, 3), ("345"))
	TEqual(t, LineOfIndex(s, 4), ("345"))
	TEqual(t, LineOfIndex(s, 5), ("345"))
	TEqual(t, LineOfIndex(s, 6), ("345"))
	TEqual(t, LineOfIndex(s, 7), ("789A"))
	TEqual(t, LineOfIndex(s, 8), ("789A"))
	TEqual(t, LineOfIndex(s, 9), ("789A"))
	TEqual(t, LineOfIndex(s, 10), ("789A"))
	TEqual(t, LineOfIndex(s, 11), ("789A"))
	TEqual(t, LineOfIndex(s, -1), (""))
	TEqual(t, LineOfIndex(s, 12), (""))
} //                                                      Test_strs_LineOfIndex_

// go test --run Test_strs_LineOffsetUTF8_
func Test_strs_LineOffsetUTF8_(t *testing.T) {
	TBegin(t)
	// LineOffsetUTF8(data []byte, lineIndex int) (byteOffset, charOffset int)
	//
	var test = func(
		data []byte, lineIndex,
		expectByteOffset, expectCharOffset int,
	) {
		var byteOffset, charOffset = LineOffsetUTF8(data, lineIndex)
		if byteOffset != expectByteOffset || charOffset != expectCharOffset {
			t.Errorf(
				"LineOffsetUTF8(%q, %d) returned (%d, %d) instead of (%d, %d)",
				string(data), lineIndex,
				byteOffset, charOffset,
				expectByteOffset, expectCharOffset,
			)
		}
	}
	{
		var data = []byte("abc\r\nАБВ\n")
		//        each Cyrillic ^ char is 2 bytes long in UTF8
		test(data, 0,
			0, 0)
		test(data, 1,
			5, 5)
		test(data, 2,
			12, 9)
		test(data, 3,
			-1, -1)
	}
	{
		var data = []byte("岸\n123")
		// Chinese char ^ is 3 bytes long in UTF8
		test(data, 0,
			0, 0)
		test(data, 1,
			4, 2)
		test(data, 3,
			-1, -1)
	}
	{
		var data = []byte("\r\r\r")
		test(data, 0,
			0, 0)
		test(data, 1,
			1, 1)
		test(data, 2,
			2, 2)
		test(data, 3,
			3, 3)
	}
} //                                                   Test_strs_LineOffsetUTF8_

// go test --run Test_strs_Padf_
func Test_strs_Padf_(t *testing.T) {
	TBegin(t)
	// Padf(minLength int, s string, args ...interface{}) string
	//
	TEqual(t, Padf(6, "%s", "abc"), "abc   ")
	//
	//TODO: more unit test cases for Padf()
} //                                                             Test_strs_Padf_

// go test --run Test_strs_ReplaceEx1_
func Test_strs_ReplaceEx1_(t *testing.T) {
	TBegin(t)
	// ReplaceEx1(s, find, repl string, count int, caseMode CaseMode) string
	//
	var s = "ABC ABC ABC"
	TEqual(t, ReplaceEx1(s, "A", "X", 0, MatchCase), ("ABC ABC ABC"))
	TEqual(t, ReplaceEx1(s, "A", "X", 1, MatchCase), ("XBC ABC ABC"))
	TEqual(t, ReplaceEx1(s, "A", "X", 2, MatchCase), ("XBC XBC ABC"))
	TEqual(t, ReplaceEx1(s, "A", "X", 3, MatchCase), ("XBC XBC XBC"))
	TEqual(t, ReplaceEx1(s, "A", "X", 4, MatchCase), ("XBC XBC XBC"))
	TEqual(t, ReplaceEx1(s, "A", "X", -1, MatchCase), ("XBC XBC XBC"))
	TEqual(t, ReplaceEx1("ABC", "A", "A", -1, MatchCase), ("ABC"))
	TEqual(t, ReplaceEx1("A AA AAA", "A", "B", 3, MatchCase), ("B BB AAA"))
	s = "aa aA Aa AA"
	TEqual(t, ReplaceEx1(s, "AA", "aa", 0, IgnoreCase), ("aa aA Aa AA"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", 0, MatchCase), ("aa aA Aa AA"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", -1, IgnoreCase), ("aa aa aa aa"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", 1, IgnoreCase), ("aa aA Aa AA"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", 1, MatchCase), ("aa aA Aa aa"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", -1, MatchCase), ("aa aA Aa aa"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", 2, IgnoreCase), ("aa aa Aa AA"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", 3, IgnoreCase), ("aa aa aa AA"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", 4, IgnoreCase), ("aa aa aa aa"))
	TEqual(t, ReplaceEx1(s, "AA", "aa", 5, IgnoreCase), ("aa aa aa aa"))
} //                                                       Test_strs_ReplaceEx1_

// go test --run Test_strs_ReplaceI_
func Test_strs_ReplaceI_(t *testing.T) {
	TBegin(t)
	// ReplaceI(s, find, repl string) string
	//
	TEqual(t, ReplaceI("A AA aaa", "A", "B"), ("B BB BBB"))
	//
	//TODO: more unit test cases for ReplaceI()
} //                                                         Test_strs_ReplaceI_

// go test --run Test_strs_ReplaceMany_
func Test_strs_ReplaceMany_(t *testing.T) {
	// ReplaceMany(
	//     s string,
	//     finds []string,
	//     repls []string,
	//     count int,
	//     caseMode CaseMode,
	//     wordMode WordMode,
	//  ) string
	//
	TBegin(t)
	//TODO: declaration comment
	//
	const IC = IgnoreCase
	const IW = IgnoreWord
	const MC = MatchCase
	const MW = MatchWord
	//
	var A = []string{"A"}
	var X = []string{"X"}
	var B = []string{"B"}
	var AA = []string{"AA"}
	var aa = []string{"aa"}
	//                                                     simplest replacements
	TEqual(t,
		ReplaceMany("X", []string{"X"}, []string{"9"}, -1, IC, MW),
		("9"),
	)
	TEqual(t,
		ReplaceMany("YZ", []string{"YZ"}, []string{"12"}, -1, IC, MW),
		("12"),
	)
	//                                            simple multi-byte replacements
	TEqual(t,
		ReplaceMany("Я", []string{"Я"}, []string{"9"}, -1, IC, MW),
		("9"),
	)
	TEqual(t,
		ReplaceMany("ЮЯ", []string{"ЮЯ"}, []string{"12"}, -1, IC, MW),
		("12"),
	)
	//                                                   replacements with blank
	TEqual(t,
		ReplaceMany("X", []string{"X"}, []string{""}, -1, IC, MW),
		(""),
	)
	TEqual(t,
		ReplaceMany("XYZ", []string{"XYZ"}, []string{""}, -1, IC, MW),
		(""),
	)
	TEqual(t,
		ReplaceMany("XYZ", []string{"Y"}, []string{""}, -1, IC, IW),
		("XZ"),
	)
	//                                                        multi-byte strings
	TEqual(t, ReplaceMany(
		"+А- +АБ- +АБВ- XYZ",
		[]string{"А", "АБ", "АБВ", "XYZ"},
		[]string{"1", "12", "123", "789"},
		-1, IC, MW,
	),
		("+1- +12- +123- 789"),
	)
	// each Cyrillic ^ char is 2 bytes long in UTF8
	TEqual(t, ReplaceMany( //                                         similar words
		"+CBA-+AB-+A-",
		[]string{"A", "AB", "CBA"},
		[]string{"1", "12", "321"},
		-1, IC, MW,
	),
		("+321-+12-+1-"),
	)
	TEqual(t, ReplaceMany(
		"+A- +AB- +ABC-",
		[]string{"A", "AB", "ABC"},
		[]string{"1", "12", "123"},
		-1, MC, MW,
	),
		("+1- +12- +123-"),
	)
	//                                                                first word
	TEqual(t, ReplaceMany(
		"A+-AB<>ABC", []string{"A"}, []string{"X"}, -1, MC, MW,
	),
		("X+-AB<>ABC"),
	)
	//                                                            word in middle
	TEqual(t, ReplaceMany(
		"ABC+-BC<>B", []string{"BC"}, []string{"XY"}, -1, MC, MW,
	),
		("ABC+-XY<>B"),
	)
	//                                                                 last word
	TEqual(t, ReplaceMany(
		"ABC+-BC<>B",
		[]string{"B"},
		[]string{"X"},
		-1, MC, MW,
	),
		("ABC+-BC<>X"),
	)
	//                                                             miscellaneous
	TEqual(t, ReplaceMany(
		"ABC DEF GHI",
		[]string{"ABC", "GHI"},
		[]string{"123", "456"},
		-1, MC, IW,
	),
		("123 DEF 456"),
	)
	var s = "ABC ABC ABC"
	TEqual(t, ReplaceMany("A AA AAA", A, B, -1, MC, IW), ("B BB BBB"))
	TEqual(t, ReplaceMany("A AA AAA", A, B, 3, MC, IW), ("B BB AAA"))
	TEqual(t, ReplaceMany("ABC", A, A, -1, MC, IW), ("ABC"))
	TEqual(t, ReplaceMany(s, A, X, -1, MC, IW), ("XBC XBC XBC"))
	TEqual(t, ReplaceMany(s, A, X, 0, MC, IW), ("ABC ABC ABC"))
	TEqual(t, ReplaceMany(s, A, X, 1, MC, IW), ("XBC ABC ABC"))
	TEqual(t, ReplaceMany(s, A, X, 2, MC, IW), ("XBC XBC ABC"))
	TEqual(t, ReplaceMany(s, A, X, 3, MC, IW), ("XBC XBC XBC"))
	TEqual(t, ReplaceMany(s, A, X, 4, MC, IW), ("XBC XBC XBC"))
	s = "aa aA Aa AA"
	TEqual(t, ReplaceMany(s, AA, aa, -1, IC, IW), ("aa aa aa aa"))
	TEqual(t, ReplaceMany(s, AA, aa, -1, MC, IW), ("aa aA Aa aa"))
	TEqual(t, ReplaceMany(s, AA, aa, 0, IC, IW), ("aa aA Aa AA"))
	TEqual(t, ReplaceMany(s, AA, aa, 0, MC, IW), ("aa aA Aa AA"))
	TEqual(t, ReplaceMany(s, AA, aa, 1, IC, IW), ("aa aA Aa AA"))
	TEqual(t, ReplaceMany(s, AA, aa, 1, MC, IW), ("aa aA Aa aa"))
	TEqual(t, ReplaceMany(s, AA, aa, 2, IC, IW), ("aa aa Aa AA"))
	TEqual(t, ReplaceMany(s, AA, aa, 3, IC, IW), ("aa aa aa AA"))
	TEqual(t, ReplaceMany(s, AA, aa, 4, IC, IW), ("aa aa aa aa"))
	TEqual(t, ReplaceMany(s, AA, aa, 5, IC, IW), ("aa aa aa aa"))
} //                                                      Test_strs_ReplaceMany_

// go test --run Test_strs_ReplaceWord_
func Test_strs_ReplaceWord_(t *testing.T) {
	TBegin(t)
	// ReplaceWord(s, find, repl string, caseMode CaseMode) string
	//
	const testStr = "A AB ABC A AB ABC A_ AB_"
	//
	// simplest replacements:
	TEqual(t, ReplaceWord("A", "A", "B", MatchCase), ("B"))
	TEqual(t, ReplaceWord("A", "A", "B", MatchCase), ("B"))
	TEqual(t, ReplaceWord("Z A", "A", "B", MatchCase), ("Z B"))
	TEqual(t, ReplaceWord("A Z", "A", "B", MatchCase), ("B Z"))
	//
	// simple replacements, considering case:
	TEqual(t, ReplaceWord("A a", "A", "B", IgnoreCase), ("B B"))
	TEqual(t, ReplaceWord("A a", "A", "B", MatchCase), ("B a"))
	//
	// don't replace words that contain search but differ:
	TEqual(t, ReplaceWord("A AB a", "A", "A", IgnoreCase), ("A AB a"))
	TEqual(t, ReplaceWord("A AB a", "A", "Z", IgnoreCase), ("Z AB Z"))
	TEqual(t, ReplaceWord("A AB a", "A", "A", MatchCase), ("A AB a"))
	TEqual(t, ReplaceWord("A AB a", "A", "Z", MatchCase), ("Z AB a"))
	//
	// other tests:
	TEqual(t, ReplaceWord(testStr, "A", "A", MatchCase),
		("A AB ABC A AB ABC A_ AB_"),
	)
	TEqual(t, ReplaceWord("A AA AAA AAAA A", "A", "Z", IgnoreCase),
		("Z AA AAA AAAA Z"),
	)
	TEqual(t, ReplaceWord(testStr, "A", "Z", IgnoreCase),
		("Z AB ABC Z AB ABC A_ AB_"),
	)
	TEqual(t, ReplaceWord(testStr, "AB", "XY", MatchCase),
		("A XY ABC A XY ABC A_ AB_"),
	)
	TEqual(t, ReplaceWord("A", "A", "ZZZ", IgnoreCase), ("ZZZ"))
	TEqual(t, ReplaceWord("A", "A", "ZZZ", MatchCase), ("ZZZ"))
	TEqual(t, ReplaceWord("A A A", "A", "A", MatchCase), ("A A A"))
	var tests = []struct {
		s        string
		find     string
		repl     string
		caseMode CaseMode
		expect   string
	}{
		{"A AA AAA AAAA A", "A", "Z", MatchCase, "Z AA AAA AAAA Z"},
		{testStr, "A", "Z", MatchCase, "Z AB ABC Z AB ABC A_ AB_"},
		{testStr, "A", "A", MatchCase, "A AB ABC A AB ABC A_ AB_"},
		{testStr, "AB", "XY", MatchCase, "A XY ABC A XY ABC A_ AB_"},
		{"A", "A", "ZZZ", MatchCase, "ZZZ"},
		{"A A A", "A", "A", MatchCase, "A A A"},
	}
	//TODO: use test() function instead of loop
	for i, test := range tests {
		var got = ReplaceWord(test.s, test.find, test.repl, test.caseMode)
		if got != test.expect {
			t.Errorf("T%d ReplaceWord(%q, %q, %q, %v)"+
				" returned %q instead of %q",
				i, test.s, test.find, test.repl,
				test.caseMode, got, test.expect,
			)
		}
	}
} //                                                      Test_strs_ReplaceWord_

// go test --run Test_strs_SetPart_
func Test_strs_SetPart_(t *testing.T) {
	TBegin(t)
	// SetPart(s, prefix, suffix, part string) string
	//
	TEqual(t,
		SetPart("name:bill;age:30;", "age:", ";", "70"),
		"name:bill;age:70;",
	)
} //                                                          Test_strs_SetPart_

// go test --run Test_strs_SetSlice_
func Test_strs_SetSlice_(t *testing.T) {
	TBegin(t)
	// SetSlice(s string, start, end int, substr string) string
	//
	var test = func(s string, start, end int, substr string,
		expect string) {
		var got = SetSlice(s, start, end, substr)
		TEqual(t, got, (expect))
	}
	test("", 0, -1, "ABC", "ABC")
} //                                                         Test_strs_SetSlice_

// go test --run Test_strs_ShowSpaces_
func Test_strs_ShowSpaces_(t *testing.T) {
	//
	// whitespace characters change to:
	TEqual(t, ShowSpaces(" "), "-")
	TEqual(t, ShowSpaces("\t"), "--->")
	TEqual(t, ShowSpaces("\r"), "CR\r")
	TEqual(t, ShowSpaces("\n"), "LF\n")
	TEqual(t, ShowSpaces("\r\n"), "CRLF\r\n")
	//
	// CR+LF combinations
	TEqual(t, ShowSpaces("\r \n"), "CR\r-LF\n")
	TEqual(t, ShowSpaces("\r\n \r \n \n \r"), "CRLF\r\n-CR\r-LF\n-LF\n-CR\r")
} //                                                       Test_strs_ShowSpaces_

// go test --run Test_strs_SkipChars_
func Test_strs_SkipChars_(t *testing.T) {
	TBegin(t)
	// SkipChars(s string, start int, chars string) int
	//
	var test = func(s string, start int, chars string, expect int) {
		var got = SkipChars(s, start, chars)
		TEqual(t, got, (expect))
	}
	test("ABC123", 0, "ABC", 3)
	test("ABC123", 6, "ABC", 6)
	//
	// if start is greater that string length, return string length
	test("ABC123", 7, "ABC", 6)
	//
	// treat a negative start as zero
	test("ABC123", -10, "ABC", 3)
} //                                                        Test_strs_SkipChars_

// go test --run Test_strs_SkipName_
func Test_strs_SkipName_(t *testing.T) {
	TBegin(t)
	// SkipName(s string, start int) int
	//
	var test = func(s string, start, expect int) {
		var got = SkipName(s, start)
		TEqual(t, got, (expect))
	}
	test("abc 123", 0, 3)
	test("abc 123", 3, 3)
} //                                                         Test_strs_SkipName_

// go test --run Test_strs_SkipSpaces_
func Test_strs_SkipSpaces_(t *testing.T) {
	TBegin(t)
	// SkipSpaces(s string, start int) int
	//
	var test = func(s string, start, expect int) {
		var got = SkipSpaces(s, start)
		TEqual(t, got, (expect))
	}
	test("ABC   123", 3, 6)
} //                                                       Test_strs_SkipSpaces_

// go test --run Test_strs_Slice_
func Test_strs_Slice_(t *testing.T) {
	TBegin(t)
	// Slice(s string, beginIndex, endIndex int) string
	//
	//TODO: use similar testing for Bytes.Slice()
	//
	// expected range values:
	TEqual(t, Slice("abc", 0, 0), (""))
	TEqual(t, Slice("abc", 0, 1), ("a"))
	TEqual(t, Slice("abc", 0, 2), ("ab"))
	TEqual(t, Slice("abc", 0, 3), ("abc"))
	//
	// endIndex -1 is acceptable: it signifies end of string
	TEqual(t, Slice("abc", 0, -1), ("abc"))
	TEqual(t, Slice("abc", 1, -1), ("bc"))
	TEqual(t, Slice("abc", 2, -1), ("c"))
	TEqual(t, Slice("abc", 3, -1), (""))
	{
		// tests with out-of-range arguments:
		DisableErrors()
		// tests with beginIndex == -1
		//
		// beginIndex should never be negative as you can use
		// zero to indicate the beginning of the string.
		//
		// if beginIndex < 0, Slice() issues a warning and treats it as 0
		TEqual(t, Slice("abc", -1, 0), (""))
		TEqual(t, Slice("abc", -1, 1), ("a"))
		TEqual(t, Slice("abc", -1, 2), ("ab"))
		TEqual(t, Slice("abc", -1, 3), ("abc"))
		TEqual(t, Slice("abc", -1, -1), ("abc"))
		//
		// same as previous if beginIndex < -1
		TEqual(t, Slice("abc", -2, 0), (""))
		TEqual(t, Slice("abc", -1024, 1), ("a"))
		TEqual(t, Slice("abc", -1048576, 2), ("ab"))
		TEqual(t, Slice("abc", -1073741824, 3), ("abc"))
		TEqual(t, Slice("abc", -2, -1), ("abc"))
		//
		// endIndex < -1 should not happen
		// a warning is issued, value is treated as -1 (end-of-string)
		TEqual(t, Slice("abc", 0, -2), ("abc"))
		TEqual(t, Slice("abc", 0, -2), ("abc"))
		TEqual(t, Slice("abc", 0, -1024), ("abc"))
		TEqual(t, Slice("abc", 0, -1048576), ("abc"))
		TEqual(t, Slice("abc", 0, -1073741824), ("abc"))
		TEqual(t, Slice("abc", -1, -2), ("abc"))
		TEqual(t, Slice("abc", -2, -2), ("abc"))
		TEqual(t, Slice("abc", -1024, -1024), ("abc"))
		TEqual(t, Slice("abc", -1048576, -1048576), ("abc"))
		TEqual(t, Slice("abc", -1073741824, -1073741824), ("abc"))
		//
		// endIndex 4 is out of range:
		TEqual(t, Slice("abc", 0, 4), ("abc"))
		EnableErrors()
	}
} //                                                            Test_strs_Slice_

// go test --run Test_strs_SplitQuoted_
func Test_strs_SplitQuoted_(t *testing.T) {
	TBegin(t)
	// SplitQuoted(s string) []string
	//
	TArrayEqual(t, []string{`"ABC"`, "123"},
		SplitQuoted(`'"ABC"' '123'`))
	TArrayEqual(t, []string{"ABC"},
		SplitQuoted("ABC"))
	TArrayEqual(t, []string{"123", "ABC"},
		SplitQuoted("123 'ABC'"))
	TArrayEqual(t, []string{"XYZ", "'ABC'", "DEF"},
		SplitQuoted("`XYZ` `'ABC'` \"DEF\""))
	TArrayEqual(t, []string{"123", "456", "ABC"},
		SplitQuoted("123 456 'ABC'"))
	TArrayEqual(t, []string{"123", "456", "A B C"},
		SplitQuoted("123   456   'A B C'"))
	var tests = []struct {
		s      string
		expect []string
	}{
		{`'"ABC"' '123'`, []string{`"ABC"`, "123"}},
		{"ABC", []string{"ABC"}},
		{"123 'ABC'", []string{"123", "ABC"}},
		{"123 456 'ABC'", []string{"123", "456", "ABC"}},
		{"123   456   'A B C'", []string{"123", "456", "A B C"}},
	}
	//TODO: use test() function instead of loop
	for i, test := range tests {
		var got = SplitQuoted(test.s)
		if !EqualStringSlices(got, test.expect) {
			t.Errorf("T%d SplitQuoted(%q) returned %q instead of %v"+LB,
				i, test.s, got, test.expect)
		}
	}
} //                                                      Test_strs_SplitQuoted_

// go test --run Test_strs_StrOneOf_
func Test_strs_StrOneOf_(t *testing.T) {
	TBegin(t)
	// StrOneOf(s string, matches ...string) bool
	//
	TEqual(t, StrOneOf("b", "a", "b", "c"), true)
	TEqual(t, StrOneOf("d", "a", "b", "c"), false)
} //                                                         Test_strs_StrOneOf_

// go test --run Test_strs_String_
func Test_strs_String_(t *testing.T) {
	TBegin(t)
	// String(val interface{}) string
	//
	var test = func(input interface{}, expect string) {
		var got = String(input)
		TEqual(t, got, (expect))
	}
	// floating-point values should not have trailing zeros
	test(float32(0.1), "0.1")
	test(float64(0.1), "0.1")
	test(float32(0.01), "0.01")
	test(float64(0.01), "0.01")
	test(float32(0.001), "0.001")
	test(float64(0.001), "0.001")
	//
	// retains zeros in integer portion?
	test(float32(1000.0), "1000")
	test(float64(1000.0), "1000")
	test(float32(1000), "1000")
	test(float64(1000), "1000")
	//
	// positive and negative integers
	test(123, "123")
	test(-123, "-123")
	// test zero values
	test(int(0), "0")
	test(float32(0), "0")
	test(float64(0), "0")
} //                                                           Test_strs_String_

// go test --run Test_strs_Substr_
func Test_strs_Substr_(t *testing.T) {
	TBegin(t)
	// Substr(s string, charIndex, charCount int) string
	//
	const IntMin = -2147483648
	const GB1 = 1073741824
	//
	TEqual(t, Substr("ABC", 0, 3), ("ABC"))
	TEqual(t, Substr("ABC", 0, 2), ("AB"))
	TEqual(t, Substr("ABC", 0, 1), ("A"))
	TEqual(t, Substr("ABC", 0, 0), (""))
	//    // ? Mid( "", -1073741824, -2147483647) = ""
	//    TEST( MidS( "", -1073741824, int(-2147483647)) == "" )
	//
	//    // ? Mid( "False", -2147483647, -2147483647) = "F"
	//    TEST( MidS( "False", -2147483647, int(-2147483647)) == "F" )
	//
	//    // ? Mid( "abc", 1, -2147483648) = ""
	//    // ? Mid( "abc", 1, -2147483647) = "a"
	//    // ? Mid( "abc", 1, -2147483646) = "ab"
	//    // ? Mid( "abc", 1, -2147483645) = "abc"
	//
	//    TEST( MidS( "abc", 1, int(-2147483648)) == "" )
	//
	//#pragma warning(pop)
	//
	//    TEST( MidS( "abc", 1, int(-2147483647)) == "a"   )
	//    TEST( MidS( "abc", 1, int(-2147483646)) == "ab"  )
	//    TEST( MidS( "abc", 1, int(-2147483645)) == "abc" )
	//
	//    // ? Mid( "abc", 1, -1073741825) = "abc"
	//    TEST( MidS( "abc", 1, int(-1073741825)) == "abc" )
	//
	//    // ? Mid( "abc", 2, -2147483647) = "b"
	//    TEST( MidS( "abc", 2, int(-2147483647)) == "b" )
	//
	//    // ? Mid( "abcdefghij", 9, -2147483647) = "i"
	//    TEST( MidS( "abcdefghij", 9, int(-2147483647)) == "i" )
	//
	//    TEST( MidS( "abc", 1, int(-2147483647)) == "a"   )
	//    TEST( MidS( "abc", 1, int(-2147483646)) == "ab"  )
	//    TEST( MidS( "abc", 1, int(-2147483645)) == "abc" )
	//
	//    TEST( MidS( "abcdefghi", -1073741824, 65536) == "" )
	//    //TEST( MidS( "abc", 1, int(-2147483648)) == "" )
	//
	//    TEST( MidS( "abc",   1, int(1))      == "a"           )
	//    TEST( MidS( "abc",   1, int(3))      == "abc"         )
	//    TEST( MidS( "abc",   1, int(300))    == "abc"         )
	//    TEST( MidS( "abc",   1, "")          == ""            )
	//    TEST( MidS( "abc",   2, int(300))    == "bc"          )
	//    TEST( MidS( "abc",   3, int(300))    == "c"           )
	//    TEST( MidS( "abc",   4, int(300))    == ""            )
	//    TEST( MidS( "abc", 777, ZERO)        == ""            )
	//    TEST( MidS( "abc", 777, int(1))      == ""            )
	//    TEST( MidS(NULL,    -1, int(-65536)) == "" )
	//    TEST( MidS(NULL,     0, int(-65536)) == "" )
	//    TEST( MidS(NULL,     1, int(1))      == "" )
	//    TEST( MidS(NULL,    10, int(-1))     == "" )
	//    TEST( MidS(NULL,    10, int(-65536)) == "" )
	//    TEST( MidS(NULL,    10, ZERO)        == "" )
	//    TEST( MidS(NULL,    10, int(1))      == "" )
	//    TEST( MidS("",       1, int(1))      == ""            )
	//
	//    // Run-time error '5': Invalid procedure call or argument
	//    //TEST( MidS( "abc", 77777777, -1 )
	//    //TEST( MidS(empty, 0, 1 )
	//    //TEST( MidS( "abc", 0, 1 )
	//    //TEST( MidS( "abc", -1, 300 )
	//
	//    // Run-time error '94': Invalid use of Null
	//    //TEST( Mid( "abc", 77777777, Null )
	var tests = []struct {
		text      string
		charIndex int
		charCount int
		expect    string
	}{
		{"АБВГ", 0, 0, ""},
		{"АБВГ", 0, 1, "А"},
		{"АБВГ", 0, 2, "АБ"},
		{"АБВГ", 0, 3, "АБВ"},
		{"АБВГ", 0, 4, "АБВГ"},
		{"АБВГ", 1, 1, "Б"},
		{"АБВГ", 1, 2, "БВ"},
		{"АБВГ", 1, 3, "БВГ"},
		{"АБВГ", 2, 1, "В"},
		{"АБВГ", 2, 2, "ВГ"},
		{"АБВГ", 3, 1, "Г"},
	}
	//TODO: use test() function instead of loop
	for i, test := range tests {
		var got = Substr(test.text, test.charIndex, test.charCount)
		if got != test.expect {
			t.Errorf("T%d Substr(%q, %d, %d) returned %q instead of %q",
				i, test.text, test.charIndex, test.charCount, got, test.expect)
		}
	}
} //                                                           Test_strs_Substr_

// go test --run Test_strs_TitleCase_
func Test_strs_TitleCase_(t *testing.T) {
	TBegin(t)
	// TitleCase(s string) string
	//
	TEqual(t, TitleCase("A TITLE"), ("A Title"))
	TEqual(t, TitleCase("a title"), ("A Title"))
	var tests = []struct {
		input  string
		expect string
	}{
		{"", ""},
		{"a", "A"},
		{"A", "A"},
		{"abc abc abc", "Abc Abc Abc"},
		{"ABC ABC ABC", "Abc Abc Abc"},
		{"АБВ АБВ АБВ", "Абв Абв Абв"},
		{"АБВ", "Абв"},
	}
	//TODO: use test() function instead of loop
	for i, test := range tests {
		var got = TitleCase(test.input)
		if got != test.expect {
			t.Errorf("T%d TitleCase(%q) returned %q instead of %q",
				i, test.input, got, test.expect)
		}
	}
} //                                                        Test_strs_TitleCase_

// go test --run Test_strs_TokenGetEx_
func Test_strs_TokenGetEx_(t *testing.T) {
	TBegin(t)
	// TokenGetEx(list string, index int, sep string, ignoreEnd bool) string
	//
	var FST = "first;second;third"
	TEqual(t, TokenGetEx(FST, 0, ";", false), ("first"))
	TEqual(t, TokenGetEx(FST, 0, ";", false), ("first"))
	TEqual(t, TokenGetEx(FST, 1, ";", false), ("second"))
	TEqual(t, TokenGetEx(FST, 2, ";", false), ("third"))
	TEqual(t, TokenGetEx(FST, 3, ";", false), (""))
	TEqual(t, TokenGetEx(FST, 0, "", false), (FST))
	TEqual(t, TokenGetEx(FST, 1, "", false), (""))
	TEqual(t, TokenGetEx(FST, 0, ";", true), (FST))
	TEqual(t, TokenGetEx(FST, 1, ";", true), ("second;third"))
	TEqual(t, TokenGetEx(FST, 2, ";", true), ("third"))
	TEqual(t, TokenGetEx(FST, 3, ";", true), (""))
	var ABC = "a;;;b;;;c"
	TEqual(t, TokenGetEx(ABC, 0, "123456789", false), (""))
	TEqual(t, TokenGetEx(ABC, 0, "1234567890", false), (""))
	TEqual(t, TokenGetEx(ABC, 0, ";;;", false), ("a"))
	TEqual(t, TokenGetEx(ABC, 1, ";;;", false), ("b"))
	TEqual(t, TokenGetEx(ABC, 2, ";;;", false), ("c"))
	TEqual(t, TokenGetEx(ABC, 0, "123456789", true), (""))
	TEqual(t, TokenGetEx(ABC, 0, "1234567890", true), (""))
	TEqual(t, TokenGetEx(ABC, 0, ";;;", true), (ABC))
	TEqual(t, TokenGetEx(ABC, 1, ";;;", true), ("b;;;c"))
	TEqual(t, TokenGetEx(ABC, 2, ";;;", true), ("c"))
} //                                                       Test_strs_TokenGetEx_

// go test --run Test_strs_TokenGet_
func Test_strs_TokenGet_(t *testing.T) {
	TBegin(t)
	// TokenGet(list string, index int, sep string) string
	//
	TEqual(t, TokenGet("first;second;third", 0, ";"), ("first"))
	TEqual(t, TokenGet("first;second;third", 1, ";"), ("second"))
	TEqual(t, TokenGet("first;second;third", 2, ";"), ("third"))
	TEqual(t, TokenGet("first;second;third", 3, ";"), (""))
	TEqual(t, TokenGet("a;;;b;;;c", 0, ";;;"), ("a"))
	TEqual(t, TokenGet("a;;;b;;;c", 1, ";;;"), ("b"))
	TEqual(t, TokenGet("a;;;b;;;c", 2, ";;;"), ("c"))
	TEqual(t, TokenGet("a;;;b;;;c", 3, ";;;"), (""))
} //                                                         Test_strs_TokenGet_

// go test --run Test_strs_WordIndex_
func Test_strs_WordIndex_(t *testing.T) {
	TBegin(t)
	// WordIndex(s, word string, caseMode CaseMode) int
	//
	TEqual(t, WordIndex("word", "WORD", IgnoreCase), (0))
	//
	// combinations with empty strings
	TEqual(t, WordIndex("", "", MatchCase), (0))
	TEqual(t, WordIndex("", "", IgnoreCase), (0))
	TEqual(t, WordIndex("", "abc", MatchCase), (-1))
	TEqual(t, WordIndex("", "abc", IgnoreCase), (-1))
	TEqual(t, WordIndex("abc", "", MatchCase), (0))
	TEqual(t, WordIndex("abc", "", IgnoreCase), (0))
	//
	// single word:
	TEqual(t, WordIndex("word", "WORD", MatchCase), (-1))
	TEqual(t, WordIndex("word", "WORD", IgnoreCase), (0))
	TEqual(t, WordIndex("word", "word", MatchCase), (0))
	TEqual(t, WordIndex("word", "words", MatchCase), (-1))
	TEqual(t, WordIndex("words", "word", MatchCase), (-1))
	TEqual(t, WordIndex("words", "word", IgnoreCase), (-1))
	//
	// word at beginning of string:
	TEqual(t, WordIndex("one two", "ONE", MatchCase), (-1))
	TEqual(t, WordIndex("one two", "ONE", IgnoreCase), (0))
	TEqual(t, WordIndex("one two", "one", MatchCase), (0))
	TEqual(t, WordIndex("ones two", "one", MatchCase), (-1))
	//
	// word at end of string:
	TEqual(t, WordIndex("one two", "TWO", MatchCase), (-1))
	TEqual(t, WordIndex("one two", "TWO", IgnoreCase), (4))
	TEqual(t, WordIndex("one two", "two", MatchCase), (4))
	TEqual(t, WordIndex("one twos", "two", MatchCase), (-1))
	//
	// word within string:
	TEqual(t, WordIndex("one two three", "TWO", MatchCase), (-1))
	TEqual(t, WordIndex("one two three", "TWO", IgnoreCase), (4))
	TEqual(t, WordIndex("one two three", "e two", MatchCase), (-1))
	TEqual(t, WordIndex("one two three", "two t", MatchCase), (-1))
	TEqual(t, WordIndex("one two three", "two", MatchCase), (4))
	TEqual(t, WordIndex("one two three", "O", MatchCase), (-1))
	TEqual(t, WordIndex("one two three", "O", IgnoreCase), (-1))
	TEqual(t, WordIndex("one two three", "o", MatchCase), (-1))
	TEqual(t, WordIndex("one two three", "o", IgnoreCase), (-1))
	TEqual(t, WordIndex("one two three", "ONE", IgnoreCase), (0))
	TEqual(t, WordIndex("one two three", "six", MatchCase), (-1))
	TEqual(t, WordIndex("ones two three", "one", MatchCase), (-1))
	var tests = []struct {
		s        string
		word     string
		caseMode CaseMode
		expect   int
	}{
		// combinations with empty strings
		{"", "", MatchCase, 0},
		{"", "", IgnoreCase, 0},
		{"", "abc", MatchCase, -1},
		{"", "abc", IgnoreCase, -1},
		{"abc", "", MatchCase, 0},
		{"abc", "", IgnoreCase, 0},
		// single
		{"word", "WORD", MatchCase, -1},
		{"word", "WORD", IgnoreCase, 0},
		{"word", "word", MatchCase, 0},
		{"word", "words", MatchCase, -1},
		{"words", "word", MatchCase, -1},
		{"words", "word", IgnoreCase, -1},
		// word at beginning of string:
		{"one two", "ONE", MatchCase, -1},
		{"one two", "ONE", IgnoreCase, 0},
		{"one two", "one", MatchCase, 0},
		{"ones two", "one", MatchCase, -1},
		// word at end of string:
		{"one two", "TWO", MatchCase, -1},
		{"one two", "TWO", IgnoreCase, 4},
		{"one two", "two", MatchCase, 4},
		{"one twos", "two", MatchCase, -1},
		// word within string:
		{"one two three", "TWO", MatchCase, -1},
		{"one two three", "TWO", IgnoreCase, 4},
		{"one two three", "e two", MatchCase, -1},
		{"one two three", "two t", MatchCase, -1},
		{"one two three", "two", MatchCase, 4},
		{"one two three", "O", MatchCase, -1},
		{"one two three", "O", IgnoreCase, -1},
		{"one two three", "o", MatchCase, -1},
		{"one two three", "o", IgnoreCase, -1},
		{"one two three", "ONE", IgnoreCase, 0},
		{"one two three", "six", MatchCase, -1},
		{"ones two three", "one", MatchCase, -1},
	}
	//TODO: use test() function instead of loop
	for i, test := range tests {
		var got = WordIndex(test.s, test.word, test.caseMode)
		if got != test.expect {
			t.Errorf("T%d WordIndex(%q, %q, %v)"+
				" returned %q instead of %q",
				i, test.s, test.word, test.caseMode, got, test.expect)
		}
	}
} //                                                        Test_strs_WordIndex_

//TODO: rename CharsOf to Runes.

//end
