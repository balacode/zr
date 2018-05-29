// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-29 09:40:32 EB8482                                [zr/strings.go]
// -----------------------------------------------------------------------------

package zr

// # String Functions
//   After(s string, find ...string) string
//   CamelCase(s string) string
//   CharsOf(s string) []rune
//   CompactSpaces(s string) string
//   ContainsI(s, substr string) bool
//   ContainsWord(s, word string, caseMode CaseMode) bool
//   CountCRLF(s string) (count, countCR, countLF int)
//   EqualStringSlices(a, b []string) bool
//   FindChar(s string, ch byte, beginIndex int) int
//   FindInSlice(s string, start, end int, substr string) int
//   First(s string, count int) string
//   GetPart(s, prefix, suffix string) string
//   IfString(condition bool, trueStr, falseStr string) string
//   IsIdentifier(s string) bool
//   IsWhiteSpace(s string) bool
//   JSUnescapeStruct(structPtr interface{})
//   JSUnescape(s string) string
//   Last(s string, count int) string
//   LineBeginIndex(s string, index int) int
//   LineBeginIndexB(s []byte, index int) int
//   LineEndIndex(s string, index int) int
//   LineEndIndexB(s []byte, index int) int
//   LineOfIndex(s string, index int) string
//   LineOffsetUTF8(data []byte, lineIndex int) (byteOffset, charOffset int)
//   Padf(minLength int, s string, args ...interface{}) string
//   ReplaceEx1(s, find, repl string, count int, caseMode CaseMode) string
//   ReplaceI(s, find, repl string, optCount ...int) string
//   ReplaceMany(
//       s string,
//       finds []string,
//       repls []string,
//       count int,
//       caseMode CaseMode,
//       wordMode WordMode,
//   ) string
//   ReplaceWord(s, find, repl string, caseMode CaseMode) string
//   SetPart(s, prefix, suffix, part string) string
//   SetSlice(s string, start, end int, substr string) string
//   ShowSpaces(s string) string
//   SkipChars(s string, start int, chars string) int
//   SkipName(s string, start int) int
//   SkipSpaces(s string, start int) int
//   Slice(s string, beginIndex, endIndex int) string
//   SplitQuoted(s string) []string
//   StrOneOf(s string, matches ...string) bool
//   String(val interface{}) string
//   Substr(s string, charIndex, charCount int) string
//       //TODO: ^beginIndex, endIndex. Why Substr when there's Slice()?
//   TitleCase(s string) string
//   TokenGet(list string, index int, sep string) string
//   TokenGetEx(list string, index int, sep string, ignoreEnd bool) string
//   WordIndex(s, word string, caseMode CaseMode) int

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// CaseMode type specifies if text comparisons are case sensitive.
// Various functions have a CaseMode parameter.
//
// There are two possible values:
//
// IgnoreCase: text is matched irrespective of letter capitalization.
// MatchCase:  case must be matched exactly.
type CaseMode uint

// IgnoreCase constant specifies that text
// comparisons should not be case sensitive.
const IgnoreCase = CaseMode(1)

// MatchCase constant specifies that text
// comparisons should be case sensitive.
const MatchCase = CaseMode(2)

// WordMode type determines if a search should match whole
// words when searching (or replacing, etc.) a string.
//
// Words are composed of alphabetic and numeric
// characters and underscores.
//
// There are two possible values:
// IgnoreWord: do not match distinct words.
// MatchWord:  match distinct words.
type WordMode uint

// IgnoreWord constant specifies that string
// searches should not match distinct words.
const IgnoreWord = WordMode(1)

// MatchWord constant specifies that string
// searches should match distinct words.
const MatchWord = WordMode(2)

// -----------------------------------------------------------------------------
// # String Functions

// After returns the part of 's' immediately after the last 'find' string.
func After(s string, find ...string) string {
	if s == "" {
		return ""
	}
	var at = -1
	for _, f := range find {
		var i = str.Index(s, f)
		if i != -1 {
			i += len(f)
			if i > at {
				at = i
			}
		}
	}
	if at == -1 {
		return ""
	}
	return s[at:]
} //                                                                       After

// CamelCase converts an identifier from underscore naming convention
// to camel case naming convention: 'some_name' becomes 'someName'.
func CamelCase(s string) string {
	var retBuf = bytes.NewBuffer(make([]byte, 0, len(s)))
	var ws = retBuf.WriteString
	var ucase bool
	for _, ch := range s {
		if ch == '_' {
			ucase = true
			continue
		}
		if ucase {
			ucase = false
			ch = unicode.ToUpper(ch)
		}
		ws(string(ch))
	}
	return retBuf.String()
} //                                                                   CamelCase

// CharsOf __
func CharsOf(s string) []rune {
	return []rune(s)
} //                                                                     CharsOf

// CompactSpaces reduces all multiple white-spaces in string to
// a single space, also converting newlines and tabs to space.
// E.g. "a\n b   c" becomes "a b c".
func CompactSpaces(s string) string {
	for _, ch := range "\a\b\f\n\r\t\v" { // <- no need to include a space here
		if str.Contains(s, string(ch)) {
			s = str.Replace(s, string(ch), " ", -1)
		}
	}
	for str.Contains(s, "  ") {
		s = str.Replace(s, "  ", " ", -1)
	}
	return s
} //                                                               CompactSpaces

// ContainsI returns true if 's' contains 'substr', ignoring case.
// It always returns true if 'substr' is a blank string.
func ContainsI(s, substr string) bool {
	s, substr = str.ToLower(s), str.ToLower(substr)
	return str.Contains(s, substr)
} //                                                                   ContainsI

// ContainsWord returns true if 's' contains 'word', provided 'word'
// is a distinct word in 's', that is, the characters before
// and after 'word' are not alphanumeric or underscores.
// It always returns true if 'word' is a blank string.
// Specify MatchCase or IgnoreCase to determine
// if the case should me matched or ignored.
func ContainsWord(s, word string, caseMode CaseMode) bool {
	return WordIndex(s, word, caseMode) != -1
} //                                                                ContainsWord

// CountCRLF returns the number of carriage returns and line feeds in
// the given string in 3 components: CR+LF count, CR count, LF count.
func CountCRLF(s string) (count, countCR, countLF int) {
	for _, ch := range s {
		if ch == '\r' {
			countCR++
		} else if ch == '\n' {
			countLF++
		}
	}
	return countCR + countLF, countCR, countLF
} //                                                                   CountCRLF

// EqualStringSlices reports whether two
// string slices a and b are identical.
func EqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, l := 0, len(a); i < l; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
} //                                                           EqualStringSlices

// FindChar __
func FindChar(s string, ch byte, beginIndex int) int {
	return beginIndex + str.IndexByte(s[beginIndex:], ch)
} //                                                                    FindChar

// FindInSlice __
func FindInSlice(s string, start, end int, substr string) int {
	if start == 0 && end == -1 {
		return str.Index(s, substr)
	}
	var sLen = len(s)
	if start < 0 || start > sLen {
		mod.Error(EInvalidArg, "start index:", start)
		if start < 0 {
			start = 0
		} else {
			return -1
		}
	}
	if end == -1 || end > sLen {
		end = sLen
	}
	if start > end {
		mod.Error("Start index", start, "> end index", end)
		return -1
	}
	return start + str.Index(s[start:end], substr)
} //                                                                 FindInSlice

// First __
func First(s string, count int) string {
	if count < 0 {
		mod.Error("Negative count:", count)
		return ""
	}
	var sLen = len(s)
	if count > sLen {
		mod.Error(EInvalidArg, "count", count, "out of range", sLen)
		count = sLen
	}
	return s[:count]
} //                                                                       First

// GetPart returns the substring between 'prefix' and 'suffix'.
// When the prefix is blank, returns the part from the beginning of 's'.
// When the suffix is blank, returns the part up to the end of 's'.
// I.e. if prefix and suffix are both blank, returns 's'.
// When either prefix or suffix is not found, returns a zero-length string.
func GetPart(s, prefix, suffix string) string {
	var at = str.Index(s, prefix)
	if at == -1 {
		return ""
	}
	s = s[at+len(prefix):]
	if suffix == "" {
		return s
	}
	at = str.Index(s, suffix)
	if at == -1 {
		return ""
	}
	return s[:at]
} //                                                                     GetPart

// IfString is a conditional IF expression: If 'condition' is true,
// returns string 'trueStr', otherwise returns string 'falseStr'.
// Not to be confused with IfStr() function, which uses Str objects.
func IfString(condition bool, trueStr, falseStr string) string {
	if condition {
		return trueStr
	}
	return falseStr
} //                                                                    IfString

// IsIdentifier checks if 's' contains only letters, numbers or underscores.
func IsIdentifier(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch != '_' && !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
} //                                                                IsIdentifier

// IsWhiteSpace returns true if all the
// characters in a string are white-spaces.
func IsWhiteSpace(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch != ' ' && ch != '\a' &&
			ch != '\b' && ch != '\f' &&
			ch != '\n' && ch != '\r' &&
			ch != '\t' && ch != '\v' {
			return false
		}
	}
	return true
} //                                                                IsWhiteSpace

// JSUnescape unescapes a string escaped with the JavaScript
// escape() function (which is deprecated).
// In such escaped strings:
// % is followed by a 2-digit hex value.
// e.g. escaped string "%25" becomes "%", "%20" becomes " ", etc.
func JSUnescape(s string) string {
	var retBuf = bytes.NewBuffer(make([]byte, 0, len(s)))
	var wr = retBuf.WriteRune
	var hexPower = 0
	var hexVal = rune(0)
	for _, ch := range s {
		if ch == '%' {
			hexPower = 2
			hexVal = 0
		} else if hexPower != 0 {
			var digit rune
			if ch >= '0' && ch <= '9' {
				digit = ch - '0'
			} else if ch >= 'a' && ch <= 'f' {
				digit = ch - 'a' + 10
			} else if ch >= 'A' && ch <= 'F' {
				digit = ch - 'A' + 10
			}
			if hexPower == 2 {
				digit *= 16
			}
			hexVal += digit
			if hexPower == 1 {
				wr(hexVal)
			}
			hexPower--
		} else {
			wr(ch)
		}
	}
	return retBuf.String()
} //                                                                  JSUnescape

// JSUnescapeStruct unescapes all the strings in a struct that have been
// escaped with the JavaScript escape() function (which is deprecated).
// In such escaped strings: % is followed by a 2-digit
// hex value. e.g. escaped string "%25" becomes "%", "%20" becomes " ", etc.
func JSUnescapeStruct(structPtr interface{}) {
	var structT = reflect.TypeOf(structPtr)
	if structT.Kind() != reflect.Ptr {
		mod.Error(EInvalidArg, "^structPtr", "is not a pointer;",
			"it is^", structT.Kind())
		return
	}
	if structT.Elem().Kind() != reflect.Struct {
		mod.Error(EInvalidArg, "^structPtr", "is not a pointer to a struct;",
			"it is^", structT.Elem().Kind())
		return
	}
	structT = structT.Elem()
	var structV = reflect.ValueOf(structPtr).Elem()
	for i := 0; i < structT.NumField(); i++ {
		var fieldV = structV.Field(i)
		var fieldK = fieldV.Kind()
		if fieldK == reflect.String {
			fieldV.SetString(JSUnescape(str.Trim(fieldV.String(), SPACES)))
		} else if fieldK == reflect.Slice {
			for rowNo := 0; rowNo < fieldV.Len(); rowNo++ {
				JSUnescapeStruct(fieldV.Index(rowNo).Addr().Interface())
			}
		}
	}
} //                                                            JSUnescapeStruct

// Last __
func Last(s string, count int) string {
	if count < 0 {
		mod.Error("Negative count:", count)
		return ""
	}
	var sLen = len(s)
	if count > sLen {
		mod.Error(EInvalidArg, "count", count, "out of range", sLen)
		count = sLen
	}
	return s[sLen-count : sLen]
} //                                                                        Last

// LineBeginIndex __
func LineBeginIndex(s string, index int) int {
	var sLen = len(s)
	if index == -1 {
		index = sLen
	}
	if index < -1 || index > sLen {
		mod.Error("Index", index, "out of range -1 to", sLen)
		if index < -1 {
			index = 0
		} else if index > sLen {
			index = sLen
		}
	}
	var cr = str.LastIndexByte(s[:index], '\r')
	var lf = str.LastIndexByte(s[:index], '\n')
	if cr > lf {
		return cr + 1
	}
	return lf + 1
} //                                                              LineBeginIndex

// LineBeginIndexB __
func LineBeginIndexB(s []byte, index int) int {
	var sLen = len(s)
	if index == -1 {
		index = sLen
	}
	if index < -1 || index > sLen {
		mod.Error("Index", index, "out of range -1 to ", sLen)
		if index < -1 {
			index = 0
		} else if index > sLen {
			index = sLen
		}
	}
	var cr = bytes.LastIndexByte(s[:index], '\r')
	var lf = bytes.LastIndexByte(s[:index], '\n')
	if cr > lf {
		return cr + 1
	}
	return lf + 1
} //                                                             LineBeginIndexB

// LineEndIndex __
func LineEndIndex(s string, index int) int {
	var sLen = len(s)
	if index == -1 {
		return sLen
	} else if index < -1 || index > sLen {
		mod.Error("Index", index, "out of range -1 to", sLen)
		if index < -1 {
			index = 0
		} else if index > sLen {
			index = sLen
		}
	}
	var cr = str.IndexByte(s[index:], '\r')
	var lf = str.IndexByte(s[index:], '\n')
	var i int
	if (cr < lf && cr != -1) || lf == -1 {
		i = cr
	} else {
		i = lf
	}
	if i == -1 {
		return sLen
	}
	return index + i
} //                                                                LineEndIndex

// LineEndIndexB __
func LineEndIndexB(s []byte, index int) int {
	var sLen = len(s)
	if index == -1 {
		index = 0
	} else if index < -1 || index > sLen {
		mod.Error("Index", index, "out of range -1 to", sLen)
		if index < -1 {
			index = 0
		} else if index > sLen {
			index = sLen
		}
	}
	var cr = bytes.IndexByte(s[index:], '\r')
	var lf = bytes.IndexByte(s[index:], '\n')
	var i = lf
	if cr != -1 && cr < lf {
		i = cr
	}
	if i == -1 {
		return sLen
	}
	return index + i
} //                                                               LineEndIndexB

// LineOfIndex __
func LineOfIndex(s string, index int) string {
	if index < 0 || index > len(s) {
		return ""
	}
	var begin = LineBeginIndex(s, index)
	var end = LineEndIndex(s, index)
	return s[begin:end]
} //                                                                 LineOfIndex

// LineOffsetUTF8 returns the byte and character offsets of
// the beginning of a line specified by 'lineIndex', within
// a UTF8-encoded slice of bytes ('data').
func LineOffsetUTF8(data []byte, lineIndex int) (byteOffset, charOffset int) {
	if lineIndex < 0 {
		mod.Error(EInvalidArg, "^lineIndex", ":", lineIndex)
		return -1, -1
	}
	if lineIndex == 0 {
		return 0, 0
	}
	var end = len(data) - 1
	var cc = 0 // number of utf8 continuation characters
	for i, b := range data {
		byteOffset++
		if cc > 0 {
			if (b & 0xC0) != 0x80 { // (b & 11 000000) == 10 000000
				Error(fmt.Sprintf(
					"Bad UTF8 continuation: %dd 0x%2Xh %8bb", b, b, b,
				))
			}
			cc--
			continue
		}
		// From RFC 2044:
		//
		// UCS-4 range (hex.)   UTF-8 octet sequence (binary)
		// 0000 0000-0000 007F  0xxxxxxx
		// 0000 0080-0000 07FF  110xxxxx 10xxxxxx
		// 0000 0800-0000 FFFF  1110xxxx 10xxxxxx 10xxxxxx
		//
		// 0001 0000-001F FFFF  11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
		// 0020 0000-03FF FFFF  111110xx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
		//
		// 0400 0000-7FFF FFFF  1111110x 10xxxxxx ... 10xxxxxx
		switch {
		case b < 128:
			charOffset++
		case (b & 0xE0) == 0xC0: // (b & 111 00000) == 110 00000
			charOffset++
			cc = 2 - 1
			continue
		case (b & 0xF0) == 0xE0: // (b & 1111 0000) == 1110 0000
			charOffset++
			cc = 3 - 1
			continue
		case (b & 0xF8) == 0xF0: // (b & 11111 000) == 11110 000
			charOffset++
			cc = 4 - 1
			continue
		case (b & 0xFC) == 0xF8: // (b & 111111 00) == 111110 00
			charOffset++
			cc = 5 - 1
			continue
		default:
			Error(fmt.Sprintf(
				"UTF8 lead byte not handled: %dd 0x%2Xh %8bb", b, b, b,
			))
		}
		if b == '\n' || (b == '\r' && (i == end || data[i+1] != '\n')) {
			//           ^handle the strange case of \r-delimited lines^
			lineIndex--
			if lineIndex <= 0 {
				return byteOffset, charOffset
			}
		}
	}
	return -1, -1
} //                                                              LineOffsetUTF8

// Padf suffixes a string with spaces to make sure it is at least
// 'minLength' characters wide. I.e. the string is left-aligned.
// If the string is wider than 'minLength', returns the string as it is.
func Padf(minLength int, s string, args ...interface{}) string {
	s = fmt.Sprintf(s, args...)
	if len(s) < minLength {
		return s + str.Repeat(" ", minLength-len(s))
	}
	return s
} //                                                                        Padf

// ReplaceEx1 __
// Specify MatchCase or IgnoreCase for case mode.
func ReplaceEx1(s, find, repl string, count int, caseMode CaseMode) string {
	if find == repl || count == 0 {
		return s // avoid allocation
	}
	var sLen = len(s)
	var findLen = len(find)
	//
	// pre-allocate a buffer (assume each char uses 2 bytes on average)
	var retBuf = bytes.NewBuffer(make([]byte, 0,
		2*int(float64(sLen)/float64(findLen)*float64(len(repl)+1))))
	var ws = retBuf.WriteString
	//
	// lowercase of 's' and 'find' used when caseMode is IgnoreCase
	var sLw, findLw string
	if caseMode == IgnoreCase {
		sLw = str.ToLower(s)
		findLw = str.ToLower(find)
	}
	var replRemain = count // number of replacements remaining
	var pos, prev = 0, 0
	for pos < sLen {
		// find the next index of 'find' in 's'
		var i int
		if caseMode == IgnoreCase {
			i = str.Index(sLw[pos:], findLw)
		} else {
			i = str.Index(s[pos:], find)
		}
		// no more matches? append the rest
		if i == -1 || replRemain == 0 {
			ws(s[prev:])
			break
		}
		pos += i // add i to pos because i is relative to slice
		// text between matches
		if pos > 0 {
			ws(s[prev:pos])
		}
		ws(repl)
		replRemain--
		prev = pos + findLen
		pos += findLen
	}
	return retBuf.String()
	//TODO: Use an array instead of bytes.Buffer. See Replace() in library.
} //                                                                  ReplaceEx1

// ReplaceI replaces 'find' with 'repl' ignoring case.
func ReplaceI(s, find, repl string, optCount ...int) string {
	var count = -1
	switch n := len(optCount); {
	case n == 1:
		count = optCount[0]
	case n > 1:
		mod.Error(EInvalidArg, "optCount", ":", optCount)
	}
	return ReplaceEx1(s, find, repl, count, IgnoreCase)
} //                                                                    ReplaceI

// ReplaceMany makes multiple string replacements in a single pass.
// Replacing is prioritized, with longest strings among 'finds'
// replaced first before shorter strings.
//
// With this function you can exchange or transpose values, which simple
// replace() functions can not do without making temporary replacements.
//
// It is much faster to call ReplaceMany() with a batch of different
// replacements (even thousands of items), than to call a simple replace()
// function repeatedly. However if you only need to replace one string,
// then standard strings.Replace() or zr.ReplaceWord() and similar
// functions run faster when making an isolated replacement.
//
// Internally, ReplaceMany() works with Unicode characters, therefore
// 's', 'finds' and 'repls' can be ANSI or UTF-8 encoded strings.
//
// Parameters:
//
// s: the string being replaced.
//
// finds: a list of strings to find. each item will be replaced
//        with the matching item in repls, so the length of
//        'finds' and 'repls' must be the same.
//
// repls: a list of replacement strings.
//
// count: maximum number of replacements to make, or -1 for unlimited.
//
// caseMode: use MatchCase for case-sensitive search string matching
//           or IgnoreCase for case-insensitive search string matching.
//
// wordMode: use MatchWord to match whole words only, or IgnoreWord to
//           replace everything. Distinct words are comprised of letters,
//           numbers and underscores.
func ReplaceMany(
	s string,
	finds []string,
	repls []string,
	count int,
	caseMode CaseMode,
	wordMode WordMode,
) string {
	// Note: this large function is split into sub-functions.
	// Processing begins after everything is declared. See run().

	if count == 0 || len(finds) == 0 {
		return s // avoid allocation
	}

	// batch __
	type batch struct {
		findLen int
		finds   []string
		repls   []string
	}

	// tree __
	type tree struct {
		find string         // what to find
		repl string         // what to replace with
		done int            // number of replacements made so far
		sub  map[rune]*tree // a 'branch' of the tree
	}

	// getBatches __
	var getBatches = func(finds, repls []string) ([]int, map[int]*batch) {
		var lengths []int
		var batches = map[int]*batch{}
		for i, find := range finds {
			var n = utf8.RuneCountInString(find)
			var b, has = batches[n]
			if !has {
				batches[n] = &batch{findLen: n}
				lengths = append(lengths, n)
			}
			b, _ = batches[n]
			if find == repls[i] {
				continue
			}
			b.finds = append(b.finds, find)
			b.repls = append(b.repls, repls[i])
		}
		sort.Sort(sort.Reverse(sort.IntSlice(lengths)))
		return lengths, batches
	}

	// getTree __
	var getTree = func(finds, repls []string, caseMode CaseMode,
	) (ret tree) {
		ret.sub = make(map[rune]*tree)
		for f, find := range finds {
			var i = 0
			var node = &ret
			var last = utf8.RuneCountInString(find)
			for _, ch := range find {
				i++ // ^ don't use the index in for, that's a byte index
				if caseMode == IgnoreCase {
					ch = unicode.ToLower(ch)
				}
				var _, exist = node.sub[ch]
				var sub *tree
				if exist {
					sub = node.sub[ch]
				} else {
					sub = &tree{sub: make(map[rune]*tree)}
				}
				if i == last {
					sub.find = find
					sub.repl = repls[f]
				}
				node.sub[ch] = sub
				node = sub
			}
		}
		return ret
	}

	// replaceMany __
	var replaceMany = func(
		s string,
		finds []string,
		repls []string,
		count int,
		caseMode CaseMode,
		wordMode WordMode,
	) string {
		var src = []rune(s)
		var srcLen = len(src)
		var root = getTree(finds, repls, caseMode)
		var node = &root // *tree pointing to current branch
		var match = 0    // <- number of matching characters
		var prev = 0
		var retBuf = bytes.NewBuffer(make([]byte, 0, int(1.5*float64(len(s)))))
		var ws = retBuf.WriteString
		for i := 0; i < srcLen; i++ {
			var ch = src[i]
			if caseMode == IgnoreCase {
				ch = unicode.ToLower(ch)
			}
			// check if the character is found under the current branch
			// if not, reset matching count and start over from root
			{
				var sub, found = node.sub[ch]
				if !found {
					node = &root
					i -= match
					match = 0
					continue
				}
				match++
				node = sub
			}
			var findLen = utf8.RuneCountInString(node.find)
			if findLen == 0 || findLen != match {
				continue
			}
			if wordMode == MatchWord {
				var l, r rune
				if (i - findLen + 1) > 0 {
					l = src[i-findLen]
				}
				if i < (srcLen - 1) {
					r = src[i+1]
				}
				if l == '_' || unicode.IsDigit(l) || unicode.IsLetter(l) ||
					r == '_' || unicode.IsDigit(r) || unicode.IsLetter(r) {
					continue
				}
			}
			if count >= 0 && node.done >= count {
				break
			}
			node.done++
			if prev < i-findLen+1 {
				ws(string(src[prev : i-findLen+1]))
			}
			ws(node.repl)
			prev = i + 1
			node = &root
			match = 0
		}
		// write the 'tail' of the string
		if prev < srcLen {
			ws(string(src[prev:]))
		}
		return retBuf.String()
		//TODO: Use an array instead of bytes.Buffer. See Replace() in library.
	}

	var run = func() string {
		// validate arguments
		const erv = ""
		if len(finds) != len(repls) {
			mod.Error(EInvalidArg, ": lengths don't match:",
				len(finds), "and", len(repls))
			return erv
		}
		if caseMode != IgnoreCase && caseMode != MatchCase {
			mod.Error(EInvalidArg,
				"^caseMode", ":", caseMode, "defaulting to 'MatchCase'")
			caseMode = MatchCase
		}
		if wordMode != IgnoreWord && wordMode != MatchWord {
			mod.Error(EInvalidArg,
				"^wordMode", ":", wordMode, "defaulting to 'IgnoreWord'")
			wordMode = IgnoreWord
		}
		var lengths, batches = getBatches(finds, repls)
		for _, n := range lengths {
			var b, _ = batches[n]
			s = replaceMany(s, b.finds, b.repls, count, caseMode, wordMode)
		}
		return s
	}
	return run()
} //                                                                 ReplaceMany

// ReplaceWord replaces a word in a string. __
// Specify MatchCase or IgnoreCase for case mode.
//
// Examples:
//     ReplaceWord("ab b c", "b", "z", MatchCase)   // returns "ab z c"
//     ReplaceWord("ab b c", "B", "Z", IgnoreCase)  // returns "ab Z c"
func ReplaceWord(s, find, repl string, caseMode CaseMode) string {
	if find == repl {
		return s // avoid allocation
	}
	var sLen = len(s)
	var findLen = len(find)
	var pos = 0
	var prev = 0
	var ret = ""
	var sLw = "" // lowercase of 's' and 'find' when caseMode is IgnoreCase
	var findLw = ""
	if caseMode == IgnoreCase {
		sLw = str.ToLower(s)
		findLw = str.ToLower(find)
	}
	var nonWord = func(ch byte) bool {
		//TODO: wrongly returns 'false' for non-Latin Unicode letters
		return !((ch >= '0' && ch <= '9') || ch == '_' ||
			(ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z'))
	}
	for {
		{ // find the next index of 'find' in 's'
			var i int
			if caseMode == IgnoreCase {
				i = str.Index(sLw[pos:], findLw)
			} else {
				i = str.Index(s[pos:], find)
			}
			if i == -1 {
				ret += s[prev:] // no more matches? append the rest
				break
			}
			pos += i
		}
		ret += s[prev:pos] // text between words
		// left word boundary:
		var c = byte(0)
		if pos > 0 {
			c = s[pos-1]
		}
		var onLeft = nonWord(c)
		// right word boundary:
		c = 0
		if pos+findLen < sLen {
			c = s[pos+findLen]
		}
		var onRight = nonWord(c)
		// an isolated word?
		if onLeft && onRight {
			ret += repl // append replacement
		} else {
			ret += find
		}
		prev = pos + findLen
		pos += findLen
	}
	return ret
} //                                                                 ReplaceWord

// SetPart __ //TODO: describe and create unit test
func SetPart(s, prefix, suffix, part string) string {
	var at = str.Index(s, prefix)
	if at == -1 {
		return s + prefix + part + suffix
	}
	var head = s[:at+len(prefix)]
	var tail = s[at+len(prefix):]
	at = str.Index(tail, suffix)
	if at == -1 {
		tail = suffix
	} else {
		tail = tail[at:]
	}
	return head + part + tail
} //                                                                     SetPart

// SetSlice __
func SetSlice(s string, start, end int, substr string) string {
	if start == 0 && end == -1 {
		return substr
	}
	var sLen = len(s)
	if start < 0 || start > sLen {
		mod.Error(EInvalidArg, "start index:", start)
		if start < 0 {
			start = 0
		} else {
			return s
		}
	}
	if end == -1 || end > sLen {
		end = sLen
	}
	if start > end {
		mod.Error("Start index", start, "> end index", end)
		return s
	}
	return s[:start] + substr + s[end:]
} //                                                                    SetSlice

// ShowSpaces replaces spaces, tabs and
// line breaks with visible placeholders.
func ShowSpaces(s string) string {
	if str.Contains(s, " ") {
		s = str.Replace(s, " ", "-", -1)
	}
	if str.Contains(s, "\t") {
		s = str.Replace(s, "\t", "--->", -1)
	}
	if str.Contains(s, CR+LF) {
		s = str.Replace(s, CR+LF, "CRLF\x00", -1)
	}
	if str.Contains(s, CR) {
		s = str.Replace(s, CR, "CR"+CR, -1)
	}
	if str.Contains(s, LF) {
		s = str.Replace(s, LF, "LF"+LF, -1)
	}
	if str.Contains(s, "CRLF") {
		s = str.Replace(s, "CRLF\x00", "CRLF"+CR+LF, -1)
	}
	return s
} //                                                                  ShowSpaces

// SkipChars __
func SkipChars(s string, start int, chars string) int {
	var n = len(s)
	if start >= n {
		return n
	}
	if start < 0 {
		start = 0
	}
	for i, a := range s[start:] {
		var has bool
		for _, b := range chars {
			if a == b {
				has = true
				break
			}
		}
		if !has {
			return start + i
		}
	}
	return n
} //                                                                   SkipChars

// SkipName __
func SkipName(s string, start int) int {
	var n = len(s)
	if start >= n {
		return start
	}
	var isPrefix = true
	for i, a := range s[start:] {
		if a != '_' && !unicode.IsLetter(a) &&
			(isPrefix || (!isPrefix && !unicode.IsNumber(a))) {
			return start + i
		}
		if isPrefix {
			isPrefix = false
		}
	}
	return n
} //                                                                   SkipName

// SkipSpaces __
func SkipSpaces(s string, start int) int {
	return SkipChars(s, start, SPACES)
} //                                                                  SkipSpaces

// Slice returns a substring starting from 'beginIndex'
// and ending just before, but not including 'endIndex'.
//
// beginIndex is the starting position. Must not be less than zero
// in which the function treats it as zero and logs an error.
//
// endIndex is the ending position, just after the last required element.
// If endIndex is -1, returns everything from beginIndex up to the end.
func Slice(s string, beginIndex, endIndex int) string {
	if beginIndex < 0 {
		mod.Error("Treating beginIndex", beginIndex, "as 0")
		beginIndex = 0
	}
	if endIndex < 0 || endIndex > len(s) {
		// -1 is acceptable for endIndex
		if endIndex != -1 {
			mod.Error("Treating endIndex", endIndex, "as len(s):", len(s))
		}
		endIndex = len(s)
	}
	return s[beginIndex:endIndex]
} //                                                                       Slice

// SplitQuoted splits s into an array of strings, using spaces as
// delimiters, but leaving the spaces within quoted substrings.
// E.g. given "'a b c' 123", returns ["A B C", 123].
// Works with single quotes ('), double quotes (") and back quotes (`).
// A substring that is started with one type of quote can contain
// quotes of other types, which are treated as normal characters
// until the substring is closed.
func SplitQuoted(s string) []string {
	const BQ = '`'
	const DQ = '"'
	const SQ = '\''
	//
	var part string
	var qmode = rune(0) // quote mode
	var next = false
	var ret []string
	for _, ch := range s + " " {
		if ch == BQ && (qmode == BQ || qmode == 0) {
			// toggle back-quote mode
			if qmode == BQ {
				qmode = 0
				next = true
			} else {
				qmode = BQ
			}
		} else if ch == DQ && (qmode == DQ || qmode == 0) {
			// toggle double-quote mode
			if qmode == DQ {
				qmode = 0
				next = true
			} else {
				qmode = DQ
			}
		} else if ch == SQ && (qmode == SQ || qmode == 0) {
			// toggle single-quote mode
			if qmode == SQ {
				qmode = 0
				next = true
			} else {
				qmode = SQ
			}
		} else if qmode == 0 && unicode.IsSpace(ch) {
			next = true
		} else {
			part += string(ch)
		}
		if next {
			next = false
			if part != "" {
				ret = append(ret, part)
				part = ""
			}
		}
	}
	return ret
} //                                                                 SplitQuoted

// StrOneOf returns true if the first string 's' matches any one
// of the strings that follow, or false otherwise. For example:
// StrOneOf("b", "a", "b", "c") returns true
// StrOneOf("d", "a", "b", "c") returns false
func StrOneOf(s string, matches ...string) bool {
	for _, match := range matches {
		if s == match {
			return true
		}
	}
	return false
} //                                                                    StrOneOf

// String converts 'val' to a string.
func String(val interface{}) string {
	switch val := val.(type) {
	case bool:
		if val {
			return "true"
		}
		return "false"
	case float32, float64:
		var ret = fmt.Sprintf("%f", val)
		if str.Contains(ret, ".") {
			ret = str.TrimRight(ret, "0")
			ret = str.TrimRight(ret, ".")
		}
		return ret
	case int:
		return strconv.Itoa(val)
	case int16,
		int32,
		int64,
		int8:
		return fmt.Sprintf("%d", val)
	case string:
		return val
	case uint,
		uint16,
		uint32,
		uint64,
		uint8:
		return fmt.Sprintf("%d", val)
	case fmt.Stringer:
		return val.String()
	case nil:
		return ""
	// pointer types:
	case *bool:
		if val != nil {
			return String(*val)
		}
	case *float32:
		if val != nil {
			return String(*val)
		}
	case *float64:
		if val != nil {
			return String(*val)
		}
	case *int:
		if val != nil {
			return String(*val)
		}
	case *int16:
		if val != nil {
			return String(*val)
		}
	case *int32:
		if val != nil {
			return String(*val)
		}
	case *int64:
		if val != nil {
			return String(*val)
		}
	case *int8:
		if val != nil {
			return String(*val)
		}
	case *string:
		if val != nil {
			return String(*val)
		}
	case *uint:
		if val != nil {
			return String(*val)
		}
	case *uint16:
		if val != nil {
			return String(*val)
		}
	case *uint32:
		if val != nil {
			return String(*val)
		}
	case *uint64:
		if val != nil {
			return String(*val)
		}
	case *uint8:
		if val != nil {
			return String(*val)
		}
	}
	mod.Error("Type", reflect.TypeOf(val), "not handled; =", val)
	return ""
} //                                                                      String

// Substr returns a substring given a string and the index (charIndex) and
// substring length (charCount). If the length is -1, returns a string
// of all characters starting from the given index.
func Substr(s string, charIndex, charCount int) string {
	var sLen = len(s)
	if charIndex == 0 && sLen == 0 && charCount == 0 {
		return ""
	}
	if charIndex < -1 || charIndex > sLen-1 {
		mod.Error("charIndex", charIndex, "< 0 || > sLen-1:", sLen-1)
		charIndex = 0
	}
	if charCount < -1 {
		mod.Error("charCount", charCount, "< -1")
		charCount = -1
	} else if charIndex+charCount > sLen {
		mod.Error("charIndex", charIndex, "+ charCount", charCount,
			"=", charIndex+charCount, "> sLen", sLen)
		charCount = sLen - charIndex
	}
	var retBuf bytes.Buffer // prepare output buffer with some space reserved
	var wr = retBuf.WriteRune
	if charCount <= -1 {
		retBuf.Grow(sLen)
	} else {
		retBuf.Grow(charCount * 3)
	}
	var i = 0
	for _, ch := range s {
		if i >= charIndex {
			if charCount == -1 {
				wr(ch)
			} else if charCount > 0 {
				charCount--
				wr(ch)
			} else {
				break
			}
		}
		i++
	}
	return retBuf.String()
} //                                                                      Substr

// TitleCase changes a string to title case.
func TitleCase(s string) string {
	var retBuf = bytes.NewBuffer(make([]byte, 0, len(s)))
	var wr = retBuf.WriteRune
	var up = true
	for _, ch := range s {
		if up {
			ch = unicode.ToUpper(ch)
		} else {
			ch = unicode.ToLower(ch)
		}
		up = !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_'
		wr(ch)
	}
	return retBuf.String()
} //                                                                   TitleCase

// TokenGet __ //TODO: rename to GetToken, GetTokenEx. create SetToken
func TokenGet(list string, index int, sep string) string {
	return TokenGetEx(list, index, sep, false)
} //                                                                    TokenGet

// TokenGetEx __
func TokenGetEx(list string, index int, sep string, ignoreEnd bool) string {
	var listLen = len(list)
	var sepLen = len(sep)
	if sepLen >= listLen {
		return ""
	}
	if sepLen == 0 {
		if index == 0 {
			return list
		}
		return ""
	}
	var curr = 0
	var next = str.Index(list, sep)
	var i = index
	for i > 0 {
		i--
		if next == -1 {
			curr = -1
			break
		}
		curr = next + sepLen
		next = str.Index(list[curr:], sep)
		if next != -1 {
			next += curr
		}
	}
	if curr == -1 {
		return ""
	}
	if next == -1 || next < curr || ignoreEnd {
		return list[curr:]
	}
	return list[curr:next]
} //                                                                  TokenGetEx

// WordIndex returns the index of the first instance of word in s,
// or -1 if word is not present in s.
// Specify MatchCase or IgnoreCase for case mode.
func WordIndex(s, word string, caseMode CaseMode) int {
	var wordLen = len(word)
	if wordLen == 0 {
		return 0
	}
	// convert strings to lower case when case-insensitive
	if caseMode == IgnoreCase {
		s = str.ToLower(s)
		word = str.ToLower(word)
	}
	var i = 0
	for {
		{ // find the next index of 'word' in 's'
			var at = str.Index(s[i:], word)
			if at == -1 {
				break // return -1
			}
			i += at
		}
		// read the runes to the left and right of word
		var lr, _ = utf8.DecodeLastRune([]byte(s[:i]))
		var rr, _ = utf8.DecodeRune([]byte(s[i+wordLen:]))
		if lr != '_' && !unicode.IsLetter(lr) && !unicode.IsNumber(lr) &&
			rr != '_' && !unicode.IsLetter(rr) && !unicode.IsNumber(rr) {
			return i
		}
		i += wordLen
	}
	return -1
} //                                                                   WordIndex

//TODO: IMPORTANT: check FindChar() usage

//end
