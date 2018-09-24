// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// ActionNameIsValid is a wrapper around the C function g_action_name_is_valid.
func ActionNameIsValid(actionName string) bool {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.g_action_name_is_valid(c_action_name)
	retGo := (bool)(retC)

	return retGo
}

// Unsupported : g_action_parse_detailed_name : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_action_print_detailed_name : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_icon_deserialize : unsupported parameter value : Blacklisted record : GVariant
