// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-17 10:58:17 4F557E                              zr/[uuid_test.go]
// -----------------------------------------------------------------------------

package zr

// # Functions
//   Test_uuid_IsUUID_
//   Test_uuid_UUID_
//
// # Helpers
//   printSampleUUIDs(count int)

//  to test all items in uuid.go use:
//      go test --run Test_uuid_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	// "crypto/rand" // used via mod.rand proxy
	"fmt"
	"testing"
)

// -----------------------------------------------------------------------------
// # Functions

// go test --run Test_uuid_IsUUID_
func Test_uuid_IsUUID_(t *testing.T) {
	//
	// func IsUUID_(value interface{}) bool
	//
	TBegin(t)
	printSampleUUIDs(0)
	var (
		uuid1   = "B2368D91-BBF8-4184-B4A4-4E3372565324"
		uuid2   = "606F02A3-87DF-406E-9E81-69911B53DD6D"
		uuid3   = "F10D9ACB-EC6D-438D-BD02-33123801E9F4"
		goodStr = NewTStringer("226369EA-1773-44AA-8DDC-C90D4A4D2571")
		badStr  = NewTStringer("BBDB0DBF-FB03-Z0F7-9682-3B15859E01CB")
	)
	tests := []struct {
		input  interface{}
		expect bool
	}{
		{"1281F572-865C-4027-BEA8-79EB65ED045", false}, // too short
		{"", false}, // too short
		//
		{"1F92B758-0B29-3BBD-8145-4B98EC6E390B", false},
		//              ^ must be 4
		{"61B7CFCF-8383-4F2C-CC6C-072A14B18A9A", false},
		{"5ECB2B92-BD9C-4B8D-786A-6808E39D78D0", false},
		{"B65FC3FA-88F7-41A5-0A4F-BC1817E021B3", false},
		//                   ^ must be 8, 9, A, or B
		{"ADG3FCD3-2188-49AB-86B6-3CC11E9EE24A", false},
		//  ^ not a hex digit
		//
		{0, false},
		{nil, false},
		{badStr, false},
		//
		{"4ED091D4-04AE-4F30-AFF6-021EE4811F9E", true},
		{"DC5B1CE8-ED30-4781-82CE-F30338C70A17", true},
		{"89B0E155-1EF1-4273-BC4B-2242CBC72899", true},
		//
		{"5A5D357B81E946009B89D8DDD00365E5", true},
		{"C046D6F7C1A84BB693EADB6E20443E0E", true},
		{"3FC169E3574141BEA3758EA5220AACC8", true},
		//
		{&uuid1, true},
		{&uuid2, true},
		{&uuid3, true},
		{goodStr, true},
	}
	for i, test := range tests {
		got := IsUUID(test.input)
		if got != test.expect {
			TFail(t,
				"#", i, " input: ", test.input,
				" returned ", got, ". must be ", test.expect, ".",
			)
		}
	}
} //                                                           Test_uuid_IsUUID_

// go test --run Test_uuid_UUID_
func Test_uuid_UUID_(t *testing.T) {
	TBegin(t)
	// UUID_() string
	//
	{
		mod.rand.Read = func([]byte) (int, error) {
			return 0, nil
		}
		const expect = "00000000-0000-4000-8000-000000000000"
		uuid := UUID()
		if uuid != expect {
			TFail(t, "Mocked rand.Read() produced '", uuid, "'.",
				" must be '", expect, "'")
		}
		mod.Reset() // restore standard functions!
	}
	// imitate rand.Read returning an error
	{
		mod.rand.Read = func([]byte) (int, error) {
			return 0, fmt.Errorf("something wrong")
		}
		uuid := UUID()
		if uuid != "" {
			TFail(t,
				"UUID() must return a zero-length string"+
					" if rand.Read() returns an error.")
		}
		mod.Reset() // restore standard functions!
	}
	m := map[string]bool{}
	for i := 0; i < 1000; i++ {
		uuid := UUID()
		// all characters must be '0'-'9', 'A'-'F' or '-'
		for i, ch := range uuid {
			if !((ch >= '0' && ch <= '9') ||
				(ch >= 'A' && ch <= 'F') ||
				(ch >= 'a' && ch <= 'f') ||
				ch == '-') {
				TFail(t, "Invalid character in UUID '", uuid, "'",
					" at index ", i, ".")
			}
		}
		// 00000000-0000-4000-0000-000000000000 <- check for '4'
		if uuid[14] != '4' {
			TFail(t, "Invalid character in UUID '", uuid, "'",
				" at index ", i, ". must be '4'.")
		}
		// 00000000-0000-0000-X000-000000000000 <- X = '8', '9', 'A', or 'B'
		if uuid[19] != '8' && uuid[19] != '9' &&
			uuid[19] != 'A' && uuid[19] != 'B' {
			TFail(t, "Invalid character in UUID '", uuid, "'",
				" at index ", i, ". must be '4'.")
		}
		// IsUUID() should always return true
		if !IsUUID(uuid) {
			TFail(t, "Generated UUID '", uuid, "' is not valid.")
		}
		// UUID must be unique
		if _, exist := m[uuid]; exist {
			TFail(t, "Generated non-unique UUID '", uuid, "'")
		}
		m[uuid] = true
	}
} //                                                             Test_uuid_UUID_

// -----------------------------------------------------------------------------
// # Helpers

// printSampleUUIDs outputs 'count' sample UUIDs to the console.
func printSampleUUIDs(count int) {
	for i := 0; i < count; i++ {
		PL(UUID())
	}
} //                                                            printSampleUUIDs

//end
