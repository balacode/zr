// -----------------------------------------------------------------------------
// ZR Library                                              zr/[settings_test.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # Methods
//   Test_sett_Settings_GetSetting_
//   Test_sett_Settings_HasSetting_
//   Test_sett_Settings_SetSetting_
//
// # Extenders
//   Test_sett_Settings_ExtendGet_
//   Test_sett_Settings_ExtendHas_
//   Test_sett_Settings_ExtendSet_

//  to test all methods use:
//      go test --run Test_sett_
//
//  to generate a test coverage report use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"
)

// -----------------------------------------------------------------------------
// # Methods

// go test --run Test_sett_Settings_GetSetting_
func Test_sett_Settings_GetSetting_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Settings) GetSetting(name string) string
	//
	{
		// zero-length name must log an error
		TBeginError()
		var o Settings
		TEqual(t, o.GetSetting(""), (""))
		TCheckError(t, EInvalidArg)
	}
	{
		// check if settings are read properly
		o := Settings{
			m: map[string]string{
				"alpha": "111",
				"beta":  "222",
				"blank": "",
			},
		}
		// existing settings are looked-up?
		TEqual(t, o.GetSetting("alpha"), ("111"))
		TEqual(t, o.GetSetting("beta"), ("222"))
		//
		// settings are case-sensitive?
		TEqual(t, o.GetSetting("ALPHA"), (""))
		TEqual(t, o.GetSetting("BETA"), (""))
		//
		// non-existent settings just return a blank string
		TEqual(t, o.GetSetting("delta"), (""))
		//
		// must store blank setting values
		TEqual(t, o.GetSetting("blank"), (""))
	}
} //                                              Test_sett_Settings_GetSetting_

// go test --run Test_sett_Settings_HasSetting_
func Test_sett_Settings_HasSetting_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Settings) HasSetting(name string) bool
	{
		// zero-length name must log an error
		TBeginError()
		var o Settings
		TEqual(t, o.GetSetting(""), (""))
		TCheckError(t, EInvalidArg)
	}
	{
		// check if settings are read properly
		o := Settings{
			m: map[string]string{
				"alpha": "111",
				"beta":  "222",
				"blank": "",
			},
		}
		TEqual(t, o.HasSetting("alpha"), (true))
		TEqual(t, o.HasSetting("beta"), (true))
		TEqual(t, o.HasSetting("blank"), (true))
		//
		// settings are case-sensitive?
		TEqual(t, o.HasSetting("ALPHA"), (false))
		TEqual(t, o.HasSetting("BETA"), (false))
		//
		TEqual(t, o.HasSetting("delta"), (false))
	}
} //                                              Test_sett_Settings_HasSetting_

// go test --run Test_sett_Settings_SetSetting_
func Test_sett_Settings_SetSetting_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Settings) SetSetting(name string, value interface{})
	{
		// zero-length name must log an error
		TBeginError()
		var o Settings
		TEqual(t, o.GetSetting(""), (""))
		TCheckError(t, EInvalidArg)
	}
	{
		// check if settings are set properly
		o := Settings{
			m: map[string]string{
				"alpha": "111",
				"beta":  "222",
				"blank": "",
			},
		}
		{
			o.SetSetting("alpha", "changed")
			TTrue(t, o.GetSetting("alpha") == "changed")
		}
		{
			o.SetSetting("delta", "new")
			TTrue(t, o.GetSetting("delta") == "new")
		}
	}
} //                                              Test_sett_Settings_SetSetting_

// -----------------------------------------------------------------------------
// # Extenders

// go test --run Test_sett_Settings_ExtendGet_
func Test_sett_Settings_ExtendGet_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Settings) ExtendGet(
	//     handler func(name, value string, exists bool) string,
	// )
	// TODO: check if extender function is called after this is set
	// TODO: must not call extender function after being set to nil
	// TODO: check if extender is passed correct arguments
	// TODO: check if extender's result is used
} //                                               Test_sett_Settings_ExtendGet_

// go test --run Test_sett_Settings_ExtendHas_
func Test_sett_Settings_ExtendHas_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Settings) ExtendHas(
	//     handler func(name, value string, exists bool) bool,
	// )
	// TODO: check if extender function is called after this is set
	// TODO: must not call extender function after being set to nil
	// TODO: check if extender is passed correct arguments
	// TODO: check if extender's result is used
} //                                               Test_sett_Settings_ExtendHas_

// go test --run Test_sett_Settings_ExtendSet_
func Test_sett_Settings_ExtendSet_(t *testing.T) {
	TBegin(t)
	//
	// (ob *Settings) ExtendSet(
	//     handler func(name string, old, value interface{}) *string,
	// )
	// TODO: check if extender function is called after this is set
	// TODO: must not call extender function after being set to nil
	// TODO: check if extender is passed correct arguments
	// TODO: check if extender's result is used
} //                                               Test_sett_Settings_ExtendSet_

// end
