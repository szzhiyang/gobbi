// Code generated - DO NOT EDIT.
// +build gobject_2.32 gobject_2.34 gobject_2.36 gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_signal_set_va_marshaller : unsupported parameter va_marshaller : no type generator for SignalCVaMarshaller (GSignalCVaMarshaller) for param va_marshaller

// GetSchar is a wrapper around the C function g_value_get_schar.
func (recv *Value) GetSchar() int8 {
	retC := C.g_value_get_schar((*C.GValue)(recv.native))
	retGo := (int8)(retC)

	return retGo
}

// SetSchar is a wrapper around the C function g_value_set_schar.
func (recv *Value) SetSchar(vChar int8) {
	c_v_char := (C.gint8)(vChar)

	C.g_value_set_schar((*C.GValue)(recv.native), c_v_char)

	return
}

// Clear is a wrapper around the C function g_weak_ref_clear.
func (recv *WeakRef) Clear() {
	C.g_weak_ref_clear((*C.GWeakRef)(recv.native))

	return
}

// Get is a wrapper around the C function g_weak_ref_get.
func (recv *WeakRef) Get() Object {
	retC := C.g_weak_ref_get((*C.GWeakRef)(recv.native))
	retGo := *ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Init is a wrapper around the C function g_weak_ref_init.
func (recv *WeakRef) Init(object *Object) {
	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	C.g_weak_ref_init((*C.GWeakRef)(recv.native), c_object)

	return
}

// Set is a wrapper around the C function g_weak_ref_set.
func (recv *WeakRef) Set(object *Object) {
	c_object := (C.gpointer)(C.NULL)
	if object != nil {
		c_object = (C.gpointer)(object.ToC())
	}

	C.g_weak_ref_set((*C.GWeakRef)(recv.native), c_object)

	return
}