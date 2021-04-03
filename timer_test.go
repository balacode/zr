// -----------------------------------------------------------------------------
// ZR Library                                                 zr/[timer_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

//  to test all items in timer.go use:
//      go test --run Test_Timer_1_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

const TIMER_VERBOSE = false

// go test --run Test_Timer_1_
func Test_Timer_1_(t *testing.T) {
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
} //                                                               Test_Timer_1_

// go test --run Test_Timer_ReportByTimeSpent_
func Test_Timer_ReportByTimeSpent_(t *testing.T) {
	TBegin(t)
	//
	var tm Timer
	for i, ms := range []int{200, 500, 100, 600, 300} {
		task := tm.Start(fmt.Sprint("task", i+1))
		time.Sleep(time.Duration(ms) * time.Millisecond)
		tm.Stop(task)
	}
	got := tm.ReportByTimeSpent()
	lines := strings.Split(got, "\n")
	TEqual(t, len(lines), 8)
	for i, line := range lines {
		PL(i, line)
	}
	// Expected output like:
	// 0     --------------------------------- SECONDS:
	// 1        0.60583: task4
	// 2        0.50635: task2
	// 3        0.30412: task5
	// 4        0.20028: task1
	// 5        0.10201: task3
	// 6        1.71859
	// 7
	//
	TTrue(t, strings.Contains(lines[0], "---"))
	TTrue(t, strings.Contains(lines[0], "SECONDS:"))
	//
	TTrue(t, strings.Contains(lines[1], "0.6"))
	TTrue(t, strings.Contains(lines[1], ": task4"))
	//
	TTrue(t, strings.Contains(lines[2], "0.5"))
	TTrue(t, strings.Contains(lines[2], ": task2"))
	//
	TTrue(t, strings.Contains(lines[3], "0.3"))
	TTrue(t, strings.Contains(lines[3], ": task5"))
	//
	TTrue(t, strings.Contains(lines[4], "0.2"))
	TTrue(t, strings.Contains(lines[4], ": task1"))
	//
	TTrue(t, strings.Contains(lines[5], "0.1"))
	TTrue(t, strings.Contains(lines[5], ": task3"))
	//
	TTrue(t, strings.Contains(lines[6], "1.7"))
	TEqual(t, lines[7], "")
} //                                               Test_Timer_ReportByTimeSpent_

// go test --run Test_Timer_Rename_
func Test_Timer_Rename_(t *testing.T) {
	TBegin(t)
	//
	var tm Timer
	for _, ms := range []int{100, 200, 300} {
		task := tm.Start(fmt.Sprint("task", ms))
		time.Sleep(time.Duration(ms) * time.Millisecond)
		tm.Stop(task)
	}
	tm.Rename("task100", "task100") // do nothing (no change)
	tm.Rename("task100", "taskA")   // rename
	tm.Rename("taskABC", "taskA")   // do nothing (non-existent task)
	tm.Rename("task200", "taskB")   // rename
	tm.Rename("task300", "taskB")   // merge
	tm.Rename("taskXYZ", "taskB")   // do nothing (non-existent task)
	//
	got := tm.ReportByTimeSpent()
	lines := strings.Split(got, "\n")
	if TIMER_VERBOSE {
		for i, line := range lines {
			PL(i, line)
		}
	}
	// Expected output (with minor timing differences):
	// 0     --------------------------------- SECONDS:
	// 1        0.52151: taskB
	// 2        0.10393: taskA
	// 3        0.62545
	// 4
	//
	TEqual(t, len(lines), 5)
	//
	TTrue(t, strings.Contains(lines[0], "---"))
	TTrue(t, strings.Contains(lines[0], "SECONDS:"))
	//
	TTrue(t, strings.Contains(lines[1], "0.5"))
	TTrue(t, strings.Contains(lines[1], ": taskB"))
	//
	TTrue(t, strings.Contains(lines[2], "0.1"))
	TTrue(t, strings.Contains(lines[2], ": taskA"))
	//
	TTrue(t, strings.Contains(lines[3], "0.6"))
	TEqual(t, lines[4], "")
} //                                                          Test_Timer_Rename_

// end
