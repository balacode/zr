// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-28 18:47:50 3CE918                                   [zr/flag.go]
// -----------------------------------------------------------------------------

package zr

import (
	"os"
	// str "strings" <- imported in strings_module.go
)

// HasFlags returns true if the command line used to start the program
// contains any of the specified flags. A flag must start with one or
// more '-' characters. The comparison is case-sensitive.
func HasFlags(flags ...string) bool {
	for _, arg := range os.Args {
		if !str.HasPrefix(arg, "-") {
			continue
		}
		arg = str.TrimLeft(arg, "-")
		for _, flag := range flags {
			flag = str.TrimLeft(flag, "-")
			if arg == flag {
				return true
			}
		}
	}
	return false
} //                                                                    HasFlags

//end
