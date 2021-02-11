// -----------------------------------------------------------------------------
// ZR Library                                        zr/[string_aligner_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

//  to test all items in string_aligner.go use:
//      go test --run Test_stra_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"fmt"
	"testing"
)

// go test --run Test_stra_StringAligner_
func Test_stra_StringAligner_(t *testing.T) {
	TBegin(t)
	//
	// TODO: declaration comment
	//
	{
		var ob StringAligner
		TEqual(t, ob.String(), (""))
	}
	{
		var ob StringAligner
		ob.Padding = 0
		ob.Write("AAA", "BBB")
		ob.Write("A", "B", "C")
		ob.Write("AA", "BB", "CC")
		ob.Write("AAAA", "BBBB", "CCCC", "DDDD")
		result := ob.String()
		const expect = "" +
			"AAA BBB\r\n" +
			"A   B   C\r\n" +
			"AA  BB  CC\r\n" +
			"AAAABBBBCCCCDDDD"
		TEqual(t, result, (expect))
		if result != expect {
			fmt.Println("result", []byte(result))
			fmt.Println("expect", []byte(expect))
		}
	}
	{
		var ob StringAligner
		ob.Padding = 2
		ob.Write("AAA", "BBB")
		ob.Write("A", "B", "C")
		ob.Write("AA", "BB", "CC")
		ob.Write("AAAA", "BBBB", "CCCC", "DDDD")
		result := ob.String()
		const expect = "" +
			"AAA   BBB\r\n" +
			"A     B     C\r\n" +
			"AA    BB    CC\r\n" +
			"AAAA  BBBB  CCCC  DDDD"
		TEqual(t, result, (expect))
		if result != expect {
			fmt.Println("result", []byte(result))
			fmt.Println("expect", []byte(expect))
		}
	}
} //                                                    Test_stra_StringAligner_

//end
