// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 724B5F                           zr/[reflect_test.go]
// -----------------------------------------------------------------------------

package zr

/*
to test all items in reflect.go use:
    go test --run Test_refl_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"testing"
)

// go test --run Test_refl_IndexSliceStructValue_
func Test_refl_IndexSliceStructValue_(t *testing.T) {
	var slice = []struct {
		A string
		B string
	}{
		{"a", "b"},
		{"aa", "bb"},
		{"aaa", "bbb"},
	}
	DisableErrors()
	// passing a nil input returns -1 and an error
	{
		var got, err = IndexSliceStructValue(nil, "A", nil)
		if err == nil {
			t.Error("check", 1, "unexpectedly returned nil error.")
		}
		if got != -1 {
			t.Error("check", 1, "returned", got, "instead of -1")
		}
	}
	// passing a zero-length fieldName returns -1 and an error
	{
		var got, err = IndexSliceStructValue(slice, "", nil)
		if err == nil {
			t.Error("check", 2, "unexpectedly returned nil error.")
		}
		if got != -1 {
			t.Error("check", 2, "returned", got, "instead of -1")
		}
	}
	// passing a nil checker() returns -1 and an error
	{
		var got, err = IndexSliceStructValue(slice, "A", nil)
		if err == nil {
			t.Error("check", 3, "unexpectedly returned nil error.")
		}
		if got != -1 {
			t.Error("check", 3, "returned", got, "instead of -1")
		}
	}
	EnableErrors()
	//
	// check if function returns expected results
	var tests = []struct {
		fieldName string
		find      string
		index     int
		callCount int
	}{
		{"A", "a", 0, 1},
		{"B", "b", 0, 1},
		{"A", "aa", 1, 2},
		{"B", "bb", 1, 2},
		{"A", "aaa", 2, 3},
		{"B", "bbb", 2, 3},
		{"A", "aaaa", -1, 3},
		{"B", "bbbb", -1, 3},
	}
	for i, test := range tests {
		var callCount int
		var checker = func(value interface{}) bool {
			callCount++
			var s, _ = value.(string)
			return s == test.find
		}
		var got, err = IndexSliceStructValue(slice, test.fieldName, checker)
		if err != nil {
			t.Error("subtest:", i, "unexpectedly returned error:", err)
		}
		if got != test.index {
			t.Error("subtest:", i, "returned", got, "instead of", test.index)
		}
		if callCount != test.callCount {
			t.Error("subtest:", i,
				"expected", test.callCount, "calls",
				"but made", callCount, "calls",
			)
		}
	}
} //                                            Test_refl_IndexSliceStructValue_

//end
