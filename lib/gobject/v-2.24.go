// Code generated - DO NOT EDIT.
// +build gobject_2.24 gobject_2.26 gobject_2.28 gobject_2.30 gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// TypeAddClassPrivate is a wrapper around the C function g_type_add_class_private.
func TypeAddClassPrivate(classType Type, privateSize uint64) {
	c_class_type := (C.GType)(classType)

	c_private_size := (C.gsize)(privateSize)

	C.g_type_add_class_private(c_class_type, c_private_size)

	return
}