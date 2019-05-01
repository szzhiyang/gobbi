// This is a generated file - DO NOT EDIT
// +build gobject_2.4 gobject_2.6 gobject_2.8 gobject_2.10 gobject_2.12 gobject_2.18 gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// ParamSpecOverride is a wrapper around the C record GParamSpecOverride.
type ParamSpecOverride struct {
	native *C.GParamSpecOverride
	// Private : parent_instance
	// Private : overridden
}

func ParamSpecOverrideNewFromC(u unsafe.Pointer) *ParamSpecOverride {
	c := (*C.GParamSpecOverride)(u)
	if c == nil {
		return nil
	}

	g := &ParamSpecOverride{native: c}

	return g
}

func (recv *ParamSpecOverride) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_param_spec_override : return type :

// Unsupported : g_signal_accumulator_true_handled : unsupported parameter dummy : no type generator for gpointer (gpointer) for param dummy

// Unsupported : g_type_add_interface_check : unsupported parameter check_data : no type generator for gpointer (gpointer) for param check_data

// Unsupported : g_type_class_peek_static : return type :

// Unsupported : g_type_default_interface_peek : return type :

// Unsupported : g_type_default_interface_ref : return type :

// Unsupported : g_type_remove_interface_check : unsupported parameter check_data : no type generator for gpointer (gpointer) for param check_data