// This is a generated file - DO NOT EDIT
// +build glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bit_trylock : return type :

// Unsupported : g_malloc0_n : no return generator

// Unsupported : g_malloc_n : no return generator

// Unsupported : g_realloc_n : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_try_malloc0_n : no return generator

// Unsupported : g_try_malloc_n : no return generator

// Unsupported : g_try_realloc_n : unsupported parameter mem : no type generator for gpointer (gpointer) for param mem

// Unsupported : g_variant_is_object_path : return type :

// Unsupported : g_variant_is_signature : return type :

// Unsupported : g_variant_type_string_scan : return type :

// Variant is a wrapper around the C record GVariant.
type Variant struct {
	native *C.GVariant
}

func VariantNewFromC(u unsafe.Pointer) *Variant {
	c := (*C.GVariant)(u)
	if c == nil {
		return nil
	}

	g := &Variant{native: c}

	return g
}

func (recv *Variant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}