// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                   [zr/logging.go]
// -----------------------------------------------------------------------------

package zr

// # Global Settings
//   GetLastLogMessage() string
//   GetShowSourceFileNames() bool
//   SetShowSourceFileNames(val bool)
//   GetVerboseMode() bool
//   SetVerboseMode(val bool)
//
// # Utility Functions
//   Assert(expect bool) bool
//   Callers(options ...interface{}) string
//   D(message string, args ...interface{})
//   DC(message string, args ...interface{})
//   DL(message string, args ...interface{})
//   DLC(message string, args ...interface{})
//   Error(args ...interface{}) error
//   GetErrorCount() int
//   Log(args ...interface{})
//   Logf(format string, args ...interface{})
//   PrintfAsync(format string, args ...interface{})
//   Timestamp() string
//   TM(messages ...string)
//   VerboseLog(args ...interface{})
//   VerboseLogf(format string, args ...interface{})
//
// # Internal Functions
//   formatArgs(format string, args ...interface{}) string
//   joinArgs(prefix string, args ...interface{}) string
//   logAsync(message string)
//   logLoopAsync()
//   removeLogOptions(args []interface{}) (ret []interface{})

import "bytes"       // standard
import "fmt"         // standard
import "os"          // standard
import "runtime"     // standard
import "strconv"     // standard
import str "strings" // standard
import "sync"        // standard
import "time"        // standard

// -----------------------------------------------------------------------------
// # Config Settings

// disableErrors disables or enables logging of errors.
var disableErrors bool

// DisableErrors __
func DisableErrors(optDisable ...bool) {
	switch n := len(optDisable); {
	case n == 1:
		disableErrors = optDisable[0]
	case n > 1:
		Error(EInvalidArg, "optDisable", ":", optDisable)
	default:
		disableErrors = true
	}
} //                                                               DisableErrors

// EnableErrors __
func EnableErrors(optEnable ...bool) {
	switch n := len(optEnable); {
	case n == 1:
		disableErrors = !optEnable[0]
	case n > 1:
		Error(EInvalidArg, "optEnable", ":", optEnable)
	default:
		disableErrors = false
	}
} //                                                                EnableErrors

// -----------------------------------------------------------------------------
// # Variables

// Config Variables

// callerPrefix specifies how Callers() prefixes each call stack entry.
// By default, each call stack entry starts on a new line and is indented.
const callerPrefix = LB + "    "

// showSourceFileNames makes Callers() display file names when set to true.
var showSourceFileNames bool

// verboseMode is global setting that turns verbose logging on or off.
var verboseMode bool

// Information Variables:

// errorCount holds the number of errors.
var errorCount int

// lastLogMessage holds the last logged entry or error.
var lastLogMessage string

// lastLogTime holds the time of the last logged entry or error.
var lastLogTime time.Time

// Call Range Types:

// HideCallers hides the call stack when passed as one of the arguments
// to Error() and Callers(). It does not interfere with other output.
type HideCallers struct{}

// MinDepth specifies the closest caller in the call stack that should
// be output by Error() and Callers(), when passed as one of the
// arguments to these functions. The function that called the current
// function is 1, etc. It does not interfere with other output.
type MinDepth int

// MaxDepth specifies the highest caller in the call stack that should
// be output by Error() and Callers(), when passed as one of the
// arguments to these functions. The function that called the current
// function is 1, etc. It does not interfere with other output.
type MaxDepth int

// Private Variables
var logChan = make(chan logArgs, 50000)
var logMutex sync.RWMutex
var logSN int

// -----------------------------------------------------------------------------
// # Async Logging Type

// logArgs __
type logArgs struct {
	msg       string
	writeFile bool
	logTime   time.Time
} //                                                                     logArgs

// -----------------------------------------------------------------------------
// # Global Settings

// GetLastLogMessage returns the last logged message.
// Log messages are commonly emitted by Error().
func GetLastLogMessage() string {
	return lastLogMessage
} //                                                           GetLastLogMessage

// GetShowSourceFileNames __
func GetShowSourceFileNames() bool {
	return showSourceFileNames
} //                                                      GetShowSourceFileNames

// SetShowSourceFileNames __
func SetShowSourceFileNames(val bool) {
	showSourceFileNames = val
} //                                                      SetShowSourceFileNames

// GetVerboseMode __
func GetVerboseMode() bool {
	return verboseMode
} //                                                              GetVerboseMode

// SetVerboseMode __
func SetVerboseMode(val bool) {
	verboseMode = val
} //                                                              SetVerboseMode

// -----------------------------------------------------------------------------
// # Utility Functions

// Assert checks if the 'expect' condition is true. If the
// condition is false, it outputs an 'ASSERTION FAILED' message
// to the standard output, including  the function,
// line number and list of functions on the call stack.
func Assert(expect bool) bool {
	if !expect {
		Error("ASSERTION FAILED")
	}
	return expect
} //                                                                      Assert

// Callers returns a human-friendly string showing the call stack with
// each calling method or function's name and line number.
// The most immediate callers  are shown first, followed by their callers,
// and so on. For brevity, 'runtime.*' and 'syscall.*' etc.
// top-level callers are not included.
func Callers(options ...interface{}) string {
	var minDepth, maxDepth = -1, -1
	for _, opt := range options {
		switch val := opt.(type) {
		case HideCallers:
			return ""
		case MinDepth:
			minDepth = int(val)
		case MaxDepth:
			maxDepth = int(val)
		}
	}
	if maxDepth == 0 {
		return ""
	}
	var retBuf = bytes.NewBuffer(make([]byte, 0, 1024))
	var ws = retBuf.WriteString
	for i, depth := 0, 0; ; i++ {
		var programCounter, filename, lineNo, _ = runtime.Caller(i)
		var funcName = runtime.FuncForPC(programCounter).Name()
		// end loop on reaching a top-level runtime function
		if funcName == "" ||
			funcName == "runtime.goexit" ||
			funcName == "runtime.main" ||
			funcName == "testing.tRunner" ||
			str.Contains(funcName, "HandlerFunc.ServeHTTP") {
			break
		}
		// skip runtime/syscall functions, but continue the loop
		if str.Contains(funcName, "zr.Callers") ||
			str.Contains(funcName, "zr.Error") ||
			str.Contains(funcName, "zr.Log") ||
			str.Contains(funcName, "zr.logAsync") ||
			str.HasPrefix(funcName, "runtime.") ||
			str.HasPrefix(funcName, "syscall.") {
			continue
		}
		// increase depth counter and skip opt-of range functions
		depth++
		if minDepth != -1 && depth < minDepth {
			continue
		}
		if maxDepth != -1 && depth > maxDepth {
			break
		}
		// let the file name's path use the right kind of OS path separator
		// (by default, the file name contains '/' on all platforms)
		if string(os.PathSeparator) != "/" {
			filename = str.Replace(filename, "/", string(os.PathSeparator), -1)
		}
		// remove parent module/function names
		if index := str.LastIndex(funcName, "/"); index != -1 {
			funcName = funcName[index+1:]
		}
		if str.Count(funcName, ".") > 1 {
			funcName = funcName[str.Index(funcName, ".")+1:]
		}
		// remove unneeded punctuation from function names
		for _, find := range []string{"(", ")", "*"} {
			if str.Contains(funcName, find) {
				funcName = str.Replace(funcName, find, "", -1)
			}
		}
		ws(callerPrefix)
		if showSourceFileNames {
			ws(fmt.Sprintf("%-30s  %4d  %-30s", funcName, lineNo, filename))
			continue
		}
		ws(fmt.Sprintf("%s:%d", funcName, lineNo))
	}
	return retBuf.String()
} //                                                                     Callers

// CallerList returns a human-friendly list of strings showing the
// call stack with each calling method or function's name and line number.
//
// The most immediate callers are listed first, followed by their callers,
// and so on. For brevity, 'runtime.*' and 'syscall.*'
// and other top-level callers are not included.
func CallerList() []string {
	var ret []string
	var i = 0
	for {
		i++
		var programCounter, filename, lineNo, _ = runtime.Caller(i)
		var funcName = runtime.FuncForPC(programCounter).Name()
		// end loop on reaching a top-level runtime function
		if funcName == "" ||
			funcName == "runtime.goexit" ||
			funcName == "runtime.main" ||
			funcName == "testing.tRunner" ||
			str.Contains(funcName, "HandlerFunc.ServeHTTP") {
			break
		}
		// skip runtime/syscall functions, but continue the loop
		if str.Contains(funcName, "zr.Callers") ||
			str.Contains(funcName, "zr.CallerList") ||
			str.Contains(funcName, "zr.Error") ||
			str.Contains(funcName, "zr.Log") ||
			str.Contains(funcName, "zr.logAsync") ||
			str.HasPrefix(funcName, "runtime.") ||
			str.HasPrefix(funcName, "syscall.") {
			continue
		}
		// let the file name's path use the right kind of OS path separator
		// (by default, the file name contains '/' on all platforms)
		if string(os.PathSeparator) != "/" {
			filename = str.Replace(filename, "/", string(os.PathSeparator), -1)
		}
		// remove parent module/function names
		if index := str.LastIndex(funcName, "/"); index != -1 {
			funcName = funcName[index+1:]
		}
		if str.Count(funcName, ".") > 1 {
			funcName = funcName[str.Index(funcName, ".")+1:]
		}
		// remove unneeded punctuation from function names
		for _, find := range []string{"(", ")", "*"} {
			if str.Contains(funcName, find) {
				funcName = str.Replace(funcName, find, "", -1)
			}
		}
		var line string
		if showSourceFileNames {
			line = fmt.Sprintf("func:%-30s  ln:%4d  file:%-30s",
				funcName, lineNo, filename)
		} else {
			line = fmt.Sprintf("%s:%d", funcName, lineNo)
		}
		ret = append(ret, line)
	}
	return ret
} //                                                                  CallerList

// D writes a formatted debug message and to the console.
// Same as fmt.Printf(), but appends a newline at the end.
func D(message string, args ...interface{}) {
	fmt.Printf(Timestamp()+message+LB, args...)
} //                                                                           D

// DC writes a formatted debug message and the call stack to the console.
func DC(message string, args ...interface{}) {
	fmt.Println(Timestamp() + fmt.Sprintf(message, args...) + Callers())
} //                                                                          DC

// DL writes a formatted debug message to log file 'run.log' saved in the
// program's current directory. The message is not output to the console.
func DL(message string, args ...interface{}) {
	AppendToTextFile("run.log", Timestamp()+fmt.Sprintf(message, args...)+LB)
} //                                                                          DL

// DLC writes a formatted debug message and the call stack to log file
// 'run.log' saved in the program's current directory.
// The message is not output to the console.
func DLC(message string, args ...interface{}) {
	AppendToTextFile("run.log",
		Timestamp()+fmt.Sprintf(message, args...)+Callers()+LB)
} //                                                                         DLC

// Error outputs an error message to the standard output and to a
// log file named 'run.log' saved in the program's current directory,
// It also outputs the call stack (names and line numbers of callers.)
// Error has no effect if disableErrors flag is set to true.
// Returns an error value initialized with the message.
func Error(args ...interface{}) error {
	errorCount++
	var msg = joinArgs("ERR: ", args...)
	lastLogMessage = msg
	if !disableErrors {
		logAsync(msg + Callers(args...))
	}
	return fmt.Errorf(msg)
} //                                                                       Error

// GetErrorCount returns the number of errors that occurred.
func GetErrorCount() int {
	return errorCount
} //                                                               GetErrorCount

// Log outputs a message string to the standard output and to a
// log file named 'run.log' saved in the program's current directory.
// It also outputs the call stack (names and line numbers of callers.)
func Log(args ...interface{}) {
	logAsync(joinArgs("", args))
} //                                                                         Log

// Logf outputs a formatted message to the standard output and to a
// log file named 'run.log' saved in the program's current directory.
// The 'format' parameter accepts a format string, followed by one or
// more optional arguments, exactly like fmt.Printf() and fmt.Errorf()
// It also outputs the call stack (names and line numbers of callers.)
func Logf(format string, args ...interface{}) {
	logAsync(formatArgs(format, args...))
} //                                                                        Logf

// PrintfAsync prints output to the standard output like fmt.Printf(),
// but asynchronously using the log loop goroutine.
// This prevents the program from being slowed down by output to console.
// (This slow-down may occur on Windows)
func PrintfAsync(format string, args ...interface{}) {
	logChan <- logArgs{msg: formatArgs(format, args...)}
	if logSN == 0 {
		go logLoopAsync()
	}
} //                                                                 PrintfAsync

// Timestamp returns a timestamp string using the current time.
// The timestamp includes the date, time with seconds
// and milliseconds, but no time zone.
func Timestamp() string {
	var ret = time.Now().String()
	if len(ret) > 24 {
		ret = ret[:24]
	}
	for len(ret) < 24 { // may need to add trailing zeros in milliseconds
		ret += "0"
	}
	return ret + " "
} //                                                                   Timestamp

// TM outputs milliseconds elapsed between calls to TM() to standard output.
// To start timing, call TM() without any arguments.
func TM(messages ...string) {
	var callLoc string
	{
		var buf = bytes.NewBuffer(make([]byte, 0, 128))
		var ws = buf.WriteString
		for i := 1; i <= 4; i++ {
			var programCounter, _, _, _ = runtime.Caller(i)
			var funcName = runtime.FuncForPC(programCounter).Name()
			ws("|")
			ws(funcName)
		}
		callLoc = buf.String()
	}
	if timings == nil {
		timings = make(map[string]time.Time, 20)
	}
	var messagesLen = len(messages)
	switch {
	case messagesLen == 0 || (messagesLen == 1 && messages[0] == ""):
		timings[callLoc] = time.Now()
		return
	case messagesLen == 1:
		var now = time.Now()
		fmt.Printf("TM % 8.2f ms: %s"+LB,
			float32(now.Sub(timings[callLoc]).Nanoseconds())/1000000,
			messages[0])
		timings[callLoc] = time.Now()
	default:
		Error("Too many values in 'messages' argument")
	}
} //                                                                          TM
var timings map[string]time.Time

// VerboseLog sends output to the log loop,
// but only when verbose mode is set to true.
func VerboseLog(args ...interface{}) {
	if !verboseMode {
		return
	}
	var msg = fmt.Sprint(args...)
	logAsync(msg)
} //                                                                  VerboseLog

// VerboseLogf outputs a formatted message to the standard output and to a
// log file named 'run.log' saved in the program's current directory,
// only when verbose mode is set to true.
// The 'format' parameter accepts a format string, followed by one or
// more optional arguments, exactly like fmt.Printf() and fmt.Errorf()
// It also outputs the call stack (names and line numbers of callers.)
func VerboseLogf(format string, args ...interface{}) {
	if !verboseMode {
		return
	}
	logAsync(formatArgs(format, args...))
} //                                                                 VerboseLogf

// -----------------------------------------------------------------------------
// # Internal Functions

// formatArgs returns a string built from a 'format' string and a list of
// variadic arguments, in a similar manner to fmt.Sprintf(). The only
// difference with fmt.Sprintf() is that this function removes special
// call log options such as HideCallers, MinDepth and MaxDepth,
// and trims white spaces from the final result.
func formatArgs(format string, args ...interface{}) string {
	args = removeLogOptions(args)
	return str.Trim(fmt.Sprintf(format, args...), SPACES)
} //                                                                  formatArgs

// joinArgs returns a string built from a list of variadic arguments 'args',
// with some minimal formatting rules described as follows:
//
// Inserts a space between each argument, unless the preceding argument
// ends with '(', or the current argument begins with ')' or ':'.
//
// If a string argument in 'args' begins with '^', then the '^' is removed
// and the argument's string is quoted in single quotes without escaping it.
//
// If a string argument in 'args' ends with '^', then the '^' is removed
// and the next argument is quoted in the same way.
func joinArgs(prefix string, args ...interface{}) string {
	args = removeLogOptions(args)
	var quoteNext bool
	var lastChar byte
	var retBuf bytes.Buffer
	var ws = retBuf.WriteString
	ws(prefix)
	for i, arg := range args {
		var s = fmt.Sprint(arg)
		var firstChar = byte(0)
		if len(s) > 0 {
			firstChar = s[0]
		}
		if i > 0 &&
			lastChar != '(' &&
			firstChar != ')' &&
			firstChar != ':' {
			ws(" ")
		}
		var q = quoteNext
		if str.HasPrefix(s, "^") {
			q = true
			s = s[1:]
		}
		quoteNext = str.HasSuffix(s, "^")
		if quoteNext {
			s = s[:len(s)-1]
		}
		if q {
			ws("'")
		}
		ws(s)
		if q {
			ws("'")
		}
		lastChar = 0
		if len(s) > 0 {
			lastChar = s[len(s)-1]
		}
	}
	return retBuf.String()
} //                                                                    joinArgs

// logAsync outputs a message to the standard output and to a log
// file named 'run.log' saved in the program's current directory.
// It also outputs the call stack (names and line numbers of callers.)
func logAsync(message string) {
	lastLogTime = time.Now()
	if disableErrors {
		return
	}
	logChan <- logArgs{
		msg:       message,
		writeFile: true,
		logTime:   lastLogTime,
	}
	if logSN == 0 {
		go logLoopAsync()
	}
} //                                                                    logAsync

// logLoopAsync handles asynchronous writing of log messages to the log file
// and the console. It receives log messages via logChan. The goroutine
// running logLoopAsync() only stops when the main() function exists.
func logLoopAsync() {
	for {
		var t = <-logChan
		logMutex.Lock()
		logSN++
		var msg = t.msg
		lastLogMessage = msg
		lastLogTime = t.logTime
		if !disableErrors {
			msg = Timestamp() + " #" + strconv.Itoa(logSN) + " " + msg
			msg = str.Trim(msg, SPACES)
			fmt.Println(msg)
			if t.writeFile {
				AppendToTextFile("run.log", msg+LB)
			}
		}
		logMutex.Unlock()
	}
} //                                                                logLoopAsync

// removeLogOptions removes all HideCallers, MinDepth and MaxDepth types
// from an interface array 'args'. The original array is not altered.
// These special types are used to control the output of Callers(),
// but should not appear in normal output.
func removeLogOptions(args []interface{}) (ret []interface{}) {
	for _, v := range args {
		switch v.(type) {
		case HideCallers, MinDepth, MaxDepth:
			continue
		default:
			ret = append(ret, v)
		}
	}
	return ret
} //                                                            removeLogOptions

//end