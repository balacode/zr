// -----------------------------------------------------------------------------
// ZR Library                                               zr/[string_stack.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

//   (ob *StringStack) ChangeTop(s string) string
//   (ob *StringStack) Pop() string
//   (ob *StringStack) Push(s string) string
//   (ob *StringStack) Top() string

// StringStack is a simple string stack class.
type StringStack struct {
	ar []string
} //                                                                 StringStack

// ChangeTop changes the topmost (most recently added)
// string in the stack. It returns the existing string.
// If the stack was empty, appends the string.
func (ob *StringStack) ChangeTop(s string) string {
	if ob == nil {
		mod.Error(ENilReceiver)
		return ""
	}
	max := len(ob.ar) - 1
	if max < 0 {
		ob.ar = append(ob.ar, s)
		return ""
	}
	ret := ob.ar[max]
	ob.ar[max] = s
	return ret
} //                                                                   ChangeTop

// Pop returns and removes the most recently added string from the stack.
func (ob *StringStack) Pop() string {
	if ob == nil {
		mod.Error(ENilReceiver)
		return ""
	}
	max := len(ob.ar) - 1
	if max < 0 {
		mod.Error("Pop() on empty stack")
		return ""
	}
	ob.ar = ob.ar[:max]
	max--
	return ob.ar[max]
} //                                                                         Pop

// Push adds string 's' to the top of the stack. It returns 's'.
func (ob *StringStack) Push(s string) string {
	if ob == nil {
		mod.Error(ENilReceiver)
		return ""
	}
	ob.ar = append(ob.ar, s)
	return s
} //                                                                        Push

// Top returns the topmost, most recently added string
// on the stack, without altering the stack.
func (ob *StringStack) Top() string {
	if ob == nil {
		mod.Error(ENilReceiver)
		return ""
	}
	max := len(ob.ar) - 1
	if max < 0 {
		mod.Error("Top() on empty stack")
		return ""
	}
	return ob.ar[max]
} //                                                                         Top

//end
