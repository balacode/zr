// -----------------------------------------------------------------------------
// ZR Library                                                   zr/[settings.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package zr

// # CONTENTS:
//   SettingsAccessor interface
//   Settings struct
//
// # Methods (ob *Settings)
//   ) GetSetting(name string) string
//   ) HasSetting(name string) bool
//   ) SetSetting(name string, value interface{})
//
// # Extenders (ob *Settings)
//   ) ExtendGet(
//       handler func(name, value string, exists bool) string,
//   )
//   ) ExtendHas(
//       handler func(name, value string, exists bool) bool,
//   )
//   ) ExtendSet(
//       handler func(name string, old, value interface{}) *string,
//   )

import (
	"fmt"
	"strings"
)

// SettingsAccessor _ _
type SettingsAccessor interface {
	GetSetting(name string) string
	HasSetting(name string) bool
	SetSetting(name string, value interface{})
	Dump()
} //                                                            SettingsAccessor

// Settings _ _
type Settings struct {
	m         map[string]string
	extendGet func(name, value string, exists bool) string
	extendHas func(name, value string, exists bool) bool
	extendSet func(name string, old, value interface{}) *string
} //                                                                    Settings

// -----------------------------------------------------------------------------
// # Methods (ob *Settings)

// Dump prints out all settings and their stored values to console.
func (ob *Settings) Dump() {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	fmt.Println("Dump:", len(ob.m), "settings")
	for name, v := range ob.m {
		fmt.Println("name:", name, "value:", v)
	}
	fmt.Println("end")
} //                                                                        Dump

// GetSetting _ _
func (ob *Settings) GetSetting(name string) string {
	const erv = ""
	if ob == nil {
		mod.Error(ENilReceiver)
		return erv
	}
	name = strings.TrimSpace(name)
	if name == "" {
		mod.Error(EInvalidArg, "^name")
		return erv
	}
	ret, exists := ob.m[name]
	if ob.extendGet != nil {
		ret = ob.extendGet(name, ret, exists)
	}
	return ret
} //                                                                  GetSetting

// HasSetting _ _
func (ob *Settings) HasSetting(name string) bool {
	const erv = false
	if ob == nil {
		mod.Error(ENilReceiver)
		return erv
	}
	name = strings.TrimSpace(name)
	if name == "" {
		mod.Error(EInvalidArg, "^name")
		return erv
	}
	v, exists := ob.m[name]
	if ob.extendSet != nil {
		return ob.extendHas(name, v, exists)
	}
	return exists
} //                                                                  HasSetting

// SetSetting _ _
func (ob *Settings) SetSetting(name string, value interface{}) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	name = strings.TrimSpace(name)
	if name == "" {
		mod.Error(EInvalidArg, "^name")
		return
	}
	if ob.m == nil {
		ob.m = map[string]string{}
	}
	s := String(value)
	if ob.extendSet != nil {
		old := String(ob.m[name])
		result := ob.extendSet(name, old, s)
		if result == nil {
			return
		}
		s = *result
	}
	ob.m[name] = s
} //                                                                  SetSetting

// -----------------------------------------------------------------------------
// # Extenders (ob *Settings)

// ExtendGet makes 'handler' process every call to GetSetting()
func (ob *Settings) ExtendGet(
	handler func(name, value string, exists bool) string,
) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	ob.extendGet = handler
} //                                                                   ExtendGet

// ExtendHas makes 'handler' process every call to HasSetting()
func (ob *Settings) ExtendHas(
	handler func(name, value string, exists bool) bool,
) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	ob.extendHas = handler
} //                                                                   ExtendHas

// ExtendSet makes 'handle' process every call to SetSetting()
func (ob *Settings) ExtendSet(
	handler func(name string, old, value interface{}) *string,
) {
	if ob == nil {
		mod.Error(ENilReceiver)
		return
	}
	ob.extendSet = handler
} //                                                                   ExtendSet

// end
