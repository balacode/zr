// -----------------------------------------------------------------------------
// ZR Library                                                      zr/[timer.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # Types
//   Timer struct
//   TimerTask struct
//
// # tmNameTask Sortable Slice
//   tmNameTask struct
//   tmNameTasks []tmNameTask
//       ) Len() int
//       ) Less(i, j int) bool
//       ) Swap(i, j int)
//
// # Methods (ob *Timer)
//   ) GetTasks() map[string]*TimerTask
//   ) Rename(taskName, newName string) string
//   ) Start(taskName string) string
//   ) Stop(taskName string)
//   ) StopLast()
//   ) Reset()
//
// # Reporting Methods (ob *Timer)
//   ) Print(prefix ...string)
//   ) ReportByTimeSpent() string
//   ) String() string
//
// # Private Method (ob *Timer)
//   ) makeTasks()

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// SUGGESTION: new methods NewTimer() and StopPrint():
// tm := NewTimer("App.GoNextLine")
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
	Mutex        sync.RWMutex
	Tasks        map[string]*TimerTask
	LastTaskName string
	PrintEvents  bool
} //                                                                       Timer

// TimerTask holds the timing statistics of a timed task.
type TimerTask struct {
	Count     int
	SerialNo  int
	StartTime time.Time
	TotalMs   float32
} //                                                                   TimerTask

// -----------------------------------------------------------------------------
// # tmNameTask Sortable Slice

type tmNameTask struct {
	name string
	task *TimerTask
}

type tmNameTasks []tmNameTask

// Len is the number of elements in the collection. (sort.Interface)
func (ob tmNameTasks) Len() int {
	return len(ob)
} //                                                                         Len

// Less reports whether the element with index
// 'i' should sort before element[j].
// (sort.Interface)
//
// This particular implementation sorts items
// by time spent in descending order.
//
func (ob tmNameTasks) Less(i, j int) bool {
	ti := ob[i].task.TotalMs
	tj := ob[j].task.TotalMs
	return ti > tj
} //                                                                        Less

// Swap swaps the elements with indexes i and j. (sort.Interface)
func (ob tmNameTasks) Swap(i, j int) {
	ob[i], ob[j] = ob[j], ob[i]
} //                                                                        Swap

// -----------------------------------------------------------------------------
// # Methods (ob *Timer)

// GetTasks returns a map of named tasks and their timing statistics.
func (ob *Timer) GetTasks() map[string]*TimerTask {
	return ob.Tasks
} //                                                                    GetTasks

// Rename renames a timed task. If there is an existing task with the same
// name as newTaskName, the old task will be merged with the existing task.
// The times of the two tasks will be combined.
//
// Returns newTaskName
//
func (ob *Timer) Rename(taskName, newName string) string {
	old, hasOld := ob.Tasks[taskName]
	if taskName == newName || taskName == "" || newName == "" {
		goto exit
	}
	if !hasOld {
		goto exit
	}
	if new, hasNew := ob.Tasks[newName]; hasNew {
		new.Count += old.Count
		if new.SerialNo > old.SerialNo {
			new.SerialNo = old.SerialNo
		}
		if new.StartTime.After(old.StartTime) {
			new.StartTime = old.StartTime
		}
		new.TotalMs += old.TotalMs
	} else {
		ob.Tasks[newName] = old
	}
	delete(ob.Tasks, taskName)
exit:
	return newName
} //                                                                      Rename

// Start begins timing the named task. Make sure you call Stop() when
// the task is complete. You can start and stop the same task multiple
// times, provided you call Stop() after every Start().
func (ob *Timer) Start(taskName string) string {
	now := time.Now()
	if ob.PrintEvents {
		fmt.Println("start", taskName+":", time.Now().String()[11:19])
	}
	ob.Mutex.Lock()
	defer ob.Mutex.Unlock()
	if ob.Tasks == nil {
		ob.makeTasks()
	}
	ob.LastTaskName = taskName
	task, exists := ob.Tasks[taskName]
	if exists {
		task.StartTime = now
		return taskName
	}
	ob.Tasks[taskName] = &TimerTask{
		SerialNo:  len(ob.Tasks) + 1,
		StartTime: now,
	}
	return taskName
} //                                                                       Start

// Stop stops timing the named task and stores the time spent in the Timer.
func (ob *Timer) Stop(taskName string) {
	now := time.Now()
	if ob.PrintEvents {
		fmt.Println("stop", taskName+":", time.Now().String()[11:19])
	}
	ob.Mutex.Lock()
	defer ob.Mutex.Unlock()
	if ob.Tasks == nil {
		ob.makeTasks()
	}
	task, exists := ob.Tasks[taskName]
	if !exists {
		Error("Never started timing^", taskName)
		PL("THERE ARE", len(ob.Tasks), "TASKS")
		return
	}
	ms := float32(now.Sub(task.StartTime).Nanoseconds()) / 1000000
	task.Count++
	task.StartTime = now
	task.TotalMs += ms
} //                                                                        Stop

// StopLast _ _
func (ob *Timer) StopLast() {
	if ob.LastTaskName == "" {
		return
	}
	ob.Stop(ob.LastTaskName)
	ob.LastTaskName = ""
} //                                                                    StopLast

// Reset clears all timing data from the timer.
func (ob *Timer) Reset() {
	ob.Mutex.Lock()
	defer ob.Mutex.Unlock()
	ob.makeTasks()
} //                                                                       Reset

// -----------------------------------------------------------------------------
// # Reporting Methods (ob *Timer)

// Print prints out a timing report to the console (i.e. standard output)
// Shows the name of each task, the total time spent on the task, the
// number of times the task was executed, and the average running time
// in seconds rounded to 4 decimal places.
func (ob *Timer) Print(prefix ...string) {
	if ob.Tasks == nil {
		ob.makeTasks()
	}
	s := ob.String()
	if len(prefix) > 0 {
		pre := strings.Join(prefix, "")
		lines := strings.Split(s, "\n")
		for i, line := range lines {
			lines[i] = pre + line
		}
		s = strings.Join(lines, "\n")
	}
	fmt.Println(s)
} //                                                                       Print

// ReportByTimeSpent returns the timing report as a string,
// with items sorted by time spent in descending order.
func (ob *Timer) ReportByTimeSpent() string {
	ob.Mutex.RLock()
	defer ob.Mutex.RUnlock()
	//
	serialMax := 0
	for _, task := range ob.Tasks {
		if task.SerialNo > serialMax {
			serialMax = task.SerialNo
		}
	}
	// sort the tasks
	var ar []tmNameTask
	for name, task := range ob.Tasks {
		ar = append(ar, tmNameTask{name, task})
	}
	sort.Sort(tmNameTasks(ar))
	//
	var buf bytes.Buffer
	ws := buf.WriteString
	ws("    --------------------------------- SECONDS:\r\n")
	sum := float64(0)
	for _, it := range ar {
		seconds := float64(it.task.TotalMs) / float64(1000)
		sum += seconds
		ws(fmt.Sprintf("%14.5f: %s\r\n", seconds, it.name))
	}
	ws(fmt.Sprintf("%14.5f\r\n", sum))
	ret := buf.String()
	return ret
} //                                                           ReportByTimeSpent

// String returns the timing report as a string,
// and implements the fmt.Stringer interface.
func (ob *Timer) String() string {
	ob.Mutex.RLock()
	defer ob.Mutex.RUnlock()
	//
	serialMax := 0
	for _, task := range ob.Tasks {
		if task.SerialNo > serialMax {
			serialMax = task.SerialNo
		}
	}
	var buf bytes.Buffer
	ws := buf.WriteString
	ws("    --------------------------------- SECONDS:\r\n")
	sum := float64(0)
	for i := 0; i <= serialMax; i++ {
		for taskName, task := range ob.Tasks {
			if task.SerialNo == i {
				seconds := float64(task.TotalMs) / float64(1000)
				sum += seconds
				ws(fmt.Sprintf("%14.5f: %s\r\n", seconds, taskName))
			}
		}
	}
	ws(fmt.Sprintf("%14.5f\r\n", sum))
	ret := buf.String()
	return ret
} //                                                                      String

// -----------------------------------------------------------------------------
// # Private Method (ob *Timer)

// makeTasks initializes ob.Tasks
func (ob *Timer) makeTasks() {
	ob.Tasks = make(map[string]*TimerTask)
} //                                                                   makeTasks

// end
