package gi

// #include <girepository.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type ArgType int

const (
	ArgType_boolean ArgType = iota
	ArgType_int8
	ArgType_uint8
	ArgType_int16
	ArgType_uint16
	ArgType_int32
	ArgType_uint32
	ArgType_int64
	ArgType_uint64
	ArgType_float
	ArgType_double
	ArgType_short
	ArgType_ushort
	ArgType_int
	ArgType_uint
	ArgType_long
	ArgType_ulong
	ArgType_ssize
	ArgType_size
	ArgType_string
	ArgType_pointer
)

type Arg struct {
	value interface{}
	typ   ArgType

	pointer bool
	array   bool
	in      bool
	out     bool
}

func (a *Arg) getValue() C.GIArgument {
	var cArg C.GIArgument

	return cArg
}

func (a *Arg) setValue(value C.GIArgument) {
	switch a.typ {
	case ArgType_uint:
		a.value = (uint)(*(*C.guint)(unsafe.Pointer(&value)))
	default:
		panic(fmt.Sprintf("Unhandle arg type, %#v", a))
	}
}