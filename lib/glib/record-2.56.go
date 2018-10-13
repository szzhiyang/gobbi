// This is a generated file - DO NOT EDIT
// +build glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bytes_new : unsupported parameter data : no param type

// Unsupported : g_bytes_new_static : unsupported parameter data : no param type

// Unsupported : g_bytes_new_take : unsupported parameter data : no param type

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no param type

// Copy is a wrapper around the C function g_date_copy.
func (recv *Date) Copy() *Date {
	retC := C.g_date_copy((*C.GDate)(recv.native))
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DateTimeNewFromIso8601 is a wrapper around the C function g_date_time_new_from_iso8601.
func DateTimeNewFromIso8601(text string, defaultTz *TimeZone) *DateTime {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_default_tz := (*C.GTimeZone)(defaultTz.ToC())

	retC := C.g_date_time_new_from_iso8601(c_text, c_default_tz)
	retGo := DateTimeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// GetLocaleForKey is a wrapper around the C function g_key_file_get_locale_for_key.
func (recv *KeyFile) GetLocaleForKey(groupName string, key string, locale string) string {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	retC := C.g_key_file_get_locale_for_key((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant