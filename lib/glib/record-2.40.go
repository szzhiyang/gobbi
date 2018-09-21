// This is a generated file - DO NOT EDIT
// +build glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// VariantDict is a wrapper around the C record GVariantDict.
type VariantDict struct {
	native *C.GVariantDict
}

func variantDictNewFromC(c *C.GVariantDict) *VariantDict {
	if c == nil {
		return nil
	}

	g := &VariantDict{native: c}

	return g
}

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant

// Unsupported : g_variant_dict_clear : no return generator

// Unsupported : g_variant_dict_contains : no return generator

// Unsupported : g_variant_dict_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_dict_init : unsupported parameter from_asv : Blacklisted record : GVariant

// Unsupported : g_variant_dict_insert : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_insert_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_dict_lookup : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_lookup_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// Ref is a wrapper around the C function g_variant_dict_ref.
func (recv *VariantDict) Ref() *VariantDict {
	retC := C.g_variant_dict_ref((*C.GVariantDict)(recv.native))
	retGo := variantDictNewFromC(retC)

	return retGo
}

// Unsupported : g_variant_dict_remove : no return generator

// Unsupported : g_variant_dict_unref : no return generator
