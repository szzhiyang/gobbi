// +build glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Utf8MakeValid is a wrapper around the C function g_utf8_make_valid.
func Utf8MakeValid(str string, len int64) {}

// UuidStringIsValid is a wrapper around the C function g_uuid_string_is_valid.
func UuidStringIsValid(str string) {}

// UuidStringRandom is a wrapper around the C function g_uuid_string_random.
func UuidStringRandom() {}