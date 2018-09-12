// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #include <glib.h>
// #include <stdlib.h>
import "C"

// AtomicIntAdd is a wrapper around the C function g_atomic_int_add.
func AtomicIntAdd(atomic int32, val int32) int32 {
	c_atomic := (C.gint)(atomic)

	c_val := (C.gint)(val)

	retC := C.g_atomic_int_add(c_atomic, c_val)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_atomic_int_compare_and_exchange : no return type

// Unsupported : g_atomic_int_dec_and_test : no return type

// AtomicIntExchangeAndAdd is a wrapper around the C function g_atomic_int_exchange_and_add.
func AtomicIntExchangeAndAdd(atomic int32, val int32) int32 {
	c_atomic := (C.gint)(atomic)

	c_val := (C.gint)(val)

	retC := C.g_atomic_int_exchange_and_add(c_atomic, c_val)
	retGo :=
		(int32)(retC)

	return retGo
}

// AtomicIntGet is a wrapper around the C function g_atomic_int_get.
func AtomicIntGet(atomic int32) int32 {
	c_atomic := (C.gint)(atomic)

	retC := C.g_atomic_int_get(c_atomic)
	retGo :=
		(int32)(retC)

	return retGo
}

// Unsupported : g_atomic_int_inc : no return type

// Unsupported : g_atomic_int_set : no return type

// Unsupported : g_atomic_pointer_compare_and_exchange : no return type

// Blacklisted : g_atomic_pointer_get

// Unsupported : g_atomic_pointer_set : no return type

// Unsupported : g_child_watch_add : unsupported parameter pid : no param type for Pid, GPid

// Unsupported : g_child_watch_add_full : unsupported parameter pid : no param type for Pid, GPid

// Unsupported : g_child_watch_source_new : unsupported parameter pid : no param type for Pid, GPid

// Unsupported : g_file_read_link : no return type

// Unsupported : g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no param type for va_list, va_list

// Unsupported : g_setenv : unsupported parameter overwrite : no param type for gboolean, gboolean

// StripContext is a wrapper around the C function g_strip_context.
func StripContext(msgid string, msgval string) string {
	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgval := C.CString(msgval)
	defer C.free(unsafe.Pointer(c_msgval))

	retC := C.g_strip_context(c_msgid, c_msgval)
	retGo :=
		C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_strsplit_set : no return type

// Unsupported : g_unichar_get_mirror_char : no return type

// Unsupported : g_unsetenv : no return type

// Unsupported : g_vasprintf : unsupported parameter string : in for string with indirection level of 2
