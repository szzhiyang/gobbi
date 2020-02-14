package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// setValue takes a C.GIArgument with a value returned from a gi
// function invocation, and sets its value in the value field.
func (a *Arg) setValue(value C.GIArgument) {
	valuePtr := unsafe.Pointer(&value)
	valueAsPtr := unsafe.Pointer(*(*C.gpointer)(unsafe.Pointer(&value)))

	switch a.typ {
	case ArgType_boolean,
		ArgType_int8, ArgType_uint8, ArgType_int16, ArgType_uint16,
		ArgType_int32, ArgType_uint32, ArgType_int64, ArgType_uint64,
		ArgType_float, ArgType_double,
		ArgType_short, ArgType_ushort, ArgType_int, ArgType_uint, ArgType_long, ArgType_ulong,
		ArgType_ssize, ArgType_size,
		ArgType_pointer:
		if a.array {
			a.setValueSimpleArray(valueAsPtr)
		} else {
			a.setValueSimple(valuePtr)
		}
	case ArgType_string:
		if a.array {
			a.setValueStringArray(valuePtr)
		} else {
			a.setValueString(valuePtr)
		}
	default:
		panic(fmt.Sprintf("Unhandled arg type, %#v", a))
	}
}

func (a *Arg) setValueSimple(valuePtr unsafe.Pointer) {
	switch a.typ {
	case ArgType_boolean:
		a.value = (*(*C.gboolean)(valuePtr)) == C.TRUE
	case ArgType_int8:
		a.value = (int8)(*(*C.gint8)(valuePtr))
	case ArgType_uint8:
		a.value = (uint8)(*(*C.guint8)(valuePtr))
	case ArgType_int16:
		a.value = (int16)(*(*C.gint16)(valuePtr))
	case ArgType_uint16:
		a.value = (uint16)(*(*C.guint16)(valuePtr))
	case ArgType_int32:
		a.value = (int32)(*(*C.gint32)(valuePtr))
	case ArgType_uint32:
		a.value = (uint32)(*(*C.guint32)(valuePtr))
	case ArgType_int64:
		a.value = (int64)(*(*C.gint64)(valuePtr))
	case ArgType_uint64:
		a.value = (uint64)(*(*C.guint64)(valuePtr))
	case ArgType_float:
		a.value = (float32)(*(*C.gfloat)(valuePtr))
	case ArgType_double:
		a.value = (float64)(*(*C.gdouble)(valuePtr))
	case ArgType_short:
		a.value = (int16)(*(*C.gshort)(valuePtr))
	case ArgType_ushort:
		a.value = (uint16)(*(*C.gushort)(valuePtr))
	case ArgType_int:
		a.value = (int)(*(*C.gint)(valuePtr))
	case ArgType_uint:
		a.value = (uint)(*(*C.guint)(valuePtr))
	case ArgType_long:
		a.value = (int64)(*(*C.glong)(valuePtr))
	case ArgType_ulong:
		a.value = (uint64)(*(*C.gulong)(valuePtr))
	case ArgType_ssize:
		a.value = (int)(*(*C.gssize)(valuePtr))
	case ArgType_size:
		a.value = (uint)(*(*C.gsize)(valuePtr))
	case ArgType_pointer:
		a.value = unsafe.Pointer(*(*C.gpointer)(valuePtr))
	}
}

func (a *Arg) setValueSimpleArray(valuePtr unsafe.Pointer) {
	if a.arrayNullTerminated {
		panic("not supported : simple null-terminated array")
	}

	switch a.typ {
	case ArgType_boolean:
		a.setValueBoolArray(valuePtr)
	case ArgType_int8:
		a.value = (*[1 << 28]int8)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_uint8:
		a.value = (*[1 << 28]uint8)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_int16:
		a.value = (*[1 << 28]int16)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_uint16:
		a.value = (*[1 << 28]uint16)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_int32:
		a.value = (*[1 << 28]int32)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_uint32:
		a.value = (*[1 << 28]uint32)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_int64:
		a.value = (*[1 << 28]int64)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_uint64:
		a.value = (*[1 << 28]uint64)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_float:
		a.value = (*[1 << 28]float32)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_double:
		a.value = (*[1 << 28]float64)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_short:
		a.value = (*[1 << 28]int16)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_ushort:
		a.value = (*[1 << 28]uint16)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_int:
		a.value = (*[1 << 28]int)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_uint:
		a.value = (*[1 << 28]uint)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_long:
		a.value = (*[1 << 28]int64)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_ulong:
		a.value = (*[1 << 28]uint64)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_ssize:
		a.value = (*[1 << 28]int)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_size:
		a.value = (*[1 << 28]uint)(valuePtr)[:a.arrayLength:a.arrayLength]
	case ArgType_pointer:
		a.value = (*[1 << 28]unsafe.Pointer)(valuePtr)[:a.arrayLength:a.arrayLength]
	default:
		panic(fmt.Sprintf("Unsupported array of type %#v", a.value))
	}
}

func (a *Arg) setValueString(valuePtr unsafe.Pointer) {
	var cString *C.gchar
	if a.out {
		cString = *(**C.gchar)(unsafe.Pointer(&a.outPtr))
	} else {
		cString = *(**C.gchar)(valuePtr)
	}

	a.value = C.GoString(cString)

	if a.transferOwnership != TransferOwnershipNone {
		C.free(unsafe.Pointer(cString))
	}
}

func (a *Arg) setValueStringArray(valuePtr unsafe.Pointer) {
	if a.arrayNullTerminated {
		a.setValueStringArrayNullTerminated(valuePtr)
	} else {
		panic("not supported : string array with length")
	}
}

func (a *Arg) setValueStringArrayNullTerminated(valuePtr unsafe.Pointer) {
	strings := []string{}

	array := *(***C.gchar)(valuePtr)

	arrayItem := array
	for *arrayItem != nil {
		goString := C.GoString(*arrayItem)
		strings = append(strings, goString)

		if a.transferOwnership == TransferOwnershipFull {
			C.free(unsafe.Pointer(*arrayItem))
		}

		arrayItem = (**C.gchar)(incrPointerPointer(unsafe.Pointer(arrayItem)))
	}

	a.value = strings

	if a.transferOwnership == TransferOwnershipFull ||
		a.transferOwnership == TransferOwnershipContainer {

		C.free(unsafe.Pointer(array))
	}
}

func (a *Arg) setValueBoolArray(valuePtr unsafe.Pointer) {
	cBooleans := (*[1 << 28]C.gboolean)(valuePtr)[:a.arrayLength:a.arrayLength]
	booleans := make([]bool, a.arrayLength, a.arrayLength)

	for i, cBoolean := range cBooleans {
		if cBoolean == C.TRUE {
			booleans[i] = true
		} else {
			booleans[i] = false
		}
	}

	a.value = booleans
}
