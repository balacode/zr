// -----------------------------------------------------------------------------
// ZR Library                                                 zr/[timer_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

//  to test all items in timer.go use:
//      go test --run Test_timr_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"strings"
	"testing"
	"time"
)

// go test --run Test_timr_
func Test_timr_(t *testing.T) {
	TBegin(t)
	//
	var tm Timer
	tm.Start("task")
	time.Sleep(300 * time.Millisecond)
	tm.Stop("task")
	//
	lines := strings.Split(tm.String(), "\n")
	TEqual(t, len(lines), 4)
	//
	// Expected output like:
	// 0:     --------------------------------- SECONDS:
	// 1:        0.31356: task
	// 2:        0.31356
	// 3:
	TTrue(t, strings.Contains(lines[0], "---"))
	TTrue(t, strings.Contains(lines[0], "SECONDS:"))
	TTrue(t, strings.Contains(lines[1], "0.3"))
	TTrue(t, strings.Contains(lines[1], ": task"))
	TTrue(t, strings.Contains(lines[2], "0.3"))
	TEqual(t, lines[3], "")
} //                                                                  Test_timr_

//end
