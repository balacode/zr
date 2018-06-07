// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 4EC42D                                  zr/[timer.go]
// -----------------------------------------------------------------------------

package zr

// # Types
//   Timer struct
//   TimerTask struct
//
// # Methods
//   (ob *Timer) GetTasks() map[string]*TimerTask
//   (ob *Timer) Print()
//   (ob *Timer) Start(taskName string)
//   (ob *Timer) Stop(taskName string)
//   (ob *Timer) String()
//   (ob *Timer) Reset()
//
// # Private Method
//   (ob *Timer) makeTasks()

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

// SUGGESTION: new methods NewTimer() and StopPrint():
// var tm = NewTimer("App.GoNextLine")
// tm.StopPrint()

// -----------------------------------------------------------------------------
// # Types

// Timer gathers timing statistics for multiple tasks,
// for example to collect the total time spent executing functions.
// Timing begins by calling Start() with the name of a timed task,
// and ends by calling Stop(taskName). You can make multiple calls to
// Start() and Stop() for the same task. Timer will accumulate the
// total time spent on each task and the number of times it was executed.
//
// To get a timing report:
// call the Print() method to output the report to the console. You can
// also get the report using the String() method, or call GetTasks() to
// obtain a map of named tasks and their timing statistics.
//
// You can reuse the same Timer by calling Reset() to clear its contents.
//
// The current version of Timer should not be used
// to time the same task running in parallel.
type Timer struct {
	Mutex sync.RWMutex
	Tasks map[string]*TimerTask
} //                                                                       Timer

// TimerTask holds the timing statistics of a timed task.
type TimerTask struct {
	Count     int
	SerialNo  int
	StartTime time.Time
	TotalMs   float32
} //                                                                   TimerTask

// -----------------------------------------------------------------------------
// # Methods

// GetTasks returns a map of named tasks and their timing statistics.
func (ob *Timer) GetTasks() map[string]*TimerTask {
	return ob.Tasks
} //                                                                    GetTasks

// Print prints out a timing report to the console (i.e. standard output)
// Shows the name of each task, the total time spent on the task,
// the number of times the task was executed, and the average running time
// in seconds rounded to 4 decimal places.
func (ob *Timer) Print() {
	if ob == nil {
		Error(ENilReceiver)
		return
	}
	if ob.Tasks == nil {
		ob.makeTasks()
	}
	var s = ob.String()
	fmt.Println(s)
} //                                                                       Print

// Start begins timing the named task. Make sure you call Stop() when
// the task is complete. You can start and stop the same task multiple
// times, provided you call Stop() after every Start().
func (ob *Timer) Start(taskName string) {
	var now = time.Now()
	if ob == nil {
		Error(ENilReceiver)
		return
	}
	ob.Mutex.Lock()
	defer ob.Mutex.Unlock()
	if ob.Tasks == nil {
		ob.makeTasks()
	}
	var task, exists = ob.Tasks[taskName]
	if exists {
		task.StartTime = now
		return
	}
	ob.Tasks[taskName] = &TimerTask{
		SerialNo:  len(ob.Tasks) + 1,
		StartTime: now,
	}
} //                                                                       Start

// Stop stops timing the named task and stores the time spent in the Timer.
func (ob *Timer) Stop(taskName string) {
	var now = time.Now()
	if ob == nil {
		Error(ENilReceiver)
		return
	}
	ob.Mutex.Lock()
	defer ob.Mutex.Unlock()
	if ob.Tasks == nil {
		ob.makeTasks()
	}
	var task, exists = ob.Tasks[taskName]
	if !exists {
		Error("Never started timing^", taskName)
		PL("THERE ARE", len(ob.Tasks), "TASKS")
		return
	}
	var ms = float32(now.Sub(task.StartTime).Nanoseconds()) / 1000000
	task.Count++
	task.StartTime = now
	task.TotalMs += ms
} //                                                                        Stop

// String returns the timing report as a string,
// and implements the Stringer interface.
func (ob *Timer) String() string {
	ob.Mutex.RLock()
	defer ob.Mutex.RUnlock()
	//
	var serialMax = 0
	for _, task := range ob.Tasks {
		if task.SerialNo > serialMax {
			serialMax = task.SerialNo
		}
	}
	var buf bytes.Buffer
	var ws = buf.WriteString
	ws("    --------------------------------- SECONDS:" + LB)
	var sum = float64(0)
	for i := 0; i <= serialMax; i++ {
		for taskName, task := range ob.Tasks {
			if task.SerialNo == i {
				var seconds = float64(task.TotalMs) / float64(1000)
				sum += seconds
				ws(fmt.Sprintf("%14.5f: %s"+LB, seconds, taskName))
			}
		}
	}
	ws(fmt.Sprintf("%14.5f"+LB, sum))
	var ret = buf.String()
	return ret
} //                                                                      String

// Reset clears all timing data from the timer.
func (ob *Timer) Reset() {
	if ob == nil {
		Error(ENilReceiver)
		return
	}
	ob.Mutex.Lock()
	defer ob.Mutex.Unlock()
	ob.makeTasks()
} //                                                                       Reset

// -----------------------------------------------------------------------------
// # Private Method

// makeTasks initializes ob.Tasks
func (ob *Timer) makeTasks() {
	ob.Tasks = make(map[string]*TimerTask)
} //                                                                   makeTasks

//end
