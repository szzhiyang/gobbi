// This is a generated file - DO NOT EDIT
// +build glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// GetUserSpecialDir is a wrapper around the C function g_get_user_special_dir.
func GetUserSpecialDir(directory UserDirectory) string {
	c_directory := (C.GUserDirectory)(directory)

	retC := C.g_get_user_special_dir(c_directory)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_once_init_enter : unsupported parameter location : no type generator for gpointer, void*

// Unsupported : g_once_init_leave : unsupported parameter location : no type generator for gpointer, void*

// Unsupported : g_regex_check_replacement : unsupported parameter has_references : no type generator for gboolean, gboolean*

// Unsupported : g_regex_escape_string : unsupported parameter string : no param type

// RegexMatchSimple is a wrapper around the C function g_regex_match_simple.
func RegexMatchSimple(pattern string, string string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) bool {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_compile_options := (C.GRegexCompileFlags)(compileOptions)

	c_match_options := (C.GRegexMatchFlags)(matchOptions)

	retC := C.g_regex_match_simple(c_pattern, c_string, c_compile_options, c_match_options)
	retGo := (bool)(retC)

	return retGo
}

// Unsupported : g_regex_split_simple : no return type

// SequenceGet is a wrapper around the C function g_sequence_get.
func SequenceGet(iter *SequenceIter) uintptr {
	c_iter := iter.toC()

	retC := C.g_sequence_get(c_iter)
	retGo := (uintptr)(retC)

	return retGo
}

// SequenceInsertBefore is a wrapper around the C function g_sequence_insert_before.
func SequenceInsertBefore(iter *SequenceIter, data uintptr) *SequenceIter {
	c_iter := iter.toC()

	c_data := (C.gpointer)(data)

	retC := C.g_sequence_insert_before(c_iter, c_data)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_sequence_move : no return generator

// Unsupported : g_sequence_move_range : no return generator

// SequenceRangeGetMidpoint is a wrapper around the C function g_sequence_range_get_midpoint.
func SequenceRangeGetMidpoint(begin *SequenceIter, end *SequenceIter) *SequenceIter {
	c_begin := begin.toC()

	c_end := end.toC()

	retC := C.g_sequence_range_get_midpoint(c_begin, c_end)
	retGo := SequenceIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_sequence_remove : no return generator

// Unsupported : g_sequence_remove_range : no return generator

// Unsupported : g_sequence_set : no return generator

// Unsupported : g_sequence_swap : no return generator

// SliceCopy is a wrapper around the C function g_slice_copy.
func SliceCopy(blockSize uint64, memBlock uintptr) uintptr {
	c_block_size := (C.gsize)(blockSize)

	c_mem_block := (C.gconstpointer)(memBlock)

	retC := C.g_slice_copy(c_block_size, c_mem_block)
	retGo := (uintptr)(retC)

	return retGo
}

// Unsupported : g_timeout_add_seconds : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_timeout_add_seconds_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// TimeoutSourceNewSeconds is a wrapper around the C function g_timeout_source_new_seconds.
func TimeoutSourceNewSeconds(interval uint32) *Source {
	c_interval := (C.guint)(interval)

	retC := C.g_timeout_source_new_seconds(c_interval)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UnicharCombiningClass is a wrapper around the C function g_unichar_combining_class.
func UnicharCombiningClass(uc rune) int32 {
	c_uc := (C.gunichar)(uc)

	retC := C.g_unichar_combining_class(c_uc)
	retGo := (int32)(retC)

	return retGo
}

// UnicharGetScript is a wrapper around the C function g_unichar_get_script.
func UnicharGetScript(ch rune) UnicodeScript {
	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_get_script(c_ch)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// UnicharIsmark is a wrapper around the C function g_unichar_ismark.
func UnicharIsmark(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_ismark(c_c)
	retGo := (bool)(retC)

	return retGo
}

// UnicharIszerowidth is a wrapper around the C function g_unichar_iszerowidth.
func UnicharIszerowidth(c rune) bool {
	c_c := (C.gunichar)(c)

	retC := C.g_unichar_iszerowidth(c_c)
	retGo := (bool)(retC)

	return retGo
}
