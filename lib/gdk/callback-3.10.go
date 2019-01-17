// This is a generated file - DO NOT EDIT
// +build gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	"sync"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void callback_windowinvalidatehandlerfuncHandler(GObject *, GdkWindow*, cairo_region_t*, gpointer);

*/
import "C"

var callbackWindowinvalidatehandlerfuncId int
var callbackWindowinvalidatehandlerfuncMap = make(map[int]WindowinvalidatehandlerfuncCallback)
var callbackWindowinvalidatehandlerfuncLock sync.RWMutex

// WindowinvalidatehandlerfuncCallback is a callback function for a 'WindowInvalidateHandlerFunc' callback.
type WindowinvalidatehandlerfuncCallback func(window *Window, region *cairo.Region)
