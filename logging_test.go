// -----------------------------------------------------------------------------
// ZR Library                                               zr/[logging_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

//  to test all items in logging.go use:
//      go test --run Test_logg_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"fmt"
	"strings"
	"testing"
)

// go test --run Test_logg_CallerList_
// only use standard testing functions in this unit test (not zr.Test...)
func Test_logg_CallerList_(t *testing.T) {
	// test without deep nesting
	{
		ar := CallerList()
		if len(ar) != 1 {
			t.Error("Expected 1 item in returned slice, but got", len(ar))
		}
		const ExpectName = "zr.Test_logg_CallerList_:"
		if len(ar) == 1 && !strings.HasPrefix(ar[0], ExpectName) {
			t.Errorf(
				"Expected %q in returned slice, but got %q",
				ExpectName,
				ar[0],
			)
		}
	}
	// test call stack with nested functions
	{
		var ar []string
		f1 := func() {
			f2 := func() {
				f3 := func() {
					ar = CallerList()
				}
				f3()
			}
			f2()
		}
		f1()
		// the returned list should contain this function and 3 nested functions
		if len(ar) != 4 {
			t.Error("Expected 4 items in returned slice, but got", len(ar))
		}
		for i, prefix := range []string{
			"Test_logg_CallerList_.func1.1.1",
			"Test_logg_CallerList_.func1.1",
			"Test_logg_CallerList_.func1",
			"zr.Test_logg_CallerList_",
		} {
			if !strings.HasPrefix(ar[i], prefix) {
				t.Errorf(
					"In returned slice at index %d: expected %q but got %q",
					i, prefix, ar[i],
				)
			}
		}
	}
} //                                                       Test_logg_CallerList_

// go test --run Test_logg_NoE_
func Test_logg_NoE_(t *testing.T) {
	const check = "xyz"
	fn := func() (string, error) {
		return "xyz", fmt.Errorf("error 123")
	}
	s := NoE(fn()).(string)
	TTrue(t, s == check)
} //                                                              Test_logg_NoE_

// end
