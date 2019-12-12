package callback

// #cgo pkg-config: gobject-2.0
//
// #include <stdlib.h>
// #include "signal.h"
import "C"

import (
	"fmt"
	"github.com/pekim/gobbi/internal/cgo"
	"sync"
	"unsafe"
)

const PackageName = "github.com/pekim/gobbi/internal/cgo/callback"

type Value C.GValue

type Handler func(return_value *Value, paramValues []Value)

type signalConnection struct {
	cConnection C.GobbiSignalConnection
	signalName  string
	handler     Handler
}

var signalConnections = make(map[int]signalConnection)
var signalConnectionsId int
var signalConnectionsMutex sync.Mutex

func ConnectSignal(instance unsafe.Pointer, signalName string, handler Handler) int {
	signalConnectionsMutex.Lock()
	defer signalConnectionsMutex.Unlock()

	id := signalConnectionsId
	signalConnectionsId++

	cSignalName := C.CString(signalName)
	defer C.free(unsafe.Pointer(cSignalName))

	cConnection := C.gobbi_connect_signal(instance, cSignalName, unsafe.Pointer(uintptr(id)))

	connection := signalConnection{
		cConnection: cConnection,
		signalName:  signalName,
		handler:     handler,
	}

	signalConnections[id] = connection

	return id
}

func DisconnectSignal(id int) {
	signalConnectionsMutex.Lock()
	defer signalConnectionsMutex.Unlock()

	connection := signalConnections[id]
	C.gobbi_disconnect_signal(connection.cConnection)

	delete(signalConnections, id)
}

//export gobbiCSignalHandler
func gobbiCSignalHandler(cReturnValue *C.GValue, nParamValues C.guint, cParamValues *C.GValue, data unsafe.Pointer) {
	id := int(uintptr(data))
	connection := signalConnections[id]

	returnValue := (*Value)(cReturnValue)
	paramValues := (*[1 << 30]Value)(unsafe.Pointer(cParamValues))[:nParamValues:nParamValues]

	if cgo.Tracing() {
		traceCall(connection, paramValues)
	}

	connection.handler(returnValue, paramValues)

	if cgo.Tracing() {
		traceReturn(connection, returnValue, paramValues)
	}
}

func traceCall(connection signalConnection, paramValues []Value) {
	args := formatTraceParams(paramValues)
	cgo.Trace(fmt.Sprintf("Signal start : %s\n%s\n", connection.signalName, args))
}

// traceReturn includes param values that might be out parameters.
// So it must be called after the handler has returned.
func traceReturn(connection signalConnection, returnValue *Value, paramValues []Value) {
	ret := ""
	if returnValue != nil {
		ret = fmt.Sprintf("   ret %v\n", *returnValue)
	}

	args := formatTraceParams(paramValues)

	cgo.Trace(fmt.Sprintf("Signal end : %s\n%s%s\n",
		connection.signalName, args, ret))
}

func formatTraceParams(paramValues []Value) string {
	if len(paramValues) > 0 {
		return fmt.Sprintf("  args %v\n", paramValues)
	}

	return ""
}
