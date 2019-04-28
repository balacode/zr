// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 16:49:21 4D8E37                                   zr/[flag.go]
// -----------------------------------------------------------------------------

package zr

import (
	"os"
	"strings"
)

// HasFlags returns true if the command line used to start the program
// contains any of the specified flags. A flag must start with one or
// more '-' characters. The comparison is case-sensitive.
func HasFlags(flags ...string) bool {
	for _, arg := range os.Args {
		if !strings.HasPrefix(arg, "-") {
			continue
		}
		arg = strings.TrimLeft(arg, "-")
		for _, flag := range flags {
			flag = strings.TrimLeft(flag, "-")
			if arg == flag {
				return true
			}
		}
	}
	return false
} //                                                                    HasFlags

//end
