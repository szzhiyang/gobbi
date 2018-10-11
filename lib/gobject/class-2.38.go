// This is a generated file - DO NOT EDIT
// +build gobject_2.38 gobject_2.42 gobject_2.44 gobject_2.46 gobject_2.54

package gobject

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unbind is a wrapper around the C function g_binding_unbind.
func (recv *Binding) Unbind() {
	C.g_binding_unbind((*C.GBinding)(recv.native))

	return
}

func (recv *Binding) Object() *Object {}

func (recv *InitiallyUnowned) Object() *Object {}

// Unsupported : g_object_new : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_valist : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_new_with_properties : unsupported parameter object_type : no type generator for GType, GType

// Unsupported : g_object_newv : unsupported parameter object_type : no type generator for GType, GType

// GetDefaultValue is a wrapper around the C function g_param_spec_get_default_value.
func (recv *ParamSpec) GetDefaultValue() *Value {
	retC := C.g_param_spec_get_default_value((*C.GParamSpec)(recv.native))
	retGo := ValueNewFromC(unsafe.Pointer(retC))

	return retGo
}

func (recv *ParamSpecBoolean) ParamSpec() *ParamSpec {}

func (recv *ParamSpecBoxed) ParamSpec() *ParamSpec {}

func (recv *ParamSpecChar) ParamSpec() *ParamSpec {}

func (recv *ParamSpecDouble) ParamSpec() *ParamSpec {}

func (recv *ParamSpecEnum) ParamSpec() *ParamSpec {}

func (recv *ParamSpecFlags) ParamSpec() *ParamSpec {}

func (recv *ParamSpecFloat) ParamSpec() *ParamSpec {}

func (recv *ParamSpecGType) ParamSpec() *ParamSpec {}

func (recv *ParamSpecInt) ParamSpec() *ParamSpec {}

func (recv *ParamSpecInt64) ParamSpec() *ParamSpec {}

func (recv *ParamSpecLong) ParamSpec() *ParamSpec {}

func (recv *ParamSpecObject) ParamSpec() *ParamSpec {}

func (recv *ParamSpecOverride) ParamSpec() *ParamSpec {}

func (recv *ParamSpecParam) ParamSpec() *ParamSpec {}

func (recv *ParamSpecPointer) ParamSpec() *ParamSpec {}

func (recv *ParamSpecString) ParamSpec() *ParamSpec {}

func (recv *ParamSpecUChar) ParamSpec() *ParamSpec {}

func (recv *ParamSpecUInt) ParamSpec() *ParamSpec {}

func (recv *ParamSpecUInt64) ParamSpec() *ParamSpec {}

func (recv *ParamSpecULong) ParamSpec() *ParamSpec {}

func (recv *ParamSpecUnichar) ParamSpec() *ParamSpec {}

func (recv *ParamSpecValueArray) ParamSpec() *ParamSpec {}

func (recv *ParamSpecVariant) ParamSpec() *ParamSpec {}

func (recv *TypeModule) Object() *Object {}
