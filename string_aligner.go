// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 339F66                         zr/[string_aligner.go]
// -----------------------------------------------------------------------------

package zr

//   StringAligner struct
//   (ob *StringAligner) Clear()
//   (ob *StringAligner) Write(row ...string)
//   (ob *StringAligner) String() string

import (
	"bytes"
)

// StringAligner left-aligns columns of strings.
type StringAligner struct {
	Values  [][]string
	Width   map[int]int
	Padding int
	Prefix  string
	Suffix  string
} //                                                               StringAligner

// Clear erases the content of the object.
func (ob *StringAligner) Clear() {
	ob.Values = [][]string{}
} //                                                                       Clear

// Write writes a row of strings to the object.
func (ob *StringAligner) Write(row ...string) {
	if ob.Width == nil {
		ob.Width = make(map[int]int, 10)
	}
	ob.Values = append(ob.Values, row)
	for i, col := range row {
		if ob.Width[i] < len(col) {
			ob.Width[i] = len(col)
		}
	}
} //                                                                       Write

// WriteBreak writes a row of strings to the object.
func (ob *StringAligner) WriteBreak() {
	ob.Values = append(ob.Values, []string{})
} //                                                                  WriteBreak

// String outputs the previously-written rows with columns aligned by spaces
// and implements the Stringer interface.
func (ob *StringAligner) String() string {
	var retBuf = bytes.NewBuffer(make([]byte, 0, 1024))
	var ws = retBuf.WriteString
	for i, row := range ob.Values {
		if i > 0 {
			ws(LB)
		}
		if ob.Prefix != "" && len(row) > 0 {
			ws(ob.Prefix)
		}
		for i, col := range row {
			ws(col)
			if i < len(row)-1 {
				var spaces = ob.Width[i] - len(col) + ob.Padding
				if spaces > 0 {
					ws(str.Repeat(" ", spaces))
				}
			}
		}
		if ob.Suffix != "" && len(row) > 0 {
			ws(ob.Suffix)
		}
	}
	return retBuf.String()
} //                                                                      String

//end
