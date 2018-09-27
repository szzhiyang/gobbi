// This is a generated file - DO NOT EDIT
// +build gdk_2.18 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.16 gdk_3.20 gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// OffscreenWindowGetEmbedder is a wrapper around the C function gdk_offscreen_window_get_embedder.
func OffscreenWindowGetEmbedder(window *Window) *Window {
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gdk_offscreen_window_get_embedder(c_window)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_offscreen_window_set_embedder : no return generator