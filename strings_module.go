// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-01-31 22:50:59 077907                         [zr/strings_module.go]
// -----------------------------------------------------------------------------

package zr

import _strings "strings" // standard

var str StringsProvider = StringsModule{}

// StringsProvider provides all the functions in the standard
// "strings" package, and in addition provides methods for
// stubbing and logging these functions.
type StringsProvider interface {
	Contains(s, substr string) bool
	ContainsAny(s, chars string) bool
	Count(s, substr string) int
	Fields(s string) []string
	HasPrefix(s, prefix string) bool
	HasSuffix(s, suffix string) bool
	Index(s, substr string) int
	IndexByte(s string, c byte) int
	Join(a []string, sep string) string
	LastIndex(s, substr string) int
	LastIndexByte(s string, c byte) int
	Repeat(s string, count int) string
	Replace(s, old, new string, n int) string
	Split(s, sep string) []string
	ToLower(s string) string
	Trim(s, trim string) string
	TrimRight(s, trim string) string
	ToUpper(s string) string
} //                                                             StringsProvider

// StringsModule provides all the functions in the standard
// "strings" package, and in addition provides methods for
// stubbing and logging these functions.
type StringsModule struct{}

// Contains is the same as strings.Contains
func (ob StringsModule) Contains(s, substr string) bool {
	return _strings.Contains(s, substr)
} //                                                                    Contains

// ContainsAny is the same as strings.ContainsAny
func (ob StringsModule) ContainsAny(s, chars string) bool {
	return _strings.ContainsAny(s, chars)
} //                                                                 ContainsAny

// Count is the same as strings.Count
func (ob StringsModule) Count(s, substr string) int {
	return _strings.Count(s, substr)
} //                                                                       Count

// Fields is the same as strings.Fields
func (ob StringsModule) Fields(s string) []string {
	return _strings.Fields(s)
} //                                                                      Fields

// HasPrefix is the same as strings.HasPrefix
func (ob StringsModule) HasPrefix(s, prefix string) bool {
	return _strings.HasPrefix(s, prefix)
} //                                                                   HasPrefix

// HasSuffix is the same as strings.HasSuffix
func (ob StringsModule) HasSuffix(s, suffix string) bool {
	return _strings.HasSuffix(s, suffix)
} //                                                                   HasSuffix

// Index is the same as strings.Index
func (ob StringsModule) Index(s, substr string) int {
	return _strings.Index(s, substr)
} //                                                                       Index

// IndexByte is the same as strings.IndexByte
func (ob StringsModule) IndexByte(s string, c byte) int {
	return _strings.IndexByte(s, c)
} //                                                                   IndexByte

// Join is the same as strings.Join
func (ob StringsModule) Join(a []string, sep string) string {
	return _strings.Join(a, sep)
} //                                                                        Join

// LastIndex is the same as strings.LastIndex
func (ob StringsModule) LastIndex(s, substr string) int {
	return _strings.LastIndex(s, substr)
} //                                                                   LastIndex

// LastIndexByte is the same as strings.LastIndexByte
func (ob StringsModule) LastIndexByte(s string, c byte) int {
	return _strings.LastIndexByte(s, c)
} //                                                               LastIndexByte

// Repeat is the same as strings.Repeat
func (ob StringsModule) Repeat(s string, count int) string {
	return _strings.Repeat(s, count)
} //                                                                      Repeat

// Replace is the same as strings.Replace
func (ob StringsModule) Replace(s, old, new string, n int) string {
	return _strings.Replace(s, old, new, n)
} //                                                                     Replace

// Split is the same as strings.Split
func (ob StringsModule) Split(s, sep string) []string {
	return _strings.Split(s, sep)
} //                                                                       Split

// ToLower is the same as strings.ToLower
func (ob StringsModule) ToLower(s string) string {
	return _strings.ToLower(s)
} //                                                                     ToLower

// Trim is the same as strings.Trim
func (ob StringsModule) Trim(s, trim string) string {
	return _strings.Trim(s, trim)
} //                                                                        Trim

// TrimRight is the same as strings.TrimRight
func (ob StringsModule) TrimRight(s, trim string) string {
	return _strings.TrimRight(s, trim)
} //                                                                   TrimRight

// ToUpper is the same as strings.ToUpper
func (ob StringsModule) ToUpper(s string) string {
	return _strings.ToUpper(s)
} //                                                                     ToUpper

//end
