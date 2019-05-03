// Code generated - DO NOT EDIT.
// +build gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void screen_compositedChangedHandler(GObject *, gpointer);

	static gulong Screen_signal_connect_composited_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "composited-changed", G_CALLBACK(screen_compositedChangedHandler), data);
	}

*/
import "C"

// SupportsInputShapes is a wrapper around the C function gdk_display_supports_input_shapes.
func (recv *Display) SupportsInputShapes() bool {
	retC := C.gdk_display_supports_input_shapes((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SupportsShapes is a wrapper around the C function gdk_display_supports_shapes.
func (recv *Display) SupportsShapes() bool {
	retC := C.gdk_display_supports_shapes((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

type signalScreenCompositedChangedDetail struct {
	callback  ScreenSignalCompositedChangedCallback
	handlerID C.gulong
}

var signalScreenCompositedChangedId int
var signalScreenCompositedChangedMap = make(map[int]signalScreenCompositedChangedDetail)
var signalScreenCompositedChangedLock sync.RWMutex

// ScreenSignalCompositedChangedCallback is a callback function for a 'composited-changed' signal emitted from a Screen.
type ScreenSignalCompositedChangedCallback func()

/*
ConnectCompositedChanged connects the callback to the 'composited-changed' signal for the Screen.

The returned value represents the connection, and may be passed to DisconnectCompositedChanged to remove it.
*/
func (recv *Screen) ConnectCompositedChanged(callback ScreenSignalCompositedChangedCallback) int {
	signalScreenCompositedChangedLock.Lock()
	defer signalScreenCompositedChangedLock.Unlock()

	signalScreenCompositedChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Screen_signal_connect_composited_changed(instance, C.gpointer(uintptr(signalScreenCompositedChangedId)))

	detail := signalScreenCompositedChangedDetail{callback, handlerID}
	signalScreenCompositedChangedMap[signalScreenCompositedChangedId] = detail

	return signalScreenCompositedChangedId
}

/*
DisconnectCompositedChanged disconnects a callback from the 'composited-changed' signal for the Screen.

The connectionID should be a value returned from a call to ConnectCompositedChanged.
*/
func (recv *Screen) DisconnectCompositedChanged(connectionID int) {
	signalScreenCompositedChangedLock.Lock()
	defer signalScreenCompositedChangedLock.Unlock()

	detail, exists := signalScreenCompositedChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScreenCompositedChangedMap, connectionID)
}

//export screen_compositedChangedHandler
func screen_compositedChangedHandler(_ *C.GObject, data C.gpointer) {
	signalScreenCompositedChangedLock.RLock()
	defer signalScreenCompositedChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScreenCompositedChangedMap[index].callback
	callback()
}

// GetActiveWindow is a wrapper around the C function gdk_screen_get_active_window.
func (recv *Screen) GetActiveWindow() *Window {
	retC := C.gdk_screen_get_active_window((*C.GdkScreen)(recv.native))
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFontOptions is a wrapper around the C function gdk_screen_get_font_options.
func (recv *Screen) GetFontOptions() *cairo.FontOptions {
	retC := C.gdk_screen_get_font_options((*C.GdkScreen)(recv.native))
	var retGo (*cairo.FontOptions)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.FontOptionsNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetResolution is a wrapper around the C function gdk_screen_get_resolution.
func (recv *Screen) GetResolution() float64 {
	retC := C.gdk_screen_get_resolution((*C.GdkScreen)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetWindowStack is a wrapper around the C function gdk_screen_get_window_stack.
func (recv *Screen) GetWindowStack() *glib.List {
	retC := C.gdk_screen_get_window_stack((*C.GdkScreen)(recv.native))
	var retGo (*glib.List)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.ListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// IsComposited is a wrapper around the C function gdk_screen_is_composited.
func (recv *Screen) IsComposited() bool {
	retC := C.gdk_screen_is_composited((*C.GdkScreen)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFontOptions is a wrapper around the C function gdk_screen_set_font_options.
func (recv *Screen) SetFontOptions(options *cairo.FontOptions) {
	c_options := (*C.cairo_font_options_t)(C.NULL)
	if options != nil {
		c_options = (*C.cairo_font_options_t)(options.ToC())
	}

	C.gdk_screen_set_font_options((*C.GdkScreen)(recv.native), c_options)

	return
}

// SetResolution is a wrapper around the C function gdk_screen_set_resolution.
func (recv *Screen) SetResolution(dpi float64) {
	c_dpi := (C.gdouble)(dpi)

	C.gdk_screen_set_resolution((*C.GdkScreen)(recv.native), c_dpi)

	return
}

// GetTypeHint is a wrapper around the C function gdk_window_get_type_hint.
func (recv *Window) GetTypeHint() WindowTypeHint {
	retC := C.gdk_window_get_type_hint((*C.GdkWindow)(recv.native))
	retGo := (WindowTypeHint)(retC)

	return retGo
}

// InputShapeCombineRegion is a wrapper around the C function gdk_window_input_shape_combine_region.
func (recv *Window) InputShapeCombineRegion(shapeRegion *cairo.Region, offsetX int32, offsetY int32) {
	c_shape_region := (*C.cairo_region_t)(C.NULL)
	if shapeRegion != nil {
		c_shape_region = (*C.cairo_region_t)(shapeRegion.ToC())
	}

	c_offset_x := (C.gint)(offsetX)

	c_offset_y := (C.gint)(offsetY)

	C.gdk_window_input_shape_combine_region((*C.GdkWindow)(recv.native), c_shape_region, c_offset_x, c_offset_y)

	return
}

// MergeChildInputShapes is a wrapper around the C function gdk_window_merge_child_input_shapes.
func (recv *Window) MergeChildInputShapes() {
	C.gdk_window_merge_child_input_shapes((*C.GdkWindow)(recv.native))

	return
}

// SetChildInputShapes is a wrapper around the C function gdk_window_set_child_input_shapes.
func (recv *Window) SetChildInputShapes() {
	C.gdk_window_set_child_input_shapes((*C.GdkWindow)(recv.native))

	return
}

// AtomInternStaticString is a wrapper around the C function gdk_atom_intern_static_string.
func AtomInternStaticString(atomName string) Atom {
	c_atom_name := C.CString(atomName)
	defer C.free(unsafe.Pointer(c_atom_name))

	retC := C.gdk_atom_intern_static_string(c_atom_name)
	retGo := *AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}