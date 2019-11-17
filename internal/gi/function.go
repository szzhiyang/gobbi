package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Function struct {
	namespace string
	funcName  string
	info      *C.GIFunctionInfo
}

func FunctionInvokerNew(namespace string, funcName string) *Function {
	cNamespace := C.CString(namespace)
	defer C.free(unsafe.Pointer(cNamespace))

	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))

	invoker := C.g_irepository_find_by_name(
		nil,
		cNamespace,
		cFuncName,
	)

	return &Function{
		namespace: namespace,
		funcName:  funcName,
		info:      invoker,
	}
}

func (fi *Function) Invoke(in []Argument, out []Argument) Argument {
	var returnValue Argument
	var err *C.GError

	// prepare in args
	var cIn *C.GIArgument
	var cInLen C.int
	if in != nil {
		cIn = (*C.GIArgument)(&in[0])
		cInLen = C.int(len(in))
	}

	// prepare out args
	var cOut *C.GIArgument
	var cOutLen C.int
	if out != nil {
		cOut = (*C.GIArgument)(&out[0])
		cOutLen = C.int(len(out))
	}

	// invoke
	invoked := C.g_function_info_invoke(
		fi.info,
		cIn, cInLen,
		cOut, cOutLen,
		(*C.GIArgument)(&returnValue),
		&err,
	) == C.TRUE

	// check error
	if err != nil {
		message := C.GoString(err.message)
		panic(message)
	}

	// verify invoke happened
	if !invoked {
		panic(fmt.Sprintf("%s.%s not called", fi.namespace, fi.funcName))
	}

	return returnValue
}
